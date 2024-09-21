package reducer

import (
	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser/lr"
)

func (Reducer *Reducer) ToParseErrorExpr(
	pe lr.ParseErrorSymbol,
) (
	ast.Expression,
	error,
) {
	Reducer.ParseErrors = append(Reducer.ParseErrors, pe.Errors...)
	return pe.ParseErrorNode, nil
}
