package reducer

import (
	"github.com/pattyshack/gt/parseutil"

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
		StartEndPos:             name.StartEnd(),
		LeadingTrailingComments: name.TakeComments(),
		Name:                    name.Value,
		Constraint:              ast.NewAnyTrait(name.End()),
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
		StartEndPos: parseutil.NewStartEndPos(name.Loc(), constraint.End()),
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
	dollar *lr.TokenValue,
	lbracket *lr.TokenValue,
	list *ast.GenericParameterList,
	rbracket *lr.TokenValue,
) (
	*ast.GenericParameterList,
	error,
) {
	list.ReduceMarkers(lbracket, rbracket)
	list.PrependToLeading(dollar.TakeTrailing())
	list.PrependToLeading(dollar.TakeLeading())
	list.StartPos = dollar.Loc()
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

func (Reducer) ImproperImplicitToGenericParameterList(
	list *ast.GenericParameterList,
	newlines *lr.TokenCount,
) (
	*ast.GenericParameterList,
	error,
) {
	return list, nil
}

func (Reducer) ImproperExplicitToGenericParameterList(
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
	return nil, nil
}

//
// GenericArgumentList
//

func (reducer *Reducer) ToGenericArguments(
	dollar *lr.TokenValue,
	lbracket *lr.TokenValue,
	list *ast.TypeExpressionList,
	rbracket *lr.TokenValue,
) (
	*ast.TypeExpressionList,
	error,
) {
	if list == nil {
		list = ast.NewTypeExpressionList()
	}
	list.ReduceMarkers(lbracket, rbracket)
	list.PrependToLeading(dollar.TakeTrailing())
	list.PrependToLeading(dollar.TakeLeading())
	list.StartPos = dollar.Loc()
	return list, nil
}

func (reducer *Reducer) NilToOptionalGenericArguments() (
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

func (reducer *Reducer) ImproperImplicitToGenericArgumentList(
	list *ast.TypeExpressionList,
	newlines *lr.TokenCount,
) (
	*ast.TypeExpressionList,
	error,
) {
	return list, nil
}

func (reducer *Reducer) ImproperExplicitToGenericArgumentList(
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
	return nil, nil
}
