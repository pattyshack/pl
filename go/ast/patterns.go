package ast

import (
	"github.com/pattyshack/gt/lexutil"

	"github.com/pattyshack/pl/errors"
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
	lexutil.StartEndPos
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

func (assign *AssignPattern) Validate(emitter *errors.Emitter) {
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
	lexutil.StartEndPos
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
// AddrDeclPattern
//

type AddrDeclPattern struct {
	IsExpr
	lexutil.StartEndPos
	LeadingTrailingComments

	IsVar   bool // true = var, false = let
	Pattern Expression
	Type    TypeExpression // use InferredTypeExpr if unspecified
}

var _ Expression = &AddrDeclPattern{}

func NewAddrDeclPattern(
	varType TokenValue,
	pattern Expression,
	typeExpr TypeExpression,
) *AddrDeclPattern {
	isVar := false
	var start Node = pattern
	if varType != nil {
		isVar = varType.Val() == "var"
		start = varType
		pattern.PrependToLeading(varType.TakeTrailing())
	}

	var end Node = pattern
	if typeExpr != nil {
		end = typeExpr
	} else {
		typeExpr = NewImplicitInferredTypeExpr(pattern.End())
	}

	expr := &AddrDeclPattern{
		StartEndPos: lexutil.NewStartEndPos(start.Loc(), end.End()),
		IsVar:       isVar,
		Pattern:     pattern,
		Type:        typeExpr,
	}

	expr.LeadingComment = start.TakeLeading()
	expr.TrailingComment = end.TakeTrailing()

	return expr
}

func (pattern *AddrDeclPattern) Walk(visitor Visitor) {
	visitor.Enter(pattern)
	pattern.Pattern.Walk(visitor)
	pattern.Type.Walk(visitor)
	visitor.Exit(pattern)
}

//
// AssignToAddrPattern
//

type AssignToAddrPattern struct {
	IsExpr
	lexutil.StartEndPos
	LeadingTrailingComments

	Pattern Expression
}

var _ Expression = &AssignToAddrPattern{}

func NewAssignToAddrPattern(
	greater TokenValue,
	pattern Expression,
) *AssignToAddrPattern {
	addrPattern := &AssignToAddrPattern{
		StartEndPos: lexutil.NewStartEndPos(greater.Loc(), pattern.End()),
		Pattern:     pattern,
	}
	addrPattern.LeadingComment = greater.TakeLeading()
	pattern.PrependToLeading(greater.TakeTrailing())
	addrPattern.TrailingComment = pattern.TakeTrailing()

	return addrPattern
}

func (pattern *AssignToAddrPattern) Walk(visitor Visitor) {
	visitor.Enter(pattern)
	pattern.Pattern.Walk(visitor)
	visitor.Exit(pattern)
}

//
// EnumPattern
//

type EnumPattern struct {
	IsExpr
	lexutil.StartEndPos
	LeadingTrailingComments

	// - identifier string matches only only that named enum value
	// - empty string matches unnamed struct enum value
	// - underscore matches any named/unnamed enum value
	EnumValue string

	Pattern Expression
}

var _ Expression = &EnumPattern{}

func (pattern *EnumPattern) Walk(visitor Visitor) {
	visitor.Enter(pattern)
	pattern.Pattern.Walk(visitor)
	visitor.Exit(pattern)
}
