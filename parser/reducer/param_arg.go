package reducer

import (
	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser/lr"
)

//
// Parameter
//

func (reducer *Reducer) namedTypedArg(
	kind ast.ParameterKind,
	name *lr.TokenValue,
	typeExpr ast.TypeExpression,
) (
	*ast.Parameter,
	error,
) {
	param := &ast.Parameter{
		StartEndPos: ast.NewStartEndPos(name.Loc(), typeExpr.End()),
		Kind:        kind,
		Name:        name.Value,
		HasEllipsis: false,
		Type:        typeExpr,
	}
	param.LeadingComment = name.TakeLeading()
	typeExpr.PrependToLeading(name.TakeTrailing())
	param.TrailingComment = typeExpr.TakeTrailing()
	return param, nil
}

func (reducer *Reducer) NamedTypedArgToProperParameterDef(
	name *lr.TokenValue,
	typeExpr ast.TypeExpression,
) (
	*ast.Parameter,
	error,
) {
	return reducer.namedTypedArg(ast.NamedTypedArgParameter, name, typeExpr)
}

func (reducer *Reducer) IgnoreTypedArgToProperParameterDef(
	underscore *lr.TokenValue,
	typeExpr ast.TypeExpression,
) (
	*ast.Parameter,
	error,
) {
	return reducer.namedTypedArg(
		ast.IgnoreTypedArgParameter,
		underscore,
		typeExpr)
}

func (reducer *Reducer) namedTypedVararg(
	kind ast.ParameterKind,
	name *lr.TokenValue,
	ellipsis *lr.TokenValue,
	typeExpr ast.TypeExpression,
) (
	*ast.Parameter,
	error,
) {
	param := &ast.Parameter{
		StartEndPos: ast.NewStartEndPos(name.Loc(), typeExpr.End()),
		Kind:        kind,
		Name:        name.Value,
		HasEllipsis: true,
		Type:        typeExpr,
	}
	param.LeadingComment = name.TakeLeading()

	typeExpr.PrependToLeading(ellipsis.TakeTrailing())
	typeExpr.PrependToLeading(ellipsis.TakeLeading())
	typeExpr.PrependToLeading(name.TakeTrailing())

	param.TrailingComment = typeExpr.TakeTrailing()
	return param, nil
}

func (reducer *Reducer) NamedTypedVarargToProperParameterDef(
	name *lr.TokenValue,
	ellipsis *lr.TokenValue,
	typeExpr ast.TypeExpression,
) (
	*ast.Parameter,
	error,
) {
	return reducer.namedTypedVararg(
		ast.NamedTypedVarargParameter,
		name,
		ellipsis,
		typeExpr)
}

func (reducer *Reducer) IgnoreTypedVarargToProperParameterDef(
	underscore *lr.TokenValue,
	ellipsis *lr.TokenValue,
	typeExpr ast.TypeExpression,
) (
	*ast.Parameter,
	error,
) {
	return reducer.namedTypedVararg(
		ast.IgnoreTypedVarargParameter,
		underscore,
		ellipsis,
		typeExpr)
}

func (reducer *Reducer) namedInferredVararg(
	kind ast.ParameterKind,
	name *lr.TokenValue,
	ellipsis *lr.TokenValue,
) (
	*ast.Parameter,
	error,
) {
	param := &ast.Parameter{
		StartEndPos: ast.NewStartEndPos(name.Loc(), ellipsis.End()),
		Kind:        kind,
		Name:        name.Value,
		HasEllipsis: true,
		Type:        nil,
	}
	param.LeadingComment = name.TakeLeading()
	trailing := name.TakeTrailing()
	trailing.Append(ellipsis.TakeLeading())
	trailing.Append(ellipsis.TakeTrailing())
	param.TrailingComment = trailing
	return param, nil
}

func (reducer *Reducer) NamedInferredVarargToProperParameterDef(
	name *lr.TokenValue,
	ellipsis *lr.TokenValue,
) (
	*ast.Parameter,
	error,
) {
	return reducer.namedInferredVararg(
		ast.NamedInferredVarargParameter,
		name,
		ellipsis)
}

func (reducer *Reducer) IgnoreInferredVarargToProperParameterDef(
	underscore *lr.TokenValue,
	ellipsis *lr.TokenValue,
) (
	*ast.Parameter,
	error,
) {
	return reducer.namedInferredVararg(
		ast.IgnoreInferredVarargParameter,
		underscore,
		ellipsis)
}

