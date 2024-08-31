package parser

type ParseErrorSymbol struct {
	Error error
	Location
}

func (pe ParseErrorSymbol) Id() SymbolId { return ParseErrorToken }

func (pe ParseErrorSymbol) Loc() Location { return pe.Location }

type CountSymbol struct {
	SymbolId
	Location
	Count int
}

func (cs CountSymbol) Id() SymbolId { return cs.SymbolId }

func (cs CountSymbol) Loc() Location { return cs.Location }

type ValueSymbol struct {
	SymbolId
	Location
	Value string
}

func (vs ValueSymbol) Id() SymbolId { return vs.SymbolId }

func (vs ValueSymbol) Loc() Location { return vs.Location }
