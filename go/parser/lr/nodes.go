package lr

import (
	"fmt"

	"github.com/pattyshack/gt/lexutil"

	"github.com/pattyshack/pl/ast"
)

type LiteralSubType string

func (t LiteralSubType) String() string { return string(t) }

const (
	DecimalInteger            = LiteralSubType("decimal integer")
	HexadecimalInteger        = LiteralSubType("hexadecimal integer")
	ZeroOPrefixedOctalInteger = LiteralSubType("0o-prefixed octal integer")
	ZeroPrefixedOctalInteger  = LiteralSubType("0-prefixed octal integer")
	BinaryInteger             = LiteralSubType("binary integer")

	DecimalFloat     = LiteralSubType("decimal float")
	HexadecimalFloat = LiteralSubType("hexadecimal float")

	SingleLineString    = LiteralSubType("single line string")
	MultiLineString     = LiteralSubType("mutli line string")
	RawSingleLineString = LiteralSubType("raw single line string")
	RawMultiLineString  = LiteralSubType("raw mutli line string")

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

type TokenCount struct {
	SymbolId
	ast.StartEndPos
	Count int
}

func (s TokenCount) Id() SymbolId {
	return s.SymbolId
}

func (s TokenCount) String() string {
	return fmt.Sprintf("[Symbol:%s Count=%d]", s.SymbolId, s.Count)
}

type TokenValue struct {
	SymbolId
	ast.StartEndPos
	ast.LeadingTrailingComments

	// The value string is optional if the value is fully determined by SymbolId.
	Value string
	// Used for classifying literal subtypes.
	SubType LiteralSubType
}

func (s TokenValue) Id() SymbolId {
	return s.SymbolId
}

func (s TokenValue) Val() string {
	return s.Value
}

func (s TokenValue) Walk(visitor ast.Visitor) {
	panic("should never be called")
}

type ParseErrorSymbol ast.ParseErrorExpr

func NewParseErrorSymbol(
	start lexutil.Location,
	end lexutil.Location,
	format string,
	args ...interface{},
) *ParseErrorSymbol {
	return &ParseErrorSymbol{
		StartEndPos: ast.NewStartEndPos(start, end),
		Error:       lexutil.NewLocationError(start, format, args...),
	}
}

func (ParseErrorSymbol) Id() SymbolId {
	return ParseErrorToken
}
