package parser

import (
	"fmt"
)

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

//
// SliceTypeExpr
//

type SliceTypeExpr struct {
	isTypeExpression
	StartPos Location
	EndPos   Location
	LeadingTrailingComments

	Value TypeExpression
}

func (slice SliceTypeExpr) Loc() Location {
	return slice.StartPos
}

func (slice SliceTypeExpr) End() Location {
	return slice.StartPos
}

func (slice SliceTypeExpr) String() string {
	return slice.TreeString("", "")
}

func (slice SliceTypeExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf("%s%s[SliceTypeExpr:\n", indent, label)
	result += slice.Value.TreeString(indent+"  ", "Value=")
	result += "\n" + indent + "]"
	return result
}

type SliceTypeExprReducerImpl struct{}

var _ SliceTypeExprReducer = &SliceTypeExprReducerImpl{}

func (SliceTypeExprReducerImpl) ToSliceTypeExpr(
	lbracket TokenValue,
	value TypeExpression,
	rbracket TokenValue,
) (
	TypeExpression,
	error,
) {
	slice := &SliceTypeExpr{
		StartPos: lbracket.Loc(),
		EndPos:   rbracket.End(),
		Value:    value,
	}

	slice.LeadingComment = lbracket.TakeLeading()
	value.PrependToLeading(lbracket.TakeTrailing())
	value.AppendToTrailing(rbracket.TakeLeading())
	slice.TrailingComment = rbracket.TakeTrailing()
	return slice, nil
}

//
// ArrayTypeExpr
//

type ArrayTypeExpr struct {
	isTypeExpression
	StartPos Location
	EndPos   Location
	LeadingTrailingComments

	Value TypeExpression
	Size  TokenValue
}

func (array ArrayTypeExpr) Loc() Location {
	return array.StartPos
}

func (array ArrayTypeExpr) End() Location {
	return array.StartPos
}

func (array ArrayTypeExpr) String() string {
	return array.TreeString("", "")
}

func (array ArrayTypeExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[ArrayTypeExpr: Size=%s\n",
		indent,
		label,
		array.Size.Value)
	result += array.Value.TreeString(indent+"  ", "Value=")
	result += "\n" + indent + "]"
	return result
}

type ArrayTypeExprReducerImpl struct{}

var _ ArrayTypeExprReducer = &ArrayTypeExprReducerImpl{}

func (ArrayTypeExprReducerImpl) ToArrayTypeExpr(
	lbracket TokenValue,
	value TypeExpression,
	comma TokenValue,
	size TokenValue,
	rbracket TokenValue,
) (
	TypeExpression,
	error,
) {
	array := &ArrayTypeExpr{
		StartPos: lbracket.Loc(),
		EndPos:   rbracket.End(),
		Value:    value,
		Size:     size,
	}

	array.LeadingComment = lbracket.TakeLeading()
	value.PrependToLeading(lbracket.TakeTrailing())
	value.AppendToTrailing(comma.TakeLeading())
	array.Size.PrependToLeading(comma.TakeTrailing())
	array.Size.AppendToTrailing(rbracket.TakeLeading())
	array.TrailingComment = rbracket.TakeTrailing()
	return array, nil
}

//
// MapTypeExpr
//

type MapTypeExpr struct {
	isTypeExpression
	StartPos Location
	EndPos   Location
	LeadingTrailingComments

	Key   TypeExpression
	Value TypeExpression
}

func (dict MapTypeExpr) Loc() Location {
	return dict.StartPos
}

func (dict MapTypeExpr) End() Location {
	return dict.StartPos
}

func (dict MapTypeExpr) String() string {
	return dict.TreeString("", "")
}

func (dict MapTypeExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf("%s%s[MapTypeExpr:\n", indent, label)
	result += dict.Key.TreeString(indent+"  ", "Key=") + "\n"
	result += dict.Value.TreeString(indent+"  ", "Value=")
	result += "\n" + indent + "]"
	return result
}

type MapTypeExprReducerImpl struct{}

var _ MapTypeExprReducer = &MapTypeExprReducerImpl{}

func (MapTypeExprReducerImpl) ToMapTypeExpr(
	lbracket TokenValue,
	key TypeExpression,
	semicolon TokenValue,
	value TypeExpression,
	rbracket TokenValue,
) (
	TypeExpression,
	error,
) {
	dict := &MapTypeExpr{
		StartPos: lbracket.Loc(),
		EndPos:   rbracket.End(),
		Key:      key,
		Value:    value,
	}

	dict.LeadingComment = lbracket.TakeLeading()
	key.PrependToLeading(lbracket.TakeTrailing())
	key.AppendToTrailing(semicolon.TakeLeading())
	value.PrependToLeading(semicolon.TakeTrailing())
	value.AppendToTrailing(rbracket.TakeLeading())
	dict.TrailingComment = rbracket.TakeTrailing()
	return dict, nil
}
