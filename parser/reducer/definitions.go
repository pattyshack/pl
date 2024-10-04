package reducer

import (
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
		StartEndPos: ast.NewStartEndPos(comments.Loc(), comments.End()),
	}
	floating.LeadingComment = comments.CommentGroups
	return floating, nil
}

//
// TypeDef
//

func (reducer *Reducer) DefinitionToTypeDef(
	typeKW *lr.TokenValue,
	name *lr.TokenValue,
	genericParameters *ast.GenericParameterList,
	baseType ast.TypeExpression,
) (
	ast.Statement,
	error,
) {
	leading := typeKW.TakeLeading()
	leading.Append(typeKW.TakeTrailing())
	leading.Append(name.TakeLeading())

	if genericParameters != nil {
		genericParameters.PrependToLeading(name.TakeTrailing())
	} else {
		baseType.PrependToLeading(name.TakeTrailing())
	}

	trailing := baseType.TakeTrailing()

	def := &ast.TypeDef{
		StartEndPos:       ast.NewStartEndPos(typeKW.Loc(), baseType.End()),
		Name:              name.Value,
		GenericParameters: genericParameters,
		BaseType:          baseType,
	}
	def.LeadingComment = leading
	def.TrailingComment = trailing

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
	leading := typeKW.TakeLeading()
	leading.Append(typeKW.TakeTrailing())
	leading.Append(name.TakeLeading())

	if genericParameters != nil {
		genericParameters.PrependToLeading(name.TakeTrailing())
	} else {
		baseType.PrependToLeading(name.TakeTrailing())
	}

	baseType.AppendToTrailing(implements.TakeLeading())
	constraint.PrependToLeading(implements.TakeTrailing())

	trailing := constraint.TakeTrailing()

	def := &ast.TypeDef{
		StartEndPos:       ast.NewStartEndPos(typeKW.Loc(), constraint.End()),
		Name:              name.Value,
		GenericParameters: genericParameters,
		BaseType:          baseType,
		Constraint:        constraint,
	}
	def.LeadingComment = leading
	def.TrailingComment = trailing

	return def, nil
}

func (reducer *Reducer) AliasToTypeDef(
	typeKW *lr.TokenValue,
	name *lr.TokenValue,
	assign *lr.TokenValue,
	baseType ast.TypeExpression,
) (
	ast.Statement,
	error,
) {
	leading := typeKW.TakeLeading()
	leading.Append(typeKW.TakeTrailing())
	leading.Append(name.TakeLeading())

	baseType.PrependToLeading(assign.TakeTrailing())
	baseType.PrependToLeading(assign.TakeLeading())
	baseType.PrependToLeading(name.TakeTrailing())

	trailing := baseType.TakeTrailing()

	def := &ast.TypeDef{
		StartEndPos: ast.NewStartEndPos(typeKW.Loc(), baseType.End()),
		Name:        name.Value,
		IsAlias:     true,
		BaseType:    baseType,
	}
	def.LeadingComment = leading
	def.TrailingComment = trailing

	return def, nil
}
