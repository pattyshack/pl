package reducer

import (
	"fmt"

	. "github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser/lr"
)

//
// (Bool/Int/Float/Rune/String)LiteralExpr
//

func (reducer *Reducer) TrueToLiteralExpr(
	value *lr.TokenValue,
) (Expression, error) {
	expr := &BoolLiteralExpr{ValuedNode: value}
	reducer.BoolLiteralExprs = append(reducer.BoolLiteralExprs, expr)
	return expr, nil
}

func (reducer *Reducer) FalseToLiteralExpr(
	value *lr.TokenValue,
) (Expression, error) {
	expr := &BoolLiteralExpr{ValuedNode: value}
	reducer.BoolLiteralExprs = append(reducer.BoolLiteralExprs, expr)
	return expr, nil
}

func (reducer *Reducer) IntegerLiteralToLiteralExpr(
	value *lr.TokenValue,
) (Expression, error) {
	expr := &IntLiteralExpr{ValuedNode: value}
	reducer.IntLiteralExprs = append(reducer.IntLiteralExprs, expr)
	return expr, nil
}

func (reducer *Reducer) FloatLiteralToLiteralExpr(
	value *lr.TokenValue,
) (Expression, error) {
	expr := &FloatLiteralExpr{ValuedNode: value}
	reducer.FloatLiteralExprs = append(reducer.FloatLiteralExprs, expr)
	return expr, nil
}

func (reducer *Reducer) RuneLiteralToLiteralExpr(
	value *lr.TokenValue,
) (Expression, error) {
	expr := &RuneLiteralExpr{ValuedNode: value}
	reducer.RuneLiteralExprs = append(reducer.RuneLiteralExprs, expr)
	return expr, nil
}

func (reducer *Reducer) StringLiteralToLiteralExpr(
	value *lr.TokenValue,
) (Expression, error) {
	expr := &StringLiteralExpr{ValuedNode: value}
	reducer.StringLiteralExprs = append(reducer.StringLiteralExprs, expr)
	return expr, nil
}

//
// NamedExpr
//

func (reducer *Reducer) IdentifierToNamedExpr(
	value *lr.TokenValue,
) (Expression, error) {
	expr := &NamedExpr{ValuedNode: value}
	reducer.NamedExprs = append(reducer.NamedExprs, expr)
	return expr, nil
}

func (reducer *Reducer) UnderscoreToNamedExpr(
	value *lr.TokenValue,
) (Expression, error) {
	expr := &NamedExpr{ValuedNode: value}
	reducer.NamedExprs = append(reducer.NamedExprs, expr)
	return expr, nil
}

//
// AccessExpr
//

