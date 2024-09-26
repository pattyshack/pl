package ast

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

	Name        string // optional
	HasEllipsis bool
	Type        TypeExpression // optional
}

var _ Node = &Parameter{}

func (param *Parameter) Walk(visitor Visitor) {
	visitor.Enter(param)
	if param.Type != nil {
		param.Type.Walk(visitor)
	}
	visitor.Exit(param)
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

func (arg *Argument) Walk(visitor Visitor) {
	visitor.Enter(arg)
	if arg.Expr != nil {
		arg.Expr.Walk(visitor)
	}
	visitor.Exit(arg)
}

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
	name TokenValue,
	assign TokenValue,
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
	ellipsis TokenValue,
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
