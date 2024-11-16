package lr

import (
	"github.com/pattyshack/gt/parseutil"

	"github.com/pattyshack/pl/ast"
)

const (
	CommentGroupTokenId = SymbolId(-4)
)

type CommentGroupToken struct {
	ast.CommentGroup
}

func (CommentGroupToken) Id() SymbolId { return CommentGroupTokenId }

type CommentGroupsTok struct {
	ast.CommentGroups
}

func (CommentGroupsTok) Id() SymbolId { return CommentGroupsToken }

type TokenCount = parseutil.TokenCount[SymbolId]

type TokenValue struct {
	parseutil.TokenValue[SymbolId]

	ast.LeadingTrailingComments
}

func (s TokenValue) Walk(visitor ast.Visitor) {
	panic("should never be called")
}

type ParseErrorSymbol ast.ParseErrorExpr

func NewParseErrorSymbol(
	pos parseutil.StartEndPos,
	format string,
	args ...interface{},
) *ParseErrorSymbol {
	return &ParseErrorSymbol{
		StartEndPos: pos,
		Error:       parseutil.NewLocationError(pos.StartPos, format, args...),
	}
}

func (ParseErrorSymbol) Id() SymbolId {
	return ParseErrorToken
}

type Token = parseutil.Token[SymbolId]
type Lexer = parseutil.Lexer[Token]
