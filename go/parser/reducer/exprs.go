package reducer

import (
	"github.com/pattyshack/gt/parseutil"

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
		StartEndPos: parseutil.NewStartEndPos(operand.Loc(), field.End()),
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
		StartEndPos: parseutil.NewStartEndPos(operand.Loc(), op.End()),
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
		StartEndPos: parseutil.NewStartEndPos(op.Loc(), operand.End()),
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
		StartEndPos: parseutil.NewStartEndPos(left.Loc(), right.End()),
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
		StartEndPos: parseutil.NewStartEndPos(expr1.Loc(), expr2.End()),
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
			ast.NewPositionalArgument(ast.NewImproperUnit(colon.Loc())),
		},
		colon,
		ast.NewPositionalArgument(ast.NewImproperUnit(colon.End())))

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
		ast.NewPositionalArgument(ast.NewImproperUnit(colon.End())))

	return &ast.ImplicitStructExpr{
		StartEndPos: parseutil.NewStartEndPos(leftExpr.Loc(), colon.End()),
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
			ast.NewPositionalArgument(ast.NewImproperUnit(colon.Loc())),
		},
		colon,
		ast.NewPositionalArgument(rightExpr))

	return &ast.ImplicitStructExpr{
		StartEndPos: parseutil.NewStartEndPos(colon.Loc(), rightExpr.End()),
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
		StartEndPos: parseutil.NewStartEndPos(leftExpr.Loc(), rightExpr.End()),
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
		ast.NewPositionalArgument(ast.NewImproperUnit(colon.End())))
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
// ParameterizedExpr
//

func (reducer *Reducer) ToParameterizedExpr(
	operand ast.Expression,
	parameters *ast.TypeExpressionList,
) (
	ast.Expression,
	error,
) {
	leading := operand.TakeLeading()

	pkg := ""
	name := ""
	invalid := false
	switch expr := operand.(type) {
	case *ast.AccessExpr:
		name = expr.Field

		pkgExpr, ok := expr.Operand.(*ast.NamedExpr)
		if !ok {
			invalid = true
		} else {
			leading.Append(pkgExpr.TakeLeading())
			leading.Append(pkgExpr.TakeTrailing())
			pkg = pkgExpr.Name
		}
	case *ast.NamedExpr:
		name = expr.Name
	default:
		invalid = true
	}

	leading.Append(operand.TakeTrailing())
	leading.Append(parameters.TakeLeading())
	leading.Append(parameters.MiddleComment)

	if invalid {
		err := reducer.Emit(
			parameters.Loc(),
			"unexpected parameterization, can only parameterize names of the form "+
				"<name> or <pkg>.<name>")

		expr := &ast.ParseErrorExpr{
			StartEndPos: parseutil.NewStartEndPos(operand.Loc(), parameters.End()),
			Error:       err,
		}
		expr.LeadingComment = leading
		expr.TrailingComment = parameters.TakeTrailing()

		return expr, nil
	}

	expr := &ast.ParameterizedExpr{
		StartEndPos: parseutil.NewStartEndPos(operand.Loc(), parameters.End()),
		Pkg:         pkg,
		Name:        name,
		Parameters:  parameters.Elements,
	}
	expr.LeadingComment = leading
	expr.TrailingComment = parameters.TakeTrailing()

	return expr, nil
}

//
// CallExpr
//

