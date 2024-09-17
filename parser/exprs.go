package parser

import (
	"fmt"

	. "github.com/pattyshack/pl/ast"
)

//
// (Bool/Int/Float/Rune/String)LiteralExpr
//

type LiteralExprReducerImpl struct {
	BoolLiteralExprs   []*BoolLiteralExpr
	IntLiteralExprs    []*IntLiteralExpr
	FloatLiteralExprs  []*FloatLiteralExpr
	RuneLiteralExprs   []*RuneLiteralExpr
	StringLiteralExprs []*StringLiteralExpr
}

var _ LiteralExprReducer = &LiteralExprReducerImpl{}

func (reducer *LiteralExprReducerImpl) TrueToLiteralExpr(
	value *TokenValue,
) (Expression, error) {
	expr := &BoolLiteralExpr{ValueNode: value}
	reducer.BoolLiteralExprs = append(reducer.BoolLiteralExprs, expr)
	return expr, nil
}

func (reducer *LiteralExprReducerImpl) FalseToLiteralExpr(
	value *TokenValue,
) (Expression, error) {
	expr := &BoolLiteralExpr{ValueNode: value}
	reducer.BoolLiteralExprs = append(reducer.BoolLiteralExprs, expr)
	return expr, nil
}

func (reducer *LiteralExprReducerImpl) IntegerLiteralToLiteralExpr(
	value *TokenValue,
) (Expression, error) {
	expr := &IntLiteralExpr{ValueNode: value}
	reducer.IntLiteralExprs = append(reducer.IntLiteralExprs, expr)
	return expr, nil
}

func (reducer *LiteralExprReducerImpl) FloatLiteralToLiteralExpr(
	value *TokenValue,
) (Expression, error) {
	expr := &FloatLiteralExpr{ValueNode: value}
	reducer.FloatLiteralExprs = append(reducer.FloatLiteralExprs, expr)
	return expr, nil
}

func (reducer *LiteralExprReducerImpl) RuneLiteralToLiteralExpr(
	value *TokenValue,
) (Expression, error) {
	expr := &RuneLiteralExpr{ValueNode: value}
	reducer.RuneLiteralExprs = append(reducer.RuneLiteralExprs, expr)
	return expr, nil
}

func (reducer *LiteralExprReducerImpl) StringLiteralToLiteralExpr(
	value *TokenValue,
) (Expression, error) {
	expr := &StringLiteralExpr{ValueNode: value}
	reducer.StringLiteralExprs = append(reducer.StringLiteralExprs, expr)
	return expr, nil
}

//
// NamedExpr
//

type NamedExprReducerImpl struct {
	NamedExprs []*NamedExpr
}

var _ NamedExprReducer = &NamedExprReducerImpl{}

func (reducer *NamedExprReducerImpl) IdentifierToNamedExpr(
	value *TokenValue,
) (Expression, error) {
	expr := &NamedExpr{ValueNode: value}
	reducer.NamedExprs = append(reducer.NamedExprs, expr)
	return expr, nil
}

func (reducer *NamedExprReducerImpl) UnderscoreToNamedExpr(
	value *TokenValue,
) (Expression, error) {
	expr := &NamedExpr{ValueNode: value}
	reducer.NamedExprs = append(reducer.NamedExprs, expr)
	return expr, nil
}

//
// AccessExpr
//

type AccessExprReducerImpl struct {
	AccessExprs []*AccessExpr
}

var _ AccessExprReducer = &AccessExprReducerImpl{}

