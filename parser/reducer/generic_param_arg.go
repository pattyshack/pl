package reducer

import (
	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser/lr"
)

//
// GenericParameter
//

func (Reducer) UnconstrainedToGenericParameter(
	name *lr.TokenValue,
) (
	*ast.GenericParameter,
	error,
) {
	return &ast.GenericParameter{
		StartEndPos:             ast.NewStartEndPos(name.Loc(), name.End()),
		LeadingTrailingComments: name.LeadingTrailingComments,
		Name:                    name.Value,
	}, nil
}

func (Reducer) ConstrainedToGenericParameter(
	name *lr.TokenValue,
	constraint ast.TypeExpression,
) (
	*ast.GenericParameter,
	error,
) {
	param := &ast.GenericParameter{
		StartEndPos: ast.NewStartEndPos(name.Loc(), constraint.End()),
		Name:        name.Value,
		Constraint:  constraint,
	}

	param.LeadingComment = name.TakeLeading()
	constraint.PrependToLeading(name.TakeTrailing())
	param.TrailingComment = constraint.TakeTrailing()
	return param, nil
}

//
// GenericParameterList
//

func (Reducer) GenericToGenericParameters(
	dollarLbracket *lr.TokenValue,
	list *ast.GenericParameterList,
	rbracket *lr.TokenValue,
) (
	*ast.GenericParameterList,
	error,
) {
	list.ReduceMarkers(dollarLbracket, rbracket)
	return list, nil
}

func (Reducer) NilToGenericParameters() (
	*ast.GenericParameterList,
	error,
) {
	return nil, nil
}

func (Reducer) AddToProperGenericParameterList(
	list *ast.GenericParameterList,
	comma *lr.TokenValue,
	parameter *ast.GenericParameter,
) (
	*ast.GenericParameterList,
	error,
) {
	list.ReduceAdd(comma, parameter)
	return list, nil
}

func (Reducer) GenericParameterToProperGenericParameterList(
	parameter *ast.GenericParameter,
) (
	*ast.GenericParameterList,
	error,
) {
	list := ast.NewGenericParameterList()
	list.Add(parameter)
	return list, nil
}

func (Reducer) ImproperToGenericParameterList(
	list *ast.GenericParameterList,
	comma *lr.TokenValue,
) (
	*ast.GenericParameterList,
	error,
) {
	list.ReduceImproper(comma)
	return list, nil
}

func (Reducer) NilToGenericParameterList() (
	*ast.GenericParameterList,
	error,
) {
	return ast.NewGenericParameterList(), nil
}

//
// GenericArgumentList
//

func (reducer *Reducer) BindingToGenericArguments(
	dollarLbracket *lr.TokenValue,
	list *ast.TypeExpressionList,
	rbracket *lr.TokenValue,
) (
	*ast.TypeExpressionList,
	error,
) {
	list.ReduceMarkers(dollarLbracket, rbracket)
	return list, nil
}

func (reducer *Reducer) NilToGenericArguments() (
	*ast.TypeExpressionList,
	error,
) {
	return nil, nil
}

func (reducer *Reducer) AddToProperGenericArgumentList(
	list *ast.TypeExpressionList,
	comma *lr.TokenValue,
	arg ast.TypeExpression,
) (
	*ast.TypeExpressionList,
	error,
) {
	list.ReduceAdd(comma, arg)
	return list, nil
}

func (reducer *Reducer) TypeExprToProperGenericArgumentList(
	arg ast.TypeExpression,
) (
	*ast.TypeExpressionList,
	error,
) {
	list := ast.NewTypeExpressionList()
	list.Add(arg)
	return list, nil
}

func (reducer *Reducer) ImproperToGenericArgumentList(
	list *ast.TypeExpressionList,
	comma *lr.TokenValue,
) (
	*ast.TypeExpressionList,
	error,
) {
	list.ReduceImproper(comma)
	return list, nil
}

func (reducer *Reducer) NilToGenericArgumentList() (
	*ast.TypeExpressionList,
	error,
) {
	return ast.NewTypeExpressionList(), nil
}
