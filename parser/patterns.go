package parser

import (
	. "github.com/pattyshack/pl/ast"
)

//
// VarPattern
//

type VarPatternReducerImpl struct{}

var _ DeclVarPatternReducer = VarPatternReducerImpl{}
var _ AssignVarPatternReducer = VarPatternReducerImpl{}

func (VarPatternReducerImpl) InferredToDeclVarPattern(
	varType *TokenValue,
	pattern Expression,
) (
	Expression,
	error,
) {
	return NewVarPattern(varType, pattern, nil), nil
}

func (VarPatternReducerImpl) TypedToDeclVarPattern(
	varType *TokenValue,
	pattern Expression,
	typeExpr TypeExpression,
) (
	Expression,
	error,
) {
	return NewVarPattern(varType, pattern, typeExpr), nil
}

func (VarPatternReducerImpl) ToAssignVarPattern(
	varType *TokenValue,
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

type CasePatternsReducerImpl struct{}

var _ CasePatternsReducer = &CasePatternsReducerImpl{}
var _ SwitchableCasePatternsReducer = &CasePatternsReducerImpl{}

func (CasePatternsReducerImpl) ToCasePatterns(
	pattern *CaseAssignPattern,
) (
	*ExpressionList,
	error,
) {
	list := NewExpressionList()
	list.Add(pattern)
	return list, nil
}

func (CasePatternsReducerImpl) SwitchableCasePatternToSwitchableCasePatterns(
	pattern Expression,
) (
	*ExpressionList,
	error,
) {
	list := NewExpressionList()
	list.Add(pattern)
	return list, nil
}

func (CasePatternsReducerImpl) AddToSwitchableCasePatterns(
	list *ExpressionList,
	comma *TokenValue,
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

type CaseAssignPatternReducerImpl struct{}

var _ CaseAssignPatternReducer = CaseAssignPatternReducerImpl{}
var _ CaseAssignExprReducer = CaseAssignPatternReducerImpl{}

func (CaseAssignPatternReducerImpl) ToCaseAssignPattern(
	assignPattern *ExpressionList,
	assign *TokenValue,
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

func (CaseAssignPatternReducerImpl) ToCaseAssignExpr(
	caseKW *TokenValue,
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

type CaseEnumPatternReducerImpl struct{}

var _ CaseEnumPatternReducer = &CaseEnumPatternReducerImpl{}

func (CaseEnumPatternReducerImpl) EnumMatchPatternToCaseEnumPattern(
	dot *TokenValue,
	enumValue *TokenValue,
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

func (CaseEnumPatternReducerImpl) EnumNondataMatchPattenToCaseEnumPattern(
	dot *TokenValue,
	enumValue *TokenValue,
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

func (CaseEnumPatternReducerImpl) EnumDeclVarPatternToCaseEnumPattern(
	varType *TokenValue,
	dot *TokenValue,
	enumValue *TokenValue,
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
