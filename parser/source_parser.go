package parser

import (
	"io"

	"github.com/pattyshack/gt/lexutil"

	"github.com/pattyshack/pl/analyze"
	"github.com/pattyshack/pl/ast"
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
	*lexutil.ErrorEmitter
	ParserOptions
}

func newSourceParser(
	fileName string,
	emitter *lexutil.ErrorEmitter,
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
		ErrorEmitter:  emitter,
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

func (parser *sourceParser) parseSource() *ast.StatementList {
	source, err := parser._parseSource()
	if err != nil {
		parser.EmitErrors(err)
		return nil
	}

	if parser.AnalyzeSyntax {
		analyze.SourceSyntax(source, parser.ErrorEmitter)
	}
	return source
}

func ParseSource(
	fileName string,
	emitter *lexutil.ErrorEmitter,
	options ParserOptions,
) *ast.StatementList {
	parser, err := newSourceParser(fileName, emitter, options)
	if err != nil {
		emitter.EmitErrors(err)
		return nil
	}

	return parser.parseSource()
}
