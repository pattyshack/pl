package reducer

import (
	"fmt"

	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser/lr"
)

//
// (Bool/Int/Float/Rune/String)LiteralExpr
//

func (reducer *Reducer) TrueToLiteralExpr(
	value *lr.TokenValue,
) (
	ast.Expression,
	error,
) {
	expr := &ast.BoolLiteralExpr{ValuedNode: value}
	reducer.BoolLiteralExprs = append(reducer.BoolLiteralExprs, expr)
	return expr, nil
}

func (reducer *Reducer) FalseToLiteralExpr(
	value *lr.TokenValue,
) (
	ast.Expression,
	error,
) {
	expr := &ast.BoolLiteralExpr{ValuedNode: value}
	reducer.BoolLiteralExprs = append(reducer.BoolLiteralExprs, expr)
	return expr, nil
}

func (reducer *Reducer) IntegerLiteralToLiteralExpr(
	value *lr.TokenValue,
) (
	ast.Expression,
	error,
) {
	expr := &ast.IntLiteralExpr{ValuedNode: value}
	reducer.IntLiteralExprs = append(reducer.IntLiteralExprs, expr)
	return expr, nil
}

func (reducer *Reducer) FloatLiteralToLiteralExpr(
	value *lr.TokenValue,
) (
	ast.Expression,
	error,
) {
	expr := &ast.FloatLiteralExpr{ValuedNode: value}
	reducer.FloatLiteralExprs = append(reducer.FloatLiteralExprs, expr)
	return expr, nil
}

func (reducer *Reducer) RuneLiteralToLiteralExpr(
	value *lr.TokenValue,
) (
	ast.Expression,
	error,
) {
	expr := &ast.RuneLiteralExpr{ValuedNode: value}
	reducer.RuneLiteralExprs = append(reducer.RuneLiteralExprs, expr)
	return expr, nil
}

func (reducer *Reducer) StringLiteralToLiteralExpr(
	value *lr.TokenValue,
) (
	ast.Expression,
	error,
) {
	expr := &ast.StringLiteralExpr{ValuedNode: value}
	reducer.StringLiteralExprs = append(reducer.StringLiteralExprs, expr)
	return expr, nil
}

//
// NamedExpr
//

func (reducer *Reducer) IdentifierToNamedExpr(
	value *lr.TokenValue,
) (
	ast.Expression,
	error,
) {
	expr := &ast.NamedExpr{ValuedNode: value}
	reducer.NamedExprs = append(reducer.NamedExprs, expr)
	return expr, nil
}

func (reducer *Reducer) UnderscoreToNamedExpr(
	value *lr.TokenValue,
) (
	ast.Expression,
	error,
) {
	expr := &ast.NamedExpr{ValuedNode: value}
	reducer.NamedExprs = append(reducer.NamedExprs, expr)
	return expr, nil
}

//
// AccessExpr
//

