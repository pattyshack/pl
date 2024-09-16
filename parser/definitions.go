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
	list.reduceAdd(TokenValue{}, def)
	return list, nil
}

func (DefinitionListReducerImpl) DefinitionToProperDefinitions(
	def Definition,
) (
	*DefinitionList,
	error,
) {
	list := NewDefinitionList()
	list.add(def)
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
		StartEndPos: newStartEndPos(comments.Loc(), comments.End()),
	}
	floating.LeadingComment = comments
	return floating, nil
}
