package parser

import (
	"io"
	"sync"

	"github.com/pattyshack/gt/lexutil"

	"github.com/pattyshack/pl/analyze"
	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/build/cfg"
	"github.com/pattyshack/pl/errors"
	"github.com/pattyshack/pl/parser/lr"
	reducerImpl "github.com/pattyshack/pl/parser/reducer"
)

type bufferLexer struct {
	*lexutil.BufferedReader[lr.Token]
	end lexutil.Location
}

func (l bufferLexer) CurrentLocation() lexutil.Location {
	peeked, err := l.Peek(1)
	if err != nil || len(peeked) == 0 {
		return l.end
	}

	return peeked[0].Loc()
}

// This splits the source token stream into definition segments, then parse
// each definition segments individually.  Each segment is terminated by either
// newlines or eof.  Note that newlines inside paren/brace scopes are
// ignored.  Parse errors are returned as part of the definition list.
type sourceParser struct {
	lexer lr.Lexer
	*reducerImpl.Reducer
	*cfg.Config
	*errors.Emitter
	ParserOptions
}

func newSourceParser(
	fileName string,
	config *cfg.Config,
	emitter *errors.Emitter,
	options ParserOptions,
) (
	*sourceParser,
	error,
) {
	reducer := reducerImpl.NewReducer(emitter)
	lexer, err := options.NewLexer(fileName, emitter, reducer)
	if err != nil {
		return nil, err
	}

	return &sourceParser{
		lexer:         lexer,
		Reducer:       reducer,
		Config:        config,
		Emitter:       emitter,
		ParserOptions: options,
	}, nil
}

func (parser *sourceParser) readDefinitionSegment() ([]lr.Token, error) {
	parenScope := 0
	braceScope := 0
	result := []lr.Token{}
	for {
		token, err := parser.lexer.Next()
		if err != nil {
			return result, err
		}

		switch token.Id() {
		case lr.NewlinesToken:
			if parenScope == 0 && braceScope == 0 {
				return result, nil
			}
		case lr.LparenToken:
			parenScope++
		case lr.RparenToken:
			parenScope--
		case lr.LbraceToken:
			braceScope++
		case lr.RbraceToken:
			braceScope--
		}

		result = append(result, token)

		if parenScope < 0 || braceScope < 0 {
			// mismatch paren/brace.  This will always result in an parse error.
			return result, nil
		}
	}
}

func (parser *sourceParser) _parseSource() (*ast.StatementList, error) {
	if parser.UseLRParseSource {
		return lr.ParseSource(parser.lexer, parser)
	}

	list := ast.NewStatementList()
	for {
		start := parser.lexer.CurrentLocation()
		segment, readErr := parser.readDefinitionSegment()
		end := parser.lexer.CurrentLocation()

		if len(segment) > 0 {
			stmt, err := lr.ParseStatement(
				&bufferLexer{
					BufferedReader: lexutil.NewBufferedReaderFromSlice(segment),
					end:            end,
				},
				parser)
			if err != nil {
				parser.EmitErrors(err)
				stmt = &ast.ParseErrorExpr{
					StartEndPos: ast.NewStartEndPos(start, end),
					Error:       err,
				}
			}

			list.Add(stmt)
		}

		if readErr != nil {
			if readErr == io.EOF {
				readErr = nil
			}

			return list, readErr
		}
	}
}

type ParsedSource struct {
	Definitions            *ast.StatementList
	Dependencies           []*ast.ImportClause
	FailedBuildConstraints bool
}

func (parser *sourceParser) parseSource() ParsedSource {
	source, err := parser._parseSource()
	if err != nil {
		parser.EmitErrors(err)
		return ParsedSource{}
	}

	result := ParsedSource{
		Definitions: source,
	}

	if !parser.DisableSyntaxAnalysis {
		deps, satisfy := analyze.Source(parser.Config, source, parser.Emitter)
		result.Dependencies = deps
		result.FailedBuildConstraints = !satisfy
	}

	return result
}

func ParseSource(
	fileName string,
	config *cfg.Config,
	emitter *errors.Emitter,
	options ParserOptions,
) ParsedSource {
	parser, err := newSourceParser(fileName, config, emitter, options)
	if err != nil {
		emitter.EmitErrors(err)
		return ParsedSource{}
	}

	return parser.parseSource()
}

func ParseSources(
	sources []string,
	config *cfg.Config,
	emitter *errors.Emitter,
	options ParserOptions,
) []ParsedSource {
	if len(sources) == 0 {
		return nil
	}

	results := make([]ParsedSource, len(sources), len(sources))

	wg := &sync.WaitGroup{}
	wg.Add(len(sources))

	for idx, src := range sources {
		go func(idx int, fileName string) {
			defer wg.Done()
			results[idx] = ParseSource(fileName, config, emitter, options)
		}(idx, src)
	}

	wg.Wait()

	return results
}
