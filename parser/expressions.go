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

//
// BinaryExpr
//

// NOTE: The op's value is the same as the op's token symbol id.
type BinaryOp SymbolId

type BinaryExpr struct {
	isExpression
	LeadingTrailingComments

	Left  Expression
	Op    BinaryOp
	Right Expression
}

func (expr *BinaryExpr) Loc() Location {
	return expr.Left.Loc()
}

func (expr *BinaryExpr) End() Location {
	return expr.Right.End()
}

type BinaryExprReducer struct {
	BinaryExprs []*BinaryExpr
}

var _ BinaryMulExprReducer = &BinaryExprReducer{}
var _ BinaryAndExprReducer = &BinaryExprReducer{}
var _ BinaryCmpExprReducer = &BinaryExprReducer{}
var _ BinaryAndExprReducer = &BinaryExprReducer{}
var _ BinaryOrExprReducer = &BinaryExprReducer{}

func (reducer *BinaryExprReducer) toBinaryExpr(
	left Expression,
	op TokenValue,
	right Expression,
) (
	Expression,
	error,
) {
	expr := &BinaryExpr{
		Left:  left,
		Op:    BinaryOp(op.SymbolId),
		Right: right,
	}

	expr.LeadingComment = left.TakeLeading()
	left.AppendToTrailing(op.TakeLeading())
	right.PrependToLeading(op.TakeTrailing())
	expr.TrailingComment = right.TakeTrailing()

	reducer.BinaryExprs = append(reducer.BinaryExprs, expr)
	return expr, nil
}

func (reducer *BinaryExprReducer) ToBinaryMulExpr(
	left Expression,
	op TokenValue,
	right Expression,
) (
	Expression,
	error,
) {
	return reducer.toBinaryExpr(left, op, right)
}

func (reducer *BinaryExprReducer) ToBinaryAddExpr(
	left Expression,
	op TokenValue,
	right Expression,
) (
	Expression,
	error,
) {
	return reducer.toBinaryExpr(left, op, right)
}

func (reducer *BinaryExprReducer) ToBinaryCmpExpr(
	left Expression,
	op TokenValue,
	right Expression,
) (
	Expression,
	error,
) {
	return reducer.toBinaryExpr(left, op, right)
}

func (reducer *BinaryExprReducer) ToBinaryAndExpr(
	left Expression,
	op TokenValue,
	right Expression,
) (
	Expression,
	error,
) {
	return reducer.toBinaryExpr(left, op, right)
}

func (reducer *BinaryExprReducer) ToBinaryOrExpr(
	left Expression,
	op TokenValue,
	right Expression,
) (
	Expression,
	error,
) {
	return reducer.toBinaryExpr(left, op, right)
}
