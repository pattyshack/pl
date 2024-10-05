package ast

import (
	"github.com/pattyshack/gt/lexutil"
)

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
var _ Validator = &AssignPattern{}

func (assign *AssignPattern) Walk(visitor Visitor) {
	visitor.Enter(assign)
	assign.Pattern.Walk(visitor)
	assign.Value.Walk(visitor)
	visitor.Exit(assign)
}

func (assign *AssignPattern) Validate(emitter *lexutil.ErrorEmitter) {
	switch assign.Kind {
	case EqualAssign, InAssign:
	default:
		emitter.Emit(
			assign.Loc(),
			"invalid ast construction. unexpected assign pattern kind (%s)",
			assign.Kind)
	}
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
// AddressPattern
//

type AddressPatternKind string

const (
	VarAddressDecl  = AddressPatternKind("var")
	LetAddressDecl  = AddressPatternKind("let")
	AssignToAddress = AddressPatternKind(">")
)

type AddressPattern struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	Kind    AddressPatternKind
	Pattern Expression
	Type    TypeExpression // optional
}

var _ Expression = &AddressPattern{}
var _ Validator = &AddressPattern{}

func NewAddressPattern(
	varType TokenValue,
	pattern Expression,
	typeExpr TypeExpression,
) *AddressPattern {
	var end Node = pattern
	if typeExpr != nil {
		end = typeExpr
	}

	expr := &AddressPattern{
		StartEndPos: NewStartEndPos(varType.Loc(), end.End()),
		Kind:        AddressPatternKind(varType.Val()),
		Pattern:     pattern,
		Type:        typeExpr,
	}

	expr.LeadingComment = varType.TakeLeading()
	expr.TrailingComment = end.TakeTrailing()

	return expr
}

func (pattern *AddressPattern) Walk(visitor Visitor) {
	visitor.Enter(pattern)
	pattern.Pattern.Walk(visitor)
	if pattern.Type != nil {
		pattern.Type.Walk(visitor)
	}
	visitor.Exit(pattern)
}

func (pattern *AddressPattern) Validate(emitter *lexutil.ErrorEmitter) {
	switch pattern.Kind {
	case VarAddressDecl, LetAddressDecl, AssignToAddress:
	default:
		emitter.Emit(
			pattern.Loc(),
			"invalid ast construction. unexpected address pattern kind (%s)",
			pattern.Kind)
	}
}

//
// EnumPattern
//

type EnumPattern struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	EnumValue string
	Pattern   Expression // optional.  either implicit struct or AddressPattern
}

var _ Expression = &EnumPattern{}

func (pattern *EnumPattern) Walk(visitor Visitor) {
	visitor.Enter(pattern)
	if pattern.Pattern != nil {
		pattern.Pattern.Walk(visitor)
	}
	visitor.Exit(pattern)
}
