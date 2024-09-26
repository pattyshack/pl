package ast

import (
	"fmt"
)

//
// VarPattern
//

type VarPatternKind string

const (
	VarDecl     = VarPatternKind("var")
	LetDecl     = VarPatternKind("let")
	AssignToVar = VarPatternKind(">")
)

type VarPattern struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	Kind       VarPatternKind
	VarPattern Expression
	Type       TypeExpression // optional
}

var _ Expression = &VarPattern{}

func NewVarPattern(
	varType TokenValue,
	pattern Expression,
	typeExpr TypeExpression,
) *VarPattern {
	var end Node = pattern
	if typeExpr != nil {
		end = typeExpr
	}

	expr := &VarPattern{
		StartEndPos: NewStartEndPos(varType.Loc(), end.End()),
		Kind:        VarPatternKind(varType.Val()),
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
		expr.Kind)
	result += expr.VarPattern.TreeString(indent+"  ", "VarPattern=") + "\n"

	if expr.Type == nil {
		result += indent + "  TypeExpr=(nil)\n"
	} else {
		result += expr.Type.TreeString(indent+"  ", "TypeExpr=") + "\n"
	}

	result += indent + "]"
	return result
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

	AssignPattern *ExpressionList
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
	result += pattern.AssignPattern.TreeString(
		indent+"  ",
		"AssignPattern=")
	result += "\n" + pattern.Value.TreeString(indent+"  ", "Value=")
	result += "\n" + indent + "]"
	return result
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
