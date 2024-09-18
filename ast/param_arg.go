package ast

import (
	"fmt"
)

//
// Parameter
//

type ParameterKind string

const (
	// NOTE: There is no implicit unnamed inferred arg variant
	NamedTypedArgParameter         = ParameterKind("named-typed-arg")
	NamedInferredArgParameter      = ParameterKind("named-inferred-arg")
	NamedTypedVarargParameter      = ParameterKind("named-typed-vararg")
	NamedInferredVarargParameter   = ParameterKind("named-inferred-vararg")
	UnnamedTypedArgParameter       = ParameterKind("unnamed-typed-arg")
	UnnamedTypedVarargParameter    = ParameterKind("unnamed-typed-vararg")
	UnnamedInferredVarargParameter = ParameterKind("unnamed-inferred-vararg")
	IgnoreTypedArgParameter        = ParameterKind("ignore-typed-arg")
	IgnoreInferredArgParameter     = ParameterKind("ignore-inferred-arg")
	IgnoreTypedVarargParameter     = ParameterKind("ignore-typed-vararg")
	IgnoreInferredVarargParameter  = ParameterKind("ignore-inferred-vararg")
)

type Parameter struct {
	StartEndPos
	LeadingTrailingComments

	Kind ParameterKind

	Name        ValuedNode // optional
	HasEllipsis bool
	Type        TypeExpression // optional
}

var _ Node = &Parameter{}

func (param Parameter) TreeString(indent string, label string) string {
	name := ""
	if param.Name != nil {
		name = param.Name.Val()
	}
	result := fmt.Sprintf(
		"%s%s[Parameter: Kind=%v Name=%s HasEllipsis=%v\n",
		indent,
		label,
		param.Kind,
		name,
		param.HasEllipsis)
	if param.Type != nil {
		result += param.Type.TreeString(indent+"  ", "Type=")
		result += "\n" + indent + "]"
	} else {
		result += "]"
	}
	return result
}

//
// Argument
//

type ArgumentKind string

const (
	PositionalArgument       = ArgumentKind("positional")
	NamedAssignmentArgument  = ArgumentKind("named-assigned")
	VarargAssignmentArgument = ArgumentKind("vararg-assigned")
	ColonExprArgument        = ArgumentKind("colon-expr")
	// Only used by patterns
	SkipPatternArgument = ArgumentKind("skip-pattern")
	// Only used by ColonExpr
	IsImplicitUnitArgument = ArgumentKind("implicit-unit")
)

type Argument struct {
	StartEndPos
	LeadingTrailingComments

	Kind ArgumentKind

	// Only set for named assignment
	OptionalName string

	// NOTE: Expr may be nil for SkipPatternArgument or IsImplicitUnit.
	Expr Expression

	HasEllipsis bool
}

var _ Node = &Argument{}

func NewPositionalArgument(expr Expression) *Argument {
	return &Argument{
		StartEndPos: NewStartEndPos(expr.Loc(), expr.End()),
		LeadingTrailingComments: LeadingTrailingComments{
			LeadingComment:  expr.TakeLeading(),
			TrailingComment: expr.TakeTrailing(),
		},
		Kind:        PositionalArgument,
		Expr:        expr,
		HasEllipsis: false,
	}
}

func NewNamedArgument(
	name ValuedNode,
	assign ValuedNode,
	expr Expression,
) *Argument {
	arg := &Argument{
		StartEndPos:  NewStartEndPos(name.Loc(), expr.End()),
		Kind:         NamedAssignmentArgument,
		OptionalName: name.Val(),
		Expr:         expr,
		HasEllipsis:  false,
	}

	arg.LeadingComment = name.TakeLeading()
	// prepend in reverse order
	expr.PrependToLeading(assign.TakeTrailing())
	expr.PrependToLeading(assign.TakeLeading())
	expr.PrependToLeading(name.TakeTrailing())
	arg.TrailingComment = expr.TakeTrailing()

	return arg
}

func NewSkipPatternArgument(
	ellipsis ValuedNode,
) *Argument {
	arg := &Argument{
		StartEndPos: NewStartEndPos(ellipsis.Loc(), ellipsis.End()),
		Kind:        SkipPatternArgument,
		Expr:        nil,
		HasEllipsis: true,
	}
	arg.LeadingComment = ellipsis.TakeLeading()
	arg.TrailingComment = ellipsis.TakeTrailing()
	return arg
}

func (arg *Argument) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[Argument: Kind=%v OptionalName=%s HasEllipsis=%v",
		indent,
		label,
		arg.Kind,
		arg.OptionalName,
		arg.HasEllipsis)
	if arg.Expr != nil {
		result += "\n"
		result += arg.Expr.TreeString(indent+"  ", "")
		result += "\n" + indent + "]"
	} else {
		result += "]"
	}

	return result
}
