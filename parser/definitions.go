package parser

import (
	"fmt"

	. "github.com/pattyshack/pl/ast"
)

//
// DefinitionList
//

type DefinitionListReducerImpl struct{}

var _ ProperDefinitionsReducer = &DefinitionListReducerImpl{}
var _ DefinitionsReducer = &DefinitionListReducerImpl{}

func (DefinitionListReducerImpl) AddToProperDefinitions(
	list *DefinitionList,
	newlines TokenCount,
	def Definition,
) (
	*DefinitionList,
	error,
) {
	list.ReduceAdd(&TokenValue{}, def)
	return list, nil
}

func (DefinitionListReducerImpl) DefinitionToProperDefinitions(
	def Definition,
) (
	*DefinitionList,
	error,
) {
	list := NewDefinitionList()
	list.Add(def)
	return list, nil
}

func (DefinitionListReducerImpl) ImproperToDefinitions(
	list *DefinitionList,
	newlines TokenCount,
) (
	*DefinitionList,
	error,
) {
	return list, nil
}

func (DefinitionListReducerImpl) NilToDefinitions() (*DefinitionList, error) {
	return NewDefinitionList(), nil
}

//
// FloatingComment
//

type FloatingCommentReducerImpl struct{}

var _ FloatingCommentReducer = &FloatingCommentReducerImpl{}

func (FloatingCommentReducerImpl) ToFloatingComment(
	comments CommentGroupsTok,
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

type PackageDefReducerImpl struct{}

var _ PackageDefReducer = &PackageDefReducerImpl{}

func (PackageDefReducerImpl) ToPackageDef(
	pkg *TokenValue,
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
			Body:        *body,
		}
		def.LeadingComment = pkg.TakeLeading()

		return def, nil
	case *ParseErrorSymbol:
		return expr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", expr))
}

//
// TypeDef
//

type TypeDefReducerImpl struct {
	TypeDefs []*TypeDef
}

var _ TypeDefReducer = &TypeDefReducerImpl{}

func (reducer *TypeDefReducerImpl) DefinitionToTypeDef(
	typeKW *TokenValue,
	name *TokenValue,
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

func (reducer *TypeDefReducerImpl) ConstrainedDefToTypeDef(
	typeKW *TokenValue,
	name *TokenValue,
	genericParameters *GenericParameterList,
	baseType TypeExpression,
	implements *TokenValue,
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

func (reducer *TypeDefReducerImpl) AliasToTypeDef(
	typeKW *TokenValue,
	name *TokenValue,
	assign *TokenValue,
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
