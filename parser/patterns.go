package parser

import (
	"fmt"

	. "github.com/pattyshack/pl/ast"
)

//
// VarPattern
//

type VarPatternKind SymbolId

type VarPattern struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	Kind       VarPatternKind // either VAR, LET, or GREATER (assign to existing)
	VarPattern Expression
	Type       TypeExpression // optional
}

var _ Expression = &VarPattern{}

func NewVarPattern(
	varType *TokenValue,
	pattern Expression,
	typeExpr TypeExpression,
) *VarPattern {
	var end Node = pattern
	if typeExpr != nil {
		end = typeExpr
	}

	expr := &VarPattern{
		StartEndPos: NewStartEndPos(varType.Loc(), end.End()),
		Kind:        VarPatternKind(varType.SymbolId),
		VarPattern:  pattern,
		Type:        typeExpr,
	}

	expr.LeadingComment = varType.TakeLeading()
	expr.TrailingComment = end.TakeTrailing()

	return expr
}

func (expr VarPattern) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[VarPattern: Kind=%s\n",
		indent,
		label,
		SymbolId(expr.Kind))
	result += expr.VarPattern.TreeString(indent+"  ", "VarPattern=") + "\n"

	if expr.Type == nil {
		result += indent + "  TypeExpr=(nil)\n"
	} else {
		result += expr.Type.TreeString(indent+"  ", "TypeExpr=") + "\n"
	}

	result += indent + "]"
	return result
}

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

// REMINDER: In post analysis, ensure AssignPattern is valid (reject case enum
// pattern).
type CaseAssignPattern struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	AssignPattern ExpressionList
	Value         Expression
}

var _ Expression = &CaseAssignPattern{}

func (pattern CaseAssignPattern) TreeString(
	indent string,
	label string,
) string {
	result := fmt.Sprintf(
		"%s%s[CaseAssignPattern:\n",
		indent,
		label)
	result += pattern.AssignPattern.TreeString(indent+"  ", "AssignPattern=")
	result += "\n" + pattern.Value.TreeString(indent+"  ", "Value=")
	result += "\n" + indent + "]"
	return result
}

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

type CaseEnumPattern struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	EnumValue  string
	VarPattern Expression // optional.  either implicit struct or VarPattern
}

var _ Expression = &CaseEnumPattern{}

func (pattern CaseEnumPattern) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[CaseEnumPattern: EnumValue=%s",
		indent,
		label,
		pattern.EnumValue)
	if pattern.VarPattern != nil {
		result += "\n" + pattern.VarPattern.TreeString(indent+"  ", "VarPattern=")
		result += "\n" + indent + "]"
	} else {
		result += "]"
	}
	return result
}

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
