package reducer

import (
	"fmt"

	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser/lr"
)

//
// DefinitionList
//

func (reducer *Reducer) AddToProperDefinitions(
	list *ast.DefinitionList,
	newlines lr.TokenCount,
	def ast.Definition,
) (
	*ast.DefinitionList,
	error,
) {
	list.ReduceAdd(&lr.TokenValue{}, def)
	return list, nil
}

func (reducer *Reducer) DefinitionToProperDefinitions(
	def ast.Definition,
) (
	*ast.DefinitionList,
	error,
) {
	list := ast.NewDefinitionList()
	list.Add(def)
	return list, nil
}

func (reducer *Reducer) ImproperToDefinitions(
	list *ast.DefinitionList,
	newlines lr.TokenCount,
) (
	*ast.DefinitionList,
	error,
) {
	return list, nil
}

func (reducer *Reducer) NilToDefinitions() (*ast.DefinitionList, error) {
	return ast.NewDefinitionList(), nil
}

//
// FloatingComment
//

func (reducer *Reducer) ToFloatingComment(
	comments lr.CommentGroupsTok,
) (
	ast.Definition,
	error,
) {
	floating := &ast.FloatingComment{
		StartEndPos: ast.NewStartEndPos(comments.Loc(), comments.End()),
	}
	floating.LeadingComment = comments.CommentGroups
	return floating, nil
}

//
// PackageDef
//

func (reducer *Reducer) ToPackageDef(
	pkg *lr.TokenValue,
	expr ast.Expression,
) (
	ast.Definition,
	error,
) {
	switch body := expr.(type) {
	case *ast.StatementsExpr:
		body.PrependToLeading(pkg.TakeTrailing())
		def := &ast.PackageDef{
			StartEndPos: ast.NewStartEndPos(pkg.Loc(), expr.End()),
			Body:        body,
		}
		def.LeadingComment = pkg.TakeLeading()

		reducer.PackageDefs = append(reducer.PackageDefs, def)
		return def, nil
	case *ast.ParseErrorNode:
		return expr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", expr))
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
	ast.Definition,
	error,
) {
	leading := typeKW.TakeLeading()
	leading.Append(typeKW.TakeTrailing())
	leading.Append(name.TakeLeading())

	var params []*ast.GenericParameter
	if genericParameters != nil {
		params = genericParameters.Elements

		leading.Append(genericParameters.TakeLeading())

		baseType.PrependToLeading(genericParameters.TakeTrailing())
		if len(genericParameters.MiddleComment.Groups) > 0 {
			baseType.PrependToLeading(genericParameters.MiddleComment)
		}
	}

	trailing := baseType.TakeTrailing()

	def := &ast.TypeDef{
		StartEndPos:       ast.NewStartEndPos(typeKW.Loc(), baseType.End()),
		Name:              name.Value,
		GenericParameters: params,
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
	ast.Definition,
	error,
) {
	leading := typeKW.TakeLeading()
	leading.Append(typeKW.TakeTrailing())
	leading.Append(name.TakeLeading())

	var params []*ast.GenericParameter
	if genericParameters != nil {
		params = genericParameters.Elements

		leading.Append(genericParameters.TakeLeading())

		baseType.PrependToLeading(genericParameters.TakeTrailing())
		if len(genericParameters.MiddleComment.Groups) > 0 {
			baseType.PrependToLeading(genericParameters.MiddleComment)
		}
	}

	baseType.AppendToTrailing(implements.TakeLeading())
	constraint.PrependToLeading(implements.TakeTrailing())

	trailing := constraint.TakeTrailing()

	def := &ast.TypeDef{
		StartEndPos:       ast.NewStartEndPos(typeKW.Loc(), constraint.End()),
		Name:              name.Value,
		GenericParameters: params,
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
	ast.Definition,
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
