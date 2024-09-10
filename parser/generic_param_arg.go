package parser

//
// GenericArgumentList
//

type GenericArgumentList = NodeList[TypeExpression]

func NewGenericArgumentList() *GenericArgumentList {
	return newNodeList[TypeExpression]("GenericArgumentList")
}

var _ Node = &GenericArgumentList{}

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