func (reducer *AccessExprReducerImpl) ToAccessExpr(
	operand Expression,
	dot *TokenValue,
	field *TokenValue,
) (
	Expression,
	error,
) {
	expr := &AccessExpr{
		StartEndPos: NewStartEndPos(operand.Loc(), field.End()),
		Operand:     operand,
		Field:       field,
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

type UnaryExprReducerImpl struct {
	UnaryExprs []*UnaryExpr
}

var _ PostfixUnaryExprReducer = &UnaryExprReducerImpl{}
var _ PrefixUnaryExprReducer = &UnaryExprReducerImpl{}
var _ UnaryOpAssignStatementReducer = &UnaryExprReducerImpl{}
var _ RecvExprReducer = &UnaryExprReducerImpl{}

func (reducer *UnaryExprReducerImpl) toPostfixUnaryExpr(
	operand Expression,
	op *TokenValue,
) *UnaryExpr {
	expr := &UnaryExpr{
		StartEndPos: NewStartEndPos(operand.Loc(), op.End()),
		IsPrefix:    false,
		Op:          UnaryOp(op.Value),
		Operand:     operand,
	}

	expr.LeadingComment = operand.TakeLeading()
	operand.AppendToTrailing(op.TakeLeading())
	expr.TrailingComment = op.TakeTrailing()

	reducer.UnaryExprs = append(reducer.UnaryExprs, expr)
	return expr
}

func (reducer *UnaryExprReducerImpl) ToPostfixUnaryExpr(
	operand Expression,
	op *TokenValue,
) (
	Expression,
	error,
) {
	return reducer.toPostfixUnaryExpr(operand, op), nil
}

func (reducer *UnaryExprReducerImpl) toPrefixUnaryExpr(
	op *TokenValue,
	operand Expression,
) Expression {
	expr := &UnaryExpr{
		StartEndPos: NewStartEndPos(op.Loc(), operand.End()),
		IsPrefix:    true,
		Op:          UnaryOp(op.Value),
		Operand:     operand,
	}

	expr.LeadingComment = op.TakeLeading()
	operand.PrependToLeading(op.TakeTrailing())
	expr.TrailingComment = operand.TakeTrailing()

	reducer.UnaryExprs = append(reducer.UnaryExprs, expr)
	return expr
}

func (reducer *UnaryExprReducerImpl) ToPrefixUnaryExpr(
	op *TokenValue,
	operand Expression,
) (
	Expression,
	error,
) {
	return reducer.toPrefixUnaryExpr(op, operand), nil
}

func (reducer *UnaryExprReducerImpl) ToRecvExpr(
	arrow *TokenValue,
	expr Expression,
) (
	Expression,
	error,
) {
	return reducer.toPrefixUnaryExpr(arrow, expr), nil
}

func (reducer *UnaryExprReducerImpl) ToUnaryOpAssignStatement(
	operand Expression,
	op *TokenValue,
) (
	Statement,
	error,
) {
	return reducer.toPostfixUnaryExpr(operand, op), nil
}

//
// BinaryExpr
//

type BinaryExprReducerImpl struct {
	BinaryExprs []*BinaryExpr
}

var _ BinaryMulExprReducer = &BinaryExprReducerImpl{}
var _ BinaryAndExprReducer = &BinaryExprReducerImpl{}
var _ BinaryCmpExprReducer = &BinaryExprReducerImpl{}
var _ BinaryAndExprReducer = &BinaryExprReducerImpl{}
var _ BinaryOrExprReducer = &BinaryExprReducerImpl{}
var _ BinaryOpAssignStatementReducer = &BinaryExprReducerImpl{}
var _ SendExprReducer = &BinaryExprReducerImpl{}
var _ ExprAssignStatementReducer = &BinaryExprReducerImpl{}
var _ SequenceExprAssignStatementReducer = &BinaryExprReducerImpl{}
var _ GlobalVarDefReducer = &BinaryExprReducerImpl{}

func (reducer *BinaryExprReducerImpl) toBinaryExpr(
	left Expression,
	op *TokenValue,
	right Expression,
) (
	Expression,
	error,
) {
	expr := &BinaryExpr{
		StartEndPos: NewStartEndPos(left.Loc(), right.End()),
		Left:        left,
		Op:          BinaryOp(op.Value),
		Right:       right,
	}

	expr.LeadingComment = left.TakeLeading()
	left.AppendToTrailing(op.TakeLeading())
	right.PrependToLeading(op.TakeTrailing())
	expr.TrailingComment = right.TakeTrailing()

	reducer.BinaryExprs = append(reducer.BinaryExprs, expr)
	return expr, nil
}

func (reducer *BinaryExprReducerImpl) ToBinaryMulExpr(
	left Expression,
	op *TokenValue,
	right Expression,
) (
	Expression,
	error,
) {
	return reducer.toBinaryExpr(left, op, right)
}

func (reducer *BinaryExprReducerImpl) ToBinaryAddExpr(
	left Expression,
	op *TokenValue,
	right Expression,
) (
	Expression,
	error,
) {
	return reducer.toBinaryExpr(left, op, right)
}

func (reducer *BinaryExprReducerImpl) ToBinaryCmpExpr(
	left Expression,
	op *TokenValue,
	right Expression,
) (
	Expression,
	error,
) {
	return reducer.toBinaryExpr(left, op, right)
}

func (reducer *BinaryExprReducerImpl) ToBinaryAndExpr(
	left Expression,
	op *TokenValue,
	right Expression,
) (
	Expression,
	error,
) {
	return reducer.toBinaryExpr(left, op, right)
}

func (reducer *BinaryExprReducerImpl) ToBinaryOrExpr(
	left Expression,
	op *TokenValue,
	right Expression,
) (
	Expression,
	error,
) {
	return reducer.toBinaryExpr(left, op, right)
}

func (reducer *BinaryExprReducerImpl) ToSendExpr(
	receiver Expression,
	arrow *TokenValue,
	expr Expression,
) (
	Expression,
	error,
) {
	return reducer.toBinaryExpr(receiver, arrow, expr)
}

func (reducer *BinaryExprReducerImpl) ToBinaryOpAssignStatement(
	address Expression,
	op *TokenValue,
	value Expression,
) (
	Statement,
	error,
) {
	return reducer.toBinaryExpr(address, op, value)
}

func (reducer *BinaryExprReducerImpl) ToExprAssignStatement(
	pattern Expression,
	assign *TokenValue,
	value Expression,
) (
	Statement,
	error,
) {
	return reducer.toBinaryExpr(pattern, assign, value)
}

func (reducer *BinaryExprReducerImpl) ToSequenceExprAssignStatement(
	pattern Expression,
	assign *TokenValue,
	value Expression,
) (
	Statement,
	error,
) {
	return reducer.toBinaryExpr(pattern, assign, value)
}

func (reducer *BinaryExprReducerImpl) DefToGlobalVarDef(
	pattern Expression,
	assign *TokenValue,
	value Expression,
) (
	Definition,
	error,
) {
	return reducer.toBinaryExpr(pattern, assign, value)
}

//
// ImplicitStructExpr
//

type ImplicitStructExprReducerImpl struct {
	ImplicitStructExprs []*ImplicitStructExpr
}

var _ ImplicitStructExprReducer = &ImplicitStructExprReducerImpl{}
var _ ImproperExprStructReducer = &ImplicitStructExprReducerImpl{}
var _ ImproperSequenceExprStructReducer = &ImplicitStructExprReducerImpl{}
var _ TuplePatternReducer = &ImplicitStructExprReducerImpl{}

func (reducer *ImplicitStructExprReducerImpl) toImplicitStructExpr(
	lparen *TokenValue,
	args *ArgumentList,
	rparen *TokenValue,
) *ImplicitStructExpr {
	args.ReduceMarkers(lparen, rparen)
	return &ImplicitStructExpr{
		ArgumentList: *args,
	}
}

func (reducer *ImplicitStructExprReducerImpl) ToImplicitStructExpr(
	lparen *TokenValue,
	args *ArgumentList,
	rparen *TokenValue,
) (
	Expression,
	error,
) {
	expr := reducer.toImplicitStructExpr(lparen, args, rparen)
	reducer.ImplicitStructExprs = append(reducer.ImplicitStructExprs, expr)
	return expr, nil
}

func (reducer *ImplicitStructExprReducerImpl) PairToImproperExprStruct(
	expr1 Expression,
	comma *TokenValue,
	expr2 Expression,
) (
	*ImplicitStructExpr,
	error,
) {
	arg1 := NewPositionalArgument(expr1)
	arg2 := NewPositionalArgument(expr2)

	list := NewArgumentList()
	list.Add(arg1)
	list.ReduceAdd(comma, arg2)

	expr := &ImplicitStructExpr{
		ArgumentList: *list,
	}

	reducer.ImplicitStructExprs = append(reducer.ImplicitStructExprs, expr)
	return expr, nil
}

func (reducer *ImplicitStructExprReducerImpl) AddToImproperExprStruct(
	structExpr *ImplicitStructExpr,
	comma *TokenValue,
	expr Expression,
) (
	*ImplicitStructExpr,
	error,
) {
	arg := NewPositionalArgument(expr)
	structExpr.ReduceAdd(comma, arg)
	return structExpr, nil
}

func (reducer *ImplicitStructExprReducerImpl) PairToImproperSequenceExprStruct(
	expr1 Expression,
	comma *TokenValue,
	expr2 Expression,
) (
	*ImplicitStructExpr,
	error,
) {
	arg1 := NewPositionalArgument(expr1)
	arg2 := NewPositionalArgument(expr2)

	list := NewArgumentList()
	list.Add(arg1)
	list.ReduceAdd(comma, arg2)

	expr := &ImplicitStructExpr{
		ArgumentList: *list,
	}

	reducer.ImplicitStructExprs = append(reducer.ImplicitStructExprs, expr)
	return expr, nil
}

func (reducer *ImplicitStructExprReducerImpl) AddToImproperSequenceExprStruct(
	structExpr *ImplicitStructExpr,
	comma *TokenValue,
	expr Expression,
) (
	*ImplicitStructExpr,
	error,
) {
	arg := NewPositionalArgument(expr)
	structExpr.ReduceAdd(comma, arg)
	return structExpr, nil
}

func (reducer *ImplicitStructExprReducerImpl) ToTuplePattern(
	lparen *TokenValue,
	list *ArgumentList,
	rparen *TokenValue,
) (
	Expression,
	error,
) {
	return reducer.toImplicitStructExpr(lparen, list, rparen), nil
}

//
// ColonExpr
//

type ColonExprReducerImpl struct {
}

var _ ColonExprReducer = &ColonExprReducerImpl{}

func (ColonExprReducerImpl) UnitUnitPairToColonExpr(
	colon *TokenValue,
) (
	*ColonExpr,
	error,
) {
	leftArg := &Argument{
		StartEndPos: NewStartEndPos(colon.Loc(), colon.End()),
		Kind:        IsImplicitUnitArgument,
	}

	rightArg := &Argument{
		StartEndPos: NewStartEndPos(colon.Loc(), colon.End()),
		Kind:        IsImplicitUnitArgument,
	}
	rightArg.LeadingComment = colon.TakeTrailing()

	args := NewNodeList[*Argument]("ColonExpr")
	args.Add(leftArg)
	args.ReduceAdd(colon, rightArg)

	return &ColonExpr{
		ArgumentList: *args,
	}, nil
}

func (ColonExprReducerImpl) ExprUnitPairToColonExpr(
	leftExpr Expression,
	colon *TokenValue,
) (
	*ColonExpr,
	error,
) {
	leftArg := &Argument{
		StartEndPos: NewStartEndPos(leftExpr.Loc(), leftExpr.End()),
		Kind:        PositionalArgument,
		Expr:        leftExpr,
	}
	leftArg.LeadingComment = leftExpr.TakeLeading()
	leftArg.TrailingComment = leftExpr.TakeTrailing()

	rightArg := &Argument{
		StartEndPos: NewStartEndPos(colon.Loc(), colon.End()),
		Kind:        IsImplicitUnitArgument,
	}
	rightArg.LeadingComment = colon.TakeTrailing()

	args := NewNodeList[*Argument]("ColonExpr")
	args.Add(leftArg)
	args.ReduceAdd(colon, rightArg)

	return &ColonExpr{
		ArgumentList: *args,
	}, nil
}

func (reducer *ColonExprReducerImpl) UnitExprPairToColonExpr(
	colon *TokenValue,
	rightExpr Expression,
) (
	*ColonExpr,
	error,
) {
	leftArg := &Argument{
		StartEndPos: NewStartEndPos(colon.Loc(), colon.End()),
		Kind:        IsImplicitUnitArgument,
	}

	rightArg := &Argument{
		StartEndPos: NewStartEndPos(rightExpr.Loc(), rightExpr.End()),
		Kind:        PositionalArgument,
		Expr:        rightExpr,
	}
	rightArg.LeadingComment = rightExpr.TakeLeading()
	rightArg.TrailingComment = rightExpr.TakeTrailing()
	rightArg.PrependToLeading(colon.TakeTrailing())

	args := NewNodeList[*Argument]("ColonExpr")
	args.Add(leftArg)
	args.ReduceAdd(colon, rightArg)

	return &ColonExpr{
		ArgumentList: *args,
	}, nil
}

func (reducer *ColonExprReducerImpl) ExprExprPairToColonExpr(
	leftExpr Expression,
	colon *TokenValue,
	rightExpr Expression,
) (
	*ColonExpr,
	error,
) {
	leftArg := &Argument{
		StartEndPos: NewStartEndPos(leftExpr.Loc(), leftExpr.End()),
		Kind:        PositionalArgument,
		Expr:        leftExpr,
	}
	leftArg.LeadingComment = leftExpr.TakeLeading()
	leftArg.TrailingComment = leftExpr.TakeTrailing()

	rightArg := &Argument{
		StartEndPos: NewStartEndPos(rightExpr.Loc(), rightExpr.End()),
		Kind:        PositionalArgument,
		Expr:        rightExpr,
	}
	rightArg.LeadingComment = rightExpr.TakeLeading()
	rightArg.TrailingComment = rightExpr.TakeTrailing()
	rightArg.PrependToLeading(colon.TakeTrailing())

	args := NewNodeList[*Argument]("ColonExpr")
	args.Add(leftArg)
	args.ReduceAdd(colon, rightArg)

	return &ColonExpr{
		ArgumentList: *args,
	}, nil
}

func (reducer *ColonExprReducerImpl) ColonExprUnitTupleToColonExpr(
	list *ColonExpr,
	colon *TokenValue,
) (
	*ColonExpr,
	error,
) {
	arg := &Argument{
		StartEndPos: NewStartEndPos(colon.Loc(), colon.End()),
		Kind:        IsImplicitUnitArgument,
	}
	arg.LeadingComment = colon.TakeTrailing()

	list.ReduceAdd(colon, arg)
	return list, nil
}

func (reducer *ColonExprReducerImpl) ColonExprExprTupleToColonExpr(
	list *ColonExpr,
	colon *TokenValue,
	expr Expression,
) (
	*ColonExpr,
	error,
) {
	arg := &Argument{
		StartEndPos: NewStartEndPos(expr.Loc(), expr.End()),
		Kind:        PositionalArgument,
		Expr:        expr,
	}
	arg.LeadingComment = expr.TakeLeading()
	arg.TrailingComment = expr.TakeTrailing()
	arg.PrependToLeading(colon.TakeTrailing())

	list.ReduceAdd(colon, arg)
	return list, nil
}

//
// CallExpr
//

type CallExprReducerImpl struct {
	CallExprs []*CallExpr
}

var _ CallExprReducer = &CallExprReducerImpl{}

func (reducer *CallExprReducerImpl) ToCallExpr(
	funcExpr Expression,
	genericArguments *GenericArgumentList,
	lparen *TokenValue,
	arguments *ArgumentList,
	rparen *TokenValue,
) (
	Expression,
	error,
) {
	if genericArguments == nil {
		genericArguments = &GenericArgumentList{}
	}

	arguments.ReduceMarkers(lparen, rparen)

	expr := &CallExpr{
		StartEndPos: NewStartEndPos(funcExpr.Loc(), rparen.End()),
		LeadingTrailingComments: LeadingTrailingComments{
			LeadingComment:  funcExpr.TakeLeading(),
			TrailingComment: arguments.TakeTrailing(),
		},
		FuncExpr:         funcExpr,
		GenericArguments: *genericArguments,
		Arguments:        *arguments,
	}

	reducer.CallExprs = append(reducer.CallExprs, expr)
	return expr, nil
}

//
// IndexExpr
//

type IndexExprReducerImpl struct {
	IndexExprs []*IndexExpr
}

var _ IndexExprReducer = &IndexExprReducerImpl{}

func (reducer *IndexExprReducerImpl) ToIndexExpr(
	accessible Expression,
	lbracket *TokenValue,
	index *Argument,
	rbracket *TokenValue,
) (
	Expression,
	error,
) {
	accessible.AppendToTrailing(lbracket.TakeLeading())
	index.PrependToLeading(lbracket.TakeTrailing())
	index.AppendToTrailing(rbracket.TakeLeading())

	expr := &IndexExpr{
		StartEndPos: NewStartEndPos(accessible.Loc(), rbracket.End()),
		Accessible:  accessible,
		Index:       *index,
	}

	expr.LeadingComment = accessible.TakeLeading()
	expr.TrailingComment = rbracket.TakeTrailing()

	reducer.IndexExprs = append(reducer.IndexExprs, expr)
	return expr, nil
}

//
// AsExpr
//

type AsExprReducerImpl struct {
	AsExprs []*AsExpr
}

var _ AsExprReducer = &AsExprReducerImpl{}

func (reducer *AsExprReducerImpl) ToAsExpr(
	accessible Expression,
	dot *TokenValue,
	as *TokenValue,
	lparen *TokenValue,
	castType TypeExpression,
	rparen *TokenValue,
) (
	Expression,
	error,
) {
	leading := accessible.TakeLeading()
	accessible.AppendToTrailing(dot.TakeLeading())
	comments := dot.TakeTrailing()
	comments.Append(as.TakeLeading())
	comments.Append(as.TakeTrailing())
	comments.Append(lparen.TakeLeading())
	comments.Append(lparen.TakeTrailing())
	castType.PrependToLeading(comments)
	castType.AppendToTrailing(rparen.TakeLeading())
	trailing := rparen.TakeTrailing()

	expr := &AsExpr{
		StartEndPos: NewStartEndPos(accessible.Loc(), rparen.End()),
		Accessible:  accessible,
		CastType:    castType,
	}
	expr.LeadingComment = leading
	expr.TrailingComment = trailing

	return expr, nil
}

//
// InitializeExpr
//

type InitializeExprReducerImpl struct {
	InitializeExprs []*InitializeExpr
}

var _ InitializeExprReducer = &InitializeExprReducerImpl{}

func (reducer *InitializeExprReducerImpl) ToInitializeExpr(
	initializable TypeExpression,
	lparen *TokenValue,
	arguments *ArgumentList,
	rparen *TokenValue,
) (
	Expression,
	error,
) {
	arguments.ReduceMarkers(lparen, rparen)

	expr := &InitializeExpr{
		StartEndPos:   NewStartEndPos(initializable.Loc(), rparen.End()),
		Initializable: initializable,
		Arguments:     *arguments,
	}

	expr.LeadingComment = expr.Initializable.TakeLeading()
	expr.TrailingComment = expr.Arguments.TakeTrailing()

	reducer.InitializeExprs = append(reducer.InitializeExprs, expr)
	return expr, nil
}

//
// IfExpr
//

type IfExprReducerImpl struct {
	IfExprs []*IfExpr
}

var _ IfExprReducer = &IfExprReducerImpl{}
var _ IfElseExprReducer = &IfExprReducerImpl{}
var _ IfElifExprReducer = &IfExprReducerImpl{}
var _ IfOnlyExprReducer = &IfExprReducerImpl{}

func (reducer *IfExprReducerImpl) UnlabelledToIfExpr(
	ifExpr *IfExpr,
) (
	Expression,
	error,
) {
	if ifExpr.ElseBranch != nil {
		ifExpr.TrailingComment = ifExpr.ElseBranch.TakeTrailing()
	} else {
		last := ifExpr.ConditionBranches[len(ifExpr.ConditionBranches)-1].Branch
		ifExpr.TrailingComment = last.TakeTrailing()
	}

	return ifExpr, nil
}

func (reducer *IfExprReducerImpl) LabelledToIfExpr(
	labelDecl *TokenValue,
	ifExpr *IfExpr,
) (
	Expression,
	error,
) {
	reducer.UnlabelledToIfExpr(ifExpr)

	ifExpr.PrependToLeading(labelDecl.TakeTrailing())
	ifExpr.PrependToLeading(labelDecl.TakeLeading())

	ifExpr.LabelDecl = labelDecl.Value
	return ifExpr, nil
}

func (reducer *IfExprReducerImpl) ElseToIfElseExpr(
	ifExpr *IfExpr,
	elseKW *TokenValue,
	branch Expression,
) (
	*IfExpr,
	error,
) {
	last := ifExpr.ConditionBranches[len(ifExpr.ConditionBranches)-1].Branch
	last.AppendToTrailing(elseKW.TakeLeading())

	branch.PrependToLeading(elseKW.TakeTrailing())

	ifExpr.EndPos = branch.End()
	ifExpr.ElseBranch = branch

	return ifExpr, nil
}

func (reducer *IfExprReducerImpl) ElifToIfElifExpr(
	ifExpr *IfExpr,
	elseKW *TokenValue,
	ifKW *TokenValue,
	condition Expression,
	branch Expression,
) (
	*IfExpr,
	error,
) {
	last := ifExpr.ConditionBranches[len(ifExpr.ConditionBranches)-1].Branch
	last.AppendToTrailing(elseKW.TakeLeading())
	last.AppendToTrailing(elseKW.TakeTrailing())
	last.AppendToTrailing(ifKW.TakeLeading())

	condition.PrependToLeading(ifKW.TakeTrailing())

	ifExpr.EndPos = branch.End()

	ifExpr.ConditionBranches = append(
		ifExpr.ConditionBranches,
		ConditionBranch{condition, branch})

	return ifExpr, nil
}

func (reducer *IfExprReducerImpl) ToIfOnlyExpr(
	ifKW *TokenValue,
	condition Expression,
	branch Expression,
) (
	*IfExpr,
	error,
) {
	condition.PrependToLeading(ifKW.TakeTrailing())
	expr := &IfExpr{
		StartEndPos:       NewStartEndPos(ifKW.Loc(), branch.End()),
		ConditionBranches: []ConditionBranch{{condition, branch}},
	}
	expr.LeadingComment = ifKW.TakeLeading()

	reducer.IfExprs = append(reducer.IfExprs, expr)
	return expr, nil
}

//
// SwitchExpr
//

type SwitchExprReducerImpl struct {
	SwitchExprs []*SwitchExpr
}

var _ SwitchExprReducer = &SwitchExprReducerImpl{}
var _ SwitchExprBodyReducer = &SwitchExprReducerImpl{}

func (reducer *SwitchExprReducerImpl) LabelledToSwitchExpr(
	labelDecl *TokenValue,
	expr Expression,
) (
	Expression,
	error,
) {
	switch switchExpr := expr.(type) {
	case *SwitchExpr:
		switchExpr.StartPos = labelDecl.Loc()
		switchExpr.PrependToLeading(labelDecl.TakeTrailing())
		switchExpr.PrependToLeading(labelDecl.TakeLeading())
		switchExpr.LabelDecl = labelDecl.Value
		return switchExpr, nil
	case *ParseErrorSymbol:
		return expr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", expr))
}

func (reducer *SwitchExprReducerImpl) ToSwitchExprBody(
	switchKW *TokenValue,
	operand Expression,
	expr Expression,
) (
	Expression,
	error,
) {
	switch branches := expr.(type) {
	case *StatementsExpr:
		trailing := branches.TakeTrailing()

		switchExpr := &SwitchExpr{
			StartEndPos: NewStartEndPos(switchKW.Loc(), branches.End()),
			Operand:     operand,
			Branches:    *branches,
		}

		switchExpr.LeadingComment = switchKW.TakeLeading()
		operand.PrependToLeading(switchKW.TakeTrailing())
		switchExpr.TrailingComment = trailing

		reducer.SwitchExprs = append(reducer.SwitchExprs, switchExpr)
		return switchExpr, nil
	case *ParseErrorSymbol:
		return expr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", expr))
}

//
// SelectExpr
//

type SelectExprReducerImpl struct {
	SelectExprs []*SelectExpr
}

var _ SelectExprReducer = &SelectExprReducerImpl{}
var _ SelectExprBodyReducer = &SelectExprReducerImpl{}

func (reducer *SelectExprReducerImpl) LabelledToSelectExpr(
	labelDecl *TokenValue,
	expr Expression,
) (
	Expression,
	error,
) {
	switch selectExpr := expr.(type) {
	case *SelectExpr:
		selectExpr.StartPos = labelDecl.Loc()
		selectExpr.PrependToLeading(labelDecl.TakeTrailing())
		selectExpr.PrependToLeading(labelDecl.TakeLeading())
		selectExpr.LabelDecl = labelDecl.Value
		return selectExpr, nil
	case *ParseErrorSymbol:
		return expr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", expr))
}

func (reducer *SelectExprReducerImpl) ToSelectExprBody(
	switchKW *TokenValue,
	expr Expression,
) (
	Expression,
	error,
) {
	switch branches := expr.(type) {
	case *StatementsExpr:
		branches.PrependToLeading(switchKW.TakeTrailing())
		trailing := branches.TakeTrailing()

		selectExpr := &SelectExpr{
			StartEndPos: NewStartEndPos(switchKW.Loc(), branches.End()),
			Branches:    *branches,
		}

		selectExpr.LeadingComment = switchKW.TakeLeading()
		selectExpr.TrailingComment = trailing

		reducer.SelectExprs = append(reducer.SelectExprs, selectExpr)
		return selectExpr, nil
	case *ParseErrorSymbol:
		return expr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", expr))
}

//
// LoopExpr
//

type LoopExprReducerImpl struct {
	LoopExprs []*LoopExpr
}

var _ LoopExprReducer = &LoopExprReducerImpl{}
var _ LoopExprBodyReducer = &LoopExprReducerImpl{}
var _ OptionalSequenceStatementReducer = &LoopExprReducerImpl{}
var _ OptionalSequenceExprReducer = &LoopExprReducerImpl{}
var _ LoopBodyReducer = &LoopExprReducerImpl{}

func (reducer *LoopExprReducerImpl) LabelledToLoopExpr(
	labelDecl *TokenValue,
	expr Expression,
) (
	Expression,
	error,
) {
	switch loop := expr.(type) {
	case *LoopExpr:
		loop.StartPos = labelDecl.Loc()
		loop.PrependToLeading(labelDecl.TakeTrailing())
		loop.PrependToLeading(labelDecl.TakeLeading())
		loop.LabelDecl = labelDecl.Value
		return loop, nil
	case *ParseErrorSymbol:
		return expr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", expr))
}

func (reducer *LoopExprReducerImpl) InfiniteToLoopExprBody(
	bodyExpr Expression,
) (
	Expression,
	error,
) {
	switch body := bodyExpr.(type) {
	case *StatementsExpr:
		leading := body.TakeLeading()
		trailing := body.TakeTrailing()

		loop := &LoopExpr{
			StartEndPos: NewStartEndPos(bodyExpr.Loc(), bodyExpr.End()),
			LoopKind:    InfiniteLoop,
			Body:        *body,
		}
		loop.LeadingComment = leading
		loop.TrailingComment = trailing

		return loop, nil
	case *ParseErrorSymbol:
		return bodyExpr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", bodyExpr))
}

func (reducer *LoopExprReducerImpl) DoWhileToLoopExprBody(
	bodyExpr Expression,
	forKW *TokenValue,
	condition Expression,
) (
	Expression,
	error,
) {
	switch body := bodyExpr.(type) {
	case *StatementsExpr:
		leading := body.TakeLeading()
		body.AppendToTrailing(forKW.TakeLeading())
		condition.PrependToLeading(forKW.TakeTrailing())
		trailing := condition.TakeTrailing()

		loop := &LoopExpr{
			StartEndPos: NewStartEndPos(bodyExpr.Loc(), bodyExpr.End()),
			LoopKind:    DoWhileLoop,
			Condition:   condition,
			Body:        *body,
		}
		loop.LeadingComment = leading
		loop.TrailingComment = trailing

		return loop, nil
	case *ParseErrorSymbol:
		return bodyExpr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", bodyExpr))
}

func (reducer *LoopExprReducerImpl) WhileToLoopExprBody(
	forKW *TokenValue,
	condition Expression,
	bodyExpr Expression,
) (
	Expression,
	error,
) {
	switch body := bodyExpr.(type) {
	case *StatementsExpr:
		leading := forKW.TakeLeading()
		condition.PrependToLeading(forKW.TakeTrailing())
		trailing := bodyExpr.TakeTrailing()

		loop := &LoopExpr{
			StartEndPos: NewStartEndPos(bodyExpr.Loc(), bodyExpr.End()),
			LoopKind:    WhileLoop,
			Condition:   condition,
			Body:        *body,
		}
		loop.LeadingComment = leading
		loop.TrailingComment = trailing

		return loop, nil
	case *ParseErrorSymbol:
		return bodyExpr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", bodyExpr))
}

func (reducer *LoopExprReducerImpl) IteratorToLoopExprBody(
	forKW *TokenValue,
	assignPattern Expression,
	in *TokenValue,
	iterator Expression,
	bodyExpr Expression,
) (
	Expression,
	error,
) {
	switch body := bodyExpr.(type) {
	case *StatementsExpr:
		leading := forKW.TakeLeading()
		assignPattern.PrependToLeading(forKW.TakeTrailing())
		assignPattern.AppendToTrailing(in.TakeLeading())
		iterator.PrependToLeading(in.TakeTrailing())
		trailing := bodyExpr.TakeTrailing()

		loop := &LoopExpr{
			StartEndPos:   NewStartEndPos(bodyExpr.Loc(), bodyExpr.End()),
			LoopKind:      IteratorLoop,
			AssignPattern: assignPattern,
			Condition:     iterator,
			Body:          *body,
		}
		loop.LeadingComment = leading
		loop.TrailingComment = trailing

		return loop, nil
	case *ParseErrorSymbol:
		return bodyExpr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", bodyExpr))
}

func (reducer *LoopExprReducerImpl) ForToLoopExprBody(
	forKW *TokenValue,
	init Statement,
	semicolon1 *TokenValue,
	condition Expression,
	semicolon2 *TokenValue,
	post Statement,
	bodyExpr Expression,
) (
	Expression,
	error,
) {
	switch body := bodyExpr.(type) {
	case *StatementsExpr:
		var last Node = body
		if post != nil {
			last = post
		}

		last.PrependToLeading(semicolon2.TakeTrailing())
		if condition == nil {
			last.PrependToLeading(semicolon2.TakeLeading())
		} else {
			condition.AppendToTrailing(semicolon2.TakeLeading())
			last = condition
		}

		last.PrependToLeading(semicolon1.TakeTrailing())
		if init == nil {
			last.PrependToLeading(semicolon1.TakeLeading())
		} else {
			init.AppendToTrailing(semicolon1.TakeLeading())
			last = init
		}

		leading := forKW.TakeLeading()
		last.PrependToLeading(forKW.TakeTrailing())
		trailing := body.TakeTrailing()

		loop := &LoopExpr{
			StartEndPos: NewStartEndPos(bodyExpr.Loc(), bodyExpr.End()),
			LoopKind:    ForLoop,
			Init:        init,
			Condition:   condition,
			Post:        post,
			Body:        *body,
		}
		loop.LeadingComment = leading
		loop.TrailingComment = trailing

		return loop, nil
	case *ParseErrorSymbol:
		return bodyExpr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", bodyExpr))
}

func (reducer *LoopExprReducerImpl) NilToOptionalSequenceStatement() (
	Statement,
	error,
) {
	return nil, nil
}

func (reducer *LoopExprReducerImpl) NilToOptionalSequenceExpr() (
	Expression,
	error,
) {
	return nil, nil
}

func (reducer *LoopExprReducerImpl) ToLoopBody(
	do *TokenValue,
	expr Expression,
) (
	Expression,
	error,
) {
	switch body := expr.(type) {
	case *StatementsExpr:
		body.StartPos = do.Loc()
		body.PrependToLeading(do.TakeTrailing())
		body.PrependToLeading(do.TakeLeading())
		return body, nil
	case *ParseErrorSymbol:
		return expr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", expr))
}
