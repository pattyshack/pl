package reducer

import (
	. "github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser/lr"
)

//
// GenericParameter
//

func (Reducer) UnconstrainedToGenericParameter(
	name *lr.TokenValue,
) (
	*GenericParameter,
	error,
) {
	return &GenericParameter{
		StartEndPos:             NewStartEndPos(name.Loc(), name.End()),
		LeadingTrailingComments: name.LeadingTrailingComments,
		Name:                    name.Value,
	}, nil
}

func (Reducer) ConstrainedToGenericParameter(
	name *lr.TokenValue,
	constraint TypeExpression,
) (
	*GenericParameter,
	error,
) {
	param := &GenericParameter{
		StartEndPos: NewStartEndPos(name.Loc(), constraint.End()),
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
	list *GenericParameterList,
	rbracket *lr.TokenValue,
) (
	*GenericParameterList,
	error,
) {
	list.ReduceMarkers(dollarLbracket, rbracket)
	return list, nil
}

func (Reducer) NilToGenericParameters() (
	*GenericParameterList,
	error,
) {
	return nil, nil
}

func (Reducer) AddToProperGenericParameterList(
	list *GenericParameterList,
	comma *lr.TokenValue,
	parameter *GenericParameter,
) (
	*GenericParameterList,
	error,
) {
	list.ReduceAdd(comma, parameter)
	return list, nil
}

func (Reducer) GenericParameterToProperGenericParameterList(
	parameter *GenericParameter,
) (
	*GenericParameterList,
	error,
) {
	list := NewGenericParameterList()
	list.Add(parameter)
	return list, nil
}

func (Reducer) ImproperToGenericParameterList(
	list *GenericParameterList,
	comma *lr.TokenValue,
) (
	*GenericParameterList,
	error,
) {
	list.ReduceImproper(comma)
	return list, nil
}

func (Reducer) NilToGenericParameterList() (
	*GenericParameterList,
	error,
) {
	return NewGenericParameterList(), nil
}

//
// GenericArgumentList
//

func (reducer *Reducer) BindingToGenericArguments(
	dollarLbracket *lr.TokenValue,
	list *GenericArgumentList,
	rbracket *lr.TokenValue,
) (
	*GenericArgumentList,
	error,
) {
	list.ReduceMarkers(dollarLbracket, rbracket)
	return list, nil
}

func (reducer *Reducer) NilToGenericArguments() (
	*GenericArgumentList,
	error,
) {
	return nil, nil
}

func (reducer *Reducer) AddToProperGenericArgumentList(
	list *GenericArgumentList,
	comma *lr.TokenValue,
	arg TypeExpression,
) (
	*GenericArgumentList,
	error,
) {
	list.ReduceAdd(comma, arg)
	return list, nil
}

func (reducer *Reducer) TypeExprToProperGenericArgumentList(
	arg TypeExpression,
) (
	*GenericArgumentList,
	error,
) {
	list := NewGenericArgumentList()
	list.Add(arg)
	return list, nil
}

func (reducer *Reducer) ImproperToGenericArgumentList(
	list *GenericArgumentList,
	comma *lr.TokenValue,
) (
	*GenericArgumentList,
	error,
) {
	list.ReduceImproper(comma)
	return list, nil
}

func (reducer *Reducer) NilToGenericArgumentList() (
	*GenericArgumentList,
	error,
) {
	return &GenericArgumentList{}, nil
}
