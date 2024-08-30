package parser

type ParseError struct {
	Error error
	Location
}

func (pe ParseError) Id() SymbolId { return ParseErrorToken }

func (pe ParseError) Loc() Location { return pe.Location }
