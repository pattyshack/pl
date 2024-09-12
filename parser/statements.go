package parser

import (
	"fmt"
)

//
// StatementsExpr
//

type StatementsExpr struct {
	isExpression

	LabelDecl string // optional
	NodeList[Statement]
}

func NewStatementList() *StatementsExpr {
	return &StatementsExpr{}
}

func (expr StatementsExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[StatementExpr: LabelDecl=%s",
		indent,
		label,
		expr.LabelDecl)
	if len(expr.Elements) == 0 {
		return result + "]"
	}

	for idx, element := range expr.Elements {
		result += "\n" + element.TreeString(
			indent+"  ",
			fmt.Sprintf("Statement%d=", idx))
	}

	result += "\n" + indent + "]"
	return result
}

type StatementsExprReducerImpl struct{}

var _ ProperStatementListReducer = &StatementsExprReducerImpl{}
var _ StatementListReducer = &StatementsExprReducerImpl{}
var _ StatementsReducer = &StatementsExprReducerImpl{}
var _ StatementsExprReducer = &StatementsExprReducerImpl{}

func (StatementsExprReducerImpl) AddImplicitToProperStatementList(
	statements *StatementsExpr,
	newlines TokenCount,
	statement Statement,
) (
	*StatementsExpr,
	error,
) {
	statements.reduceAdd(TokenValue{}, statement)
	return statements, nil
}

func (StatementsExprReducerImpl) AddExplicitToProperStatementList(
	statements *StatementsExpr,
	semicolon TokenValue,
	statement Statement,
) (
	*StatementsExpr,
	error,
) {
	statements.reduceAdd(semicolon, statement)
	return statements, nil
}

func (StatementsExprReducerImpl) StatementToProperStatementList(
	statement Statement,
) (
	*StatementsExpr,
	error,
) {
	list := NewStatementList()
	list.add(statement)
	return list, nil
}

func (StatementsExprReducerImpl) ImproperImplicitToStatementList(
	statements *StatementsExpr,
	newlines TokenCount,
) (
	*StatementsExpr,
	error,
) {
	return statements, nil
}

func (StatementsExprReducerImpl) ImproperExplicitToStatementList(
	statements *StatementsExpr,
	semicolon TokenValue,
) (
	*StatementsExpr,
	error,
) {
	statements.reduceImproper(semicolon)
	return statements, nil
}

func (StatementsExprReducerImpl) NilToStatementList() (*StatementsExpr, error) {
	return NewStatementList(), nil
}

func (StatementsExprReducerImpl) ToStatements(
	lbrace TokenValue,
	list *StatementsExpr,
	rbrace TokenValue,
) (
	Expression,
	error,
) {
	list.reduceMarkers(lbrace, rbrace)
	return list, nil
}

func (StatementsExprReducerImpl) LabelledToStatementsExpr(
	labelDecl TokenValue,
	statementsExprOrParseError Expression,
) (
	Expression,
	error,
) {
	switch expr := statementsExprOrParseError.(type) {
	case *ParseErrorSymbol:
		return expr, nil
	case *StatementsExpr:
		expr.StartPos = labelDecl.Loc()
		expr.LabelDecl = labelDecl.Value
		expr.PrependToLeading(labelDecl.TakeTrailing())
		expr.PrependToLeading(labelDecl.TakeLeading())
		return expr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", statementsExprOrParseError))
}
