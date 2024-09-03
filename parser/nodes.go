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
// separated by single newlines (ignoring spaces).
type CommentGroup []TokenValue

func (CommentGroup) Id() SymbolId { return commentGroupToken }

func (cg CommentGroup) Loc() Location { return cg[0].Loc() }
func (cg CommentGroup) End() Location { return cg[len(cg)-1].End() }

// A symbol's trailing comment groups are comment groups that are on the same
// line as the symbol and immediately follows the symbol.
//
// A symbol's leading comment groups are any comment group immediately before
// the symbol that do not belong to previous symbols.
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
// /* other group 6 */ OTHER_SYMBOL
//
// SYMBOL's leading comment groups are {1, 2, 3}.  SYMBOL's trailing comment
// groups are {4, 5}. Comment groups {0, 6} belong to other symbols.
type CommentGroups struct {
	Groups []CommentGroup
}

func (CommentGroups) Id() SymbolId { return CommentGroupsToken }

func (cgs CommentGroups) Loc() Location {
	return cgs.Groups[0].Loc()
}

func (cgs CommentGroups) End() Location {
	return cgs.Groups[len(cgs.Groups)-1].End()
}

func (cgs *CommentGroups) Prepend(other CommentGroups) {
	cgs.Groups = append(other.Groups, cgs.Groups...)
}

func (cgs *CommentGroups) Append(other CommentGroups) {
	cgs.Groups = append(cgs.Groups, other.Groups...)
}

type ParseErrorSymbol struct {
	StartPos        Location
	EndPos          Location
	LeadingComment  CommentGroups
	TrailingComment CommentGroups

	Error error
}

func (s ParseErrorSymbol) Id() SymbolId { return ParseErrorToken }

func (s ParseErrorSymbol) Loc() Location { return s.StartPos }
func (s ParseErrorSymbol) End() Location { return s.EndPos }

func (s ParseErrorSymbol) String() string {
	return s.Error.Error() + " " + s.StartPos.String()
}

type TokenCount struct {
	SymbolId
	StartPos Location
	EndPos   Location
	Count    int
}

func (s TokenCount) Id() SymbolId { return s.SymbolId }

func (s TokenCount) Loc() Location { return s.StartPos }
func (s TokenCount) End() Location { return s.EndPos }

type TokenValue struct {
	SymbolId
	StartPos        Location
	EndPos          Location
	LeadingComment  CommentGroups
	TrailingComment CommentGroups

	// The value string is optional if the value is fully determined by SymbolId.
	Value string
	// Used for classifying literal subtypes.
	SubType string
}

func (s TokenValue) Id() SymbolId { return s.SymbolId }

func (s TokenValue) Loc() Location { return s.StartPos }
func (s TokenValue) End() Location { return s.EndPos }
