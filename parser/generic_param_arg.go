package parser

import (
	"fmt"
)

//
// GenericParameter
//

type GenericParameter struct {
	StartEndPos
	LeadingTrailingComments

	Name string

	Constraint TypeExpression // optional
}

var _ Node = &GenericParameter{}

func (param GenericParameter) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[GenericParameter: Name=%s",
		indent,
		label,
		param.Name)

	if param.Constraint == nil {
		result += "]"
	} else {
		result += "\n" + param.Constraint.TreeString(indent+"  ", "Constraint=")
		result += "\n" + indent + "]"
	}

	return result
}

type GenericParameterReducerImpl struct{}

var _ GenericParameterReducer = &GenericParameterReducerImpl{}

func (GenericParameterReducerImpl) UnconstrainedToGenericParameter(
	name TokenValue,
) (
	*GenericParameter,
	error,
) {
	return &GenericParameter{
		StartEndPos:             newStartEndPos(name.Loc(), name.End()),
		LeadingTrailingComments: name.LeadingTrailingComments,
		Name:                    name.Value,
	}, nil
}

func (GenericParameterReducerImpl) ConstrainedToGenericParameter(
	name TokenValue,
	constraint TypeExpression,
) (
	*GenericParameter,
	error,
) {
	param := &GenericParameter{
		StartEndPos: newStartEndPos(name.Loc(), constraint.End()),
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

type GenericParameterList = NodeList[*GenericParameter]

var _ Node = &GenericParameterList{}

func NewGenericParameterList() *GenericParameterList {
	return newNodeList[*GenericParameter]("GenericParameterList")
}

type GenericParameterListReducerImpl struct{}

var _ GenericParametersReducer = &GenericParameterListReducerImpl{}
var _ ProperGenericParameterListReducer = &GenericParameterListReducerImpl{}
var _ GenericParameterListReducer = &GenericParameterListReducerImpl{}

func (GenericParameterListReducerImpl) GenericToGenericParameters(
	dollarLbracket TokenValue,
	list *GenericParameterList,
	rbracket TokenValue,
) (
	*GenericParameterList,
	error,
) {
	list.reduceMarkers(dollarLbracket, rbracket)
	return list, nil
}

func (GenericParameterListReducerImpl) NilToGenericParameters() (
	*GenericParameterList,
	error,
) {
	return nil, nil
}

func (GenericParameterListReducerImpl) AddToProperGenericParameterList(
	list *GenericParameterList,
	comma TokenValue,
	parameter *GenericParameter,
) (
	*GenericParameterList,
	error,
) {
	list.reduceAdd(comma, parameter)
	return list, nil
}

func (GenericParameterListReducerImpl) GenericParameterToProperGenericParameterList(
	parameter *GenericParameter,
) (
	*GenericParameterList,
	error,
) {
	list := NewGenericParameterList()
	list.add(parameter)
	return list, nil
}

func (GenericParameterListReducerImpl) ImproperToGenericParameterList(
	list *GenericParameterList,
	comma TokenValue,
) (
	*GenericParameterList,
	error,
) {
	list.reduceImproper(comma)
	return list, nil
}

func (GenericParameterListReducerImpl) NilToGenericParameterList() (
	*GenericParameterList,
	error,
) {
	return NewGenericParameterList(), nil
}

//
// GenericArgumentList
//

type GenericArgumentList = NodeList[TypeExpression]

var _ Node = &GenericArgumentList{}

func NewGenericArgumentList() *GenericArgumentList {
	return newNodeList[TypeExpression]("GenericArgumentList")
}

type GenericArgumentListReducerImpl struct{}

var _ GenericArgumentsReducer = &GenericArgumentListReducerImpl{}
var _ ProperGenericArgumentListReducer = &GenericArgumentListReducerImpl{}
var _ GenericArgumentListReducer = &GenericArgumentListReducerImpl{}

func (reducer *GenericArgumentListReducerImpl) BindingToGenericArguments(
	dollarLbracket TokenValue,
	list *GenericArgumentList,
	rbracket TokenValue,
) (
	*GenericArgumentList,
	error,
) {
	list.reduceMarkers(dollarLbracket, rbracket)
	return list, nil
}

func (reducer *GenericArgumentListReducerImpl) NilToGenericArguments() (
	*GenericArgumentList,
	error,
) {
	return nil, nil
}

func (reducer *GenericArgumentListReducerImpl) AddToProperGenericArgumentList(
	list *GenericArgumentList,
	comma TokenValue,
	arg TypeExpression,
) (
	*GenericArgumentList,
	error,
) {
	list.reduceAdd(comma, arg)
	return list, nil
}

func (reducer *GenericArgumentListReducerImpl) TypeExprToProperGenericArgumentList(
	arg TypeExpression,
) (
	*GenericArgumentList,
	error,
) {
	list := NewGenericArgumentList()
	list.add(arg)
	return list, nil
}

func (reducer *GenericArgumentListReducerImpl) ImproperToGenericArgumentList(
	list *GenericArgumentList,
	comma TokenValue,
) (
	*GenericArgumentList,
	error,
) {
	list.reduceImproper(comma)
	return list, nil
}

func (reducer *GenericArgumentListReducerImpl) NilToGenericArgumentList() (
	*GenericArgumentList,
	error,
) {
	return &GenericArgumentList{}, nil
}
