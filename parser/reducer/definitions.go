package reducer

import (
	"fmt"

	. "github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser/lr"
)

//
// DefinitionList
//

func (reducer *Reducer) AddToProperDefinitions(
	list *DefinitionList,
	newlines lr.TokenCount,
	def Definition,
) (
	*DefinitionList,
	error,
) {
	list.ReduceAdd(&lr.TokenValue{}, def)
	return list, nil
}

func (reducer *Reducer) DefinitionToProperDefinitions(
	def Definition,
) (
	*DefinitionList,
	error,
) {
	list := NewDefinitionList()
	list.Add(def)
	return list, nil
}

func (reducer *Reducer) ImproperToDefinitions(
	list *DefinitionList,
	newlines lr.TokenCount,
) (
	*DefinitionList,
	error,
) {
	return list, nil
}

func (reducer *Reducer) NilToDefinitions() (*DefinitionList, error) {
	return NewDefinitionList(), nil
}

//
// FloatingComment
//

func (reducer *Reducer) ToFloatingComment(
	comments lr.CommentGroupsTok,
) (
	Definition,
	error,
) {
	floating := &FloatingComment{
		StartEndPos: NewStartEndPos(comments.Loc(), comments.End()),
	}
	floating.LeadingComment = comments.CommentGroups
	return floating, nil
}

//
// PackageDef
//

func (reducer *Reducer) ToPackageDef(
	pkg *lr.TokenValue,
	expr Expression,
) (
	Definition,
	error,
) {
	switch body := expr.(type) {
	case *StatementsExpr:
		body.PrependToLeading(pkg.TakeTrailing())
		def := &PackageDef{
			StartEndPos: NewStartEndPos(pkg.Loc(), expr.End()),
			Body:        body,
		}
		def.LeadingComment = pkg.TakeLeading()

		return def, nil
	case *ParseErrorNode:
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
	genericParameters *GenericParameterList,
	baseType TypeExpression,
) (
	Definition,
	error,
) {
	leading := typeKW.TakeLeading()
	leading.Append(typeKW.TakeTrailing())
	leading.Append(name.TakeLeading())

	if genericParameters != nil {
		genericParameters.PrependToLeading(name.TakeTrailing())
	} else {
		genericParameters = NewGenericParameterList()
		baseType.PrependToLeading(name.TakeTrailing())
	}

	trailing := baseType.TakeTrailing()

	def := &TypeDef{
		StartEndPos:       NewStartEndPos(typeKW.Loc(), baseType.End()),
		Name:              name.Value,
		GenericParameters: *genericParameters,
		BaseType:          baseType,
	}
	def.LeadingComment = leading
	def.TrailingComment = trailing

	return def, nil
}

func (reducer *Reducer) ConstrainedDefToTypeDef(
	typeKW *lr.TokenValue,
	name *lr.TokenValue,
	genericParameters *GenericParameterList,
	baseType TypeExpression,
	implements *lr.TokenValue,
	constraint TypeExpression,
) (
	Definition,
	error,
) {
	leading := typeKW.TakeLeading()
	leading.Append(typeKW.TakeTrailing())
	leading.Append(name.TakeLeading())

	if genericParameters != nil {
		genericParameters.PrependToLeading(name.TakeTrailing())
	} else {
		genericParameters = NewGenericParameterList()
		baseType.PrependToLeading(name.TakeTrailing())
	}

	baseType.AppendToTrailing(implements.TakeLeading())
	constraint.PrependToLeading(implements.TakeTrailing())

	trailing := constraint.TakeTrailing()

	def := &TypeDef{
		StartEndPos:       NewStartEndPos(typeKW.Loc(), constraint.End()),
		Name:              name.Value,
		GenericParameters: *genericParameters,
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
	baseType TypeExpression,
) (
	Definition,
	error,
) {
	leading := typeKW.TakeLeading()
	leading.Append(typeKW.TakeTrailing())
	leading.Append(name.TakeLeading())

	baseType.PrependToLeading(assign.TakeTrailing())
	baseType.PrependToLeading(assign.TakeLeading())
	baseType.PrependToLeading(name.TakeTrailing())

	trailing := baseType.TakeTrailing()

	def := &TypeDef{
		StartEndPos: NewStartEndPos(typeKW.Loc(), baseType.End()),
		Name:        name.Value,
		IsAlias:     true,
		BaseType:    baseType,
	}
	def.LeadingComment = leading
	def.TrailingComment = trailing

	return def, nil
}
