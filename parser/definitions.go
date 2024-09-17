package parser

import (
	"fmt"
)

//
// DefinitionList
//

type DefinitionList = NodeList[Definition]

func NewDefinitionList() *DefinitionList {
	return newNodeList[Definition]("DefinitionList")
}

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
	list.ReduceAdd(TokenValue{}, def)
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

type FloatingComment struct {
	isDefinition
	StartEndPos
	LeadingTrailingComments
}

var _ Definition = &FloatingComment{}

func (FloatingComment) TreeString(indent string, label string) string {
	return fmt.Sprintf("%s%s[FloatingComment]", indent, label)
}

type FloatingCommentReducerImpl struct{}

var _ FloatingCommentReducer = &FloatingCommentReducerImpl{}

func (FloatingCommentReducerImpl) ToFloatingComment(
	comments CommentGroups,
) (
	Definition,
	error,
) {
	floating := &FloatingComment{
		StartEndPos: NewStartEndPos(comments.Loc(), comments.End()),
	}
	floating.LeadingComment = comments
	return floating, nil
}

//
// PackageDef
//

type PackageDef struct {
	isDefinition
	StartEndPos
	LeadingTrailingComments

	Body StatementsExpr
}

var _ Definition = &PackageDef{}

func (def PackageDef) TreeString(indent string, label string) string {
	result := fmt.Sprintf("%s%s[PackageDef:\n", indent, label)
	result += def.Body.TreeString(indent+"  ", "Body=")
	result += "\n" + indent + "]"
	return result
}

type PackageDefReducerImpl struct{}

var _ PackageDefReducer = &PackageDefReducerImpl{}

func (PackageDefReducerImpl) ToPackageDef(
	pkg TokenValue,
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

type TypeDef struct {
	isDefinition
	StartEndPos
	LeadingTrailingComments

	Name              string
	IsAlias           bool
	GenericParameters GenericParameterList // optional
	BaseType          TypeExpression
	Constraint        TypeExpression // optional
}

var _ Definition = &TypeDef{}

func (def TypeDef) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[TypeDef: Name=%s IsAlias=%v\n",
		indent,
		label,
		def.Name,
		def.IsAlias)
	result += "\n" + indent + "]"
	return result
}

type TypeDefReducerImpl struct {
	TypeDefs []*TypeDef
}

var _ TypeDefReducer = &TypeDefReducerImpl{}

func (reducer *TypeDefReducerImpl) DefinitionToTypeDef(
	typeKW TokenValue,
	name TokenValue,
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
	typeKW TokenValue,
	name TokenValue,
	genericParameters *GenericParameterList,
	baseType TypeExpression,
	implements TokenValue,
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
	typeKW TokenValue,
	name TokenValue,
	assign TokenValue,
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
