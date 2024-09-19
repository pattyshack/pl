package parser

import (
	"io"

	"github.com/pattyshack/gt/lexutil"

	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser/lexer"
	"github.com/pattyshack/pl/parser/lr"
	reducerImpl "github.com/pattyshack/pl/parser/reducer"
)

type ParserOptions struct {
	lexer.LexerOptions

	UseLRParseSource bool
}

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
// newlines or eof.  Note that newlines inside lparen/rparen scopes are
// ignored (newlines inside lbrace/rbrace scopes are handled by scoped lexer).
// Parse errors are returned as part of the definition list.
type sourceParser struct {
	lexer   lr.Lexer
	reducer *reducerImpl.Reducer
}

func newSourceParser(
	lexer lr.Lexer,
	reducer *reducerImpl.Reducer,
) *sourceParser {
	return &sourceParser{
		lexer:   lexer,
		reducer: reducer,
	}
}

func (parser *sourceParser) readDefinitionSegment() ([]lr.Token, error) {
	parenScope := 0
	result := []lr.Token{}
	for {
		token, err := parser.lexer.Next()
		if err != nil {
			return result, err
		}

		switch token.Id() {
		case lr.NewlinesToken:
			if parenScope == 0 {
				return result, nil
			}
		case lr.LparenToken:
			parenScope++
		case lr.RparenToken:
			parenScope--
		}

		result = append(result, token)

		if parenScope < 0 {
			// mismatch paren.  this will always result in an parse error.
			return result, nil
		}
	}
}

func (parser *sourceParser) parse() (*ast.DefinitionList, error) {
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
				def, err = parser.reducer.ToParseErrorExpr(
					lr.ParseErrorSymbol{
						ast.NewParseErrorNode(ast.NewStartEndPos(start, end), err),
					})
				if err != nil {
					return list, err
				}
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

func ParseSource(
	fileName string,
	content io.Reader,
	options ParserOptions,
) (
	*reducerImpl.Reducer,
	*ast.DefinitionList,
	error,
) {
	reducer := reducerImpl.NewReducer()
	lexer := lexer.NewLexer(fileName, content, options.LexerOptions, reducer)

	var list *ast.DefinitionList
	var err error
	if options.UseLRParseSource {
		list, err = lr.ParseSource(lexer, reducer)
	} else {
		parser := newSourceParser(lexer, reducer)
		list, err = parser.parse()
	}

	return reducer, list, err
}
