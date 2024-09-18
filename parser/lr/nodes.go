package lr

import (
	"fmt"

	. "github.com/pattyshack/pl/ast"
)

const (
	DecimalInteger            = "decimal integer literal"
	HexadecimalInteger        = "hexadecimal integer literal"
	ZeroOPrefixedOctalInteger = "0o-prefixed octal integer literal"
	ZeroPrefixedOctalInteger  = "0-prefixed octal integer literal"
	BinaryInteger             = "binary integer literal"

	SingleLineString    = "single line string literal"
	MultiLineString     = "mutli line string literal"
	RawSingleLineString = "raw single line string literal"
	RawMultiLineString  = "raw mutli line string literal"

	DecimalFloat     = "decimal float literal"
	HexadecimalFloat = "hexadecimal float literal"

	CommentGroupTokenId = SymbolId(-4)
)

type CommentGroupToken struct {
	CommentGroup
}

func (CommentGroupToken) Id() SymbolId { return CommentGroupTokenId }

type CommentGroupsTok struct {
	CommentGroups
}

func (CommentGroupsTok) Id() SymbolId { return CommentGroupsToken }

type TokenCount struct {
	SymbolId
	StartEndPos
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
	StartEndPos
	LeadingTrailingComments

	// The value string is optional if the value is fully determined by SymbolId.
	Value string
	// Used for classifying literal subtypes.
	SubType string
}

func (s TokenValue) Id() SymbolId {
	return s.SymbolId
}

func (s TokenValue) Val() string {
	return s.Value
}

func (s TokenValue) TreeString(indent string, label string) string {
	return fmt.Sprintf(
		"%s%s[TokenValue: SymbolId=%s Value=%s]",
		indent,
		label,
		s.SymbolId,
		s.Value)
}

type ParseErrorSymbol struct {
	*ParseErrorNode
}

func (ParseErrorSymbol) Id() SymbolId {
	return ParseErrorToken
}
