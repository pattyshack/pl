package reducer

import (
	. "github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser/lr"
)

//
// FieldDef
//

func (Reducer) NamedToFieldDef(
	name *lr.TokenValue,
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

func (Reducer) UnnamedToFieldDef(
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

func (Reducer) DefaultEnumFieldDefToTypeProperty(
	defaultKW *lr.TokenValue,
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

func (reducer *Reducer) PaddingFieldDefToTypeProperty(
	underscore *lr.TokenValue,
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

func (Reducer) AddToProperImplicitTypeProperties(
	list *TypePropertyList,
	comma *lr.TokenValue,
	property TypeProperty,
) (
	*TypePropertyList,
	error,
) {
	list.ReduceAdd(comma, property)
	return list, nil
}

func (Reducer) TypePropertyToProperImplicitTypeProperties(
	property TypeProperty,
) (
	*TypePropertyList,
	error,
) {
	list := NewTypePropertyList()
	list.Add(property)
	return list, nil
}

func (Reducer) ImproperToImplicitTypeProperties(
	list *TypePropertyList,
	comma *lr.TokenValue,
) (
	*TypePropertyList,
	error,
) {
	list.ReduceImproper(comma)
	return list, nil
}

func (Reducer) NilToImplicitTypeProperties() (
	*TypePropertyList,
	error,
) {
	return NewTypePropertyList(), nil
}

func (Reducer) AddImplicitToProperExplicitTypeProperties(
	list *TypePropertyList,
	newlines lr.TokenCount,
	property TypeProperty,
) (
	*TypePropertyList,
	error,
) {
	list.ReduceAdd(&lr.TokenValue{}, property)
	return list, nil
}

func (Reducer) AddExplicitToProperExplicitTypeProperties(
	list *TypePropertyList,
	comma *lr.TokenValue,
	property TypeProperty,
) (
	*TypePropertyList,
	error,
) {
	list.ReduceAdd(comma, property)
	return list, nil
}

func (Reducer) TypePropertyToProperExplicitTypeProperties(
	property TypeProperty,
) (
	*TypePropertyList,
	error,
) {
	list := NewTypePropertyList()
	list.Add(property)
	return list, nil
}

func (Reducer) ImproperImplicitToExplicitTypeProperties(
	list *TypePropertyList,
	newlines lr.TokenCount,
) (
	*TypePropertyList,
	error,
) {
	return list, nil
}

func (Reducer) ImproperExplicitToExplicitTypeProperties(
	list *TypePropertyList,
	comma *lr.TokenValue,
) (
	*TypePropertyList,
	error,
) {
	list.ReduceImproper(comma)
	return list, nil
}

func (Reducer) NilToExplicitTypeProperties() (
	*TypePropertyList,
	error,
) {
	return NewTypePropertyList(), nil
}

func (Reducer) PairToProperImplicitEnumTypeProperties(
	property1 TypeProperty,
	or *lr.TokenValue,
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

func (Reducer) AddToProperImplicitEnumTypeProperties(
	list *TypePropertyList,
	or *lr.TokenValue,
	property TypeProperty,
) (
	*TypePropertyList,
	error,
) {
	list.ReduceAdd(or, property)
	return list, nil
}

func (Reducer) ImproperToImplicitEnumTypeProperties(
	list *TypePropertyList,
	newlines lr.TokenCount,
) (
	*TypePropertyList,
	error,
) {
	return list, nil
}

func (Reducer) ExplicitPairToProperExplicitEnumTypeProperties(
	property1 TypeProperty,
	or *lr.TokenValue,
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

func (Reducer) ImplicitPairToProperExplicitEnumTypeProperties(
	property1 TypeProperty,
	newlines lr.TokenCount,
	property2 TypeProperty,
) (
	*TypePropertyList,
	error,
) {
	list := NewTypePropertyList()
	list.Add(property1)
	list.ReduceAdd(&lr.TokenValue{}, property2)
	return list, nil
}

func (Reducer) ExplicitAddToProperExplicitEnumTypeProperties(
	list *TypePropertyList,
	or *lr.TokenValue,
	property TypeProperty,
) (
	*TypePropertyList,
	error,
) {
	list.ReduceAdd(or, property)
	return list, nil
}

func (Reducer) ImplicitAddToProperExplicitEnumTypeProperties(
	list *TypePropertyList,
	newlines lr.TokenCount,
	property TypeProperty,
) (
	*TypePropertyList,
	error,
) {
	list.ReduceAdd(&lr.TokenValue{}, property)
	return list, nil
}

func (Reducer) ImproperToExplicitEnumTypeProperties(
	list *TypePropertyList,
	newlines lr.TokenCount,
) (
	*TypePropertyList,
	error,
) {
	return list, nil
}
