package parser

type TypeArgumentList = NodeList[TypeExpression]

var _ Node = &TypeArgumentList{}

type TypeArgumentListReducer struct{}

var _ GenericTypeArgumentsReducer = &TypeArgumentListReducer{}
var _ ProperTypeArgumentsReducer = &TypeArgumentListReducer{}
var _ TypeArgumentsReducer = &TypeArgumentListReducer{}

func (reducer *TypeArgumentListReducer) BindingToGenericTypeArguments(
	dollarLbracket TokenValue,
	list *TypeArgumentList,
	rbracket TokenValue,
) (
	*TypeArgumentList,
	error,
) {
	list.reduceMarkers(dollarLbracket, rbracket)
	return list, nil
}

func (reducer *TypeArgumentListReducer) NilToGenericTypeArguments() (
	*TypeArgumentList,
	error,
) {
	return nil, nil
}

func (reducer *TypeArgumentListReducer) AddToProperTypeArguments(
	list *TypeArgumentList,
	comma TokenValue,
	arg TypeExpression,
) (
	*TypeArgumentList,
	error,
) {
	list.reduceAdd(comma, arg)
	return list, nil
}

func (reducer *TypeArgumentListReducer) TypeExprToProperTypeArguments(
	arg TypeExpression,
) (
	*TypeArgumentList,
	error,
) {
	return newNodeList[TypeExpression]("TypeArgumentList", arg), nil
}

func (reducer *TypeArgumentListReducer) ImproperToTypeArguments(
	list *TypeArgumentList,
	comma TokenValue,
) (
	*TypeArgumentList,
	error,
) {
	list.reduceImproper(comma)
	return list, nil
}

func (reducer *TypeArgumentListReducer) NilToTypeArguments() (
	*TypeArgumentList,
	error,
) {
	return nil, nil
}
