package reducer

import (
	. "github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser"
)

func (Reducer *Reducer) ToParseErrorExpr(
	pe parser.ParseErrorSymbol,
) (
	Expression,
	error,
) {
	Reducer.ParseErrors = append(Reducer.ParseErrors, pe.ParseErrorNode)
	return pe.ParseErrorNode, nil
}

func (Reducer *Reducer) ToParseErrorTypeExpr(
	pe parser.ParseErrorSymbol,
) (
	TypeExpression,
	error,
) {
	Reducer.ParseErrors = append(Reducer.ParseErrors, pe.ParseErrorNode)
	return pe.ParseErrorNode, nil
}
