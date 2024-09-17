package parser

import (
	"fmt"
)

//
// FieldDef
//

type FieldDefKind string

const (
	NamedFieldDef              = FieldDefKind("named-field-def")
	UnnamedFieldDef            = FieldDefKind("unnamed-field-def")
	NamedDefaultEnumFieldDef   = FieldDefKind("named-default-enum-default-def")
	UnnamedDefaultEnumFieldDef = FieldDefKind("unnamed-default-enum-default-def")
	PaddingFieldDef            = FieldDefKind("padding-field-def")
)

type FieldDef struct {
	IsTypeProp
	StartEndPos
	LeadingTrailingComments

	Kind FieldDefKind

	Name string // Could be identifier, underscore, or ""
	Type TypeExpression
}

var _ TypeProperty = &FieldDef{}

func (def FieldDef) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[FieldDef: Kind=%s Name=%s\n",
		indent,
		label,
		def.Kind,
		def.Name)
	result += def.Type.TreeString(indent+"  ", "Type=")
	result += "\n" + indent + "]"
	return result
}

type FieldDefReducerImpl struct{}

var _ FieldDefReducer = &FieldDefReducerImpl{}
var _ TypePropertyReducer = &FieldDefReducerImpl{}

func (FieldDefReducerImpl) NamedToFieldDef(
	name *TokenValue,
	typeExpr TypeExpression,
) (
	*FieldDef,
	error,
) {
	def := &FieldDef{
		StartEndPos: NewStartEndPos(name.Loc(), typeExpr.End()),
		Kind:        NamedFieldDef,
		Name:        name.Value,
		Type:        typeExpr,
	}

	def.LeadingComment = name.TakeLeading()
	typeExpr.PrependToLeading(name.TakeTrailing())
	def.TrailingComment = typeExpr.TakeTrailing()

	return def, nil
}

func (FieldDefReducerImpl) UnnamedToFieldDef(
	typeExpr TypeExpression,
) (
	*FieldDef,
	error,
) {
	def := &FieldDef{
		StartEndPos: NewStartEndPos(typeExpr.Loc(), typeExpr.End()),
		Kind:        UnnamedFieldDef,
		Type:        typeExpr,
	}

	def.LeadingComment = typeExpr.TakeLeading()
	def.TrailingComment = typeExpr.TakeTrailing()

	return def, nil
}

func (FieldDefReducerImpl) DefaultEnumFieldDefToTypeProperty(
	defaultKW *TokenValue,
	def *FieldDef,
) (
	TypeProperty,
	error,
) {
	if def.Name == "" {
		def.Kind = UnnamedDefaultEnumFieldDef
	} else {
		def.Kind = NamedDefaultEnumFieldDef
	}

	def.PrependToLeading(defaultKW.TakeTrailing())
	def.PrependToLeading(defaultKW.TakeLeading())

	return def, nil
}

func (reducer FieldDefReducerImpl) PaddingFieldDefToTypeProperty(
	underscore *TokenValue,
	typeExpr TypeExpression,
) (
	TypeProperty,
	error,
) {
	def, err := reducer.NamedToFieldDef(underscore, typeExpr)
	def.Kind = PaddingFieldDef
	return def, err
}

//
// TypePropertyList
//

type TypePropertyList = NodeList[TypeProperty]

func NewTypePropertyList() *TypePropertyList {
	return NewNodeList[TypeProperty]("TypePropertyList")
}

var _ Node = &TypePropertyList{}

type TypePropertyListReducerImpl struct{}

var _ ProperImplicitTypePropertiesReducer = &TypePropertyListReducerImpl{}
var _ ImplicitTypePropertiesReducer = &TypePropertyListReducerImpl{}
var _ ProperExplicitTypePropertiesReducer = &TypePropertyListReducerImpl{}
var _ ExplicitTypePropertiesReducer = &TypePropertyListReducerImpl{}
var _ ProperImplicitEnumTypePropertiesReducer = &TypePropertyListReducerImpl{}
var _ ImplicitEnumTypePropertiesReducer = &TypePropertyListReducerImpl{}
var _ ProperExplicitEnumTypePropertiesReducer = &TypePropertyListReducerImpl{}
var _ ExplicitEnumTypePropertiesReducer = &TypePropertyListReducerImpl{}

func (TypePropertyListReducerImpl) AddToProperImplicitTypeProperties(
	list *TypePropertyList,
	comma *TokenValue,
	property TypeProperty,
) (
	*TypePropertyList,
	error,
) {
	list.ReduceAdd(comma, property)
	return list, nil
}

