package parser

type IntegerSubType string

const (
	DecimalInteger            = IntegerSubType("decimal")
	HexadecimalInteger        = IntegerSubType("hexadecimal")
	ZeroOPrefixedOctalInteger = IntegerSubType("0o-prefixed octal")
	ZeroPrefixedOctalInteger  = IntegerSubType("0-prefixed octal")
	BinaryInteger             = IntegerSubType("binary")
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

type IntegerLiteralSymbol struct {
	Location
	Value string
	IntegerSubType
}

func (s IntegerLiteralSymbol) Id() SymbolId { return IntegerLiteralToken }

func (s IntegerLiteralSymbol) Loc() Location { return s.Location }
