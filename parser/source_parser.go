package parser

import (
	"io"

	"github.com/pattyshack/gt/lexutil"

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
	lexer   lr.Lexer
	reducer *reducerImpl.Reducer
	ParserOptions
}

func newSourceParser(
	fileName string,
	options ParserOptions,
) (
	*sourceParser,
	error,
) {
	reducer := reducerImpl.NewReducer()
	lexer, err := options.NewLexer(fileName, reducer)
	if err != nil {
		return nil, err
	}

	return &sourceParser{
		lexer:         lexer,
		reducer:       reducer,
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
		return lr.ParseSource(parser.lexer, parser.reducer)
	}

	list := ast.NewDefinitionList()
	for {
		start := parser.lexer.CurrentLocation()
		segment, err := parser.readDefinitionSegment()
		end := parser.lexer.CurrentLocation()

		if len(segment) > 0 {
			def, err := lr.ParseDefinition(
				&bufferLexer{
					BufferedReader: lexutil.NewBufferedReaderFromSlice(segment),
					end:            end,
				},
				parser.reducer)
			if err != nil {
				parser.reducer.ParseErrors = append(parser.reducer.ParseErrors, err)
				def = ast.NewParseErrorNode(ast.NewStartEndPos(start, end), err)
			}

			if len(list.Elements) == 0 {
				list.Add(def)
			} else {
				list.ReduceAdd(&lr.TokenValue{}, def)
			}
		}

		if err != nil {
			if err == io.EOF {
				err = nil
			}

			return list, err
		}
	}
}

func (parser *sourceParser) parseSource() (
	*reducerImpl.Reducer,
	*ast.DefinitionList,
	error,
) {
	source, err := parser._parseSource()
	if err != nil {
		return nil, nil, err
	}

	return parser.reducer, source, nil
}

func (parser *sourceParser) parseExpr() (
	*reducerImpl.Reducer,
	ast.Expression,
	error,
) {
	expr, err := lr.ParseExpr(parser.lexer, parser.reducer)
	if err != nil {
		return nil, nil, err
	}

	return parser.reducer, expr, nil
}

func (parser *sourceParser) parseTypeExpr() (
	*reducerImpl.Reducer,
	ast.TypeExpression,
	error,
) {
	typeExpr, err := lr.ParseTypeExpr(parser.lexer, parser.reducer)
	if err != nil {
		return nil, nil, err
	}

	return parser.reducer, typeExpr, nil
}

func (parser *sourceParser) parseStatement() (
	*reducerImpl.Reducer,
	ast.Statement,
	error,
) {
	stmt, err := lr.ParseStatement(parser.lexer, parser.reducer)
	if err != nil {
		return nil, nil, err
	}

	return parser.reducer, stmt, nil
}

func (parser *sourceParser) parseDefinition() (
	*reducerImpl.Reducer,
	ast.Definition,
	error,
) {
	def, err := lr.ParseDefinition(parser.lexer, parser.reducer)
	if err != nil {
		return nil, nil, err
	}

	return parser.reducer, def, nil
}

func ParseExpr(
	fileName string,
	options ParserOptions,
) (
	*reducerImpl.Reducer,
	ast.Expression,
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
	*reducerImpl.Reducer,
	ast.TypeExpression,
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
	*reducerImpl.Reducer,
	ast.Statement,
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
	*reducerImpl.Reducer,
	ast.Definition,
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
	*reducerImpl.Reducer,
	*ast.DefinitionList,
	error,
) {
	parser, err := newSourceParser(fileName, options)
	if err != nil {
		return nil, nil, err
	}

	return parser.parseSource()
}
