package reducer

import (
	. "github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser"
)

//
// GenericParameter
//

func (Reducer) UnconstrainedToGenericParameter(
	name *parser.TokenValue,
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
	name *parser.TokenValue,
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
	dollarLbracket *parser.TokenValue,
	list *GenericParameterList,
	rbracket *parser.TokenValue,
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
	comma *parser.TokenValue,
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
	comma *parser.TokenValue,
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
	dollarLbracket *parser.TokenValue,
	list *GenericArgumentList,
	rbracket *parser.TokenValue,
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
	comma *parser.TokenValue,
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
	comma *parser.TokenValue,
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
