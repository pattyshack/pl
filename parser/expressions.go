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
// AccessExpr
//

type AccessExpr struct {
	isExpression

	LeadingTrailingComments

	Operand Expression
	Field   TokenValue
}

func (expr AccessExpr) Loc() Location {
	return expr.Operand.Loc()
}

func (expr AccessExpr) End() Location {
	return expr.Field.End()
}

func (expr AccessExpr) String() string {
	return expr.TreeString("", "")
}

func (expr AccessExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[AccessExpr: Field=%s\n", indent, label, expr.Field.Value)
	result += expr.Operand.TreeString(indent+"  ", "Operand=")
	result += "\n" + indent + "]"
	return result
}

type AccessExprReducerImpl struct {
	AccessExprs []*AccessExpr
}

var _ AccessExprReducer = &AccessExprReducerImpl{}

func (reducer *AccessExprReducerImpl) ToAccessExpr(
	operand Expression,
	dot TokenValue,
	field TokenValue,
) (
	Expression,
	error,
) {
	expr := &AccessExpr{
		Operand: operand,
		Field:   field,
	}

	expr.LeadingComment = operand.TakeLeading()
	operand.AppendToTrailing(dot.TakeLeading())
	field.PrependToLeading(dot.TakeTrailing())
	expr.TrailingComment = field.TakeTrailing()

	reducer.AccessExprs = append(reducer.AccessExprs, expr)
	return expr, nil
}

//
// UnaryExpr
//

// NOTE: The op's value is the same as the op's token symbol id.
type UnaryOp SymbolId

type UnaryExpr struct {
	isExpression

	StartPos Location
	EndPos   Location
	LeadingTrailingComments

	IsPrefix bool

	Op      UnaryOp
	Operand Expression
}

func (expr UnaryExpr) Loc() Location {
	return expr.StartPos
}

func (expr UnaryExpr) End() Location {
	return expr.EndPos
}

func (expr UnaryExpr) String() string {
	return expr.TreeString("", "")
}

func (expr UnaryExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[UnaryExpr: IsPrefix=%v Op=%s\n",
		indent,
		label,
		expr.IsPrefix,
		SymbolId(expr.Op))
	result += expr.Operand.TreeString(indent+"  ", "Operand=")
	result += "\n" + indent + "]"
	return result
}

type UnaryExprReducer struct {
	UnaryExprs []*UnaryExpr
}

var _ PostfixUnaryExprReducer = &UnaryExprReducer{}
var _ PrefixUnaryExprReducer = &UnaryExprReducer{}

func (reducer *UnaryExprReducer) ToPostfixUnaryExpr(
	operand Expression,
	op TokenValue,
) (
	Expression,
	error,
) {
	expr := &UnaryExpr{
		StartPos: operand.Loc(),
		EndPos:   op.End(),
		IsPrefix: false,
		Op:       UnaryOp(op.SymbolId),
		Operand:  operand,
	}

	expr.LeadingComment = operand.TakeLeading()
	operand.AppendToTrailing(op.TakeLeading())
	expr.TrailingComment = op.TakeTrailing()

	reducer.UnaryExprs = append(reducer.UnaryExprs, expr)
	return expr, nil
}

func (reducer *UnaryExprReducer) ToPrefixUnaryExpr(
	op TokenValue,
	operand Expression,
) (
	Expression,
	error,
) {
	expr := &UnaryExpr{
		StartPos: op.Loc(),
		EndPos:   operand.End(),
		IsPrefix: true,
		Op:       UnaryOp(op.SymbolId),
		Operand:  operand,
	}

	expr.LeadingComment = op.TakeLeading()
	operand.PrependToLeading(op.TakeTrailing())
	expr.TrailingComment = operand.TakeTrailing()

	reducer.UnaryExprs = append(reducer.UnaryExprs, expr)
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

func (expr BinaryExpr) Loc() Location {
	return expr.Left.Loc()
}

func (expr BinaryExpr) End() Location {
	return expr.Right.End()
}

func (expr BinaryExpr) String() string {
	return expr.TreeString("", "")
}

func (expr BinaryExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[BinaryExpr: Op=%s\n", indent, label, SymbolId(expr.Op))
	result += expr.Left.TreeString(indent+"  ", "Left=") + "\n"
	result += expr.Right.TreeString(indent+"  ", "Right=")
	result += "\n" + indent + "]"
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

//
// Argument
//

type ArgumentKind int

const (
	PositionalArgument = ArgumentKind(iota)
	NamedAssignmentArgument
	VarargAssignmentArgument
	ColonExprArgument
	// Only used by patterns
	SkipPatternArgument
	// Only used by ColonExpr
	IsImplicitUnitArgument
)

type Argument struct {
	StartPos Location
	EndPos   Location
	LeadingTrailingComments

	Kind ArgumentKind

	// Only set for named assignment
	OptionalName string

	// NOTE: Expr may be nil for SkipPatternArgument or IsImplicitUnit.
	Expr Expression

	HasEllipsis bool
}

var _ Node = &Argument{}

func (arg *Argument) Loc() Location {
	return arg.StartPos
}

func (arg *Argument) End() Location {
	return arg.StartPos
}

func (arg *Argument) String() string {
	return arg.TreeString("", "")
}

func (arg *Argument) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[Argument: Kind=%v OptionalName=%s HasEllipsis=%v",
		indent,
		label,
		arg.Kind,
		arg.OptionalName,
		arg.HasEllipsis)
	if arg.Expr == nil {
		result += "\n"
		result += arg.Expr.TreeString(indent+"  ", "")
		result += "\n" + indent + "]"
	} else {
		result += "]"
	}

	return result
}

type ArgumentReducerImpl struct {
}

var _ ArgumentReducer = ArgumentReducerImpl{}

func (ArgumentReducerImpl) PositionalToArgument(
	expr Expression,
) (
	*Argument,
	error,
) {
	return &Argument{
		StartPos: expr.Loc(),
		EndPos:   expr.End(),
		LeadingTrailingComments: LeadingTrailingComments{
			LeadingComment:  expr.TakeLeading(),
			TrailingComment: expr.TakeTrailing(),
		},
		Kind:        PositionalArgument,
		Expr:        expr,
		HasEllipsis: false,
	}, nil
}

func (ArgumentReducerImpl) ColonExprToArgument(
	expr *ColonExpr,
) (
	*Argument,
	error,
) {
	return &Argument{
		StartPos: expr.Loc(),
		EndPos:   expr.End(),
		LeadingTrailingComments: LeadingTrailingComments{
			LeadingComment:  expr.TakeLeading(),
			TrailingComment: expr.TakeTrailing(),
		},
		Kind:        ColonExprArgument,
		Expr:        expr,
		HasEllipsis: false,
	}, nil
}

func (ArgumentReducerImpl) NamedAssignmentToArgument(
	name TokenValue,
	assign TokenValue,
	expr Expression,
) (
	*Argument,
	error,
) {
	arg := &Argument{
		StartPos:    name.Loc(),
		EndPos:      expr.End(),
		Kind:        NamedAssignmentArgument,
		Expr:        expr,
		HasEllipsis: false,
	}

	arg.LeadingComment = name.TakeLeading()
	// prepend in reverse order
	expr.PrependToLeading(assign.TakeTrailing())
	expr.PrependToLeading(assign.TakeLeading())
	expr.PrependToLeading(name.TakeTrailing())
	arg.TrailingComment = expr.TakeTrailing()

	return arg, nil
}

func (ArgumentReducerImpl) VarargAssignmentToArgument(
	expr Expression,
	ellipsis TokenValue,
) (
	*Argument,
	error,
) {
	arg := &Argument{
		StartPos:    expr.Loc(),
		EndPos:      ellipsis.End(),
		Kind:        VarargAssignmentArgument,
		Expr:        expr,
		HasEllipsis: true,
	}

	arg.LeadingComment = expr.TakeLeading()
	expr.AppendToTrailing(ellipsis.TakeLeading())
	arg.TrailingComment = ellipsis.TakeTrailing()
	return arg, nil
}

func (ArgumentReducerImpl) SkipPatternToArgument(
	ellipsis TokenValue,
) (
	*Argument,
	error,
) {
	return &Argument{
		StartPos:                ellipsis.Loc(),
		EndPos:                  ellipsis.End(),
		LeadingTrailingComments: ellipsis.LeadingTrailingComments,
		Kind:                    SkipPatternArgument,
		Expr:                    nil,
		HasEllipsis:             true,
	}, nil
}

//
// Argument list
//

type ArgumentList = NodeList[*Argument]

type ArgumentListReducer struct{}

var _ ProperArgumentsReducer = &ArgumentListReducer{}
var _ ArgumentsReducer = &ArgumentListReducer{}

func (reducer *ArgumentListReducer) AddToProperArguments(
	list *ArgumentList,
	comma TokenValue,
	arg *Argument,
) (
	*ArgumentList,
	error,
) {
	list.reduceAdd(comma, arg)
	return list, nil
}

func (reducer *ArgumentListReducer) ArgumentToProperArguments(
	arg *Argument,
) (
	*ArgumentList,
	error,
) {
	return newNodeList[*Argument]("ArgumentList", arg), nil
}

func (reducer *ArgumentListReducer) ImproperToArguments(
	list *ArgumentList,
	comma TokenValue,
) (
	*ArgumentList,
	error,
) {
	list.reduceImproper(comma)
	return list, nil
}

func (reducer *ArgumentListReducer) NilToArguments() (*ArgumentList, error) {
	return nil, nil
}

//
// ImplicitStructExpr
//

type ImplicitStructExpr struct {
	isExpression
	ArgumentList

	// An improper struct is the a comma separated list of expressions without
	// left/right paren.  e.g., return 1, 2, 3
	IsImproper bool
}

func (expr *ImplicitStructExpr) String() string {
	return expr.TreeString("", "")
}

func (expr *ImplicitStructExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[ImplicitStructExpr: IsImproper=%v\n",
		indent,
		label,
		expr.IsImproper)
	result += expr.ArgumentList.TreeString(indent+"  ", "")
	result += "\n" + indent + "]\n"
	return result
}

type ImplicitStructExprReducerImpl struct {
	ImplicitStructExprs []*ImplicitStructExpr
}

var _ ImplicitStructExprReducer = &ImplicitStructExprReducerImpl{}

func (reducer *ImplicitStructExprReducerImpl) ToImplicitStructExpr(
	lparen TokenValue,
	args *ArgumentList,
	rparen TokenValue,
) (
	Expression,
	error,
) {
	args.reduceMarkers(lparen, rparen)
	args.ListType = "ImplicitStructExpr"

	expr := &ImplicitStructExpr{
		ArgumentList: *args,
	}

	reducer.ImplicitStructExprs = append(reducer.ImplicitStructExprs, expr)
	return expr, nil
}

//
// ColonExpr
//

type ColonExpr struct {
	isExpression
	ArgumentList
}

type ColonExprReducerImpl struct {
}

var _ ColonExprReducer = &ColonExprReducerImpl{}

func (ColonExprReducerImpl) UnitUnitPairToColonExpr(
	colon TokenValue,
) (
	*ColonExpr,
	error,
) {
	leftArg := &Argument{
		StartPos: colon.Loc(),
		EndPos:   colon.End(),
		Kind:     IsImplicitUnitArgument,
	}

	rightArg := &Argument{
		StartPos: colon.Loc(),
		EndPos:   colon.End(),
		Kind:     IsImplicitUnitArgument,
	}
	rightArg.LeadingComment = colon.TakeTrailing()

	args := newNodeList[*Argument]("ColonExpr", leftArg)
	args.reduceAdd(colon, rightArg)

	return &ColonExpr{
		ArgumentList: *args,
	}, nil
}

func (ColonExprReducerImpl) ExprUnitPairToColonExpr(
	leftExpr Expression,
	colon TokenValue,
) (
	*ColonExpr,
	error,
) {
	leftArg := &Argument{
		StartPos: leftExpr.Loc(),
		EndPos:   leftExpr.End(),
		Kind:     PositionalArgument,
		Expr:     leftExpr,
	}
	leftArg.LeadingComment = leftExpr.TakeLeading()
	leftArg.TrailingComment = leftExpr.TakeTrailing()

	rightArg := &Argument{
		StartPos: colon.Loc(),
		EndPos:   colon.End(),
		Kind:     IsImplicitUnitArgument,
	}
	rightArg.LeadingComment = colon.TakeTrailing()

	args := newNodeList[*Argument]("ColonExpr", leftArg)
	args.reduceAdd(colon, rightArg)

	return &ColonExpr{
		ArgumentList: *args,
	}, nil
}

func (reducer *ColonExprReducerImpl) UnitExprPairToColonExpr(
	colon TokenValue,
	rightExpr Expression,
) (
	*ColonExpr,
	error,
) {
	leftArg := &Argument{
		StartPos: colon.Loc(),
		EndPos:   colon.End(),
		Kind:     IsImplicitUnitArgument,
	}

	rightArg := &Argument{
		StartPos: rightExpr.Loc(),
		EndPos:   rightExpr.End(),
		Kind:     PositionalArgument,
		Expr:     rightExpr,
	}
	rightArg.LeadingComment = rightExpr.TakeLeading()
	rightArg.TrailingComment = rightExpr.TakeTrailing()
	rightArg.PrependToLeading(colon.TakeTrailing())

	args := newNodeList[*Argument]("ColonExpr", leftArg)
	args.reduceAdd(colon, rightArg)

	return &ColonExpr{
		ArgumentList: *args,
	}, nil
}

func (reducer *ColonExprReducerImpl) ExprExprPairToColonExpr(
	leftExpr Expression,
	colon TokenValue,
	rightExpr Expression,
) (
	*ColonExpr,
	error,
) {
	leftArg := &Argument{
		StartPos: leftExpr.Loc(),
		EndPos:   leftExpr.End(),
		Kind:     PositionalArgument,
		Expr:     leftExpr,
	}
	leftArg.LeadingComment = leftExpr.TakeLeading()
	leftArg.TrailingComment = leftExpr.TakeTrailing()

	rightArg := &Argument{
		StartPos: rightExpr.Loc(),
		EndPos:   rightExpr.End(),
		Kind:     PositionalArgument,
		Expr:     rightExpr,
	}
	rightArg.LeadingComment = rightExpr.TakeLeading()
	rightArg.TrailingComment = rightExpr.TakeTrailing()
	rightArg.PrependToLeading(colon.TakeTrailing())

	args := newNodeList[*Argument]("ColonExpr", leftArg)
	args.reduceAdd(colon, rightArg)

	return &ColonExpr{
		ArgumentList: *args,
	}, nil
}

func (reducer *ColonExprReducerImpl) ColonExprUnitTupleToColonExpr(
	list *ColonExpr,
	colon TokenValue,
) (
	*ColonExpr,
	error,
) {
	arg := &Argument{
		StartPos: colon.Loc(),
		EndPos:   colon.End(),
		Kind:     IsImplicitUnitArgument,
	}
	arg.LeadingComment = colon.TakeTrailing()

	list.reduceAdd(colon, arg)
	return list, nil
}

func (reducer *ColonExprReducerImpl) ColonExprExprTupleToColonExpr(
	list *ColonExpr,
	colon TokenValue,
	expr Expression,
) (
	*ColonExpr,
	error,
) {
	arg := &Argument{
		StartPos: expr.Loc(),
		EndPos:   expr.End(),
		Kind:     PositionalArgument,
		Expr:     expr,
	}
	arg.LeadingComment = expr.TakeLeading()
	arg.TrailingComment = expr.TakeTrailing()
	arg.PrependToLeading(colon.TakeTrailing())

	list.reduceAdd(colon, arg)
	return list, nil
}

//
// CallExpr
//

type CallExpr struct {
	isExpression
	StartPos Location
	EndPos   Location
	LeadingTrailingComments

	FuncExpr      Expression
	TypeArguments TypeArgumentList
	Arguments     ArgumentList
}

func (expr CallExpr) Loc() Location {
	return expr.StartPos
}

func (expr CallExpr) End() Location {
	return expr.EndPos
}

func (expr CallExpr) String() string {
	return expr.TreeString("", "")
}

func (expr CallExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf("%s%s[CallExpr\n", indent, label)
	result += expr.FuncExpr.TreeString(indent+"  ", "FuncExpr=") + "\n"
	if len(expr.TypeArguments.Elements) > 0 {
		result += expr.TypeArguments.TreeString(indent+"  ", "TypeArguments=")
		result += "\n"
	}
	result += expr.Arguments.TreeString(indent+"  ", "Arguments=")
	result += indent + "]"

	return result
}

var _ Expression = &CallExpr{}

type CallExprReducerImpl struct {
	CallExprs []*CallExpr
}

var _ CallExprReducer = &CallExprReducerImpl{}

func (reducer *CallExprReducerImpl) ToCallExpr(
	funcExpr Expression,
	typeArguments *TypeArgumentList,
	lparen TokenValue,
	arguments *ArgumentList,
	rparen TokenValue,
) (
	Expression,
	error,
) {
	arguments.reduceMarkers(lparen, rparen)

	expr := &CallExpr{
		StartPos: funcExpr.Loc(),
		EndPos:   rparen.End(),
		LeadingTrailingComments: LeadingTrailingComments{
			LeadingComment:  funcExpr.TakeLeading(),
			TrailingComment: arguments.TakeTrailing(),
		},
		FuncExpr:      funcExpr,
		TypeArguments: *typeArguments,
		Arguments:     *arguments,
	}

	reducer.CallExprs = append(reducer.CallExprs, expr)
	return expr, nil
}
