package reducer

import (
	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser/lr"
)

//
// ParseErrorExpr
//

func (reducer *Reducer) ToParseErrorExpr(
	pe *lr.ParseErrorSymbol,
) (
	ast.Expression,
	error,
) {
	expr := ast.ParseErrorExpr(*pe)
	return &expr, nil
}

//
// (Bool/Int/Float/Rune/String)LiteralExpr
//

func (reducer *Reducer) TrueToLiteralExpr(
	value *lr.TokenValue,
) (
	ast.Expression,
	error,
) {
	expr := &ast.LiteralExpr{
		StartEndPos:             value.StartEndPos,
		LeadingTrailingComments: value.LeadingTrailingComments,
		Kind:                    ast.BoolLiteral,
		Value:                   value.Value,
	}
	return expr, nil
}

func (reducer *Reducer) FalseToLiteralExpr(
	value *lr.TokenValue,
) (
	ast.Expression,
	error,
) {
	expr := &ast.LiteralExpr{
		StartEndPos:             value.StartEndPos,
		LeadingTrailingComments: value.LeadingTrailingComments,
		Kind:                    ast.BoolLiteral,
		Value:                   value.Value,
	}
	return expr, nil
}

func (reducer *Reducer) IntegerLiteralToLiteralExpr(
	value *lr.TokenValue,
) (
	ast.Expression,
	error,
) {
	expr := &ast.LiteralExpr{
		StartEndPos:             value.StartEndPos,
		LeadingTrailingComments: value.LeadingTrailingComments,
		Kind:                    ast.IntLiteral,
		Value:                   value.Value,
	}
	return expr, nil
}

func (reducer *Reducer) FloatLiteralToLiteralExpr(
	value *lr.TokenValue,
) (
	ast.Expression,
	error,
) {
	expr := &ast.LiteralExpr{
		StartEndPos:             value.StartEndPos,
		LeadingTrailingComments: value.LeadingTrailingComments,
		Kind:                    ast.FloatLiteral,
		Value:                   value.Value,
	}
	return expr, nil
}

func (reducer *Reducer) RuneLiteralToLiteralExpr(
	value *lr.TokenValue,
) (
	ast.Expression,
	error,
) {
	expr := &ast.LiteralExpr{
		StartEndPos:             value.StartEndPos,
		LeadingTrailingComments: value.LeadingTrailingComments,
		Kind:                    ast.RuneLiteral,
		Value:                   value.Value,
	}
	return expr, nil
}

func (reducer *Reducer) StringLiteralToLiteralExpr(
	value *lr.TokenValue,
) (
	ast.Expression,
	error,
) {
	expr := &ast.LiteralExpr{
		StartEndPos:             value.StartEndPos,
		LeadingTrailingComments: value.LeadingTrailingComments,
		Kind:                    ast.StringLiteral,
		Value:                   value.Value,
	}
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
	expr := &ast.NamedExpr{
		StartEndPos:             value.StartEndPos,
		LeadingTrailingComments: value.LeadingTrailingComments,
		Name:                    value.Value,
	}
	return expr, nil
}

