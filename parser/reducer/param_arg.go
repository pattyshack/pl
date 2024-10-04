package reducer

import (
	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser/lr"
)

//
// Parameter
//

func (reducer *Reducer) NamedArgToParameter(
	name *lr.TokenValue,
	typeExpr ast.TypeExpression,
) (
	*ast.Parameter,
	error,
) {
	param := &ast.Parameter{
		StartEndPos: ast.NewStartEndPos(name.Loc(), typeExpr.End()),
		Kind:        ast.ArgParameter,
		Name:        name.Value,
		Type:        typeExpr,
	}
	param.LeadingComment = name.TakeLeading()
	param.LeadingComment.Append(name.TakeTrailing())
	param.TrailingComment = typeExpr.TakeLeading()
	return param, nil
}

func (reducer *Reducer) NamedReceiverToParameter(
	name *lr.TokenValue,
	greater *lr.TokenValue,
	typeExpr ast.TypeExpression,
) (
	*ast.Parameter,
	error,
) {
	param := &ast.Parameter{
		StartEndPos: ast.NewStartEndPos(name.Loc(), typeExpr.End()),
		Kind:        ast.ReceiverParameter,
		Name:        name.Value,
		Type:        typeExpr,
	}
	leading := name.TakeLeading()
	leading.Append(name.TakeTrailing())
	leading.Append(greater.TakeLeading())

	param.LeadingComment = leading
	typeExpr.PrependToLeading(greater.TakeTrailing())
	param.TrailingComment = typeExpr.TakeLeading()
	return param, nil
}

func (reducer *Reducer) NamedVarargToParameter(
	name *lr.TokenValue,
	ellipsis *lr.TokenValue,
	typeExpr ast.TypeExpression,
) (
	*ast.Parameter,
	error,
) {
	param := &ast.Parameter{
		StartEndPos: ast.NewStartEndPos(name.Loc(), typeExpr.End()),
		Kind:        ast.VarargParameter,
		Name:        name.Value,
		Type:        typeExpr,
	}
	leading := name.TakeLeading()
	leading.Append(name.TakeTrailing())
	leading.Append(ellipsis.TakeLeading())

	param.LeadingComment = leading
	typeExpr.PrependToLeading(ellipsis.TakeTrailing())
	param.TrailingComment = typeExpr.TakeLeading()
	return param, nil
}

func (reducer *Reducer) IgnoreArgToParameter(
	underscore *lr.TokenValue,
	typeExpr ast.TypeExpression,
) (
	*ast.Parameter,
	error,
) {
	param := &ast.Parameter{
		StartEndPos: ast.NewStartEndPos(underscore.Loc(), typeExpr.End()),
		Kind:        ast.ArgParameter,
		Type:        typeExpr,
	}
	param.LeadingComment = underscore.TakeLeading()
	param.LeadingComment.Append(underscore.TakeTrailing())
	param.TrailingComment = typeExpr.TakeLeading()
	return param, nil
}

func (reducer *Reducer) IgnoreReceiverToParameter(
	underscore *lr.TokenValue,
	greater *lr.TokenValue,
	typeExpr ast.TypeExpression,
) (
	*ast.Parameter,
	error,
) {
	param := &ast.Parameter{
		StartEndPos: ast.NewStartEndPos(underscore.Loc(), typeExpr.End()),
		Kind:        ast.ReceiverParameter,
		Type:        typeExpr,
	}
	leading := underscore.TakeLeading()
	leading.Append(underscore.TakeTrailing())
	leading.Append(greater.TakeLeading())

	param.LeadingComment = leading
	typeExpr.PrependToLeading(greater.TakeTrailing())
	param.TrailingComment = typeExpr.TakeLeading()
	return param, nil
}

func (reducer *Reducer) IgnoreVarargToParameter(
	underscore *lr.TokenValue,
	ellipsis *lr.TokenValue,
	typeExpr ast.TypeExpression,
) (
	*ast.Parameter,
	error,
) {
	param := &ast.Parameter{
		StartEndPos: ast.NewStartEndPos(underscore.Loc(), typeExpr.End()),
		Kind:        ast.VarargParameter,
		Type:        typeExpr,
	}
	leading := underscore.TakeLeading()
	leading.Append(underscore.TakeTrailing())
	leading.Append(ellipsis.TakeLeading())

	param.LeadingComment = leading
	typeExpr.PrependToLeading(ellipsis.TakeTrailing())
	param.TrailingComment = typeExpr.TakeLeading()
	return param, nil
}

func (reducer *Reducer) UnnamedArgToParameter(
	typeExpr ast.TypeExpression,
) (
	*ast.Parameter,
	error,
) {
	param := &ast.Parameter{
		StartEndPos:             typeExpr.StartEnd(),
		LeadingTrailingComments: typeExpr.TakeComments(),
		Kind:                    ast.ArgParameter,
		Type:                    typeExpr,
	}
	return param, nil
}

func (reducer *Reducer) UnnamedReceiverToParameter(
	greater *lr.TokenValue,
	typeExpr ast.TypeExpression,
) (
	*ast.Parameter,
	error,
) {
	param := &ast.Parameter{
		StartEndPos: ast.NewStartEndPos(greater.Loc(), typeExpr.End()),
		Kind:        ast.ReceiverParameter,
		Type:        typeExpr,
	}
	param.LeadingComment = greater.TakeLeading()
	typeExpr.PrependToLeading(greater.TakeTrailing())
	param.TrailingComment = typeExpr.TakeLeading()
	return param, nil
}

func (reducer *Reducer) UnnamedVarargToParameter(
	ellipsis *lr.TokenValue,
	typeExpr ast.TypeExpression,
) (
	*ast.Parameter,
	error,
) {
	param := &ast.Parameter{
		StartEndPos: ast.NewStartEndPos(ellipsis.Loc(), typeExpr.End()),
		Kind:        ast.VarargParameter,
		Type:        typeExpr,
	}
	param.LeadingComment = ellipsis.TakeLeading()
	typeExpr.PrependToLeading(ellipsis.TakeTrailing())
	param.TrailingComment = typeExpr.TakeLeading()
	return param, nil
}

//
// ParameterList
//

func (reducer *Reducer) AddToProperParameterList(
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

func (reducer *Reducer) ParameterToProperParameterList(
	parameter *ast.Parameter,
) (
	*ast.ParameterList,
	error,
) {
	list := ast.NewParameterList()
	list.Add(parameter)
	return list, nil
}

func (reducer *Reducer) ImproperToParameterList(
	list *ast.ParameterList,
	comma *lr.TokenValue,
) (
	*ast.ParameterList,
	error,
) {
	list.ReduceImproper(comma)
	return list, nil
}

func (reducer *Reducer) NilToParameterList() (
	*ast.ParameterList,
	error,
) {
	return nil, nil
}

func (reducer *Reducer) ToParameters(
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
	return ast.NewVariadicArgument(expr, ellipsis), nil
}

func (Reducer) SkipPatternToArgument(
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