func (reducer *Reducer) ToAccessExpr(
	operand Expression,
	dot *lr.TokenValue,
	field *lr.TokenValue,
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

func (reducer *Reducer) toPostfixUnaryExpr(
	operand Expression,
	op *lr.TokenValue,
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

func (reducer *Reducer) ToPostfixUnaryExpr(
	operand Expression,
	op *lr.TokenValue,
) (
	Expression,
	error,
) {
	return reducer.toPostfixUnaryExpr(operand, op), nil
}

func (reducer *Reducer) toPrefixUnaryExpr(
	op *lr.TokenValue,
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

func (reducer *Reducer) ToPrefixUnaryExpr(
	op *lr.TokenValue,
	operand Expression,
) (
	Expression,
	error,
) {
	return reducer.toPrefixUnaryExpr(op, operand), nil
}

func (reducer *Reducer) ToRecvExpr(
	arrow *lr.TokenValue,
	expr Expression,
) (
	Expression,
	error,
) {
	return reducer.toPrefixUnaryExpr(arrow, expr), nil
}

func (reducer *Reducer) ToUnaryOpAssignStatement(
	operand Expression,
	op *lr.TokenValue,
) (
	Statement,
	error,
) {
	return reducer.toPostfixUnaryExpr(operand, op), nil
}

//
// BinaryExpr
//

func (reducer *Reducer) toBinaryExpr(
	left Expression,
	op *lr.TokenValue,
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

func (reducer *Reducer) ToBinaryMulExpr(
	left Expression,
	op *lr.TokenValue,
	right Expression,
) (
	Expression,
	error,
) {
	return reducer.toBinaryExpr(left, op, right)
}

func (reducer *Reducer) ToBinaryAddExpr(
	left Expression,
	op *lr.TokenValue,
	right Expression,
) (
	Expression,
	error,
) {
	return reducer.toBinaryExpr(left, op, right)
}

func (reducer *Reducer) ToBinaryCmpExpr(
	left Expression,
	op *lr.TokenValue,
	right Expression,
) (
	Expression,
	error,
) {
	return reducer.toBinaryExpr(left, op, right)
}

func (reducer *Reducer) ToBinaryAndExpr(
	left Expression,
	op *lr.TokenValue,
	right Expression,
) (
	Expression,
	error,
) {
	return reducer.toBinaryExpr(left, op, right)
}

func (reducer *Reducer) ToBinaryOrExpr(
	left Expression,
	op *lr.TokenValue,
	right Expression,
) (
	Expression,
	error,
) {
	return reducer.toBinaryExpr(left, op, right)
}

func (reducer *Reducer) ToSendExpr(
	receiver Expression,
	arrow *lr.TokenValue,
	expr Expression,
) (
	Expression,
	error,
) {
	return reducer.toBinaryExpr(receiver, arrow, expr)
}

func (reducer *Reducer) ToBinaryOpAssignStatement(
	address Expression,
	op *lr.TokenValue,
	value Expression,
) (
	Statement,
	error,
) {
	return reducer.toBinaryExpr(address, op, value)
}

func (reducer *Reducer) ToExprAssignStatement(
	pattern Expression,
	assign *lr.TokenValue,
	value Expression,
) (
	Statement,
	error,
) {
	return reducer.toBinaryExpr(pattern, assign, value)
}

func (reducer *Reducer) ToSequenceExprAssignStatement(
	pattern Expression,
	assign *lr.TokenValue,
	value Expression,
) (
	Statement,
	error,
) {
	return reducer.toBinaryExpr(pattern, assign, value)
}

func (reducer *Reducer) DefToGlobalVarDef(
	pattern Expression,
	assign *lr.TokenValue,
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

func (reducer *Reducer) toImplicitStructExpr(
	lparen *lr.TokenValue,
	args *ArgumentList,
	rparen *lr.TokenValue,
) *ImplicitStructExpr {
	args.ReduceMarkers(lparen, rparen)
	return &ImplicitStructExpr{
		ArgumentList: *args,
	}
}

func (reducer *Reducer) ToImplicitStructExpr(
	lparen *lr.TokenValue,
	args *ArgumentList,
	rparen *lr.TokenValue,
) (
	Expression,
	error,
) {
	expr := reducer.toImplicitStructExpr(lparen, args, rparen)
	reducer.ImplicitStructExprs = append(reducer.ImplicitStructExprs, expr)
	return expr, nil
}

func (reducer *Reducer) PairToImproperExprStruct(
	expr1 Expression,
	comma *lr.TokenValue,
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

func (reducer *Reducer) AddToImproperExprStruct(
	structExpr *ImplicitStructExpr,
	comma *lr.TokenValue,
	expr Expression,
) (
	*ImplicitStructExpr,
	error,
) {
	arg := NewPositionalArgument(expr)
	structExpr.ReduceAdd(comma, arg)
	return structExpr, nil
}

func (reducer *Reducer) PairToImproperSequenceExprStruct(
	expr1 Expression,
	comma *lr.TokenValue,
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

func (reducer *Reducer) AddToImproperSequenceExprStruct(
	structExpr *ImplicitStructExpr,
	comma *lr.TokenValue,
	expr Expression,
) (
	*ImplicitStructExpr,
	error,
) {
	arg := NewPositionalArgument(expr)
	structExpr.ReduceAdd(comma, arg)
	return structExpr, nil
}

func (reducer *Reducer) ToTuplePattern(
	lparen *lr.TokenValue,
	list *ArgumentList,
	rparen *lr.TokenValue,
) (
	Expression,
	error,
) {
	return reducer.toImplicitStructExpr(lparen, list, rparen), nil
}

//
// ColonExpr
//

func (Reducer) UnitUnitPairToColonExpr(
	colon *lr.TokenValue,
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

func (Reducer) ExprUnitPairToColonExpr(
	leftExpr Expression,
	colon *lr.TokenValue,
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

func (reducer *Reducer) UnitExprPairToColonExpr(
	colon *lr.TokenValue,
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

func (reducer *Reducer) ExprExprPairToColonExpr(
	leftExpr Expression,
	colon *lr.TokenValue,
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

func (reducer *Reducer) ColonExprUnitTupleToColonExpr(
	list *ColonExpr,
	colon *lr.TokenValue,
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

func (reducer *Reducer) ColonExprExprTupleToColonExpr(
	list *ColonExpr,
	colon *lr.TokenValue,
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

func (reducer *Reducer) ToCallExpr(
	funcExpr Expression,
	genericArguments *GenericArgumentList,
	lparen *lr.TokenValue,
	arguments *ArgumentList,
	rparen *lr.TokenValue,
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

func (reducer *Reducer) ToIndexExpr(
	accessible Expression,
	lbracket *lr.TokenValue,
	index *Argument,
	rbracket *lr.TokenValue,
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

func (reducer *Reducer) ToAsExpr(
	accessible Expression,
	dot *lr.TokenValue,
	as *lr.TokenValue,
	lparen *lr.TokenValue,
	castType TypeExpression,
	rparen *lr.TokenValue,
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

func (reducer *Reducer) ToInitializeExpr(
	initializable TypeExpression,
	lparen *lr.TokenValue,
	arguments *ArgumentList,
	rparen *lr.TokenValue,
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

func (reducer *Reducer) UnlabelledToIfExpr(
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

func (reducer *Reducer) LabelledToIfExpr(
	labelDecl *lr.TokenValue,
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

func (reducer *Reducer) ElseToIfElseExpr(
	ifExpr *IfExpr,
	elseKW *lr.TokenValue,
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

func (reducer *Reducer) ElifToIfElifExpr(
	ifExpr *IfExpr,
	elseKW *lr.TokenValue,
	ifKW *lr.TokenValue,
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

func (reducer *Reducer) ToIfOnlyExpr(
	ifKW *lr.TokenValue,
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

func (reducer *Reducer) LabelledToSwitchExpr(
	labelDecl *lr.TokenValue,
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
	case *ParseErrorNode:
		return expr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", expr))
}

func (reducer *Reducer) ToSwitchExprBody(
	switchKW *lr.TokenValue,
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
	case *ParseErrorNode:
		return expr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", expr))
}

//
// SelectExpr
//

func (reducer *Reducer) LabelledToSelectExpr(
	labelDecl *lr.TokenValue,
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
	case *ParseErrorNode:
		return expr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", expr))
}

func (reducer *Reducer) ToSelectExprBody(
	switchKW *lr.TokenValue,
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
	case *ParseErrorNode:
		return expr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", expr))
}

//
// LoopExpr
//

func (reducer *Reducer) LabelledToLoopExpr(
	labelDecl *lr.TokenValue,
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
	case *ParseErrorNode:
		return expr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", expr))
}

func (reducer *Reducer) InfiniteToLoopExprBody(
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
	case *ParseErrorNode:
		return bodyExpr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", bodyExpr))
}

func (reducer *Reducer) DoWhileToLoopExprBody(
	bodyExpr Expression,
	forKW *lr.TokenValue,
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
	case *ParseErrorNode:
		return bodyExpr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", bodyExpr))
}

func (reducer *Reducer) WhileToLoopExprBody(
	forKW *lr.TokenValue,
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
	case *ParseErrorNode:
		return bodyExpr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", bodyExpr))
}

func (reducer *Reducer) IteratorToLoopExprBody(
	forKW *lr.TokenValue,
	assignPattern Expression,
	in *lr.TokenValue,
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
	case *ParseErrorNode:
		return bodyExpr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", bodyExpr))
}

func (reducer *Reducer) ForToLoopExprBody(
	forKW *lr.TokenValue,
	init Statement,
	semicolon1 *lr.TokenValue,
	condition Expression,
	semicolon2 *lr.TokenValue,
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
	case *ParseErrorNode:
		return bodyExpr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", bodyExpr))
}

func (reducer *Reducer) NilToOptionalSequenceStatement() (
	Statement,
	error,
) {
	return nil, nil
}

func (reducer *Reducer) NilToOptionalSequenceExpr() (
	Expression,
	error,
) {
	return nil, nil
}

func (reducer *Reducer) ToLoopBody(
	do *lr.TokenValue,
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
	case *ParseErrorNode:
		return expr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", expr))
}