func (reducer *Reducer) UnderscoreToNamedExpr(
	value *lr.TokenValue,
) (
	ast.Expression,
	error,
) {
	expr := &ast.NamedExpr{
		StartEndPos:             value.StartEndPos,
		LeadingTrailingComments: value.LeadingTrailingComments,
		Name:                    value.Value,
	}
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
		Field:       field.Value,
	}

	expr.LeadingComment = operand.TakeLeading()
	operand.AppendToTrailing(dot.TakeLeading())
	operand.AppendToTrailing(dot.TakeTrailing())
	operand.AppendToTrailing(field.TakeLeading())
	expr.TrailingComment = field.TakeTrailing()

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
) *ast.BinaryExpr {
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

//
// ImplicitStructExpr
//

func (reducer *Reducer) toImplicitStructExpr(
	lparen *lr.TokenValue,
	args *ast.ArgumentList,
	rparen *lr.TokenValue,
) *ast.ImplicitStructExpr {
	if args == nil {
		args = ast.NewArgumentList()
	}
	args.ReduceMarkers(lparen, rparen)
	args.LeadingComment.Append(args.MiddleComment)

	return &ast.ImplicitStructExpr{
		StartEndPos:             args.StartEndPos,
		LeadingTrailingComments: args.TakeComments(),
		Kind:                    ast.ProperImplicitStruct,
		Arguments:               args.Elements,
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
	list := []*ast.Argument{ast.NewPositionalArgument(expr1)}
	list = ast.ReduceAdd(list, comma, ast.NewPositionalArgument(expr2))

	expr := &ast.ImplicitStructExpr{
		StartEndPos: ast.NewStartEndPos(expr1.Loc(), expr2.End()),
		Kind:        ast.ImproperImplicitStruct,
		Arguments:   list,
	}

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
	structExpr.EndPos = expr.End()
	structExpr.Arguments = ast.ReduceAdd(
		structExpr.Arguments,
		comma,
		ast.NewPositionalArgument(expr))
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
	*ast.ImplicitStructExpr,
	error,
) {
	args := ast.ReduceAdd(
		[]*ast.Argument{
			ast.NewPositionalArgument(
				&ast.ImplicitStructExpr{
					StartEndPos: colon.StartEnd(),
					Kind:        ast.ImproperImplicitStruct,
				}),
		},
		colon,
		ast.NewPositionalArgument(
			&ast.ImplicitStructExpr{
				StartEndPos: colon.StartEnd(),
				Kind:        ast.ImproperImplicitStruct,
			}))

	return &ast.ImplicitStructExpr{
		StartEndPos: colon.StartEnd(),
		Kind:        ast.ColonImplicitStruct,
		Arguments:   args,
	}, nil
}

func (Reducer) ExprUnitPairToColonExpr(
	leftExpr ast.Expression,
	colon *lr.TokenValue,
) (
	*ast.ImplicitStructExpr,
	error,
) {
	args := ast.ReduceAdd(
		[]*ast.Argument{ast.NewPositionalArgument(leftExpr)},
		colon,
		ast.NewPositionalArgument(
			&ast.ImplicitStructExpr{
				StartEndPos: colon.StartEnd(),
				Kind:        ast.ImproperImplicitStruct,
			}))

	return &ast.ImplicitStructExpr{
		StartEndPos: ast.NewStartEndPos(leftExpr.Loc(), colon.End()),
		Kind:        ast.ColonImplicitStruct,
		Arguments:   args,
	}, nil
}

func (reducer *Reducer) UnitExprPairToColonExpr(
	colon *lr.TokenValue,
	rightExpr ast.Expression,
) (
	*ast.ImplicitStructExpr,
	error,
) {
	args := ast.ReduceAdd(
		[]*ast.Argument{
			ast.NewPositionalArgument(
				&ast.ImplicitStructExpr{
					StartEndPos: colon.StartEnd(),
					Kind:        ast.ImproperImplicitStruct,
				}),
		},
		colon,
		ast.NewPositionalArgument(rightExpr))

	return &ast.ImplicitStructExpr{
		StartEndPos: ast.NewStartEndPos(colon.Loc(), rightExpr.End()),
		Kind:        ast.ColonImplicitStruct,
		Arguments:   args,
	}, nil
}

func (reducer *Reducer) ExprExprPairToColonExpr(
	leftExpr ast.Expression,
	colon *lr.TokenValue,
	rightExpr ast.Expression,
) (
	*ast.ImplicitStructExpr,
	error,
) {
	args := ast.ReduceAdd(
		[]*ast.Argument{ast.NewPositionalArgument(leftExpr)},
		colon,
		ast.NewPositionalArgument(rightExpr))

	return &ast.ImplicitStructExpr{
		StartEndPos: ast.NewStartEndPos(leftExpr.Loc(), rightExpr.End()),
		Kind:        ast.ColonImplicitStruct,
		Arguments:   args,
	}, nil
}

func (reducer *Reducer) ColonExprUnitTupleToColonExpr(
	colonExpr *ast.ImplicitStructExpr,
	colon *lr.TokenValue,
) (
	*ast.ImplicitStructExpr,
	error,
) {
	colonExpr.Arguments = ast.ReduceAdd(
		colonExpr.Arguments,
		colon,
		ast.NewPositionalArgument(
			&ast.ImplicitStructExpr{
				StartEndPos: colon.StartEnd(),
				Kind:        ast.ImproperImplicitStruct,
			}))
	colonExpr.EndPos = colon.End()
	return colonExpr, nil
}

func (reducer *Reducer) ColonExprExprTupleToColonExpr(
	colonExpr *ast.ImplicitStructExpr,
	colon *lr.TokenValue,
	expr ast.Expression,
) (
	*ast.ImplicitStructExpr,
	error,
) {
	colonExpr.Arguments = ast.ReduceAdd(
		colonExpr.Arguments,
		colon,
		ast.NewPositionalArgument(expr))
	colonExpr.EndPos = expr.End()
	return colonExpr, nil
}

//
// CallExpr
//

func (reducer *Reducer) ToCallExpr(
	funcExpr ast.Expression,
	genericArguments *ast.TypeExpressionList,
	lparen *lr.TokenValue,
	arguments *ast.ArgumentList,
	rparen *lr.TokenValue,
) (
	ast.Expression,
	error,
) {
	leading := funcExpr.TakeLeading()

	var prev ast.Node = funcExpr
	if genericArguments != nil {
		prev = genericArguments
	}

	if arguments == nil {
		arguments = ast.NewArgumentList()
	}
	arguments.ReduceMarkers(lparen, rparen)

	prev.AppendToTrailing(arguments.TakeLeading())
	trailing := arguments.TakeTrailing()

	expr := &ast.CallExpr{
		StartEndPos:      ast.NewStartEndPos(funcExpr.Loc(), rparen.End()),
		FuncExpr:         funcExpr,
		GenericArguments: genericArguments,
		Arguments:        arguments.Elements,
	}
	expr.LeadingComment = leading
	expr.TrailingComment = trailing

	return expr, nil
}

//
// IndexExpr
//

func (reducer *Reducer) ToIndexExpr(
	accessible ast.Expression,
	lbracket *lr.TokenValue,
	index ast.Expression,
	rbracket *lr.TokenValue,
) (
	ast.Expression,
	error,
) {
	args := []ast.Expression{}
	implicitStruct, ok := index.(*ast.ImplicitStructExpr)
	if ok && implicitStruct.Kind == ast.ColonImplicitStruct {
		for _, arg := range implicitStruct.Arguments {
			if arg.Kind != ast.PositionalArgument {
				panic("Should never happen")
			}

			expr := arg.Expr
			expr.PrependToLeading(arg.TakeLeading())
			expr.AppendToTrailing(arg.TakeTrailing())

			args = append(args, expr)
		}
	} else {
		args = append(args, index)
	}

	accessible.AppendToTrailing(lbracket.TakeLeading())
	index.PrependToLeading(lbracket.TakeTrailing())
	index.AppendToTrailing(rbracket.TakeLeading())

	expr := &ast.IndexExpr{
		StartEndPos: ast.NewStartEndPos(accessible.Loc(), rbracket.End()),
		Accessible:  accessible,
		IndexArgs:   args,
	}

	expr.LeadingComment = accessible.TakeLeading()
	expr.TrailingComment = rbracket.TakeTrailing()

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
		Arguments:     arguments.Elements,
	}

	expr.LeadingComment = initializable.TakeLeading()
	initializable.AppendToTrailing(arguments.TakeLeading())
	initializable.AppendToTrailing(arguments.MiddleComment)
	expr.TrailingComment = arguments.TakeTrailing()

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
	branches := ifExpr.ConditionBranches
	ifExpr.TrailingComment = branches[len(branches)-1].TakeTrailing()
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
	branch *ast.StatementsExpr,
) (
	*ast.IfExpr,
	error,
) {
	last := ifExpr.ConditionBranches[len(ifExpr.ConditionBranches)-1]
	last.AppendToTrailing(elseKW.TakeLeading())

	cb := &ast.ConditionBranchStmt{
		StartEndPos:     ast.NewStartEndPos(elseKW.Loc(), branch.End()),
		IsDefaultBranch: true,
		Branch:          branch,
	}
	cb.LeadingComment = elseKW.TakeTrailing()
	cb.LeadingComment.Append(branch.TakeLeading())
	cb.TrailingComment = branch.TakeTrailing()

	ifExpr.EndPos = branch.End()
	ifExpr.ConditionBranches = append(ifExpr.ConditionBranches, cb)
	return ifExpr, nil
}

func (reducer *Reducer) ElifToIfElifExpr(
	ifExpr *ast.IfExpr,
	elseKW *lr.TokenValue,
	ifKW *lr.TokenValue,
	condition ast.Expression,
	branch *ast.StatementsExpr,
) (
	*ast.IfExpr,
	error,
) {
	last := ifExpr.ConditionBranches[len(ifExpr.ConditionBranches)-1]
	last.AppendToTrailing(elseKW.TakeLeading())

	cb := &ast.ConditionBranchStmt{
		StartEndPos: ast.NewStartEndPos(elseKW.Loc(), branch.End()),
		Condition:   condition,
		Branch:      branch,
	}
	cb.LeadingComment = elseKW.TakeTrailing()
	cb.LeadingComment.Append(ifKW.TakeLeading())
	cb.LeadingComment.Append(ifKW.TakeTrailing())
	cb.LeadingComment.Append(condition.TakeLeading())
	cb.TrailingComment = branch.TakeTrailing()

	ifExpr.EndPos = branch.End()
	ifExpr.ConditionBranches = append(ifExpr.ConditionBranches, cb)
	return ifExpr, nil
}

func (reducer *Reducer) ToIfOnlyExpr(
	ifKW *lr.TokenValue,
	condition ast.Expression,
	branch *ast.StatementsExpr,
) (
	*ast.IfExpr,
	error,
) {
	leading := ifKW.TakeLeading()

	cb := &ast.ConditionBranchStmt{
		StartEndPos: ast.NewStartEndPos(ifKW.Loc(), branch.End()),
		Condition:   condition,
		Branch:      branch,
	}
	cb.LeadingComment = ifKW.TakeTrailing()
	cb.LeadingComment.Append(condition.TakeLeading())
	cb.TrailingComment = branch.TakeTrailing()

	expr := &ast.IfExpr{
		StartEndPos:       ast.NewStartEndPos(ifKW.Loc(), branch.End()),
		ConditionBranches: []*ast.ConditionBranchStmt{cb},
	}
	expr.LeadingComment = leading

	return expr, nil
}

//
// SwitchExpr
//

func (reducer *Reducer) groupBranches(
	stmts *ast.StatementsExpr,
) []*ast.ConditionBranchStmt {
	branches := make([]*ast.ConditionBranchStmt, 0, len(stmts.Statements))
	var current *ast.ConditionBranchStmt
	for _, stmt := range stmts.Statements {
		branch, ok := stmt.(*ast.ConditionBranchStmt)
		if !ok {
			if current != nil {
				current.Branch.Statements = append(current.Branch.Statements, stmt)
				current.Branch.EndPos = stmt.End()
				current.EndPos = stmt.End()
			} else {
				reducer.Emit(
					"%s: unexpected statement. "+
						"expected case or default branch statement",
					stmt.Loc())
			}
			continue
		}

		current = branch
		branches = append(branches, branch)

		// Flatten nested case statements
		for {
			if len(current.Branch.Statements) == 0 {
				break
			}

			if len(current.Branch.Statements) != 1 {
				panic("should never happen")
			}

			subStmt := current.Branch.Statements[0]
			nestedBranch, ok := subStmt.(*ast.ConditionBranchStmt)
			if !ok {
				current.Branch.EndPos = subStmt.End()
				current.EndPos = subStmt.End()
				break
			}

			current.Branch.Statements = nil
			current = nestedBranch
			branches = append(branches, nestedBranch)
		}
	}

	return branches
}

func (reducer *Reducer) LabelledToSwitchExpr(
	labelDecl *lr.TokenValue,
	switchExpr *ast.SwitchExpr,
) (
	ast.Expression,
	error,
) {
	switchExpr.StartPos = labelDecl.Loc()
	switchExpr.PrependToLeading(labelDecl.TakeTrailing())
	switchExpr.PrependToLeading(labelDecl.TakeLeading())
	switchExpr.LabelDecl = labelDecl.Value
	return switchExpr, nil
}

func (reducer *Reducer) ToSwitchExprBody(
	switchKW *lr.TokenValue,
	operand ast.Expression,
	stmts *ast.StatementsExpr,
) (
	*ast.SwitchExpr,
	error,
) {
	branches := reducer.groupBranches(stmts)

	switchExpr := &ast.SwitchExpr{
		StartEndPos:       ast.NewStartEndPos(switchKW.Loc(), stmts.End()),
		Operand:           operand,
		ConditionBranches: branches,
	}

	switchExpr.LeadingComment = switchKW.TakeLeading()
	operand.PrependToLeading(switchKW.TakeTrailing())
	operand.AppendToTrailing(stmts.TakeLeading())
	switchExpr.TrailingComment = stmts.TakeTrailing()

	return switchExpr, nil
}

//
// SelectExpr
//

func (reducer *Reducer) LabelledToSelectExpr(
	labelDecl *lr.TokenValue,
	selectExpr *ast.SelectExpr,
) (
	ast.Expression,
	error,
) {
	selectExpr.StartPos = labelDecl.Loc()
	selectExpr.PrependToLeading(labelDecl.TakeTrailing())
	selectExpr.PrependToLeading(labelDecl.TakeLeading())
	selectExpr.LabelDecl = labelDecl.Value
	return selectExpr, nil
}

func (reducer *Reducer) ToSelectExprBody(
	selectKW *lr.TokenValue,
	stmts *ast.StatementsExpr,
) (
	*ast.SelectExpr,
	error,
) {
	branches := reducer.groupBranches(stmts)

	for _, branch := range branches {
		expr, ok := branch.Condition.(*ast.CasePatterns)
		if !ok {
			continue
		}

		if len(expr.Patterns) == 1 {
			branch.Condition = expr.Patterns[0]
		}
	}

	selectExpr := &ast.SelectExpr{
		StartEndPos:       ast.NewStartEndPos(selectKW.Loc(), stmts.End()),
		ConditionBranches: branches,
	}

	selectExpr.LeadingComment = selectKW.TakeLeading()
	selectExpr.LeadingComment.Append(selectKW.TakeTrailing())
	selectExpr.LeadingComment.Append(stmts.TakeLeading())
	selectExpr.TrailingComment = stmts.TakeTrailing()

	return selectExpr, nil
}

//
// LoopExpr
//

func (reducer *Reducer) LabelledToLoopExpr(
	labelDecl *lr.TokenValue,
	loop *ast.LoopExpr,
) (
	ast.Expression,
	error,
) {
	loop.StartPos = labelDecl.Loc()
	loop.PrependToLeading(labelDecl.TakeTrailing())
	loop.PrependToLeading(labelDecl.TakeLeading())
	loop.LabelDecl = labelDecl.Value
	return loop, nil
}

func (reducer *Reducer) InfiniteToLoopExprBody(
	body *ast.StatementsExpr,
) (
	*ast.LoopExpr,
	error,
) {
	leading := body.TakeLeading()
	trailing := body.TakeTrailing()

	loop := &ast.LoopExpr{
		StartEndPos: ast.NewStartEndPos(body.Loc(), body.End()),
		LoopKind:    ast.InfiniteLoop,
		Body:        body,
	}
	loop.LeadingComment = leading
	loop.TrailingComment = trailing

	return loop, nil
}

func (reducer *Reducer) DoWhileToLoopExprBody(
	body *ast.StatementsExpr,
	forKW *lr.TokenValue,
	condition ast.Expression,
) (
	*ast.LoopExpr,
	error,
) {
	leading := body.TakeLeading()
	body.AppendToTrailing(forKW.TakeLeading())
	condition.PrependToLeading(forKW.TakeTrailing())
	trailing := condition.TakeTrailing()

	loop := &ast.LoopExpr{
		StartEndPos: ast.NewStartEndPos(body.Loc(), body.End()),
		LoopKind:    ast.DoWhileLoop,
		Condition:   condition,
		Body:        body,
	}
	loop.LeadingComment = leading
	loop.TrailingComment = trailing

	return loop, nil
}

func (reducer *Reducer) WhileToLoopExprBody(
	forKW *lr.TokenValue,
	condition ast.Expression,
	body *ast.StatementsExpr,
) (
	*ast.LoopExpr,
	error,
) {
	leading := forKW.TakeLeading()
	condition.PrependToLeading(forKW.TakeTrailing())
	trailing := body.TakeTrailing()

	loop := &ast.LoopExpr{
		StartEndPos: ast.NewStartEndPos(body.Loc(), body.End()),
		LoopKind:    ast.WhileLoop,
		Condition:   condition,
		Body:        body,
	}
	loop.LeadingComment = leading
	loop.TrailingComment = trailing

	return loop, nil
}

func (reducer *Reducer) IteratorToLoopExprBody(
	forKW *lr.TokenValue,
	assignPattern ast.Expression,
	in *lr.TokenValue,
	iterator ast.Expression,
	body *ast.StatementsExpr,
) (
	*ast.LoopExpr,
	error,
) {
	leading := forKW.TakeLeading()
	trailing := body.TakeTrailing()

	loop := &ast.LoopExpr{
		StartEndPos: ast.NewStartEndPos(body.Loc(), body.End()),
		LoopKind:    ast.IteratorLoop,
		Condition:   reducer.toAssignPattern(assignPattern, in, iterator),
		Body:        body,
	}
	loop.LeadingComment = leading
	loop.TrailingComment = trailing

	return loop, nil
}

func (reducer *Reducer) ForToLoopExprBody(
	forKW *lr.TokenValue,
	init ast.Statement,
	semicolon1 *lr.TokenValue,
	condition ast.Expression,
	semicolon2 *lr.TokenValue,
	post ast.Statement,
	body *ast.StatementsExpr,
) (
	*ast.LoopExpr,
	error,
) {
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
		StartEndPos: ast.NewStartEndPos(body.Loc(), body.End()),
		LoopKind:    ast.ForLoop,
		Init:        init,
		Condition:   condition,
		Post:        post,
		Body:        body,
	}
	loop.LeadingComment = leading
	loop.TrailingComment = trailing

	return loop, nil
}

func (reducer *Reducer) NilToOptionalStatement() (
	ast.Statement,
	error,
) {
	return nil, nil
}

func (reducer *Reducer) NilToOptionalExpr() (
	ast.Expression,
	error,
) {
	return nil, nil
}

func (reducer *Reducer) toLoopBody(
	kw *lr.TokenValue,
	body *ast.StatementsExpr,
) *ast.StatementsExpr {
	body.StartPos = kw.Loc()
	body.PrependToLeading(kw.TakeTrailing())
	body.PrependToLeading(kw.TakeLeading())
	return body
}

func (reducer *Reducer) ToRepeatLoopBody(
	repeat *lr.TokenValue,
	expr *ast.StatementsExpr,
) (
	*ast.StatementsExpr,
	error,
) {
	return reducer.toLoopBody(repeat, expr), nil
}

func (reducer *Reducer) ToForLoopBody(
	do *lr.TokenValue,
	expr *ast.StatementsExpr,
) (
	*ast.StatementsExpr,
	error,
) {
	return reducer.toLoopBody(do, expr), nil
}