func (reducer *Reducer) UnnamedTypedArgToParameterDecl(
	typeExpr ast.TypeExpression,
) (
	*ast.Parameter,
	error,
) {
	param := &ast.Parameter{
		StartEndPos: ast.NewStartEndPos(typeExpr.Loc(), typeExpr.End()),
		Kind:        ast.UnnamedTypedArgParameter,
		Name:        "",
		HasEllipsis: false,
		Type:        typeExpr,
	}
	param.LeadingComment = typeExpr.TakeLeading()
	param.TrailingComment = typeExpr.TakeTrailing()
	return param, nil
}

func (reducer *Reducer) UnnamedInferredVarargToParameterDecl(
	ellipsis *lr.TokenValue,
) (
	*ast.Parameter,
	error,
) {
	param := &ast.Parameter{
		StartEndPos: ast.NewStartEndPos(ellipsis.Loc(), ellipsis.End()),
		Kind:        ast.UnnamedInferredVarargParameter,
		Name:        "",
		HasEllipsis: true,
		Type:        nil,
	}
	param.LeadingTrailingComments = ellipsis.LeadingTrailingComments
	return param, nil
}

func (reducer *Reducer) UnnamedTypedVarargToParameterDecl(
	ellipsis *lr.TokenValue,
	typeExpr ast.TypeExpression,
) (
	*ast.Parameter,
	error,
) {
	param := &ast.Parameter{
		StartEndPos: ast.NewStartEndPos(ellipsis.Loc(), typeExpr.End()),
		Kind:        ast.UnnamedTypedVarargParameter,
		Name:        "",
		HasEllipsis: true,
		Type:        typeExpr,
	}
	param.LeadingComment = ellipsis.TakeLeading()
	typeExpr.PrependToLeading(ellipsis.TakeTrailing())
	param.TrailingComment = typeExpr.TakeTrailing()
	return param, nil
}

func (reducer *Reducer) namedInferredArg(
	kind ast.ParameterKind,
	name *lr.TokenValue,
) (
	*ast.Parameter,
	error,
) {
	param := &ast.Parameter{
		StartEndPos: ast.NewStartEndPos(name.Loc(), name.End()),
		Kind:        kind,
		Name:        name.Value,
		HasEllipsis: false,
		Type:        nil,
	}
	param.LeadingComment = name.TakeLeading()
	param.TrailingComment = name.TakeTrailing()
	return param, nil
}

func (reducer *Reducer) NamedInferredArgToParameterDef(
	name *lr.TokenValue,
) (
	*ast.Parameter,
	error,
) {
	return reducer.namedInferredArg(ast.NamedInferredArgParameter, name)
}

func (reducer *Reducer) IgnoreInferredArgToParameterDef(
	underscore *lr.TokenValue,
) (
	*ast.Parameter,
	error,
) {
	return reducer.namedInferredArg(ast.IgnoreInferredArgParameter, underscore)
}

//
// ParameterList
//

func (reducer *Reducer) AddToProperParameterDeclList(
	list *ast.ParameterList,
	comma *lr.TokenValue,
	parameter *ast.Parameter,
) (
	*ast.ParameterList,
	error,
) {
	list.ReduceAdd(comma, parameter)
	return list, nil
}

func (reducer *Reducer) ParameterDeclToProperParameterDeclList(
	parameter *ast.Parameter,
) (
	*ast.ParameterList,
	error,
) {
	list := ast.NewParameterList()
	list.Add(parameter)
	return list, nil
}

func (reducer *Reducer) ImproperToParameterDeclList(
	list *ast.ParameterList,
	comma *lr.TokenValue,
) (
	*ast.ParameterList,
	error,
) {
	list.ReduceImproper(comma)
	return list, nil
}

func (reducer *Reducer) NilToParameterDeclList() (
	*ast.ParameterList,
	error,
) {
	return nil, nil
}

func (reducer *Reducer) ToParameterDecls(
	lparen *lr.TokenValue,
	list *ast.ParameterList,
	rparen *lr.TokenValue,
) (
	*ast.ParameterList,
	error,
) {
	if list == nil {
		list = ast.NewParameterList()
	}
	list.ReduceMarkers(lparen, rparen)
	return list, nil
}

func (reducer *Reducer) AddToProperParameterDefList(
	list *ast.ParameterList,
	comma *lr.TokenValue,
	parameter *ast.Parameter,
) (
	*ast.ParameterList,
	error,
) {
	list.ReduceAdd(comma, parameter)
	return list, nil
}

func (reducer *Reducer) ParameterDefToProperParameterDefList(
	parameter *ast.Parameter,
) (
	*ast.ParameterList,
	error,
) {
	list := ast.NewParameterList()
	list.Add(parameter)
	return list, nil
}

