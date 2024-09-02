package parser

type IntegerLiteralSubType string

type StringLiteralSubType string

const (
	DecimalInteger     = IntegerLiteralSubType("decimal integer literal")
	HexadecimalInteger = IntegerLiteralSubType(
		"hexadecimal integer literal")
	ZeroOPrefixedOctalInteger = IntegerLiteralSubType(
		"0o-prefixed octal integer literal")
	ZeroPrefixedOctalInteger = IntegerLiteralSubType(
		"0-prefixed octal integer literal")
	BinaryInteger = IntegerLiteralSubType("binary integer literal")

	SingleLineString    = StringLiteralSubType("single line string literal")
	MultiLineString     = StringLiteralSubType("mutli line string literal")
	RawSingleLineString = StringLiteralSubType("raw single line string literal")
	RawMultiLineString  = StringLiteralSubType("raw mutli line string literal")
)

type ParseErrorSymbol struct {
	Location
	Error error
}

func (s ParseErrorSymbol) Id() SymbolId { return ParseErrorToken }

func (s ParseErrorSymbol) Loc() Location { return s.Location }

func (s ParseErrorSymbol) String() string {
	return s.Error.Error() + " " + s.Location.String()
}

type CountSymbol struct {
	SymbolId
	Location
	Count int
}

func (s CountSymbol) Id() SymbolId { return s.SymbolId }

func (s CountSymbol) Loc() Location { return s.Location }

type ValueSymbol struct {
	SymbolId
	Location
	Value string
}

func (s ValueSymbol) Id() SymbolId { return s.SymbolId }

func (s ValueSymbol) Loc() Location { return s.Location }

type LiteralSymbol[T any] struct {
	SymbolId
	Location
	Value   string
	SubType T
}

func (s LiteralSymbol[T]) Id() SymbolId { return s.SymbolId }

func (s LiteralSymbol[T]) Loc() Location { return s.Location }

type IntegerLiteralSymbol = LiteralSymbol[IntegerLiteralSubType]
type RuneLiteralSymbol = LiteralSymbol[struct{}]
type StringLiteralSymbol = LiteralSymbol[StringLiteralSubType]
