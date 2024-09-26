package reducer

import (
	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser/lr"
)

//
// VarPattern
//

func (reducer *Reducer) InferredToDeclVarPattern(
	varType *lr.TokenValue,
	pattern ast.Expression,
) (
	ast.Expression,
	error,
) {
	return ast.NewVarPattern(varType, pattern, nil), nil
}

func (Reducer) TypedToDeclVarPattern(
	varType *lr.TokenValue,
	pattern ast.Expression,
	typeExpr ast.TypeExpression,
) (
	ast.Expression,
	error,
) {
	return ast.NewVarPattern(varType, pattern, typeExpr), nil
}

//
// ExpressionList
//

func (Reducer) ToCasePatterns(
	pattern *ast.CaseAssignPattern,
) (
	*ast.ExpressionList,
	error,
) {
	list := ast.NewExpressionList()
	list.Add(pattern)
	return list, nil
}

func (Reducer) SwitchableCasePatternToSwitchableCasePatterns(
	pattern ast.Expression,
) (
	*ast.ExpressionList,
	error,
) {
	list := ast.NewExpressionList()
	list.Add(pattern)
	return list, nil
}

func (Reducer) AddToSwitchableCasePatterns(
	list *ast.ExpressionList,
	comma *lr.TokenValue,
	pattern ast.Expression,
) (
	*ast.ExpressionList,
	error,
) {
	list.ReduceAdd(comma, pattern)
	return list, nil
}

//
// CaseAssignPattern
//

func (Reducer) ToCaseAssignPattern(
	assignPatterns *ast.ExpressionList,
	assign *lr.TokenValue,
	value ast.Expression,
) (
	*ast.CaseAssignPattern,
	error,
) {
	leading := assignPatterns.TakeLeading()

	assignPatterns.AppendToTrailing(assign.TakeLeading())
	value.PrependToLeading(assign.TakeTrailing())

	trailing := value.TakeTrailing()

	pattern := &ast.CaseAssignPattern{
		StartEndPos:   ast.NewStartEndPos(assignPatterns.Loc(), value.End()),
		AssignPattern: assignPatterns,
		Value:         value,
	}
	pattern.LeadingComment = leading
	pattern.TrailingComment = trailing

	return pattern, nil
}

func (Reducer) ToCaseAssignExpr(
	caseKW *lr.TokenValue,
	pattern *ast.CaseAssignPattern,
) (
	ast.Expression,
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
	varPattern ast.Expression,
) (
	ast.Expression,
	error,
) {
	leading := dot.TakeLeading()
	leading.Append(dot.TakeTrailing())
	leading.Append(enumValue.TakeLeading())
	varPattern.PrependToLeading(enumValue.TakeTrailing())
	trailing := varPattern.TakeTrailing()

	pattern := &ast.CaseEnumPattern{
		StartEndPos: ast.NewStartEndPos(dot.Loc(), enumValue.End()),
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
	ast.Expression,
	error,
) {
	leading := dot.TakeLeading()
	leading.Append(dot.TakeTrailing())
	leading.Append(enumValue.TakeLeading())
	trailing := enumValue.TakeTrailing()

	pattern := &ast.CaseEnumPattern{
		StartEndPos: ast.NewStartEndPos(dot.Loc(), enumValue.End()),
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
	tuplePattern ast.Expression,
) (
	ast.Expression,
	error,
) {
	leading := varType.TakeLeading()
	leading.Append(varType.TakeTrailing())
	leading.Append(dot.TakeLeading())
	leading.Append(dot.TakeTrailing())
	leading.Append(enumValue.TakeTrailing())

	varPattern := ast.NewVarPattern(varType, tuplePattern, nil)
	varPattern.PrependToLeading(enumValue.TakeTrailing())

	trailing := tuplePattern.TakeTrailing()

	pattern := &ast.CaseEnumPattern{
		StartEndPos: ast.NewStartEndPos(varType.Loc(), enumValue.End()),
		EnumValue:   enumValue.Value,
		VarPattern:  varPattern,
	}
	pattern.LeadingComment = leading
	pattern.TrailingComment = trailing
	return pattern, nil
}
