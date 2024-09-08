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
	return &TypeArgumentList{}, nil
}

//
// SliceTypeExpr
//

type SliceTypeExpr struct {
	isTypeExpression
	StartEndPos
	LeadingTrailingComments

	Value TypeExpression
}

func (slice SliceTypeExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf("%s%s[SliceTypeExpr:\n", indent, label)
	result += slice.Value.TreeString(indent+"  ", "Value=")
	result += "\n" + indent + "]"
	return result
}

type SliceTypeExprReducerImpl struct {
	SliceTypeExprs []*SliceTypeExpr
}

var _ SliceTypeExprReducer = &SliceTypeExprReducerImpl{}

func (reducer *SliceTypeExprReducerImpl) ToSliceTypeExpr(
	lbracket TokenValue,
	value TypeExpression,
	rbracket TokenValue,
) (
	TypeExpression,
	error,
) {
	slice := &SliceTypeExpr{
		StartEndPos: newStartEndPos(lbracket.Loc(), rbracket.End()),
		Value:       value,
	}

	slice.LeadingComment = lbracket.TakeLeading()
	value.PrependToLeading(lbracket.TakeTrailing())
	value.AppendToTrailing(rbracket.TakeLeading())
	slice.TrailingComment = rbracket.TakeTrailing()

	reducer.SliceTypeExprs = append(reducer.SliceTypeExprs, slice)
	return slice, nil
}

//
// ArrayTypeExpr
//

type ArrayTypeExpr struct {
	isTypeExpression
	StartEndPos
	LeadingTrailingComments

	Value TypeExpression
	Size  TokenValue
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

type ArrayTypeExprReducerImpl struct {
	ArrayTypeExprs []*ArrayTypeExpr
}

var _ ArrayTypeExprReducer = &ArrayTypeExprReducerImpl{}

func (reducer *ArrayTypeExprReducerImpl) ToArrayTypeExpr(
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
		StartEndPos: newStartEndPos(lbracket.Loc(), rbracket.End()),
		Value:       value,
		Size:        size,
	}

	array.LeadingComment = lbracket.TakeLeading()
	value.PrependToLeading(lbracket.TakeTrailing())
	value.AppendToTrailing(comma.TakeLeading())
	array.Size.PrependToLeading(comma.TakeTrailing())
	array.Size.AppendToTrailing(rbracket.TakeLeading())
	array.TrailingComment = rbracket.TakeTrailing()

	reducer.ArrayTypeExprs = append(reducer.ArrayTypeExprs, array)
	return array, nil
}

//
// MapTypeExpr
//

type MapTypeExpr struct {
	isTypeExpression
	StartEndPos
	LeadingTrailingComments

	Key   TypeExpression
	Value TypeExpression
}

func (dict MapTypeExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf("%s%s[MapTypeExpr:\n", indent, label)
	result += dict.Key.TreeString(indent+"  ", "Key=") + "\n"
	result += dict.Value.TreeString(indent+"  ", "Value=")
	result += "\n" + indent + "]"
	return result
}

type MapTypeExprReducerImpl struct {
	MapTypeExprs []*MapTypeExpr
}

var _ MapTypeExprReducer = &MapTypeExprReducerImpl{}

func (reducer *MapTypeExprReducerImpl) ToMapTypeExpr(
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
		StartEndPos: newStartEndPos(lbracket.Loc(), rbracket.End()),
		Key:         key,
		Value:       value,
	}

	dict.LeadingComment = lbracket.TakeLeading()
	key.PrependToLeading(lbracket.TakeTrailing())
	key.AppendToTrailing(semicolon.TakeLeading())
	value.PrependToLeading(semicolon.TakeTrailing())
	value.AppendToTrailing(rbracket.TakeLeading())
	dict.TrailingComment = rbracket.TakeTrailing()

	reducer.MapTypeExprs = append(reducer.MapTypeExprs, dict)
	return dict, nil
}

