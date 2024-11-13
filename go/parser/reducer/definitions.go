package reducer

import (
	"github.com/pattyshack/gt/lexutil"

	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser/lr"
)

//
// FloatingComment
//

func (reducer *Reducer) ToFloatingComment(
	comments lr.CommentGroupsTok,
) (
	ast.Statement,
	error,
) {
	floating := &ast.FloatingComment{
		StartEndPos: lexutil.NewStartEndPos(comments.Loc(), comments.End()),
	}
	floating.LeadingComment = comments.CommentGroups
	return floating, nil
}

//
// TypeDef
//

func (reducer *Reducer) toTypeDef(
	typeKW *lr.TokenValue,
	name *lr.TokenValue,
	genericParameters *ast.GenericParameterList,
	baseType ast.TypeExpression,
	constraint ast.TypeExpression,
) *ast.TypeDef {
	leading := typeKW.TakeLeading()
	leading.Append(typeKW.TakeTrailing())
	leading.Append(name.TakeLeading())

	if genericParameters != nil {
		genericParameters.PrependToLeading(name.TakeTrailing())
	} else {
		genericParameters = ast.NewImplicitGenericParameterList(name.End())
		baseType.PrependToLeading(name.TakeTrailing())
	}

	def := &ast.TypeDef{
		StartEndPos:       lexutil.NewStartEndPos(typeKW.Loc(), baseType.End()),
		Name:              name.Value,
		GenericParameters: genericParameters,
		BaseType:          baseType,
		Constraint:        constraint,
	}
	def.LeadingComment = leading

	return def
}

func (reducer *Reducer) DefinitionToTypeDef(
	typeKW *lr.TokenValue,
	name *lr.TokenValue,
	genericParameters *ast.GenericParameterList,
	baseType ast.TypeExpression,
) (
	ast.Statement,
	error,
) {
	def := reducer.toTypeDef(
		typeKW,
		name,
		genericParameters,
		baseType,
		ast.NewAnyTrait(baseType.End()))
	def.TrailingComment = baseType.TakeTrailing()

	return def, nil
}

func (reducer *Reducer) ConstrainedDefToTypeDef(
	typeKW *lr.TokenValue,
	name *lr.TokenValue,
	genericParameters *ast.GenericParameterList,
	baseType ast.TypeExpression,
	implements *lr.TokenValue,
	constraint ast.TypeExpression,
) (
	ast.Statement,
	error,
) {
	baseType.AppendToTrailing(implements.TakeLeading())
	constraint.PrependToLeading(implements.TakeTrailing())

	def := reducer.toTypeDef(
		typeKW,
		name,
		genericParameters,
		baseType,
		constraint)
	def.TrailingComment = constraint.TakeTrailing()

	return def, nil
}

func (reducer *Reducer) ToAliasDef(
	aliasKW *lr.TokenValue,
	alias *lr.TokenValue,
	genericParameters *ast.GenericParameterList,
	value ast.TypeExpression,
) (
	ast.Statement,
	error,
) {
	leading := aliasKW.TakeLeading()
	leading.Append(aliasKW.TakeTrailing())
	leading.Append(alias.TakeLeading())

	if genericParameters != nil {
		genericParameters.PrependToLeading(alias.TakeTrailing())
	} else {
		genericParameters = ast.NewImplicitGenericParameterList(alias.End())

		value.PrependToLeading(alias.TakeTrailing())
	}

	def := &ast.AliasDef{
		StartEndPos:       lexutil.NewStartEndPos(aliasKW.Loc(), value.End()),
		Alias:             alias.Value,
		GenericParameters: genericParameters,
		Value:             value,
	}
	def.LeadingComment = leading
	def.TrailingComment = value.TakeTrailing()

	return def, nil
}
