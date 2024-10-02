package ast

//
// AssignPattern
//

type AssignKind string

const (
	EqualAssign = AssignKind("=")
	InAssign    = AssignKind("in")
)

type AssignPattern struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	Kind AssignKind

	Pattern Expression
	Value   Expression
}

var _ Expression = &AssignPattern{}

func (assign *AssignPattern) Walk(visitor Visitor) {
	visitor.Enter(assign)
	assign.Pattern.Walk(visitor)
	assign.Value.Walk(visitor)
	visitor.Exit(assign)
}

//
// CasePatterns
//

type CasePatterns struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	Patterns []Expression
}

var _ Expression = &CasePatterns{}

func (cond *CasePatterns) Walk(visitor Visitor) {
	visitor.Enter(cond)
	for _, pattern := range cond.Patterns {
		pattern.Walk(visitor)
	}
	visitor.Exit(cond)
}

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
// EnumPattern
//

type EnumPattern struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	EnumValue string
	Pattern   Expression // optional.  either implicit struct or VarPattern
}

var _ Expression = &EnumPattern{}

func (pattern *EnumPattern) Walk(visitor Visitor) {
	visitor.Enter(pattern)
	if pattern.Pattern != nil {
		pattern.Pattern.Walk(visitor)
	}
	visitor.Exit(pattern)
}