//
// InferredTypeExpr
//

type InferredTypeExpr struct {
	isTypeExpression
	TokenValue
}

func (expr InferredTypeExpr) TreeString(indent string, label string) string {
	return fmt.Sprintf(
		"%s%s[InferredTypeExpr: SymbolId=%s]",
		indent,
		label,
		expr.SymbolId)
}

type InferredTypeExprReducerImpl struct {
	InferredTypeExprs []*InferredTypeExpr
}

var _ InferredTypeExprReducer = &InferredTypeExprReducerImpl{}

func (reducer *InferredTypeExprReducerImpl) DotToInferredTypeExpr(
	dot TokenValue,
) (
	TypeExpression,
	error,
) {
	expr := &InferredTypeExpr{TokenValue: dot}
	reducer.InferredTypeExprs = append(reducer.InferredTypeExprs, expr)
	return expr, nil
}

func (reducer *InferredTypeExprReducerImpl) UnderscoreToInferredTypeExpr(
	underscore TokenValue,
) (
	TypeExpression,
	error,
) {
	expr := &InferredTypeExpr{TokenValue: underscore}
	reducer.InferredTypeExprs = append(reducer.InferredTypeExprs, expr)
	return expr, nil
}

//
// NamedTypeExpr
//

type NamedTypeExpr struct {
	isTypeExpression
	StartEndPos
	LeadingTrailingComments

	Pkg string // optional.  "" = local

	Name TokenValue

	TypeArguments TypeArgumentList
}

func (expr NamedTypeExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[NamedTypeExpr: Pkg=%s Name=%s",
		indent,
		label,
		expr.Pkg,
		expr.Name.Value)

	if len(expr.TypeArguments.Elements) == 0 {
		return result + "]"
	}

	result += expr.TypeArguments.TreeString(indent+"  ", "TypeArguments=")
	result += "\n" + indent + "]"
	return result
}

type NamedTypeExprReducerImpl struct {
	NamedTypeExprs []*NamedTypeExpr
}

var _ NamedTypeExprReducer = &NamedTypeExprReducerImpl{}

func (reducer *NamedTypeExprReducerImpl) toNamedTypeExpr(
	name TokenValue,
	typeArguments *TypeArgumentList,
) *NamedTypeExpr {
	var endPos Location
	var trailing CommentGroups
	if typeArguments == nil {
		typeArguments = &TypeArgumentList{}
		endPos = name.End()
		trailing = name.TakeTrailing()
	} else {
		endPos = typeArguments.End()
		trailing = typeArguments.TakeTrailing()
	}
	named := &NamedTypeExpr{
		StartEndPos: newStartEndPos(name.Loc(), endPos),
		LeadingTrailingComments: LeadingTrailingComments{
			TrailingComment: trailing,
		},
		Name:          name,
		TypeArguments: *typeArguments,
	}

	reducer.NamedTypeExprs = append(reducer.NamedTypeExprs, named)
	return named
}

func (reducer *NamedTypeExprReducerImpl) LocalToNamedTypeExpr(
	name TokenValue,
	typeArguments *TypeArgumentList,
) (
	TypeExpression,
	error,
) {
	named := reducer.toNamedTypeExpr(name, typeArguments)
	named.LeadingComment = name.TakeLeading()
	return named, nil
}

func (reducer *NamedTypeExprReducerImpl) ExternalToNamedTypeExpr(
	pkg TokenValue,
	dot TokenValue,
	name TokenValue,
	typeArguments *TypeArgumentList,
) (
	TypeExpression,
	error,
) {
	name.PrependToLeading(dot.TakeTrailing())
	name.PrependToLeading(dot.TakeLeading())
	name.PrependToLeading(pkg.TakeTrailing())

	named := reducer.toNamedTypeExpr(name, typeArguments)
	named.StartPos = pkg.Loc()
	named.Pkg = pkg.Value
	named.LeadingComment = pkg.TakeLeading()
	return named, nil
}
