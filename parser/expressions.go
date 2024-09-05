package parser

import (
	"fmt"
)

//
// (Bool/Int/Float/Rune/String)LiteralExpr
//

type BoolLiteralExpr struct {
	isExpression
	TokenValue
}

func (expr BoolLiteralExpr) String() string {
	return expr.TreeString("", "")
}

func (expr BoolLiteralExpr) TreeString(indent string, label string) string {
	return fmt.Sprintf("%s%s[BoolLiteralExpr: %s]", indent, label, expr.Value)
}

type IntLiteralExpr struct {
	isExpression
	TokenValue
}

func (expr IntLiteralExpr) String() string {
	return expr.TreeString("", "")
}

func (expr IntLiteralExpr) TreeString(indent string, label string) string {
	return fmt.Sprintf("%s%s[IntLiteralExpr: %s]", indent, label, expr.Value)
}

type FloatLiteralExpr struct {
	isExpression
	TokenValue
}

func (expr FloatLiteralExpr) String() string {
	return expr.TreeString("", "")
}

func (expr FloatLiteralExpr) TreeString(indent string, label string) string {
	return fmt.Sprintf("%s%s[FloatLiteralExpr: %s]", indent, label, expr.Value)
}

type RuneLiteralExpr struct {
	isExpression
	TokenValue
}

func (expr RuneLiteralExpr) String() string {
	return expr.TreeString("", "")
}

func (expr RuneLiteralExpr) TreeString(indent string, label string) string {
	return fmt.Sprintf("%s%s[RuneLiteralExpr: %s]", indent, label, expr.Value)
}

type StringLiteralExpr struct {
	isExpression
	TokenValue
}

func (expr StringLiteralExpr) String() string {
	return expr.TreeString("", "")
}

func (expr StringLiteralExpr) TreeString(indent string, label string) string {
	return fmt.Sprintf("%s%s[StringLiteralExpr: %s]", indent, label, expr.Value)
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

func (expr IdentifierExpr) String() string {
	return expr.TreeString("", "")
}

func (expr IdentifierExpr) TreeString(indent string, label string) string {
	return fmt.Sprintf("%s%s[IdentifierExpr: %s]", indent, label, expr.Value)
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

func (expr BinaryExpr) String() string {
	return expr.TreeString("", "")
}

func (expr BinaryExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[BinaryExpr: Op=%s\n", indent, label, SymbolId(expr.Op))
	result += expr.Left.TreeString(indent+"  ", "Left=") + "\n"
	result += expr.Right.TreeString(indent+"  ", "Right=") + "\n"
	result += fmt.Sprintf("%s]", indent)
	return result
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
