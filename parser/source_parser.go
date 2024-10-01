package parser

import (
	"io"

	"github.com/pattyshack/gt/lexutil"

	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser/lr"
	reducerImpl "github.com/pattyshack/pl/parser/reducer"
	"github.com/pattyshack/pl/util"
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
	ParserOptions
}

func newSourceParser(
	fileName string,
	options ParserOptions,
) (
	*sourceParser,
	error,
) {
	emitter := &util.ErrorEmitter{}
	reducer := reducerImpl.NewReducer(emitter)
	lexer, err := options.NewLexer(fileName, emitter, reducer)
	if err != nil {
		return nil, err
	}

	return &sourceParser{
		lexer:         lexer,
		Reducer:       reducer,
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

func (parser *sourceParser) _parseSource() (*ast.DefinitionList, error) {
	if parser.UseLRParseSource {
		return lr.ParseSource(parser.lexer, parser)
	}

	list := ast.NewDefinitionList()
	for {
		start := parser.lexer.CurrentLocation()
		segment, readErr := parser.readDefinitionSegment()
		end := parser.lexer.CurrentLocation()

		if len(segment) > 0 {
			def, err := lr.ParseDefinition(
				&bufferLexer{
					BufferedReader: lexutil.NewBufferedReaderFromSlice(segment),
					end:            end,
				},
				parser)
			if err != nil {
				parser.EmitErrors(err)
				def = &ast.ParseErrorExpr{
					StartEndPos: ast.NewStartEndPos(start, end),
					Error:       err,
				}
			}

			if len(list.Elements) == 0 {
				list.Add(def)
			} else {
				list.ReduceAdd(&lr.TokenValue{}, def)
			}
		}

		if readErr != nil {
			if readErr == io.EOF {
				readErr = nil
			}

			return list, readErr
		}
	}
}

func (parser *sourceParser) analyze(node ast.Node) {
	passes := [][]util.Pass{
		{
			detectUnreachableStatements(),
			detectUnexpectedStatements(),
			detectUnexpectedArguments(),
			detectUnexpectedDefaultBranches(),
		},
	}

	parser.EmitErrors(util.Process(node, passes, 0)...)
}

func (parser *sourceParser) parseSource() (
	*ast.DefinitionList,
	[]error,
	error,
) {
	source, err := parser._parseSource()
	if err != nil {
		return nil, nil, err
	}

	parser.analyze(source)
	return source, parser.Errors(), nil
}

func (parser *sourceParser) parseExpr() (
	ast.Expression,
	[]error,
	error,
) {
	expr, err := lr.ParseExpr(parser.lexer, parser)
	if err != nil {
		return nil, nil, err
	}

	parser.analyze(expr)
	return expr, parser.Errors(), nil
}

func (parser *sourceParser) parseTypeExpr() (
	ast.TypeExpression,
	[]error,
	error,
) {
	typeExpr, err := lr.ParseTypeExpr(parser.lexer, parser)
	if err != nil {
		return nil, nil, err
	}

	parser.analyze(typeExpr)
	return typeExpr, parser.Errors(), nil
}

func (parser *sourceParser) parseStatement() (
	ast.Statement,
	[]error,
	error,
) {
	stmt, err := lr.ParseStatement(parser.lexer, parser)
	if err != nil {
		return nil, nil, err
	}

	parser.analyze(stmt)
	return stmt, parser.Errors(), nil
}

func (parser *sourceParser) parseDefinition() (
	ast.Definition,
	[]error,
	error,
) {
	def, err := lr.ParseDefinition(parser.lexer, parser)
	if err != nil {
		return nil, nil, err
	}

	parser.analyze(def)
	return def, parser.Errors(), nil
}

func ParseExpr(
	fileName string,
	options ParserOptions,
) (
	ast.Expression,
	[]error,
	error,
) {
	parser, err := newSourceParser(fileName, options)
	if err != nil {
		return nil, nil, err
	}

	return parser.parseExpr()
}

func ParseTypeExpr(
	fileName string,
	options ParserOptions,
) (
	ast.TypeExpression,
	[]error,
	error,
) {
	parser, err := newSourceParser(fileName, options)
	if err != nil {
		return nil, nil, err
	}

	return parser.parseTypeExpr()
}

func ParseStatement(
	fileName string,
	options ParserOptions,
) (
	ast.Statement,
	[]error,
	error,
) {
	parser, err := newSourceParser(fileName, options)
	if err != nil {
		return nil, nil, err
	}

	return parser.parseStatement()
}

func ParseDefinition(
	fileName string,
	options ParserOptions,
) (
	ast.Definition,
	[]error,
	error,
) {
	parser, err := newSourceParser(fileName, options)
	if err != nil {
		return nil, nil, err
	}

	return parser.parseDefinition()
}

func ParseSource(
	fileName string,
	options ParserOptions,
) (
	*ast.DefinitionList,
	[]error,
	error,
) {
	parser, err := newSourceParser(fileName, options)
	if err != nil {
		return nil, nil, err
	}

	return parser.parseSource()
}
