package ast

import (
	"fmt"
)

type NodeList[T Node] struct {
	StartEndPos

	// The leading / trailing comments from the list's start/end marker tokens.
	LeadingTrailingComments

	// MiddleComment is only used when the list is empty, but contains comments.
	MiddleComment CommentGroups

	Elements []T

	// Only used for TreeString() / String()
	ListType string
}

func (list NodeList[T]) ElementsString(indent string) string {
	result := ""
	for idx, element := range list.Elements {
		result += "\n" + element.TreeString(
			indent,
			fmt.Sprintf("Element%d=", idx))
	}
	return result
}

func (list NodeList[T]) TreeString(indent string, label string) string {
	result := fmt.Sprintf("%s%s[%s:", indent, label, list.ListType)
	if len(list.Elements) == 0 {
		return result + "]"
	}

	result += list.ElementsString(indent + "  ")
	result += "\n" + indent + "]"
	return result
}

func NewNodeList[T Node](listType string) *NodeList[T] {
	return &NodeList[T]{
		ListType: listType,
	}
}

func (list *NodeList[T]) Add(element T) {
	if len(list.Elements) == 0 {
		list.StartPos = element.Loc()
	}
	list.EndPos = element.End()
	list.Elements = append(list.Elements, element)
}

func (list *NodeList[T]) ReduceAdd(separator ValueNode, element T) {
	prev := list.Elements[len(list.Elements)-1]
	prev.AppendToTrailing(separator.TakeLeading())
	prev.AppendToTrailing(separator.TakeTrailing())

	list.Add(element)
}

func (list *NodeList[T]) ReduceImproper(separator ValueNode) {
	list.EndPos = separator.End()

	lastElement := list.Elements[len(list.Elements)-1]
	lastElement.AppendToTrailing(separator.TakeLeading())
	lastElement.AppendToTrailing(separator.TakeTrailing())
}

func (list *NodeList[T]) ReduceMarkers(start ValueNode, end ValueNode) {
	list.StartPos = start.Loc()
	list.EndPos = end.End()

	list.LeadingComment = start.TakeLeading()
	if len(list.Elements) > 0 {
		list.Elements[0].PrependToLeading(start.TakeTrailing())
		list.Elements[len(list.Elements)-1].AppendToTrailing(end.TakeLeading())
	} else {
		list.MiddleComment = start.TakeTrailing()
		list.MiddleComment.Append(end.TakeLeading())
	}
	list.TrailingComment = end.TakeTrailing()
}

//
// DefinitionList
//

type DefinitionList = NodeList[Definition]

func NewDefinitionList() *DefinitionList {
	return NewNodeList[Definition]("DefinitionList")
}

//
// ExpressionList
//

type ExpressionList = NodeList[Expression]

func NewExpressionList() *ExpressionList {
	return NewNodeList[Expression]("ExpressionList")
}

//
// TypePropertyList
//

type TypePropertyList = NodeList[TypeProperty]

func NewTypePropertyList() *TypePropertyList {
	return NewNodeList[TypeProperty]("TypePropertyList")
}

//
// GenericArgumentList
//

type GenericArgumentList = NodeList[TypeExpression]

func NewGenericArgumentList() *GenericArgumentList {
	return NewNodeList[TypeExpression]("GenericArgumentList")
}

//
// ParameterList
//

type ParameterList = NodeList[*Parameter]

func NewParameterList() *ParameterList {
	return NewNodeList[*Parameter]("ParameterList")
}

//
// Argument list
//

type ArgumentList = NodeList[*Argument]

func NewArgumentList() *ArgumentList {
	return NewNodeList[*Argument]("ArgumentList")
}

//
// GenericParameterList
//

type GenericParameterList = NodeList[*GenericParameter]

func NewGenericParameterList() *GenericParameterList {
	return NewNodeList[*GenericParameter]("GenericParameterList")
}
