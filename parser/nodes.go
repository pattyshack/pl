package parser

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
)

// A comment group is a single block comment, or a group of line comments
// separated by single newlines (and spaces).
type CommentGroup []ValueSymbol

// Trailing comment groups are comment groups that immediately follow a
// symbol, and are on the same line as the symbol.
//
// For example,
//
//	OTHER_SYMBOL /* group 0 */
//	// group 1 line 1
//	// group 1 line 2
//
//	// group 2
//	/* group 3 */ SYMBOL /* group 4 */ // group 5 line 1
//	// group 5 line 2
//
// // other group 6
//
// SYMBOL's leading comment groups are {1, 2, 3}.  SYMBOL's trailing comment
// groups are {4, 5}. Comment groups {0, 6} belong to other symbols.
type CommentGroups struct {
	LeadingComments []CommentGroup

	TrailingComments []CommentGroup
}

type ParseErrorSymbol struct {
	StartPos Location
	EndPos   Location
	CommentGroups

	Error error
}

func (s ParseErrorSymbol) Id() SymbolId { return ParseErrorToken }

func (s ParseErrorSymbol) Loc() Location { return s.StartPos }
func (s ParseErrorSymbol) End() Location { return s.EndPos }

func (s ParseErrorSymbol) String() string {
	return s.Error.Error() + " " + s.StartPos.String()
}

type CountSymbol struct {
	SymbolId
	StartPos Location
	EndPos   Location
	Count    int
}

func (s CountSymbol) Id() SymbolId { return s.SymbolId }

func (s CountSymbol) Loc() Location { return s.StartPos }
func (s CountSymbol) End() Location { return s.EndPos }

type ValueSymbol struct {
	SymbolId
	StartPos Location
	EndPos   Location
	CommentGroups

	Value   string
	SubType string
}

func (s ValueSymbol) Id() SymbolId { return s.SymbolId }

func (s ValueSymbol) Loc() Location { return s.StartPos }
func (s ValueSymbol) End() Location { return s.EndPos }
