package reducer

import (
	. "github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser/lr"
)

func (Reducer *Reducer) ToParseErrorExpr(
	pe lr.ParseErrorSymbol,
) (
	Expression,
	error,
) {
	Reducer.ParseErrors = append(Reducer.ParseErrors, pe.ParseErrorNode)
	return pe.ParseErrorNode, nil
}

func (Reducer *Reducer) ToParseErrorTypeExpr(
	pe lr.ParseErrorSymbol,
) (
	TypeExpression,
	error,
) {
	Reducer.ParseErrors = append(Reducer.ParseErrors, pe.ParseErrorNode)
	return pe.ParseErrorNode, nil
}
