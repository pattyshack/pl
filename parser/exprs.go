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

func (expr BoolLiteralExpr) TreeString(indent string, label string) string {
	return fmt.Sprintf("%s%s[BoolLiteralExpr: %s]", indent, label, expr.Value)
}

type IntLiteralExpr struct {
	isExpression
	TokenValue
}

func (expr IntLiteralExpr) TreeString(indent string, label string) string {
	return fmt.Sprintf("%s%s[IntLiteralExpr: %s]", indent, label, expr.Value)
}

type FloatLiteralExpr struct {
	isExpression
	TokenValue
}

func (expr FloatLiteralExpr) TreeString(indent string, label string) string {
	return fmt.Sprintf("%s%s[FloatLiteralExpr: %s]", indent, label, expr.Value)
}

type RuneLiteralExpr struct {
	isExpression
	TokenValue
}

func (expr RuneLiteralExpr) TreeString(indent string, label string) string {
	return fmt.Sprintf("%s%s[RuneLiteralExpr: %s]", indent, label, expr.Value)
}

type StringLiteralExpr struct {
	isExpression
	TokenValue
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
// NamedExpr
//

type NamedExpr struct {
	isExpression
	TokenValue
}

func (expr NamedExpr) TreeString(indent string, label string) string {
	return fmt.Sprintf("%s%s[NamedExpr: %s]", indent, label, expr.Value)
}

type NamedExprReducerImpl struct {
	NamedExprs []*NamedExpr
}

var _ NamedExprReducer = &NamedExprReducerImpl{}

func (reducer *NamedExprReducerImpl) IdentifierToNamedExpr(
	value TokenValue,
) (Expression, error) {
	expr := &NamedExpr{TokenValue: value}
	reducer.NamedExprs = append(reducer.NamedExprs, expr)
	return expr, nil
}

func (reducer *NamedExprReducerImpl) UnderscoreToNamedExpr(
	value TokenValue,
) (Expression, error) {
	expr := &NamedExpr{TokenValue: value}
	reducer.NamedExprs = append(reducer.NamedExprs, expr)
	return expr, nil
}

//
// AccessExpr
//

type AccessExpr struct {
	isExpression
	StartEndPos
	LeadingTrailingComments

	Operand Expression
	Field   TokenValue
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
		StartEndPos: newStartEndPos(operand.Loc(), field.End()),
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

// NOTE: The op's value is the same as the op's token symbol id.
type UnaryOp SymbolId

type UnaryExpr struct {
	isExpression

	StartEndPos
	LeadingTrailingComments

	IsPrefix bool

	Op      UnaryOp
	Operand Expression
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

type UnaryExprReducerImpl struct {
	UnaryExprs []*UnaryExpr
}

var _ PostfixUnaryExprReducer = &UnaryExprReducerImpl{}
var _ PrefixUnaryExprReducer = &UnaryExprReducerImpl{}

func (reducer *UnaryExprReducerImpl) ToPostfixUnaryExpr(
	operand Expression,
	op TokenValue,
) (
	Expression,
	error,
) {
	expr := &UnaryExpr{
		StartEndPos: newStartEndPos(operand.Loc(), op.End()),
		IsPrefix:    false,
		Op:          UnaryOp(op.SymbolId),
		Operand:     operand,
	}

	expr.LeadingComment = operand.TakeLeading()
	operand.AppendToTrailing(op.TakeLeading())
	expr.TrailingComment = op.TakeTrailing()

	reducer.UnaryExprs = append(reducer.UnaryExprs, expr)
	return expr, nil
}

func (reducer *UnaryExprReducerImpl) ToPrefixUnaryExpr(
	op TokenValue,
	operand Expression,
) (
	Expression,
	error,
) {
	expr := &UnaryExpr{
		StartEndPos: newStartEndPos(op.Loc(), operand.End()),
		IsPrefix:    true,
		Op:          UnaryOp(op.SymbolId),
		Operand:     operand,
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
	StartEndPos
	LeadingTrailingComments

	Left  Expression
	Op    BinaryOp
	Right Expression
}

func (expr BinaryExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[BinaryExpr: Op=%s\n", indent, label, SymbolId(expr.Op))
	result += expr.Left.TreeString(indent+"  ", "Left=") + "\n"
	result += expr.Right.TreeString(indent+"  ", "Right=")
	result += "\n" + indent + "]"
	return result
}

type BinaryExprReducerImpl struct {
	BinaryExprs []*BinaryExpr
}

var _ BinaryMulExprReducer = &BinaryExprReducerImpl{}
var _ BinaryAndExprReducer = &BinaryExprReducerImpl{}
var _ BinaryCmpExprReducer = &BinaryExprReducerImpl{}
var _ BinaryAndExprReducer = &BinaryExprReducerImpl{}
var _ BinaryOrExprReducer = &BinaryExprReducerImpl{}

func (reducer *BinaryExprReducerImpl) toBinaryExpr(
	left Expression,
	op TokenValue,
	right Expression,
) (
	Expression,
	error,
) {
	expr := &BinaryExpr{
		StartEndPos: newStartEndPos(left.Loc(), right.End()),
		Left:        left,
		Op:          BinaryOp(op.SymbolId),
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
	op TokenValue,
	right Expression,
) (
	Expression,
	error,
) {
	return reducer.toBinaryExpr(left, op, right)
}

func (reducer *BinaryExprReducerImpl) ToBinaryAddExpr(
	left Expression,
	op TokenValue,
	right Expression,
) (
	Expression,
	error,
) {
	return reducer.toBinaryExpr(left, op, right)
}

func (reducer *BinaryExprReducerImpl) ToBinaryCmpExpr(
	left Expression,
	op TokenValue,
	right Expression,
) (
	Expression,
	error,
) {
	return reducer.toBinaryExpr(left, op, right)
}

func (reducer *BinaryExprReducerImpl) ToBinaryAndExpr(
	left Expression,
	op TokenValue,
	right Expression,
) (
	Expression,
	error,
) {
	return reducer.toBinaryExpr(left, op, right)
}

func (reducer *BinaryExprReducerImpl) ToBinaryOrExpr(
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
// ImplicitStructExpr
//

type ImplicitStructExpr struct {
	isExpression
	ArgumentList

	// An improper struct is the a comma separated list of expressions without
	// left/right paren.  e.g., return 1, 2, 3
	IsImproper bool
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
		StartEndPos: newStartEndPos(colon.Loc(), colon.End()),
		Kind:        IsImplicitUnitArgument,
	}

	rightArg := &Argument{
		StartEndPos: newStartEndPos(colon.Loc(), colon.End()),
		Kind:        IsImplicitUnitArgument,
	}
	rightArg.LeadingComment = colon.TakeTrailing()

	args := newNodeList[*Argument]("ColonExpr")
	args.add(leftArg)
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
		StartEndPos: newStartEndPos(leftExpr.Loc(), leftExpr.End()),
		Kind:        PositionalArgument,
		Expr:        leftExpr,
	}
	leftArg.LeadingComment = leftExpr.TakeLeading()
	leftArg.TrailingComment = leftExpr.TakeTrailing()

	rightArg := &Argument{
		StartEndPos: newStartEndPos(colon.Loc(), colon.End()),
		Kind:        IsImplicitUnitArgument,
	}
	rightArg.LeadingComment = colon.TakeTrailing()

	args := newNodeList[*Argument]("ColonExpr")
	args.add(leftArg)
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
		StartEndPos: newStartEndPos(colon.Loc(), colon.End()),
		Kind:        IsImplicitUnitArgument,
	}

	rightArg := &Argument{
		StartEndPos: newStartEndPos(rightExpr.Loc(), rightExpr.End()),
		Kind:        PositionalArgument,
		Expr:        rightExpr,
	}
	rightArg.LeadingComment = rightExpr.TakeLeading()
	rightArg.TrailingComment = rightExpr.TakeTrailing()
	rightArg.PrependToLeading(colon.TakeTrailing())

	args := newNodeList[*Argument]("ColonExpr")
	args.add(leftArg)
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
		StartEndPos: newStartEndPos(leftExpr.Loc(), leftExpr.End()),
		Kind:        PositionalArgument,
		Expr:        leftExpr,
	}
	leftArg.LeadingComment = leftExpr.TakeLeading()
	leftArg.TrailingComment = leftExpr.TakeTrailing()

	rightArg := &Argument{
		StartEndPos: newStartEndPos(rightExpr.Loc(), rightExpr.End()),
		Kind:        PositionalArgument,
		Expr:        rightExpr,
	}
	rightArg.LeadingComment = rightExpr.TakeLeading()
	rightArg.TrailingComment = rightExpr.TakeTrailing()
	rightArg.PrependToLeading(colon.TakeTrailing())

	args := newNodeList[*Argument]("ColonExpr")
	args.add(leftArg)
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
		StartEndPos: newStartEndPos(colon.Loc(), colon.End()),
		Kind:        IsImplicitUnitArgument,
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
		StartEndPos: newStartEndPos(expr.Loc(), expr.End()),
		Kind:        PositionalArgument,
		Expr:        expr,
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
	StartEndPos
	LeadingTrailingComments

	FuncExpr         Expression
	GenericArguments GenericArgumentList
	Arguments        ArgumentList
}

func (expr CallExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf("%s%s[CallExpr\n", indent, label)
	result += expr.FuncExpr.TreeString(indent+"  ", "FuncExpr=") + "\n"
	if len(expr.GenericArguments.Elements) > 0 {
		result += expr.GenericArguments.TreeString(indent+"  ", "GenericArguments=")
		result += "\n"
	}
	result += expr.Arguments.TreeString(indent+"  ", "Arguments=")
	result += "\n" + indent + "]"

	return result
}

var _ Expression = &CallExpr{}

type CallExprReducerImpl struct {
	CallExprs []*CallExpr
}

var _ CallExprReducer = &CallExprReducerImpl{}

func (reducer *CallExprReducerImpl) ToCallExpr(
	funcExpr Expression,
	genericArguments *GenericArgumentList,
	lparen TokenValue,
	arguments *ArgumentList,
	rparen TokenValue,
) (
	Expression,
	error,
) {
	if genericArguments == nil {
		genericArguments = &GenericArgumentList{}
	}

	arguments.reduceMarkers(lparen, rparen)

	expr := &CallExpr{
		StartEndPos: newStartEndPos(funcExpr.Loc(), rparen.End()),
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

type IndexExpr struct {
	isExpression
	StartEndPos
	LeadingTrailingComments

	Accessible Expression
	Index      Argument
}

func (expr IndexExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf("%s%s[IndexExpr\n", indent, label)
	result += expr.Accessible.TreeString(indent+"  ", "Accessible=")
	result += "\n" + expr.Index.TreeString(indent+"  ", "Argument=")
	result += "\n" + indent + "]"
	return result
}

type IndexExprReducerImpl struct {
	IndexExprs []*IndexExpr
}

var _ IndexExprReducer = &IndexExprReducerImpl{}

func (reducer *IndexExprReducerImpl) ToIndexExpr(
	accessible Expression,
	lbracket TokenValue,
	index *Argument,
	rbracket TokenValue,
) (
	Expression,
	error,
) {
	accessible.AppendToTrailing(lbracket.TakeLeading())
	index.PrependToLeading(lbracket.TakeTrailing())
	index.AppendToTrailing(rbracket.TakeLeading())

	expr := &IndexExpr{
		StartEndPos: newStartEndPos(accessible.Loc(), rbracket.End()),
		Accessible:  accessible,
		Index:       *index,
	}

	expr.LeadingComment = accessible.TakeLeading()
	expr.TrailingComment = rbracket.TakeTrailing()

	reducer.IndexExprs = append(reducer.IndexExprs, expr)
	return expr, nil
}

//
// InitializeExpr
//

type InitializeExpr struct {
	isExpression
	StartEndPos
	LeadingTrailingComments

	Initializable TypeExpression
	Arguments     ArgumentList
}

func (expr InitializeExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf("%s%s[InitializeExpr:\n", indent, label)
	result += expr.Initializable.TreeString(indent+"  ", "Initialiable=") + "\n"
	result += expr.Arguments.TreeString(indent+"  ", "Arguments=")
	result += "\n" + indent + "]"
	return result
}

type InitializeExprReducerImpl struct {
	InitializeExprs []*InitializeExpr
}

var _ InitializeExprReducer = &InitializeExprReducerImpl{}

func (reducer *InitializeExprReducerImpl) ToInitializeExpr(
	initializable TypeExpression,
	lparen TokenValue,
	arguments *ArgumentList,
	rparen TokenValue,
) (
	Expression,
	error,
) {
	arguments.reduceMarkers(lparen, rparen)

	expr := &InitializeExpr{
		StartEndPos:   newStartEndPos(initializable.Loc(), rparen.End()),
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

type ConditionBranch struct {
	Condition Expression
	Branch    Expression
}

func (cb ConditionBranch) TreeString(indent string, label string) string {
	result := fmt.Sprintf("%s%s[ConditionBranch:\n", indent, label)
	result += cb.Condition.TreeString(indent+"  ", "Condition=") + "\n"
	result += cb.Branch.TreeString(indent+"  ", "Branch=")
	result += "\n" + indent + "]"
	return result
}

type IfExpr struct {
	isExpression
	StartEndPos
	LeadingTrailingComments

	LabelDecl         string // optional
	ConditionBranches []ConditionBranch
	ElseBranch        Expression // optional
}

var _ Expression = &IfExpr{}

func (expr IfExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[IfExpr: LabelDecl=%s\n",
		indent,
		label,
		expr.LabelDecl)
	for idx, condBranch := range expr.ConditionBranches {
		branchLabel := fmt.Sprintf("Branch%d=", idx)
		result += condBranch.TreeString(indent+"  ", branchLabel) + "\n"
	}

	if expr.ElseBranch != nil {
		result += expr.ElseBranch.TreeString(indent+"  ", "ElseBranch=") + "\n"
	}

	result += indent + "]"
	return result
}

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
	labelDecl TokenValue,
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
	elseKW TokenValue,
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
	elseKW TokenValue,
	ifKW TokenValue,
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
	ifKW TokenValue,
	condition Expression,
	branch Expression,
) (
	*IfExpr,
	error,
) {
	condition.PrependToLeading(ifKW.TakeTrailing())
	expr := &IfExpr{
		StartEndPos:       newStartEndPos(ifKW.Loc(), branch.End()),
		ConditionBranches: []ConditionBranch{{condition, branch}},
	}
	expr.LeadingComment = ifKW.TakeLeading()

	reducer.IfExprs = append(reducer.IfExprs, expr)
	return expr, nil
}