func (reducer *Reducer) ToAccessExpr(
	operand ast.Expression,
	dot *lr.TokenValue,
	field *lr.TokenValue,
) (
	ast.Expression,
	error,
) {
	expr := &ast.AccessExpr{
		StartEndPos: ast.NewStartEndPos(operand.Loc(), field.End()),
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
	operand ast.Expression,
	op *lr.TokenValue,
) *ast.UnaryExpr {
	expr := &ast.UnaryExpr{
		StartEndPos: ast.NewStartEndPos(operand.Loc(), op.End()),
		IsPrefix:    false,
		Op:          ast.UnaryOp(op.Value),
		Operand:     operand,
	}

	expr.LeadingComment = operand.TakeLeading()
	operand.AppendToTrailing(op.TakeLeading())
	expr.TrailingComment = op.TakeTrailing()

	reducer.UnaryExprs = append(reducer.UnaryExprs, expr)
	return expr
}

func (reducer *Reducer) ToPostfixUnaryExpr(
	operand ast.Expression,
	op *lr.TokenValue,
) (
	ast.Expression,
	error,
) {
	return reducer.toPostfixUnaryExpr(operand, op), nil
}

func (reducer *Reducer) toPrefixUnaryExpr(
	op *lr.TokenValue,
	operand ast.Expression,
) ast.Expression {
	expr := &ast.UnaryExpr{
		StartEndPos: ast.NewStartEndPos(op.Loc(), operand.End()),
		IsPrefix:    true,
		Op:          ast.UnaryOp(op.Value),
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
	operand ast.Expression,
) (
	ast.Expression,
	error,
) {
	return reducer.toPrefixUnaryExpr(op, operand), nil
}

func (reducer *Reducer) ToRecvExpr(
	arrow *lr.TokenValue,
	expr ast.Expression,
) (
	ast.Expression,
	error,
) {
	return reducer.toPrefixUnaryExpr(arrow, expr), nil
}

//
// BinaryExpr
//

func (reducer *Reducer) toBinaryExpr(
	left ast.Expression,
	op *lr.TokenValue,
	right ast.Expression,
) ast.Expression {
	expr := &ast.BinaryExpr{
		StartEndPos: ast.NewStartEndPos(left.Loc(), right.End()),
		Left:        left,
		Op:          ast.BinaryOp(op.Value),
		Right:       right,
	}

	expr.LeadingComment = left.TakeLeading()
	left.AppendToTrailing(op.TakeLeading())
	right.PrependToLeading(op.TakeTrailing())
	expr.TrailingComment = right.TakeTrailing()

	reducer.BinaryExprs = append(reducer.BinaryExprs, expr)
	return expr
}

func (reducer *Reducer) ToBinaryMulExpr(
	left ast.Expression,
	op *lr.TokenValue,
	right ast.Expression,
) (
	ast.Expression,
	error,
) {
	return reducer.toBinaryExpr(left, op, right), nil
}

func (reducer *Reducer) ToBinaryAddExpr(
	left ast.Expression,
	op *lr.TokenValue,
	right ast.Expression,
) (
	ast.Expression,
	error,
) {
	return reducer.toBinaryExpr(left, op, right), nil
}

func (reducer *Reducer) ToBinaryCmpExpr(
	left ast.Expression,
	op *lr.TokenValue,
	right ast.Expression,
) (
	ast.Expression,
	error,
) {
	return reducer.toBinaryExpr(left, op, right), nil
}

func (reducer *Reducer) ToBinaryAndExpr(
	left ast.Expression,
	op *lr.TokenValue,
	right ast.Expression,
) (
	ast.Expression,
	error,
) {
	return reducer.toBinaryExpr(left, op, right), nil
}

func (reducer *Reducer) ToBinaryOrExpr(
	left ast.Expression,
	op *lr.TokenValue,
	right ast.Expression,
) (
	ast.Expression,
	error,
) {
	return reducer.toBinaryExpr(left, op, right), nil
}

func (reducer *Reducer) ToSendExpr(
	receiver ast.Expression,
	arrow *lr.TokenValue,
	expr ast.Expression,
) (
	ast.Expression,
	error,
) {
	return reducer.toBinaryExpr(receiver, arrow, expr), nil
}

func (reducer *Reducer) ToBinaryAssignOpExpr(
	address ast.Expression,
	op *lr.TokenValue,
	value ast.Expression,
) (
	ast.Expression,
	error,
) {
	return reducer.toBinaryExpr(address, op, value), nil
}

func (reducer *Reducer) ToExprAssignStatement(
	pattern ast.Expression,
	assign *lr.TokenValue,
	value ast.Expression,
) (
	ast.Statement,
	error,
) {
	return reducer.toBinaryExpr(pattern, assign, value), nil
}

func (reducer *Reducer) DefToGlobalVarDef(
	pattern ast.Expression,
	assign *lr.TokenValue,
	value ast.Expression,
) (
	ast.Definition,
	error,
) {
	return reducer.toBinaryExpr(pattern, assign, value), nil
}

//
// ImplicitStructExpr
//

func (reducer *Reducer) toImplicitStructExpr(
	lparen *lr.TokenValue,
	args *ast.ArgumentList,
	rparen *lr.TokenValue,
) *ast.ImplicitStructExpr {
	args.ReduceMarkers(lparen, rparen)
	return &ast.ImplicitStructExpr{
		ArgumentList: *args,
	}
}

func (reducer *Reducer) ToImplicitStructExpr(
	lparen *lr.TokenValue,
	args *ast.ArgumentList,
	rparen *lr.TokenValue,
) (
	ast.Expression,
	error,
) {
	expr := reducer.toImplicitStructExpr(lparen, args, rparen)
	reducer.ImplicitStructExprs = append(reducer.ImplicitStructExprs, expr)
	return expr, nil
}

func (reducer *Reducer) PairToImproperExprStruct(
	expr1 ast.Expression,
	comma *lr.TokenValue,
	expr2 ast.Expression,
) (
	*ast.ImplicitStructExpr,
	error,
) {
	arg1 := ast.NewPositionalArgument(expr1)
	arg2 := ast.NewPositionalArgument(expr2)

	list := ast.NewArgumentList()
	list.Add(arg1)
	list.ReduceAdd(comma, arg2)

	expr := &ast.ImplicitStructExpr{
		ArgumentList: *list,
	}

	reducer.ImplicitStructExprs = append(reducer.ImplicitStructExprs, expr)
	return expr, nil
}

func (reducer *Reducer) AddToImproperExprStruct(
	structExpr *ast.ImplicitStructExpr,
	comma *lr.TokenValue,
	expr ast.Expression,
) (
	*ast.ImplicitStructExpr,
	error,
) {
	arg := ast.NewPositionalArgument(expr)
	structExpr.ReduceAdd(comma, arg)
	return structExpr, nil
}

func (reducer *Reducer) ToTuplePattern(
	lparen *lr.TokenValue,
	list *ast.ArgumentList,
	rparen *lr.TokenValue,
) (
	ast.Expression,
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
	*ast.ColonExpr,
	error,
) {
	leftArg := &ast.Argument{
		StartEndPos: ast.NewStartEndPos(colon.Loc(), colon.End()),
		Kind:        ast.IsImplicitUnitArgument,
	}

	rightArg := &ast.Argument{
		StartEndPos: ast.NewStartEndPos(colon.Loc(), colon.End()),
		Kind:        ast.IsImplicitUnitArgument,
	}
	rightArg.LeadingComment = colon.TakeTrailing()

	args := ast.NewNodeList[*ast.Argument]("ColonExpr")
	args.Add(leftArg)
	args.ReduceAdd(colon, rightArg)

	return &ast.ColonExpr{
		ArgumentList: *args,
	}, nil
}

func (Reducer) ExprUnitPairToColonExpr(
	leftExpr ast.Expression,
	colon *lr.TokenValue,
) (
	*ast.ColonExpr,
	error,
) {
	leftArg := &ast.Argument{
		StartEndPos: ast.NewStartEndPos(leftExpr.Loc(), leftExpr.End()),
		Kind:        ast.PositionalArgument,
		Expr:        leftExpr,
	}
	leftArg.LeadingComment = leftExpr.TakeLeading()
	leftArg.TrailingComment = leftExpr.TakeTrailing()

	rightArg := &ast.Argument{
		StartEndPos: ast.NewStartEndPos(colon.Loc(), colon.End()),
		Kind:        ast.IsImplicitUnitArgument,
	}
	rightArg.LeadingComment = colon.TakeTrailing()

	args := ast.NewNodeList[*ast.Argument]("ColonExpr")
	args.Add(leftArg)
	args.ReduceAdd(colon, rightArg)

	return &ast.ColonExpr{
		ArgumentList: *args,
	}, nil
}

func (reducer *Reducer) UnitExprPairToColonExpr(
	colon *lr.TokenValue,
	rightExpr ast.Expression,
) (
	*ast.ColonExpr,
	error,
) {
	leftArg := &ast.Argument{
		StartEndPos: ast.NewStartEndPos(colon.Loc(), colon.End()),
		Kind:        ast.IsImplicitUnitArgument,
	}

	rightArg := &ast.Argument{
		StartEndPos: ast.NewStartEndPos(rightExpr.Loc(), rightExpr.End()),
		Kind:        ast.PositionalArgument,
		Expr:        rightExpr,
	}
	rightArg.LeadingComment = rightExpr.TakeLeading()
	rightArg.TrailingComment = rightExpr.TakeTrailing()
	rightArg.PrependToLeading(colon.TakeTrailing())

	args := ast.NewNodeList[*ast.Argument]("ColonExpr")
	args.Add(leftArg)
	args.ReduceAdd(colon, rightArg)

	return &ast.ColonExpr{
		ArgumentList: *args,
	}, nil
}

func (reducer *Reducer) ExprExprPairToColonExpr(
	leftExpr ast.Expression,
	colon *lr.TokenValue,
	rightExpr ast.Expression,
) (
	*ast.ColonExpr,
	error,
) {
	leftArg := &ast.Argument{
		StartEndPos: ast.NewStartEndPos(leftExpr.Loc(), leftExpr.End()),
		Kind:        ast.PositionalArgument,
		Expr:        leftExpr,
	}
	leftArg.LeadingComment = leftExpr.TakeLeading()
	leftArg.TrailingComment = leftExpr.TakeTrailing()

	rightArg := &ast.Argument{
		StartEndPos: ast.NewStartEndPos(rightExpr.Loc(), rightExpr.End()),
		Kind:        ast.PositionalArgument,
		Expr:        rightExpr,
	}
	rightArg.LeadingComment = rightExpr.TakeLeading()
	rightArg.TrailingComment = rightExpr.TakeTrailing()
	rightArg.PrependToLeading(colon.TakeTrailing())

	args := ast.NewNodeList[*ast.Argument]("ColonExpr")
	args.Add(leftArg)
	args.ReduceAdd(colon, rightArg)

	return &ast.ColonExpr{
		ArgumentList: *args,
	}, nil
}

func (reducer *Reducer) ColonExprUnitTupleToColonExpr(
	list *ast.ColonExpr,
	colon *lr.TokenValue,
) (
	*ast.ColonExpr,
	error,
) {
	arg := &ast.Argument{
		StartEndPos: ast.NewStartEndPos(colon.Loc(), colon.End()),
		Kind:        ast.IsImplicitUnitArgument,
	}
	arg.LeadingComment = colon.TakeTrailing()

	list.ReduceAdd(colon, arg)
	return list, nil
}

func (reducer *Reducer) ColonExprExprTupleToColonExpr(
	list *ast.ColonExpr,
	colon *lr.TokenValue,
	expr ast.Expression,
) (
	*ast.ColonExpr,
	error,
) {
	arg := &ast.Argument{
		StartEndPos: ast.NewStartEndPos(expr.Loc(), expr.End()),
		Kind:        ast.PositionalArgument,
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
	funcExpr ast.Expression,
	genericArguments *ast.GenericArgumentList,
	lparen *lr.TokenValue,
	arguments *ast.ArgumentList,
	rparen *lr.TokenValue,
) (
	ast.Expression,
	error,
) {
	if genericArguments == nil {
		genericArguments = ast.NewGenericArgumentList()
	}

	arguments.ReduceMarkers(lparen, rparen)

	expr := &ast.CallExpr{
		StartEndPos: ast.NewStartEndPos(funcExpr.Loc(), rparen.End()),
		LeadingTrailingComments: ast.LeadingTrailingComments{
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
	accessible ast.Expression,
	lbracket *lr.TokenValue,
	index *ast.Argument,
	rbracket *lr.TokenValue,
) (
	ast.Expression,
	error,
) {
	accessible.AppendToTrailing(lbracket.TakeLeading())
	index.PrependToLeading(lbracket.TakeTrailing())
	index.AppendToTrailing(rbracket.TakeLeading())

	expr := &ast.IndexExpr{
		StartEndPos: ast.NewStartEndPos(accessible.Loc(), rbracket.End()),
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
	accessible ast.Expression,
	dot *lr.TokenValue,
	as *lr.TokenValue,
	lparen *lr.TokenValue,
	castType ast.TypeExpression,
	rparen *lr.TokenValue,
) (
	ast.Expression,
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

	expr := &ast.AsExpr{
		StartEndPos: ast.NewStartEndPos(accessible.Loc(), rparen.End()),
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
	initializable ast.TypeExpression,
	lparen *lr.TokenValue,
	arguments *ast.ArgumentList,
	rparen *lr.TokenValue,
) (
	ast.Expression,
	error,
) {
	arguments.ReduceMarkers(lparen, rparen)

	expr := &ast.InitializeExpr{
		StartEndPos:   ast.NewStartEndPos(initializable.Loc(), rparen.End()),
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
	ifExpr *ast.IfExpr,
) (
	ast.Expression,
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
	ifExpr *ast.IfExpr,
) (
	ast.Expression,
	error,
) {
	reducer.UnlabelledToIfExpr(ifExpr)

	ifExpr.PrependToLeading(labelDecl.TakeTrailing())
	ifExpr.PrependToLeading(labelDecl.TakeLeading())

	ifExpr.LabelDecl = labelDecl.Value
	return ifExpr, nil
}

func (reducer *Reducer) ElseToIfElseExpr(
	ifExpr *ast.IfExpr,
	elseKW *lr.TokenValue,
	branch ast.Expression,
) (
	*ast.IfExpr,
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
	ifExpr *ast.IfExpr,
	elseKW *lr.TokenValue,
	ifKW *lr.TokenValue,
	condition ast.Expression,
	branch ast.Expression,
) (
	*ast.IfExpr,
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
		ast.ConditionBranch{condition, branch})

	return ifExpr, nil
}

func (reducer *Reducer) ToIfOnlyExpr(
	ifKW *lr.TokenValue,
	condition ast.Expression,
	branch ast.Expression,
) (
	*ast.IfExpr,
	error,
) {
	condition.PrependToLeading(ifKW.TakeTrailing())
	expr := &ast.IfExpr{
		StartEndPos:       ast.NewStartEndPos(ifKW.Loc(), branch.End()),
		ConditionBranches: []ast.ConditionBranch{{condition, branch}},
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
	expr ast.Expression,
) (
	ast.Expression,
	error,
) {
	switch switchExpr := expr.(type) {
	case *ast.SwitchExpr:
		switchExpr.StartPos = labelDecl.Loc()
		switchExpr.PrependToLeading(labelDecl.TakeTrailing())
		switchExpr.PrependToLeading(labelDecl.TakeLeading())
		switchExpr.LabelDecl = labelDecl.Value
		return switchExpr, nil
	case *ast.ParseErrorNode:
		return expr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", expr))
}

func (reducer *Reducer) ToSwitchExprBody(
	switchKW *lr.TokenValue,
	operand ast.Expression,
	expr ast.Expression,
) (
	ast.Expression,
	error,
) {
	switch branches := expr.(type) {
	case *ast.StatementsExpr:
		trailing := branches.TakeTrailing()

		switchExpr := &ast.SwitchExpr{
			StartEndPos: ast.NewStartEndPos(switchKW.Loc(), branches.End()),
			Operand:     operand,
			Branches:    branches,
		}

		switchExpr.LeadingComment = switchKW.TakeLeading()
		operand.PrependToLeading(switchKW.TakeTrailing())
		switchExpr.TrailingComment = trailing

		reducer.SwitchExprs = append(reducer.SwitchExprs, switchExpr)
		return switchExpr, nil
	case *ast.ParseErrorNode:
		return expr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", expr))
}

//
// SelectExpr
//

func (reducer *Reducer) LabelledToSelectExpr(
	labelDecl *lr.TokenValue,
	expr ast.Expression,
) (
	ast.Expression,
	error,
) {
	switch selectExpr := expr.(type) {
	case *ast.SelectExpr:
		selectExpr.StartPos = labelDecl.Loc()
		selectExpr.PrependToLeading(labelDecl.TakeTrailing())
		selectExpr.PrependToLeading(labelDecl.TakeLeading())
		selectExpr.LabelDecl = labelDecl.Value
		return selectExpr, nil
	case *ast.ParseErrorNode:
		return expr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", expr))
}

func (reducer *Reducer) ToSelectExprBody(
	switchKW *lr.TokenValue,
	expr ast.Expression,
) (
	ast.Expression,
	error,
) {
	switch branches := expr.(type) {
	case *ast.StatementsExpr:
		branches.PrependToLeading(switchKW.TakeTrailing())
		trailing := branches.TakeTrailing()

		selectExpr := &ast.SelectExpr{
			StartEndPos: ast.NewStartEndPos(switchKW.Loc(), branches.End()),
			Branches:    branches,
		}

		selectExpr.LeadingComment = switchKW.TakeLeading()
		selectExpr.TrailingComment = trailing

		reducer.SelectExprs = append(reducer.SelectExprs, selectExpr)
		return selectExpr, nil
	case *ast.ParseErrorNode:
		return expr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", expr))
}

//
// LoopExpr
//

func (reducer *Reducer) LabelledToLoopExpr(
	labelDecl *lr.TokenValue,
	expr ast.Expression,
) (
	ast.Expression,
	error,
) {
	switch loop := expr.(type) {
	case *ast.LoopExpr:
		loop.StartPos = labelDecl.Loc()
		loop.PrependToLeading(labelDecl.TakeTrailing())
		loop.PrependToLeading(labelDecl.TakeLeading())
		loop.LabelDecl = labelDecl.Value
		return loop, nil
	case *ast.ParseErrorNode:
		return expr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", expr))
}

func (reducer *Reducer) InfiniteToLoopExprBody(
	bodyExpr ast.Expression,
) (
	ast.Expression,
	error,
) {
	switch body := bodyExpr.(type) {
	case *ast.StatementsExpr:
		leading := body.TakeLeading()
		trailing := body.TakeTrailing()

		loop := &ast.LoopExpr{
			StartEndPos: ast.NewStartEndPos(bodyExpr.Loc(), bodyExpr.End()),
			LoopKind:    ast.InfiniteLoop,
			Body:        body,
		}
		loop.LeadingComment = leading
		loop.TrailingComment = trailing

		return loop, nil
	case *ast.ParseErrorNode:
		return bodyExpr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", bodyExpr))
}

func (reducer *Reducer) DoWhileToLoopExprBody(
	bodyExpr ast.Expression,
	forKW *lr.TokenValue,
	condition ast.Expression,
) (
	ast.Expression,
	error,
) {
	switch body := bodyExpr.(type) {
	case *ast.StatementsExpr:
		leading := body.TakeLeading()
		body.AppendToTrailing(forKW.TakeLeading())
		condition.PrependToLeading(forKW.TakeTrailing())
		trailing := condition.TakeTrailing()

		loop := &ast.LoopExpr{
			StartEndPos: ast.NewStartEndPos(bodyExpr.Loc(), bodyExpr.End()),
			LoopKind:    ast.DoWhileLoop,
			Condition:   condition,
			Body:        body,
		}
		loop.LeadingComment = leading
		loop.TrailingComment = trailing

		return loop, nil
	case *ast.ParseErrorNode:
		return bodyExpr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", bodyExpr))
}

func (reducer *Reducer) WhileToLoopExprBody(
	forKW *lr.TokenValue,
	condition ast.Expression,
	bodyExpr ast.Expression,
) (
	ast.Expression,
	error,
) {
	switch body := bodyExpr.(type) {
	case *ast.StatementsExpr:
		leading := forKW.TakeLeading()
		condition.PrependToLeading(forKW.TakeTrailing())
		trailing := bodyExpr.TakeTrailing()

		loop := &ast.LoopExpr{
			StartEndPos: ast.NewStartEndPos(bodyExpr.Loc(), bodyExpr.End()),
			LoopKind:    ast.WhileLoop,
			Condition:   condition,
			Body:        body,
		}
		loop.LeadingComment = leading
		loop.TrailingComment = trailing

		return loop, nil
	case *ast.ParseErrorNode:
		return bodyExpr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", bodyExpr))
}

func (reducer *Reducer) IteratorToLoopExprBody(
	forKW *lr.TokenValue,
	assignPattern ast.Expression,
	in *lr.TokenValue,
	iterator ast.Expression,
	bodyExpr ast.Expression,
) (
	ast.Expression,
	error,
) {
	switch body := bodyExpr.(type) {
	case *ast.StatementsExpr:
		leading := forKW.TakeLeading()
		trailing := bodyExpr.TakeTrailing()

		loop := &ast.LoopExpr{
			StartEndPos: ast.NewStartEndPos(bodyExpr.Loc(), bodyExpr.End()),
			LoopKind:    ast.IteratorLoop,
			Condition:   reducer.toBinaryExpr(assignPattern, in, iterator),
			Body:        body,
		}
		loop.LeadingComment = leading
		loop.TrailingComment = trailing

		return loop, nil
	case *ast.ParseErrorNode:
		return bodyExpr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", bodyExpr))
}

func (reducer *Reducer) ForToLoopExprBody(
	forKW *lr.TokenValue,
	init ast.Statement,
	semicolon1 *lr.TokenValue,
	condition ast.Expression,
	semicolon2 *lr.TokenValue,
	post ast.Statement,
	bodyExpr ast.Expression,
) (
	ast.Expression,
	error,
) {
	switch body := bodyExpr.(type) {
	case *ast.StatementsExpr:
		var last ast.Node = body
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

		loop := &ast.LoopExpr{
			StartEndPos: ast.NewStartEndPos(bodyExpr.Loc(), bodyExpr.End()),
			LoopKind:    ast.ForLoop,
			Init:        init,
			Condition:   condition,
			Post:        post,
			Body:        body,
		}
		loop.LeadingComment = leading
		loop.TrailingComment = trailing

		return loop, nil
	case *ast.ParseErrorNode:
		return bodyExpr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", bodyExpr))
}

func (reducer *Reducer) NilToOptionalStatement() (
	ast.Statement,
	error,
) {
	return nil, nil
}

func (reducer *Reducer) NilToOptionalSimpleExpr() (
	ast.Expression,
	error,
) {
	return nil, nil
}

func (reducer *Reducer) toLoopBody(
	kw *lr.TokenValue,
	expr ast.Expression,
) ast.Expression {
	switch body := expr.(type) {
	case *ast.StatementsExpr:
		body.StartPos = kw.Loc()
		body.PrependToLeading(kw.TakeTrailing())
		body.PrependToLeading(kw.TakeLeading())
		return body
	case *ast.ParseErrorNode:
		return expr
	}

	panic(fmt.Sprintf("Unexpected expression: %v", expr))
}

func (reducer *Reducer) ToRepeatLoopBody(
	repeat *lr.TokenValue,
	expr ast.Expression,
) (
	ast.Expression,
	error,
) {
	return reducer.toLoopBody(repeat, expr), nil
}

func (reducer *Reducer) ToForLoopBody(
	do *lr.TokenValue,
	expr ast.Expression,
) (
	ast.Expression,
	error,
) {
	return reducer.toLoopBody(do, expr), nil
}
