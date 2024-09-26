package ast

import (
	"fmt"
)

func elementsTreeString[T Node](
	list []T,
	indent string,
	elementType string,
) string {
	result := ""
	for idx, element := range list {
		result += "\n" + element.TreeString(
			indent,
			fmt.Sprintf("%s%d=", elementType, idx))
	}
	return result
}

func ListTreeString[T Node](
	list []T,
	indent string,
	label string,
	elementType string,
) string {
	result := fmt.Sprintf("%s%s[%sList:", indent, label, elementType)
	if len(list) == 0 {
		return result + "]"
	}

	result += elementsTreeString(list, indent+"  ", elementType)
	result += "\n" + indent + "]"
	return result
}

type NodeList[T Node] struct {
	StartEndPos

	// The leading / trailing comments from the list's start/end marker tokens.
	LeadingTrailingComments

	// MiddleComment is only used when the list is empty, but contains comments.
	MiddleComment CommentGroups

	Elements []T

	// Only used by TreeString
	ElementType string
}

func (list *NodeList[T]) Walk(visitor Visitor) {
	visitor.Enter(list)
	for _, element := range list.Elements {
		element.Walk(visitor)
	}
	visitor.Exit(list)
}

func (list *NodeList[T]) TreeString(indent string, label string) string {
	return ListTreeString(list.Elements, indent, label, list.ElementType)
}

func (list *NodeList[T]) Add(element T) {
	if len(list.Elements) == 0 {
		list.StartPos = element.Loc()
	}
	list.EndPos = element.End()
	list.Elements = append(list.Elements, element)
}

func (list *NodeList[T]) ReduceAdd(separator TokenValue, element T) {
	prev := list.Elements[len(list.Elements)-1]
	prev.AppendToTrailing(separator.TakeLeading())
	prev.AppendToTrailing(separator.TakeTrailing())

	list.Add(element)
}

func (list *NodeList[T]) ReduceImproper(separator TokenValue) {
	list.EndPos = separator.End()

	lastElement := list.Elements[len(list.Elements)-1]
	lastElement.AppendToTrailing(separator.TakeLeading())
	lastElement.AppendToTrailing(separator.TakeTrailing())
}

func (list *NodeList[T]) ReduceMarkers(start TokenValue, end TokenValue) {
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
	return &NodeList[Definition]{ElementType: "Definition"}
}

//
// ExpressionList
//

type ExpressionList = NodeList[Expression]

func NewExpressionList() *ExpressionList {
	return &NodeList[Expression]{ElementType: "Expression"}
}

//
// TypePropertyList
//

type TypePropertyList = NodeList[TypeProperty]

func NewTypePropertyList() *TypePropertyList {
	return &NodeList[TypeProperty]{ElementType: "TypeProperty"}
}

//
// TypeExpressionList
//

type TypeExpressionList = NodeList[TypeExpression]

func NewTypeExpressionList() *TypeExpressionList {
	return &NodeList[TypeExpression]{ElementType: "TypeExpression"}
}

//
// StatementList
//

type StatementList = NodeList[Statement]

func NewStatementList() *StatementList {
	return &NodeList[Statement]{ElementType: "Statement"}
}

//
// ParameterList
//

type ParameterList = NodeList[*Parameter]

func NewParameterList() *ParameterList {
	return &NodeList[*Parameter]{ElementType: "Parameter"}
}

//
// Argument list
//

type ArgumentList = NodeList[*Argument]

func NewArgumentList() *ArgumentList {
	return &NodeList[*Argument]{ElementType: "Argument"}
}

//
// GenericParameterList
//

type GenericParameterList = NodeList[*GenericParameter]

func NewGenericParameterList() *GenericParameterList {
	return &NodeList[*GenericParameter]{ElementType: "GenericParameter"}
}

//
// ImportClauseList
//

type ImportClauseList = NodeList[*ImportClause]

func NewImportClauseList() *ImportClauseList {
	return &NodeList[*ImportClause]{ElementType: "ImportClause"}
}

//
// ConditionBranchList
//

type ConditionBranchList = NodeList[*ConditionBranch]

func NewConditionBranchList() *ConditionBranchList {
	return &NodeList[*ConditionBranch]{ElementType: "ConditionBranch"}
}