func (reducer *Reducer) ToCallExpr(
	funcExpr ast.Expression,
	lparen *lr.TokenValue,
	arguments *ast.ArgumentList,
	rparen *lr.TokenValue,
) (
	ast.Expression,
	error,
) {
	leading := funcExpr.TakeLeading()

	if arguments == nil {
		arguments = ast.NewArgumentList()
	}
	arguments.ReduceMarkers(lparen, rparen)

	funcExpr.AppendToTrailing(arguments.TakeLeading())
	funcExpr.AppendToTrailing(arguments.MiddleComment)
	trailing := arguments.TakeTrailing()

	expr := &ast.CallExpr{
		StartEndPos: parseutil.NewStartEndPos(funcExpr.Loc(), rparen.End()),
		FuncExpr:    funcExpr,
		Arguments:   arguments.Elements,
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
			if arg.Kind != ast.SingularArgument || arg.Name != "" {
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
		StartEndPos: parseutil.NewStartEndPos(accessible.Loc(), rbracket.End()),
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
		StartEndPos: parseutil.NewStartEndPos(accessible.Loc(), rparen.End()),
		Accessible:  accessible,
		CastType:    castType,
	}
	expr.LeadingComment = leading
	expr.TrailingComment = trailing

	return expr, nil
}

//
// MakeExpr
//

func (Reducer) SizeProperToMakeExpr(
	expr *ast.MakeExpr,
	rparen *lr.TokenValue,
) (
	*ast.MakeExpr,
	error,
) {
	end := expr.Size
	if expr.Capacity != nil {
		end = expr.Capacity
	}
	end.AppendToTrailing(rparen.TakeLeading())
	expr.TrailingComment = rparen.TakeTrailing()
	expr.EndPos = rparen.End()
	return expr, nil
}

func (Reducer) SizeImproperToMakeExpr(
	expr *ast.MakeExpr,
	comma *lr.TokenValue,
	rparen *lr.TokenValue,
) (
	*ast.MakeExpr,
	error,
) {
	end := expr.Size
	if expr.Capacity != nil {
		end = expr.Capacity
	}
	end.AppendToTrailing(comma.TakeLeading())
	end.AppendToTrailing(comma.TakeTrailing())
	end.AppendToTrailing(rparen.TakeLeading())
	expr.TrailingComment = rparen.TakeTrailing()
	expr.EndPos = rparen.End()
	return expr, nil
}

func (Reducer) ValueProperToMakeExpr(
	expr *ast.MakeExpr,
	comma *lr.TokenValue,
	value ast.Expression,
	rparen *lr.TokenValue,
) (
	*ast.MakeExpr,
	error,
) {
	expr.Value = value

	end := expr.Size
	if expr.Capacity != nil {
		end = expr.Capacity
	}
	end.AppendToTrailing(comma.TakeLeading())
	end.AppendToTrailing(comma.TakeTrailing())

	value.AppendToTrailing(rparen.TakeLeading())
	expr.TrailingComment = rparen.TakeTrailing()
	expr.EndPos = rparen.End()
	return expr, nil
}

func (Reducer) ValueImproperToMakeExpr(
	expr *ast.MakeExpr,
	comma1 *lr.TokenValue,
	value ast.Expression,
	comma2 *lr.TokenValue,
	rparen *lr.TokenValue,
) (
	*ast.MakeExpr,
	error,
) {
	expr.Value = value

	end := expr.Size
	if expr.Capacity != nil {
		end = expr.Capacity
	}
	end.AppendToTrailing(comma1.TakeLeading())
	end.AppendToTrailing(comma1.TakeTrailing())

	value.AppendToTrailing(comma2.TakeLeading())
	value.AppendToTrailing(comma2.TakeTrailing())
	value.AppendToTrailing(rparen.TakeLeading())
	expr.TrailingComment = rparen.TakeTrailing()
	expr.EndPos = rparen.End()
	return expr, nil
}

func (Reducer) SizeToMakeExprSize(
	expr *ast.MakeExpr,
	size ast.Expression,
) (
	*ast.MakeExpr,
	error,
) {
	expr.Size = size
	return expr, nil
}

func (Reducer) SizeCapacityToMakeExprSize(
	expr *ast.MakeExpr,
	size ast.Expression,
	colon *lr.TokenValue,
	capacity ast.Expression,
) (
	*ast.MakeExpr,
	error,
) {
	size.AppendToTrailing(colon.TakeLeading())
	capacity.PrependToLeading(colon.TakeTrailing())
	expr.Size = size
	expr.Capacity = capacity
	return expr, nil
}

func (Reducer) ToMakeExprHead(
	makeKW *lr.TokenValue,
	lparen *lr.TokenValue,
	typeExpr ast.TypeExpression,
	comma *lr.TokenValue,
) (
	*ast.MakeExpr,
	error,
) {
	leading := makeKW.TakeLeading()
	leading.Append(makeKW.TakeTrailing())
	leading.Append(lparen.TakeLeading())
	typeExpr.PrependToLeading(lparen.TakeTrailing())
	typeExpr.AppendToTrailing(comma.TakeLeading())
	typeExpr.AppendToTrailing(comma.TakeTrailing())

	expr := &ast.MakeExpr{
		VariableSizedType: typeExpr,
	}
	expr.StartPos = makeKW.Loc()
	expr.LeadingComment = leading

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
		StartEndPos:   parseutil.NewStartEndPos(initializable.Loc(), rparen.End()),
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
		StartEndPos:     parseutil.NewStartEndPos(elseKW.Loc(), branch.End()),
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
		StartEndPos: parseutil.NewStartEndPos(elseKW.Loc(), branch.End()),
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
		StartEndPos: parseutil.NewStartEndPos(ifKW.Loc(), branch.End()),
		Condition:   condition,
		Branch:      branch,
	}
	cb.LeadingComment = ifKW.TakeTrailing()
	cb.LeadingComment.Append(condition.TakeLeading())
	cb.TrailingComment = branch.TakeTrailing()

	expr := &ast.IfExpr{
		StartEndPos:       parseutil.NewStartEndPos(ifKW.Loc(), branch.End()),
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
					stmt.Loc(),
					"unexpected statement. expected case or default branch statement")
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

func (reducer *Reducer) ToSwitchExprBody(
	switchKW *lr.TokenValue,
	operand ast.Expression,
	stmts *ast.StatementsExpr,
) (
	ast.ControlFlowExpr,
	error,
) {
	branches := reducer.groupBranches(stmts)

	switchExpr := &ast.SwitchExpr{
		StartEndPos:       parseutil.NewStartEndPos(switchKW.Loc(), stmts.End()),
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

func (reducer *Reducer) ToSelectExprBody(
	selectKW *lr.TokenValue,
	stmts *ast.StatementsExpr,
) (
	ast.ControlFlowExpr,
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
		StartEndPos:       parseutil.NewStartEndPos(selectKW.Loc(), stmts.End()),
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

func (reducer *Reducer) InfiniteToLoopExprBody(
	body *ast.StatementsExpr,
) (
	ast.ControlFlowExpr,
	error,
) {
	leading := body.TakeLeading()
	trailing := body.TakeTrailing()

	loop := &ast.LoopExpr{
		StartEndPos: parseutil.NewStartEndPos(body.Loc(), body.End()),
		Kind:        ast.InfiniteLoop,
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
	ast.ControlFlowExpr,
	error,
) {
	leading := body.TakeLeading()
	body.AppendToTrailing(forKW.TakeLeading())
	condition.PrependToLeading(forKW.TakeTrailing())
	trailing := condition.TakeTrailing()

	loop := &ast.LoopExpr{
		StartEndPos: parseutil.NewStartEndPos(body.Loc(), body.End()),
		Kind:        ast.DoWhileLoop,
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
	ast.ControlFlowExpr,
	error,
) {
	leading := forKW.TakeLeading()
	condition.PrependToLeading(forKW.TakeTrailing())
	trailing := body.TakeTrailing()

	loop := &ast.LoopExpr{
		StartEndPos: parseutil.NewStartEndPos(body.Loc(), body.End()),
		Kind:        ast.WhileLoop,
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
	ast.ControlFlowExpr,
	error,
) {
	leading := forKW.TakeLeading()
	trailing := body.TakeTrailing()

	loop := &ast.LoopExpr{
		StartEndPos: parseutil.NewStartEndPos(body.Loc(), body.End()),
		Kind:        ast.IteratorLoop,
		Condition:   reducer.toAssignPattern(assignPattern, in, iterator),
		Body:        body,
	}
	loop.LeadingComment = leading
	loop.TrailingComment = trailing

	return loop, nil
}

func (reducer *Reducer) ForToLoopExprBody(
	forKW *lr.TokenValue,
	condition ast.Expression,
	semicolon *lr.TokenValue,
	post ast.Expression,
	body *ast.StatementsExpr,
) (
	ast.ControlFlowExpr,
	error,
) {
	leading := forKW.TakeLeading()

	condition.PrependToLeading(forKW.TakeTrailing())
	condition.AppendToTrailing(semicolon.TakeLeading())
	post.PrependToLeading(semicolon.TakeTrailing())

	trailing := body.TakeTrailing()

	loop := &ast.LoopExpr{
		StartEndPos: parseutil.NewStartEndPos(body.Loc(), body.End()),
		Kind:        ast.ForLoop,
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

func (reducer *Reducer) ToLoopBody(
	do *lr.TokenValue,
	body *ast.StatementsExpr,
) (
	*ast.StatementsExpr,
	error,
) {
	body.StartPos = do.Loc()
	body.PrependToLeading(do.TakeTrailing())
	body.PrependToLeading(do.TakeLeading())
	return body, nil
}

//
// ControlFlowExpr
//

func (Reducer) LabelledToControlFlowExpr(
	label *lr.TokenValue,
	at *lr.TokenValue,
	expr ast.ControlFlowExpr,
) (
	ast.ControlFlowExpr,
	error,
) {
	expr.SetLabel(label.Value)
	expr.PrependToLeading(at.TakeTrailing())
	expr.PrependToLeading(at.TakeLeading())
	expr.PrependToLeading(label.TakeTrailing())
	expr.PrependToLeading(label.TakeLeading())
	return expr, nil
}
