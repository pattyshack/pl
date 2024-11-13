package reducer

import (
	"github.com/pattyshack/gt/lexutil"

	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser/lr"
)

//
// ast.FieldDef
//

func (Reducer) unnamedField(
	typeExpr ast.TypeExpression,
) *ast.FieldDef {
	def := &ast.FieldDef{
		StartEndPos: lexutil.NewStartEndPos(typeExpr.Loc(), typeExpr.End()),
		Type:        typeExpr,
	}

	def.LeadingComment = typeExpr.TakeLeading()
	def.TrailingComment = typeExpr.TakeTrailing()

	return def
}

func (Reducer) namedField(
	name *lr.TokenValue,
	typeExpr ast.TypeExpression,
) *ast.FieldDef {
	def := &ast.FieldDef{
		StartEndPos: lexutil.NewStartEndPos(name.Loc(), typeExpr.End()),
		Name:        name.Value,
		Type:        typeExpr,
	}

	def.LeadingComment = name.TakeLeading()
	typeExpr.PrependToLeading(name.TakeTrailing())
	def.TrailingComment = typeExpr.TakeTrailing()

	return def
}

func (reducer Reducer) UnnamedFieldToTypeProperty(
	typeExpr ast.TypeExpression,
) (
	ast.TypeProperty,
	error,
) {
	sig, ok := typeExpr.(*ast.FuncSignature)
	if ok && sig.Name != "" {
		return sig, nil
	}

	return reducer.unnamedField(typeExpr), nil
}

func (reducer Reducer) NamedFieldToTypeProperty(
	name *lr.TokenValue,
	typeExpr ast.TypeExpression,
) (
	ast.TypeProperty,
	error,
) {
	return reducer.namedField(name, typeExpr), nil
}

func (reducer Reducer) PaddingFieldToTypeProperty(
	underscore *lr.TokenValue,
	typeExpr ast.TypeExpression,
) (
	ast.TypeProperty,
	error,
) {
	return reducer.namedField(underscore, typeExpr), nil
}

func (reducer Reducer) VarTypeUnnamedFieldToTypeProperty(
	varType *lr.TokenValue,
	typeExpr ast.TypeExpression,
) (
	ast.TypeProperty,
	error,
) {
	def := reducer.unnamedField(typeExpr)
	def.Qualifier = ast.FieldDefQualifier(varType.Value)

	def.PrependToLeading(varType.TakeTrailing())
	def.PrependToLeading(varType.TakeLeading())

	return def, nil
}

func (reducer Reducer) VarTypeNamedFieldToTypeProperty(
	varType *lr.TokenValue,
	name *lr.TokenValue,
	typeExpr ast.TypeExpression,
) (
	ast.TypeProperty,
	error,
) {
	def := reducer.namedField(name, typeExpr)
	def.Qualifier = ast.FieldDefQualifier(varType.Value)

	def.PrependToLeading(varType.TakeTrailing())
	def.PrependToLeading(varType.TakeLeading())

	return def, nil
}

func (reducer Reducer) DefaultNamedEnumFieldToTypeProperty(
	defaultKW *lr.TokenValue,
	name *lr.TokenValue,
	typeExpr ast.TypeExpression,
) (
	ast.TypeProperty,
	error,
) {
	def := reducer.namedField(name, typeExpr)
	def.Qualifier = ast.DefaultFieldDefQualifier

	def.PrependToLeading(defaultKW.TakeTrailing())
	def.PrependToLeading(defaultKW.TakeLeading())

	return def, nil
}

func (reducer Reducer) DefaultUnnamedEnumFieldToTypeProperty(
	defaultKW *lr.TokenValue,
	typeExpr ast.TypeExpression,
) (
	ast.TypeProperty,
	error,
) {
	def := reducer.unnamedField(typeExpr)
	def.Qualifier = ast.DefaultFieldDefQualifier

	def.PrependToLeading(defaultKW.TakeTrailing())
	def.PrependToLeading(defaultKW.TakeLeading())

	return def, nil
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

func (Reducer) ImproperImplicitToImplicitTypeProperties(
	list *ast.TypePropertyList,
	newlines lr.TokenCount,
) (
	*ast.TypePropertyList,
	error,
) {
	return list, nil
}

func (Reducer) ImproperExplicitToImplicitTypeProperties(
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
