package parser

import (
	"fmt"
)

//
// VarPattern
//

type VarPatternKind SymbolId

type VarPattern struct {
	isExpression
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
// CasePatternList
//

type CasePatternList = NodeList[CasePattern]

func NewCasePatternList() *CasePatternList {
	return newNodeList[CasePattern]("CasePatternList")
}

var _ Node = &CasePatternList{}

type CasePatternsReducerImpl struct{}

var _ CasePatternsReducer = &CasePatternsReducerImpl{}
var _ SwitchableCasePatternsReducer = &CasePatternsReducerImpl{}

func (CasePatternsReducerImpl) ToCasePatterns(
	pattern *CaseAssignPattern,
) (
	*CasePatternList,
	error,
) {
	list := NewCasePatternList()
	list.Add(pattern)
	return list, nil
}

func (CasePatternsReducerImpl) SwitchableCasePatternToSwitchableCasePatterns(
	pattern CasePattern,
) (
	*CasePatternList,
	error,
) {
	list := NewCasePatternList()
	list.Add(pattern)
	return list, nil
}

func (CasePatternsReducerImpl) AddToSwitchableCasePatterns(
	list *CasePatternList,
	comma *TokenValue,
	pattern CasePattern,
) (
	*CasePatternList,
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
	isExpression
	StartEndPos
	LeadingTrailingComments

	AssignPattern CasePatternList
	Value         Expression
}

var _ CasePattern = &CaseAssignPattern{}

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
	assignPattern *CasePatternList,
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
	isCasePattern
	StartEndPos
	LeadingTrailingComments

	EnumValue  string
	VarPattern Expression // optional.  either implicit struct or VarPattern
}

var _ CasePattern = &CaseEnumPattern{}

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
	CasePattern,
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
	CasePattern,
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
	CasePattern,
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
