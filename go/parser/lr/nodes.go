package lr

import (
	"github.com/pattyshack/gt/lexutil"

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

type TokenCount = lexutil.TokenCount[SymbolId]

type TokenValue struct {
	lexutil.TokenValue[SymbolId]

	ast.LeadingTrailingComments
}

func (s TokenValue) Walk(visitor ast.Visitor) {
	panic("should never be called")
}

type ParseErrorSymbol ast.ParseErrorExpr

func NewParseErrorSymbol(
	pos lexutil.StartEndPos,
	format string,
	args ...interface{},
) *ParseErrorSymbol {
	return &ParseErrorSymbol{
		StartEndPos: pos,
		Error:       lexutil.NewLocationError(pos.StartPos, format, args...),
	}
}

func (ParseErrorSymbol) Id() SymbolId {
	return ParseErrorToken
}
