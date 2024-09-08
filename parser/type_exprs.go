package parser

import (
	"fmt"
)

//
// TypeArgumentList
//

type TypeArgumentList = NodeList[TypeExpression]

var _ Node = &TypeArgumentList{}

type TypeArgumentListReducerImpl struct{}

var _ GenericTypeArgumentsReducer = &TypeArgumentListReducerImpl{}
var _ ProperTypeArgumentsReducer = &TypeArgumentListReducerImpl{}
var _ TypeArgumentsReducer = &TypeArgumentListReducerImpl{}

func (reducer *TypeArgumentListReducerImpl) BindingToGenericTypeArguments(
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

func (reducer *TypeArgumentListReducerImpl) NilToGenericTypeArguments() (
	*TypeArgumentList,
	error,
) {
	return nil, nil
}

func (reducer *TypeArgumentListReducerImpl) AddToProperTypeArguments(
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

func (reducer *TypeArgumentListReducerImpl) TypeExprToProperTypeArguments(
	arg TypeExpression,
) (
	*TypeArgumentList,
	error,
) {
	return newNodeList[TypeExpression]("TypeArgumentList", arg), nil
}

func (reducer *TypeArgumentListReducerImpl) ImproperToTypeArguments(
	list *TypeArgumentList,
	comma TokenValue,
) (
	*TypeArgumentList,
	error,
) {
	list.reduceImproper(comma)
	return list, nil
}

func (reducer *TypeArgumentListReducerImpl) NilToTypeArguments() (
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

//
// UnaryTypeExpr
//

type UnaryTypeOp SymbolId

type UnaryTypeExpr struct {
	isTypeExpression
	StartEndPos
	LeadingTrailingComments

	Op      UnaryTypeOp
	Operand TypeExpression
}

func (expr UnaryTypeExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[UnaryTypeExpr: Op=%s\n",
		indent,
		label,
		SymbolId(expr.Op))
	result += expr.Operand.TreeString(indent+"  ", "Operand=")
	result += "\n" + indent + "]"
	return result
}

type UnaryTypeExprReducerImpl struct {
	UnaryTypeExprs []*UnaryTypeExpr
}

var _ PrefixUnaryTypeExprReducer = &UnaryTypeExprReducerImpl{}

func (reducer *UnaryTypeExprReducerImpl) ToPrefixUnaryTypeExpr(
	op TokenValue,
	operand TypeExpression,
) (
	TypeExpression,
	error,
) {
	expr := &UnaryTypeExpr{
		StartEndPos: newStartEndPos(op.Loc(), operand.End()),
		Op:          UnaryTypeOp(op.SymbolId),
		Operand:     operand,
	}

	expr.LeadingComment = op.TakeLeading()
	operand.PrependToLeading(op.TakeTrailing())
	expr.TrailingComment = operand.TakeTrailing()

	reducer.UnaryTypeExprs = append(reducer.UnaryTypeExprs, expr)
	return expr, nil
}

//
// BinaryTypeExpr
//

type BinaryTypeOp SymbolId

type BinaryTypeExpr struct {
	isTypeExpression
	StartEndPos
	LeadingTrailingComments

	Left  TypeExpression
	Op    BinaryTypeOp
	Right TypeExpression
}

func (expr BinaryTypeExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[BinaryTypeExpr: Op=%s\n", indent, label, SymbolId(expr.Op))
	result += expr.Left.TreeString(indent+"  ", "Left=") + "\n"
	result += expr.Right.TreeString(indent+"  ", "Right=")
	result += "\n" + indent + "]"
	return result
}

type BinaryTypeExprReducerImpl struct {
	BinaryTypeExprs []*BinaryTypeExpr
}

var _ BinaryTypeExprReducer = &BinaryTypeExprReducerImpl{}

func (reducer *BinaryTypeExprReducerImpl) ToBinaryTypeExpr(
	left TypeExpression,
	op TokenValue,
	right TypeExpression,
) (
	TypeExpression,
	error,
) {
	expr := &BinaryTypeExpr{
		StartEndPos: newStartEndPos(left.Loc(), right.End()),
		Left:        left,
		Op:          BinaryTypeOp(op.SymbolId),
		Right:       right,
	}

	expr.LeadingComment = left.TakeLeading()
	left.AppendToTrailing(op.TakeLeading())
	right.PrependToLeading(op.TakeTrailing())
	expr.TrailingComment = right.TakeTrailing()

	reducer.BinaryTypeExprs = append(reducer.BinaryTypeExprs, expr)
	return expr, nil
}

//
// Parameter
//

type ParameterKind int

const (
	// NOTE: There is no implicit unnamed inferred arg variant
	NamedTypedArgParameter = ParameterKind(iota)
	NamedInferredArgParameter
	NamedTypedVarargParameter
	NamedInferredVarargParameter
	UnnamedTypedArgParameter
	UnnamedTypedVarargParameter
	UnnamedInferredVarargParameter
	IgnoreTypedArgParameter
	IgnoreInferredArgParameter
	IgnoreTypedVarargParameter
	IgnoreInferredVarargParameter
)

type Parameter struct {
	StartEndPos
	LeadingTrailingComments

	Kind ParameterKind

	Name        *TokenValue // optional
	HasEllipsis bool
	Type        TypeExpression // optional
}

var _ Node = &Parameter{}

func (param Parameter) TreeString(indent string, label string) string {
	name := ""
	if param.Name != nil {
		name = param.Name.Value
	}
	result := fmt.Sprintf(
		"%s%s[Parameter: Kind=%v Name=%s HasEllipsis=%v\n",
		indent,
		label,
		param.Kind,
		name,
		param.HasEllipsis)
	if param.Type != nil {
		result += param.Type.TreeString(indent+"  ", "Type=")
		result += "\n" + indent + "]"
	} else {
		result += "]"
	}
	return result
}

type ParameterReducerImpl struct{}

var _ ProperParameterDefReducer = &ParameterReducerImpl{}
var _ ParameterDeclReducer = &ParameterReducerImpl{}
var _ ParameterDefReducer = &ParameterReducerImpl{}

func (reducer *ParameterReducerImpl) namedTypedArg(
	kind ParameterKind,
	name *TokenValue,
	typeExpr TypeExpression,
) (
	*Parameter,
	error,
) {
	param := &Parameter{
		StartEndPos: newStartEndPos(name.Loc(), typeExpr.End()),
		Kind:        kind,
		Name:        name,
		HasEllipsis: false,
		Type:        typeExpr,
	}
	param.LeadingComment = name.TakeLeading()
	param.TrailingComment = typeExpr.TakeTrailing()
	return param, nil
}

func (reducer *ParameterReducerImpl) NamedTypedArgToProperParameterDef(
	name TokenValue,
	typeExpr TypeExpression,
) (
	*Parameter,
	error,
) {
	return reducer.namedTypedArg(NamedTypedArgParameter, &name, typeExpr)
}

func (reducer *ParameterReducerImpl) IgnoreTypedArgToProperParameterDef(
	underscore TokenValue,
	typeExpr TypeExpression,
) (
	*Parameter,
	error,
) {
	return reducer.namedTypedArg(IgnoreTypedArgParameter, &underscore, typeExpr)
}

func (reducer *ParameterReducerImpl) namedTypedVararg(
	kind ParameterKind,
	name *TokenValue,
	ellipsis TokenValue,
	typeExpr TypeExpression,
) (
	*Parameter,
	error,
) {
	param := &Parameter{
		StartEndPos: newStartEndPos(name.Loc(), typeExpr.End()),
		Kind:        kind,
		Name:        name,
		HasEllipsis: true,
		Type:        typeExpr,
	}
	param.LeadingComment = name.TakeLeading()
	name.AppendToTrailing(ellipsis.TakeLeading())
	typeExpr.PrependToLeading(ellipsis.TakeTrailing())
	param.TrailingComment = typeExpr.TakeTrailing()
	return param, nil
}

func (reducer *ParameterReducerImpl) NamedTypedVarargToProperParameterDef(
	name TokenValue,
	ellipsis TokenValue,
	typeExpr TypeExpression,
) (
	*Parameter,
	error,
) {
	return reducer.namedTypedVararg(
		NamedTypedVarargParameter,
		&name,
		ellipsis,
		typeExpr)
}

func (reducer *ParameterReducerImpl) IgnoreTypedVarargToProperParameterDef(
	underscore TokenValue,
	ellipsis TokenValue,
	typeExpr TypeExpression,
) (
	*Parameter,
	error,
) {
	return reducer.namedTypedVararg(
		IgnoreTypedVarargParameter,
		&underscore,
		ellipsis,
		typeExpr)
}

func (reducer *ParameterReducerImpl) namedInferredVararg(
	kind ParameterKind,
	name *TokenValue,
	ellipsis TokenValue,
) (
	*Parameter,
	error,
) {
	param := &Parameter{
		StartEndPos: newStartEndPos(name.Loc(), ellipsis.End()),
		Kind:        kind,
		Name:        name,
		HasEllipsis: true,
		Type:        nil,
	}
	param.LeadingComment = name.TakeLeading()
	name.AppendToTrailing(ellipsis.TakeLeading())
	param.TrailingComment = ellipsis.TakeTrailing()
	return param, nil
}

func (reducer *ParameterReducerImpl) NamedInferredVarargToProperParameterDef(
	name TokenValue,
	ellipsis TokenValue,
) (
	*Parameter,
	error,
) {
	return reducer.namedInferredVararg(
		NamedInferredVarargParameter,
		&name,
		ellipsis)
}

func (reducer *ParameterReducerImpl) IgnoreInferredVarargToProperParameterDef(
	underscore TokenValue,
	ellipsis TokenValue,
) (
	*Parameter,
	error,
) {
	return reducer.namedInferredVararg(
		IgnoreInferredVarargParameter,
		&underscore,
		ellipsis)
}

func (reducer *ParameterReducerImpl) UnnamedTypedArgToParameterDecl(
	typeExpr TypeExpression,
) (
	*Parameter,
	error,
) {
	param := &Parameter{
		StartEndPos: newStartEndPos(typeExpr.Loc(), typeExpr.End()),
		Kind:        UnnamedTypedArgParameter,
		Name:        nil,
		HasEllipsis: false,
		Type:        typeExpr,
	}
	param.LeadingComment = typeExpr.TakeLeading()
	param.TrailingComment = typeExpr.TakeTrailing()
	return param, nil
}

func (reducer *ParameterReducerImpl) UnnamedInferredVarargToParameterDecl(
	ellipsis TokenValue,
) (
	*Parameter,
	error,
) {
	param := &Parameter{
		StartEndPos: newStartEndPos(ellipsis.Loc(), ellipsis.End()),
		Kind:        UnnamedInferredVarargParameter,
		Name:        nil,
		HasEllipsis: true,
		Type:        nil,
	}
	param.LeadingTrailingComments = ellipsis.LeadingTrailingComments
	return param, nil
}

func (reducer *ParameterReducerImpl) UnnamedTypedVarargToParameterDecl(
	ellipsis TokenValue,
	typeExpr TypeExpression,
) (
	*Parameter,
	error,
) {
	param := &Parameter{
		StartEndPos: newStartEndPos(ellipsis.Loc(), typeExpr.End()),
		Kind:        UnnamedTypedVarargParameter,
		Name:        nil,
		HasEllipsis: true,
		Type:        typeExpr,
	}
	param.LeadingComment = ellipsis.TakeLeading()
	typeExpr.PrependToLeading(ellipsis.TakeTrailing())
	param.TrailingComment = typeExpr.TakeTrailing()
	return param, nil
}

func (reducer *ParameterReducerImpl) namedInferredArg(
	kind ParameterKind,
	name *TokenValue,
) (
	*Parameter,
	error,
) {
	param := &Parameter{
		StartEndPos: newStartEndPos(name.Loc(), name.End()),
		Kind:        kind,
		Name:        name,
		HasEllipsis: false,
		Type:        nil,
	}
	param.LeadingComment = name.TakeLeading()
	param.TrailingComment = name.TakeTrailing()
	return param, nil
}

func (reducer *ParameterReducerImpl) NamedInferredArgToParameterDef(
	name TokenValue,
) (
	*Parameter,
	error,
) {
	return reducer.namedInferredArg(NamedInferredArgParameter, &name)
}

func (reducer *ParameterReducerImpl) IgnoreInferredArgToParameterDef(
	underscore TokenValue,
) (
	*Parameter,
	error,
) {
	return reducer.namedInferredArg(IgnoreInferredArgParameter, &underscore)
}

//
// ParameterList
//

type ParameterList = NodeList[*Parameter]

var _ Node = &ParameterList{}

type ParameterListReducerImpl struct{}

var _ ProperParameterDeclsReducer = &ParameterListReducerImpl{}
var _ ParameterDeclsReducer = &ParameterListReducerImpl{}
var _ ProperParameterDefsReducer = &ParameterListReducerImpl{}
var _ ParameterDefsReducer = &ParameterListReducerImpl{}

func (reducer *ParameterListReducerImpl) AddToProperParameterDecls(
	list *ParameterList,
	comma TokenValue,
	parameter *Parameter,
) (
	*ParameterList,
	error,
) {
	list.reduceAdd(comma, parameter)
	return list, nil
}

func (reducer *ParameterListReducerImpl) ParameterDeclToProperParameterDecls(
	parameter *Parameter,
) (
	*ParameterList,
	error,
) {
	return newNodeList[*Parameter]("ParameterList", parameter), nil
}

func (reducer *ParameterListReducerImpl) ImproperToParameterDecls(
	list *ParameterList,
	comma TokenValue,
) (
	*ParameterList,
	error,
) {
	list.reduceImproper(comma)
	return list, nil
}

func (reducer *ParameterListReducerImpl) NilToParameterDecls() (
	*ParameterList,
	error,
) {
	return &ParameterList{ListType: "ParameterList"}, nil
}

func (reducer *ParameterListReducerImpl) AddToProperParameterDefs(
	list *ParameterList,
	comma TokenValue,
	parameter *Parameter,
) (
	*ParameterList,
	error,
) {
	list.reduceAdd(comma, parameter)
	return list, nil
}

func (reducer *ParameterListReducerImpl) ParameterDefToProperParameterDefs(
	parameter *Parameter,
) (
	*ParameterList,
	error,
) {
	return newNodeList[*Parameter]("ParameterList", parameter), nil
}

func (reducer *ParameterListReducerImpl) ImproperToParameterDefs(
	list *ParameterList,
	comma TokenValue,
) (
	*ParameterList,
	error,
) {
	list.reduceImproper(comma)
	return list, nil
}

func (reducer *ParameterListReducerImpl) NilToParameterDefs() (
	*ParameterList,
	error,
) {
	return &ParameterList{ListType: "ParameterList"}, nil
}
