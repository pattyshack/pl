package ast

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

func (pattern *VarPattern) Walk(visitor Visitor) {
	visitor.Enter(pattern)
	pattern.VarPattern.Walk(visitor)
	if pattern.Type != nil {
		pattern.Type.Walk(visitor)
	}
	visitor.Exit(pattern)
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

func (pattern *CaseAssignPattern) Walk(visitor Visitor) {
	visitor.Enter(pattern)
	pattern.AssignPattern.Walk(visitor)
	pattern.Value.Walk(visitor)
	visitor.Exit(pattern)
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

func (pattern *CaseEnumPattern) Walk(visitor Visitor) {
	visitor.Enter(pattern)
	if pattern.VarPattern != nil {
		pattern.VarPattern.Walk(visitor)
	}
	visitor.Exit(pattern)
}