func (reducer *Reducer) ImproperToParameterDefList(
	list *ast.ParameterList,
	comma *lr.TokenValue,
) (
	*ast.ParameterList,
	error,
) {
	list.ReduceImproper(comma)
	return list, nil
}

func (reducer *Reducer) NilToParameterDefList() (
	*ast.ParameterList,
	error,
) {
	return nil, nil
}

func (reducer *Reducer) ToParameterDefs(
	lparen *lr.TokenValue,
	list *ast.ParameterList,
	rparen *lr.TokenValue,
) (
	*ast.ParameterList,
	error,
) {
	if list == nil {
		list = ast.NewParameterList()
	}
	list.ReduceMarkers(lparen, rparen)
	return list, nil
}

//
// Argument
//

func (Reducer) PositionalToArgument(
	expr ast.Expression,
) (
	*ast.Argument,
	error,
) {
	return ast.NewPositionalArgument(expr), nil
}

func (Reducer) ColonExprToArgument(
	expr *ast.ImplicitStructExpr,
) (
	*ast.Argument,
	error,
) {
	return ast.NewPositionalArgument(expr), nil
}

func (Reducer) NamedAssignmentToArgument(
	name *lr.TokenValue,
	assign *lr.TokenValue,
	expr ast.Expression,
) (
	*ast.Argument,
	error,
) {
	return ast.NewNamedArgument(name, assign, expr), nil
}

func (Reducer) VarargAssignmentToArgument(
	expr ast.Expression,
	ellipsis *lr.TokenValue,
) (
	*ast.Argument,
	error,
) {
	arg := &ast.Argument{
		StartEndPos: ast.NewStartEndPos(expr.Loc(), ellipsis.End()),
		Kind:        ast.VarargAssignmentArgument,
		Expr:        expr,
		HasEllipsis: true,
	}

	arg.LeadingComment = expr.TakeLeading()
	expr.AppendToTrailing(ellipsis.TakeLeading())
	arg.TrailingComment = ellipsis.TakeTrailing()
	return arg, nil
}

func (Reducer) SkipPatternToArgument(
	ellipsis *lr.TokenValue,
) (
	*ast.Argument,
	error,
) {
	return ast.NewSkipPatternArgument(ellipsis), nil
}

func (Reducer) PositionalToFieldVarPattern(
	expr ast.Expression,
) (
	*ast.Argument,
	error,
) {
	return ast.NewPositionalArgument(expr), nil
}

func (Reducer) NamedAssignmentToFieldVarPattern(
	name *lr.TokenValue,
	assign *lr.TokenValue,
	expr ast.Expression,
) (
	*ast.Argument,
	error,
) {
	return ast.NewNamedArgument(name, assign, expr), nil
}

func (Reducer) SkipPatternToFieldVarPattern(
	ellipsis *lr.TokenValue,
) (
	*ast.Argument,
	error,
) {
	return ast.NewSkipPatternArgument(ellipsis), nil
}

//
// Argument list
//

func (reducer *Reducer) AddToProperArguments(
	list *ast.ArgumentList,
	comma *lr.TokenValue,
	arg *ast.Argument,
) (
	*ast.ArgumentList,
	error,
) {
	list.ReduceAdd(comma, arg)
	return list, nil
}

func (reducer *Reducer) ArgumentToProperArguments(
	arg *ast.Argument,
) (
	*ast.ArgumentList,
	error,
) {
	list := ast.NewArgumentList()
	list.Add(arg)
	return list, nil
}

func (reducer *Reducer) ImproperToArguments(
	list *ast.ArgumentList,
	comma *lr.TokenValue,
) (
	*ast.ArgumentList,
	error,
) {
	list.ReduceImproper(comma)
	return list, nil
}

func (reducer *Reducer) NilToArguments() (
	*ast.ArgumentList,
	error,
) {
	return nil, nil
}

func (reducer *Reducer) FieldVarPatternToProperFieldVarPatterns(
	pattern *ast.Argument,
) (
	*ast.ArgumentList,
	error,
) {
	list := ast.NewArgumentList()
	list.Add(pattern)
	return list, nil
}

func (reducer *Reducer) AddToProperFieldVarPatterns(
	list *ast.ArgumentList,
	comma *lr.TokenValue,
	pattern *ast.Argument,
) (
	*ast.ArgumentList,
	error,
) {
	list.ReduceAdd(comma, pattern)
	return list, nil
}

func (reducer *Reducer) ImproperToFieldVarPatterns(
	list *ast.ArgumentList,
	comma *lr.TokenValue,
) (
	*ast.ArgumentList,
	error,
) {
	list.ReduceImproper(comma)
	return list, nil
}
