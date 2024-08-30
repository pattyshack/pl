package parser

type ParseErrorSymbol struct {
	Error error
	Location
}

func (pe ParseErrorSymbol) Id() SymbolId { return ParseErrorToken }

func (pe ParseErrorSymbol) Loc() Location { return pe.Location }

type ValueSymbol struct {
	SymbolId
	Location
	Value string
}

func (vs ValueSymbol) Id() SymbolId { return vs.SymbolId }

func (vs ValueSymbol) Loc() Location { return vs.Location }
