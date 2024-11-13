package lexer

import (
	"io"

	"github.com/pattyshack/gt/lexutil"

	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/errors"
	"github.com/pattyshack/pl/parser/lr"
)

type scope struct {
	firstLbrace bool
	endOfScope  bool
}

type ScopedLexer struct {
	*errors.Emitter

	buffered *lexutil.BufferedReader[lr.Token]
	base     lr.Lexer
	reducer  lr.Reducer

	stack []scope
}

func NewLexer(
	sourceFileName string,
	sourceContent io.Reader,
	emitter *errors.Emitter,
	reducer lr.Reducer,
	options LexerOptions,
) *ScopedLexer {
	base := NewBasicLexer(sourceFileName, sourceContent, emitter, options)
	return &ScopedLexer{
		Emitter: emitter,
		buffered: lexutil.NewBufferedReader(
			lexutil.NewLexerReader[lr.Token](base),
			10),
		base:    base,
		reducer: reducer,
	}
}

func (lexer *ScopedLexer) CurrentLocation() lexutil.Location {
	peeked, err := lexer.buffered.Peek(1)
	if err != nil || len(peeked) == 0 {
		return lexer.base.CurrentLocation()
	}

	return peeked[0].Loc()
}

func (lexer *ScopedLexer) isCurrentScopeFirstLbrace() bool {
	if len(lexer.stack) == 0 {
		return false
	}

	return lexer.stack[len(lexer.stack)-1].firstLbrace
}

func (lexer *ScopedLexer) consumeCurrentScopeFirstLbrace() {
	if len(lexer.stack) == 0 {
		return
	}

	lexer.stack[len(lexer.stack)-1].firstLbrace = false
}

func (lexer *ScopedLexer) isEndOfCurrentScope() bool {
	if len(lexer.stack) == 0 {
		return false
	}

	return lexer.stack[len(lexer.stack)-1].endOfScope
}

func (lexer *ScopedLexer) endCurrentScope() {
	if len(lexer.stack) == 0 { // rbrace mismatch
		return
	}

	lexer.stack[len(lexer.stack)-1].endOfScope = true
}

func (lexer *ScopedLexer) pushScope() int {
	lexer.stack = append(
		lexer.stack,
		scope{
			firstLbrace: true,
			endOfScope:  false,
		})

	return len(lexer.stack)
}

func (lexer *ScopedLexer) popScope() {
	lexer.stack = lexer.stack[:len(lexer.stack)-1]
}

func (lexer *ScopedLexer) Next() (lr.Token, error) {
	if lexer.isEndOfCurrentScope() {
		lexer.popScope()
		return nil, io.EOF
	}

	peeked, err := lexer.buffered.Peek(1)
	if err != nil || len(peeked) == 0 {
		return lexer.buffered.Next()
	}

	switch peeked[0].Id() {
	case lr.LbraceToken:
		if lexer.isCurrentScopeFirstLbrace() {
			lexer.consumeCurrentScopeFirstLbrace()
			return lexer.buffered.Next()
		}

		currentLevel := lexer.pushScope()
		pos := lexer.CurrentLocation()
		block, parseErr := lr.ParseStatements(lexer, lexer.reducer)
		if parseErr == nil {
			if currentLevel != len(lexer.stack)+1 { // sanity check
				panic("Should never happen")
			}
			return &lr.Symbol{
				SymbolId_:      lr.StatementsType,
				StatementsExpr: block,
			}, nil
		}

		lexer.EmitErrors(parseErr)
		startEnd := lexutil.NewStartEndPos(pos, lexer.CurrentLocation())
		stmts := &ast.StatementsExpr{
			StartEndPos: startEnd,
			Statements: []ast.Statement{
				&ast.ParseErrorExpr{
					StartEndPos: startEnd,
					Error:       parseErr,
				},
			},
		}

		for len(lexer.stack) >= currentLevel {
			if lexer.isEndOfCurrentScope() {
				lexer.popScope()
				continue
			}

			peeked, err := lexer.buffered.Peek(1)
			if err != nil || len(peeked) == 0 {
				break
			}

			switch peeked[0].Id() {
			case lr.LbraceToken:
				curr := lexer.CurrentLocation()
				lexer.pushScope()
				_, parseErr = lr.ParseStatements(lexer, lexer.reducer)
				if parseErr != nil {
					lexer.EmitErrors(parseErr)
					stmts.Statements = append(
						stmts.Statements,
						&ast.ParseErrorExpr{
							StartEndPos: lexutil.NewStartEndPos(curr, lexer.CurrentLocation()),
							Error:       parseErr,
						})
				}
			case lr.RbraceToken:
				lexer.popScope()
				fallthrough
			default:
				_, err := lexer.buffered.Discard(1)
				if err != nil {
					panic("Should never happen")
				}
			}
		}

		stmts.EndPos = lexer.CurrentLocation()
		return &lr.Symbol{
			SymbolId_:      lr.StatementsType,
			StatementsExpr: stmts,
		}, nil
	case lr.RbraceToken:
		lexer.endCurrentScope()
	}

	return lexer.buffered.Next()
}
