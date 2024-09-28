package reducer

import (
	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser/lr"
)

//
// ast.FieldDef
//

func (Reducer) NamedToFieldDef(
	name *lr.TokenValue,
	typeExpr ast.TypeExpression,
) (
	*ast.FieldDef,
	error,
) {
	def := &ast.FieldDef{
		StartEndPos: ast.NewStartEndPos(name.Loc(), typeExpr.End()),
		Kind:        ast.NamedFieldDef,
		Name:        name.Value,
		Type:        typeExpr,
	}

	def.LeadingComment = name.TakeLeading()
	typeExpr.PrependToLeading(name.TakeTrailing())
	def.TrailingComment = typeExpr.TakeTrailing()

	return def, nil
}

func (Reducer) UnnamedToFieldDef(
	typeExpr ast.TypeExpression,
) (
	*ast.FieldDef,
	error,
) {
	def := &ast.FieldDef{
		StartEndPos: ast.NewStartEndPos(typeExpr.Loc(), typeExpr.End()),
		Kind:        ast.UnnamedFieldDef,
		Type:        typeExpr,
	}

	def.LeadingComment = typeExpr.TakeLeading()
	def.TrailingComment = typeExpr.TakeTrailing()

	return def, nil
}

func (Reducer) DefaultEnumFieldDefToTypeProperty(
	defaultKW *lr.TokenValue,
	def *ast.FieldDef,
) (
	ast.TypeProperty,
	error,
) {
	if def.Name == "" {
		def.Kind = ast.UnnamedDefaultEnumFieldDef
	} else {
		def.Kind = ast.NamedDefaultEnumFieldDef
	}

	def.PrependToLeading(defaultKW.TakeTrailing())
	def.PrependToLeading(defaultKW.TakeLeading())

	return def, nil
}

func (reducer *Reducer) PaddingFieldDefToTypeProperty(
	underscore *lr.TokenValue,
	typeExpr ast.TypeExpression,
) (
	ast.TypeProperty,
	error,
) {
	def, err := reducer.NamedToFieldDef(underscore, typeExpr)
	def.Kind = ast.PaddingFieldDef
	return def, err
}

//
// ast.TypePropertyList
//

func (Reducer) AddToProperImplicitTypeProperties(
	list *ast.TypePropertyList,
	comma *lr.TokenValue,
	property ast.TypeProperty,
) (
	*ast.TypePropertyList,
	error,
) {
	list.ReduceAdd(comma, property)
	return list, nil
}

func (Reducer) TypePropertyToProperImplicitTypeProperties(
	property ast.TypeProperty,
) (
	*ast.TypePropertyList,
	error,
) {
	list := ast.NewTypePropertyList()
	list.Add(property)
	return list, nil
}

func (Reducer) ImproperToImplicitTypeProperties(
	list *ast.TypePropertyList,
	comma *lr.TokenValue,
) (
	*ast.TypePropertyList,
	error,
) {
	list.ReduceImproper(comma)
	return list, nil
}

func (Reducer) NilToImplicitTypeProperties() (
	*ast.TypePropertyList,
	error,
) {
	return nil, nil
}

func (Reducer) AddImplicitToProperExplicitTypeProperties(
	list *ast.TypePropertyList,
	newlines lr.TokenCount,
	property ast.TypeProperty,
) (
	*ast.TypePropertyList,
	error,
) {
	list.Add(property)
	return list, nil
}

func (Reducer) AddExplicitToProperExplicitTypeProperties(
	list *ast.TypePropertyList,
	comma *lr.TokenValue,
	property ast.TypeProperty,
) (
	*ast.TypePropertyList,
	error,
) {
	list.ReduceAdd(comma, property)
	return list, nil
}

func (Reducer) TypePropertyToProperExplicitTypeProperties(
	property ast.TypeProperty,
) (
	*ast.TypePropertyList,
	error,
) {
	list := ast.NewTypePropertyList()
	list.Add(property)
	return list, nil
}

func (Reducer) ImproperImplicitToExplicitTypeProperties(
	list *ast.TypePropertyList,
	newlines lr.TokenCount,
) (
	*ast.TypePropertyList,
	error,
) {
	return list, nil
}

func (Reducer) ImproperExplicitToExplicitTypeProperties(
	list *ast.TypePropertyList,
	comma *lr.TokenValue,
) (
	*ast.TypePropertyList,
	error,
) {
	list.ReduceImproper(comma)
	return list, nil
}

func (Reducer) NilToExplicitTypeProperties() (
	*ast.TypePropertyList,
	error,
) {
	return nil, nil
}

func (Reducer) PairToProperImplicitEnumTypeProperties(
	property1 ast.TypeProperty,
	or *lr.TokenValue,
	property2 ast.TypeProperty,
) (
	*ast.TypePropertyList,
	error,
) {
	list := ast.NewTypePropertyList()
	list.Add(property1)
	list.ReduceAdd(or, property2)
	return list, nil
}

func (Reducer) AddToProperImplicitEnumTypeProperties(
	list *ast.TypePropertyList,
	or *lr.TokenValue,
	property ast.TypeProperty,
) (
	*ast.TypePropertyList,
	error,
) {
	list.ReduceAdd(or, property)
	return list, nil
}

func (Reducer) ImproperToImplicitEnumTypeProperties(
	list *ast.TypePropertyList,
	newlines lr.TokenCount,
) (
	*ast.TypePropertyList,
	error,
) {
	return list, nil
}

func (Reducer) ExplicitPairToProperExplicitEnumTypeProperties(
	property1 ast.TypeProperty,
	or *lr.TokenValue,
	property2 ast.TypeProperty,
) (
	*ast.TypePropertyList,
	error,
) {
	list := ast.NewTypePropertyList()
	list.Add(property1)
	list.ReduceAdd(or, property2)
	return list, nil
}

func (Reducer) ImplicitPairToProperExplicitEnumTypeProperties(
	property1 ast.TypeProperty,
	newlines lr.TokenCount,
	property2 ast.TypeProperty,
) (
	*ast.TypePropertyList,
	error,
) {
	list := ast.NewTypePropertyList()
	list.Add(property1)
	list.Add(property2)
	return list, nil
}

func (Reducer) ExplicitAddToProperExplicitEnumTypeProperties(
	list *ast.TypePropertyList,
	or *lr.TokenValue,
	property ast.TypeProperty,
) (
	*ast.TypePropertyList,
	error,
) {
	list.ReduceAdd(or, property)
	return list, nil
}

func (Reducer) ImplicitAddToProperExplicitEnumTypeProperties(
	list *ast.TypePropertyList,
	newlines lr.TokenCount,
	property ast.TypeProperty,
) (
	*ast.TypePropertyList,
	error,
) {
	list.Add(property)
	return list, nil
}

func (Reducer) ImproperToExplicitEnumTypeProperties(
	list *ast.TypePropertyList,
	newlines lr.TokenCount,
) (
	*ast.TypePropertyList,
	error,
) {
	return list, nil
}
