package parser

import (
	"fmt"
)

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
		StartEndPos: NewStartEndPos(lbracket.Loc(), rbracket.End()),
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
		StartEndPos: NewStartEndPos(lbracket.Loc(), rbracket.End()),
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
		StartEndPos: NewStartEndPos(lbracket.Loc(), rbracket.End()),
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

	GenericArguments GenericArgumentList
}

func (expr NamedTypeExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[NamedTypeExpr: Pkg=%s Name=%s",
		indent,
		label,
		expr.Pkg,
		expr.Name.Value)

	if len(expr.GenericArguments.Elements) == 0 {
		return result + "]"
	}

	result += "\n" + expr.GenericArguments.TreeString(
		indent+"  ",
		"GenericArguments=")
	result += "\n" + indent + "]"
	return result
}

type NamedTypeExprReducerImpl struct {
	NamedTypeExprs []*NamedTypeExpr
}

var _ NamedTypeExprReducer = &NamedTypeExprReducerImpl{}

func (reducer *NamedTypeExprReducerImpl) toNamedTypeExpr(
	name TokenValue,
	genericArguments *GenericArgumentList,
) *NamedTypeExpr {
	var endPos Location
	var trailing CommentGroups
	if genericArguments == nil {
		genericArguments = &GenericArgumentList{}
		endPos = name.End()
		trailing = name.TakeTrailing()
	} else {
		endPos = genericArguments.End()
		trailing = genericArguments.TakeTrailing()
	}
	named := &NamedTypeExpr{
		StartEndPos: NewStartEndPos(name.Loc(), endPos),
		LeadingTrailingComments: LeadingTrailingComments{
			TrailingComment: trailing,
		},
		Name:             name,
		GenericArguments: *genericArguments,
	}

	reducer.NamedTypeExprs = append(reducer.NamedTypeExprs, named)
	return named
}

func (reducer *NamedTypeExprReducerImpl) LocalToNamedTypeExpr(
	name TokenValue,
	genericArguments *GenericArgumentList,
) (
	TypeExpression,
	error,
) {
	named := reducer.toNamedTypeExpr(name, genericArguments)
	named.LeadingComment = name.TakeLeading()
	return named, nil
}

