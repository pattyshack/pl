package ast

type NodeList[T Node] struct {
	StartEndPos

	// The leading / trailing comments from the list's start/end marker tokens.
	LeadingTrailingComments

	// MiddleComment is only used when the list is empty, but contains comments.
	MiddleComment CommentGroups

	Elements []T
}

func (list *NodeList[T]) Walk(visitor Visitor) {
	visitor.Enter(list)
	for _, element := range list.Elements {
		element.Walk(visitor)
	}
	visitor.Exit(list)
}

func (list *NodeList[T]) Add(element T) {
	if len(list.Elements) == 0 {
		list.StartPos = element.Loc()
	}
	list.EndPos = element.End()
	list.Elements = append(list.Elements, element)
}

func ReduceAdd[T Node](list []T, seperator TokenValue, element T) []T {
	prev := list[len(list)-1]
	prev.AppendToTrailing(seperator.TakeLeading())
	prev.AppendToTrailing(seperator.TakeTrailing())

	return append(list, element)
}

func (list *NodeList[T]) ReduceAdd(separator TokenValue, element T) {
	list.Elements = ReduceAdd(list.Elements, separator, element)
	list.EndPos = element.End()
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
// StatementList
//

// NOTE: StatementList is only used for top-level statements.  StatementExpr
// uses []Statement directly.
type StatementList = NodeList[Statement]

func NewStatementList() *StatementList {
	return &NodeList[Statement]{}
}

//
// TypeExpressionList
//

type TypeExpressionList = NodeList[TypeExpression]

func NewTypeExpressionList() *TypeExpressionList {
	return &NodeList[TypeExpression]{}
}

//
// ParameterList
//

type ParameterList = NodeList[*Parameter]

func NewParameterList() *ParameterList {
	return &NodeList[*Parameter]{}
}

//
// GenericParameterList
//

type GenericParameterList = NodeList[*GenericParameter]

func NewGenericParameterList() *GenericParameterList {
	return &NodeList[*GenericParameter]{}
}

//
// Intermediate representations only used during parsing.  Not real ast list
//

type TypePropertyList = NodeList[TypeProperty]

func NewTypePropertyList() *TypePropertyList {
	return &NodeList[TypeProperty]{}
}

type ConditionBranchStmtList = NodeList[*ConditionBranchStmt]

func NewConditionBranchStmtList() *ConditionBranchStmtList {
	return &NodeList[*ConditionBranchStmt]{}
}

type ImportClauseList = NodeList[*ImportClause]

func NewImportClauseList() *ImportClauseList {
	return &NodeList[*ImportClause]{}
}

type ArgumentList = NodeList[*Argument]

func NewArgumentList() *ArgumentList {
	return &NodeList[*Argument]{}
}

type ExpressionList = NodeList[Expression]

func NewExpressionList() *ExpressionList {
	return &NodeList[Expression]{}
}
