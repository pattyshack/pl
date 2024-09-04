package parser

//
// (Bool/Int/Float/Rune/String)LiteralExpr
//

type BoolLiteralExpr struct {
	isExpression
	TokenValue
}

type IntLiteralExpr struct {
	isExpression
	TokenValue
}

type FloatLiteralExpr struct {
	isExpression
	TokenValue
}

type RuneLiteralExpr struct {
	isExpression
	TokenValue
}

type StringLiteralExpr struct {
	isExpression
	TokenValue
}

type LiteralExprReducerImpl struct {
	BoolLiteralExprs   []*BoolLiteralExpr
	IntLiteralExprs    []*IntLiteralExpr
	FloatLiteralExprs  []*FloatLiteralExpr
	RuneLiteralExprs   []*RuneLiteralExpr
	StringLiteralExprs []*StringLiteralExpr
}

var _ LiteralExprReducer = &LiteralExprReducerImpl{}

func (reducer *LiteralExprReducerImpl) TrueToLiteralExpr(
	value TokenValue,
) (Expression, error) {
	expr := &BoolLiteralExpr{TokenValue: value}
	reducer.BoolLiteralExprs = append(reducer.BoolLiteralExprs, expr)
	return expr, nil
}

func (reducer *LiteralExprReducerImpl) FalseToLiteralExpr(
	value TokenValue,
) (Expression, error) {
	expr := &BoolLiteralExpr{TokenValue: value}
	reducer.BoolLiteralExprs = append(reducer.BoolLiteralExprs, expr)
	return expr, nil
}

func (reducer *LiteralExprReducerImpl) IntegerLiteralToLiteralExpr(
	value TokenValue,
) (Expression, error) {
	expr := &IntLiteralExpr{TokenValue: value}
	reducer.IntLiteralExprs = append(reducer.IntLiteralExprs, expr)
	return expr, nil
}

func (reducer *LiteralExprReducerImpl) FloatLiteralToLiteralExpr(
	value TokenValue,
) (Expression, error) {
	expr := &FloatLiteralExpr{TokenValue: value}
	reducer.FloatLiteralExprs = append(reducer.FloatLiteralExprs, expr)
	return expr, nil
}

func (reducer *LiteralExprReducerImpl) RuneLiteralToLiteralExpr(
	value TokenValue,
) (Expression, error) {
	expr := &RuneLiteralExpr{TokenValue: value}
	reducer.RuneLiteralExprs = append(reducer.RuneLiteralExprs, expr)
	return expr, nil
}

func (reducer *LiteralExprReducerImpl) StringLiteralToLiteralExpr(
	value TokenValue,
) (Expression, error) {
	expr := &StringLiteralExpr{TokenValue: value}
	reducer.StringLiteralExprs = append(reducer.StringLiteralExprs, expr)
	return expr, nil
}

//
// IdentifierExpr
//

type IdentifierExpr struct {
	isExpression
	TokenValue
}

type IdentifierExprReducerImpl struct {
	IdentifierExprs []*IdentifierExpr
}

var _ IdentifierExprReducer = &IdentifierExprReducerImpl{}

func (reducer *IdentifierExprReducerImpl) ToIdentifierExpr(
	value TokenValue,
) (Expression, error) {
	expr := &IdentifierExpr{TokenValue: value}
	reducer.IdentifierExprs = append(reducer.IdentifierExprs, expr)
	return expr, nil
}