func (reducer *NamedTypeExprReducerImpl) ExternalToNamedTypeExpr(
	pkg TokenValue,
	dot TokenValue,
	name TokenValue,
	genericArguments *GenericArgumentList,
) (
	TypeExpression,
	error,
) {
	name.PrependToLeading(dot.TakeTrailing())
	name.PrependToLeading(dot.TakeLeading())
	name.PrependToLeading(pkg.TakeTrailing())

	named := reducer.toNamedTypeExpr(name, genericArguments)
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
		StartEndPos: NewStartEndPos(op.Loc(), operand.End()),
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
		StartEndPos: NewStartEndPos(left.Loc(), right.End()),
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
// Struct / Enum / Trait PropertiesType
//

type PropertiesKind string

const (
	StructKind = PropertiesKind("struct")
	EnumKind   = PropertiesKind("enum")
	TraitKind  = PropertiesKind("trait")
)

type PropertiesTypeExpr struct {
	isTypeExpression
	StartEndPos
	LeadingTrailingComments

	Kind PropertiesKind

	IsImplicit bool

	Properties TypePropertyList
}

var _ TypeExpression = &PropertiesTypeExpr{}

func (expr PropertiesTypeExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[PropertiesTypeExpr: Kind=%s IsImplicit=%v\n",
		indent,
		label,
		expr.Kind,
		expr.IsImplicit)
	result += expr.Properties.TreeString(indent+"  ", "Properties=")
	result += "\n" + indent + "]"
	return result
}

type PropertiesTypeExprReducer struct {
	StructTypeExprs []*PropertiesTypeExpr
	EnumTypeExprs   []*PropertiesTypeExpr
	TraitTypeExprs  []*PropertiesTypeExpr
}

var _ ImplicitStructTypeExprReducer = &PropertiesTypeExprReducer{}
var _ ExplicitStructTypeExprReducer = &PropertiesTypeExprReducer{}
var _ TraitTypeExprReducer = &PropertiesTypeExprReducer{}
var _ ImplicitEnumTypeExprReducer = &PropertiesTypeExprReducer{}
var _ ExplicitEnumTypeExprReducer = &PropertiesTypeExprReducer{}

func (PropertiesTypeExprReducer) implicit(
	kind PropertiesKind,
	lparen TokenValue,
	properties *TypePropertyList,
	rparen TokenValue,
) *PropertiesTypeExpr {
	properties.ReduceMarkers(lparen, rparen)
	expr := &PropertiesTypeExpr{
		StartEndPos: NewStartEndPos(lparen.Loc(), rparen.End()),
		Kind:        kind,
		IsImplicit:  true,
		Properties:  *properties,
	}
	expr.LeadingComment = properties.TakeLeading()
	expr.TrailingComment = properties.TakeTrailing()

	return expr
}

func (PropertiesTypeExprReducer) explicit(
	kind PropertiesKind,
	kw TokenValue,
	lparen TokenValue,
	properties *TypePropertyList,
	rparen TokenValue,
) *PropertiesTypeExpr {
	properties.ReduceMarkers(lparen, rparen)

	properties.PrependToLeading(kw.TakeTrailing())
	trailing := properties.TakeTrailing()

	expr := &PropertiesTypeExpr{
		StartEndPos: NewStartEndPos(kw.Loc(), rparen.End()),
		Kind:        kind,
		IsImplicit:  false,
		Properties:  *properties,
	}
	expr.LeadingComment = kw.TakeLeading()
	expr.TrailingComment = trailing

	return expr
}

func (reducer *PropertiesTypeExprReducer) ToImplicitStructTypeExpr(
	lparen TokenValue,
	properties *TypePropertyList,
	rparen TokenValue,
) (
	TypeExpression,
	error,
) {
	expr := reducer.implicit(StructKind, lparen, properties, rparen)
	reducer.StructTypeExprs = append(reducer.StructTypeExprs, expr)
	return expr, nil
}

func (reducer *PropertiesTypeExprReducer) ToExplicitStructTypeExpr(
	structKW TokenValue,
	lparen TokenValue,
	properties *TypePropertyList,
	rparen TokenValue,
) (
	TypeExpression,
	error,
) {
	expr := reducer.explicit(StructKind, structKW, lparen, properties, rparen)
	reducer.StructTypeExprs = append(reducer.StructTypeExprs, expr)
	return expr, nil
}

func (reducer *PropertiesTypeExprReducer) ToTraitTypeExpr(
	trait TokenValue,
	lparen TokenValue,
	properties *TypePropertyList,
	rparen TokenValue,
) (
	TypeExpression,
	error,
) {
	expr := reducer.explicit(TraitKind, trait, lparen, properties, rparen)
	reducer.TraitTypeExprs = append(reducer.TraitTypeExprs, expr)
	return expr, nil
}

func (reducer *PropertiesTypeExprReducer) ToImplicitEnumTypeExpr(
	lparen TokenValue,
	properties *TypePropertyList,
	rparen TokenValue,
) (
	TypeExpression,
	error,
) {
	expr := reducer.implicit(EnumKind, lparen, properties, rparen)
	reducer.EnumTypeExprs = append(reducer.EnumTypeExprs, expr)
	return expr, nil
}

func (reducer *PropertiesTypeExprReducer) ToExplicitEnumTypeExpr(
	enum TokenValue,
	lparen TokenValue,
	properties *TypePropertyList,
	rparen TokenValue,
) (
	TypeExpression,
	error,
) {
	expr := reducer.explicit(EnumKind, enum, lparen, properties, rparen)
	reducer.EnumTypeExprs = append(reducer.EnumTypeExprs, expr)
	return expr, nil
}