func (TypePropertyListReducerImpl) TypePropertyToProperImplicitTypeProperties(
	property TypeProperty,
) (
	*TypePropertyList,
	error,
) {
	list := NewTypePropertyList()
	list.Add(property)
	return list, nil
}

func (TypePropertyListReducerImpl) ImproperToImplicitTypeProperties(
	list *TypePropertyList,
	comma *TokenValue,
) (
	*TypePropertyList,
	error,
) {
	list.ReduceImproper(comma)
	return list, nil
}

func (TypePropertyListReducerImpl) NilToImplicitTypeProperties() (
	*TypePropertyList,
	error,
) {
	return NewTypePropertyList(), nil
}

func (TypePropertyListReducerImpl) AddImplicitToProperExplicitTypeProperties(
	list *TypePropertyList,
	newlines TokenCount,
	property TypeProperty,
) (
	*TypePropertyList,
	error,
) {
	list.ReduceAdd(&TokenValue{}, property)
	return list, nil
}

func (TypePropertyListReducerImpl) AddExplicitToProperExplicitTypeProperties(
	list *TypePropertyList,
	comma *TokenValue,
	property TypeProperty,
) (
	*TypePropertyList,
	error,
) {
	list.ReduceAdd(comma, property)
	return list, nil
}

func (TypePropertyListReducerImpl) TypePropertyToProperExplicitTypeProperties(
	property TypeProperty,
) (
	*TypePropertyList,
	error,
) {
	list := NewTypePropertyList()
	list.Add(property)
	return list, nil
}

func (TypePropertyListReducerImpl) ImproperImplicitToExplicitTypeProperties(
	list *TypePropertyList,
	newlines TokenCount,
) (
	*TypePropertyList,
	error,
) {
	return list, nil
}

func (TypePropertyListReducerImpl) ImproperExplicitToExplicitTypeProperties(
	list *TypePropertyList,
	comma *TokenValue,
) (
	*TypePropertyList,
	error,
) {
	list.ReduceImproper(comma)
	return list, nil
}

func (TypePropertyListReducerImpl) NilToExplicitTypeProperties() (
	*TypePropertyList,
	error,
) {
	return NewTypePropertyList(), nil
}

func (TypePropertyListReducerImpl) PairToProperImplicitEnumTypeProperties(
	property1 TypeProperty,
	or *TokenValue,
	property2 TypeProperty,
) (
	*TypePropertyList,
	error,
) {
	list := NewTypePropertyList()
	list.Add(property1)
	list.ReduceAdd(or, property2)
	return list, nil
}

func (TypePropertyListReducerImpl) AddToProperImplicitEnumTypeProperties(
	list *TypePropertyList,
	or *TokenValue,
	property TypeProperty,
) (
	*TypePropertyList,
	error,
) {
	list.ReduceAdd(or, property)
	return list, nil
}

func (TypePropertyListReducerImpl) ImproperToImplicitEnumTypeProperties(
	list *TypePropertyList,
	newlines TokenCount,
) (
	*TypePropertyList,
	error,
) {
	return list, nil
}

func (TypePropertyListReducerImpl) ExplicitPairToProperExplicitEnumTypeProperties(
	property1 TypeProperty,
	or *TokenValue,
	property2 TypeProperty,
) (
	*TypePropertyList,
	error,
) {
	list := NewTypePropertyList()
	list.Add(property1)
	list.ReduceAdd(or, property2)
	return list, nil
}

func (TypePropertyListReducerImpl) ImplicitPairToProperExplicitEnumTypeProperties(
	property1 TypeProperty,
	newlines TokenCount,
	property2 TypeProperty,
) (
	*TypePropertyList,
	error,
) {
	list := NewTypePropertyList()
	list.Add(property1)
	list.ReduceAdd(&TokenValue{}, property2)
	return list, nil
}

func (TypePropertyListReducerImpl) ExplicitAddToProperExplicitEnumTypeProperties(
	list *TypePropertyList,
	or *TokenValue,
	property TypeProperty,
) (
	*TypePropertyList,
	error,
) {
	list.ReduceAdd(or, property)
	return list, nil
}

func (TypePropertyListReducerImpl) ImplicitAddToProperExplicitEnumTypeProperties(
	list *TypePropertyList,
	newlines TokenCount,
	property TypeProperty,
) (
	*TypePropertyList,
	error,
) {
	list.ReduceAdd(&TokenValue{}, property)
	return list, nil
}

func (TypePropertyListReducerImpl) ImproperToExplicitEnumTypeProperties(
	list *TypePropertyList,
	newlines TokenCount,
) (
	*TypePropertyList,
	error,
) {
	return list, nil
}
