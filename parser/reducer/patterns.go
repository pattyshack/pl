package reducer

import (
	. "github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser/lr"
)

//
// VarPattern
//

func (reducer *Reducer) InferredToDeclVarPattern(
	varType *lr.TokenValue,
	pattern Expression,
) (
	Expression,
	error,
) {
	return NewVarPattern(varType, pattern, nil), nil
}

func (Reducer) TypedToDeclVarPattern(
	varType *lr.TokenValue,
	pattern Expression,
	typeExpr TypeExpression,
) (
	Expression,
	error,
) {
	return NewVarPattern(varType, pattern, typeExpr), nil
}

func (Reducer) ToAssignVarPattern(
	varType *lr.TokenValue,
	pattern Expression,
) (
	Expression,
	error,
) {
	return NewVarPattern(varType, pattern, nil), nil
}

//
// ExpressionList
//

func (Reducer) ToCasePatterns(
	pattern *CaseAssignPattern,
) (
	*ExpressionList,
	error,
) {
	list := NewExpressionList()
	list.Add(pattern)
	return list, nil
}

func (Reducer) SwitchableCasePatternToSwitchableCasePatterns(
	pattern Expression,
) (
	*ExpressionList,
	error,
) {
	list := NewExpressionList()
	list.Add(pattern)
	return list, nil
}

func (Reducer) AddToSwitchableCasePatterns(
	list *ExpressionList,
	comma *lr.TokenValue,
	pattern Expression,
) (
	*ExpressionList,
	error,
) {
	list.ReduceAdd(comma, pattern)
	return list, nil
}

//
// CaseAssignPattern
//

func (Reducer) ToCaseAssignPattern(
	assignPattern *ExpressionList,
	assign *lr.TokenValue,
	value Expression,
) (
	*CaseAssignPattern,
	error,
) {
	leading := assignPattern.TakeLeading()
	assignPattern.AppendToTrailing(assign.TakeLeading())
	value.PrependToLeading(assign.TakeTrailing())

	pattern := &CaseAssignPattern{
		StartEndPos:   NewStartEndPos(assignPattern.Loc(), value.End()),
		AssignPattern: *assignPattern,
		Value:         value,
	}
	pattern.LeadingComment = leading
	pattern.TrailingComment = value.TakeTrailing()

	return pattern, nil
}

func (Reducer) ToCaseAssignExpr(
	caseKW *lr.TokenValue,
	pattern *CaseAssignPattern,
) (
	Expression,
	error,
) {
	pattern.PrependToLeading(caseKW.TakeTrailing())
	pattern.PrependToLeading(caseKW.TakeLeading())
	return pattern, nil
}

//
// CaseEnumPattern
//

func (Reducer) EnumMatchPatternToCaseEnumPattern(
	dot *lr.TokenValue,
	enumValue *lr.TokenValue,
	varPattern Expression,
) (
	Expression,
	error,
) {
	leading := dot.TakeLeading()
	leading.Append(dot.TakeTrailing())
	leading.Append(enumValue.TakeLeading())
	varPattern.PrependToLeading(enumValue.TakeTrailing())
	trailing := varPattern.TakeTrailing()

	pattern := &CaseEnumPattern{
		StartEndPos: NewStartEndPos(dot.Loc(), enumValue.End()),
		EnumValue:   enumValue.Value,
		VarPattern:  varPattern,
	}
	pattern.LeadingComment = leading
	pattern.TrailingComment = trailing
	return pattern, nil
}

func (Reducer) EnumNondataMatchPattenToCaseEnumPattern(
	dot *lr.TokenValue,
	enumValue *lr.TokenValue,
) (
	Expression,
	error,
) {
	leading := dot.TakeLeading()
	leading.Append(dot.TakeTrailing())
	leading.Append(enumValue.TakeLeading())
	trailing := enumValue.TakeTrailing()

	pattern := &CaseEnumPattern{
		StartEndPos: NewStartEndPos(dot.Loc(), enumValue.End()),
		EnumValue:   enumValue.Value,
	}
	pattern.LeadingComment = leading
	pattern.TrailingComment = trailing
	return pattern, nil
}

func (Reducer) EnumDeclVarPatternToCaseEnumPattern(
	varType *lr.TokenValue,
	dot *lr.TokenValue,
	enumValue *lr.TokenValue,
	tuplePattern Expression,
) (
	Expression,
	error,
) {
	leading := varType.TakeLeading()
	leading.Append(varType.TakeTrailing())
	leading.Append(dot.TakeLeading())
	leading.Append(dot.TakeTrailing())
	leading.Append(enumValue.TakeTrailing())

	varPattern := NewVarPattern(varType, tuplePattern, nil)
	varPattern.PrependToLeading(enumValue.TakeTrailing())

	trailing := tuplePattern.TakeTrailing()

	pattern := &CaseEnumPattern{
		StartEndPos: NewStartEndPos(varType.Loc(), enumValue.End()),
		EnumValue:   enumValue.Value,
		VarPattern:  varPattern,
	}
	pattern.LeadingComment = leading
	pattern.TrailingComment = trailing
	return pattern, nil
}
