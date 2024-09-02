package parser

type IntegerLiteralSubType string
type StringLiteralSubType string
type FloatLiteralSubType string

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

	DecimalFloat     = FloatLiteralSubType("decimal float literal")
	HexadecimalFloat = FloatLiteralSubType("hexadecimal float literal")
)

type ParseErrorSymbol struct {
	StartPos Location
	EndPos   Location
	Error    error
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
	Value    string
}

func (s ValueSymbol) Id() SymbolId { return s.SymbolId }

func (s ValueSymbol) Loc() Location { return s.StartPos }
func (s ValueSymbol) End() Location { return s.EndPos }

type LiteralSymbol[T any] struct {
	SymbolId
	StartPos Location
	EndPos   Location
	Value    string
	SubType  T
}

func (s LiteralSymbol[T]) Id() SymbolId { return s.SymbolId }

func (s LiteralSymbol[T]) Loc() Location { return s.StartPos }
func (s LiteralSymbol[T]) End() Location { return s.EndPos }

type IntegerLiteralSymbol = LiteralSymbol[IntegerLiteralSubType]
type RuneLiteralSymbol = LiteralSymbol[struct{}]
type StringLiteralSymbol = LiteralSymbol[StringLiteralSubType]
type FloatLiteralSymbol = LiteralSymbol[FloatLiteralSubType]
