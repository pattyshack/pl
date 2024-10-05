package ast

//
// FieldDef
//

type FieldDef struct {
	IsTypeProp
	StartEndPos
	LeadingTrailingComments

	IsDefault bool

	Name string // Could be identifier, underscore, or ""
	Type TypeExpression
}

func (def *FieldDef) IsUnnamed() bool {
	return def.Name == ""
}

func (def *FieldDef) IsPadding() bool {
	return def.Name == "_"
}

var _ TypeProperty = &FieldDef{}

func (def *FieldDef) Walk(visitor Visitor) {
	visitor.Enter(def)
	def.Type.Walk(visitor)
	visitor.Exit(def)
}

//
// GenericParameter
//

type GenericParameter struct {
	StartEndPos
	LeadingTrailingComments

	Name string

	Constraint TypeExpression // optional
}

var _ Node = &GenericParameter{}

func (param *GenericParameter) Walk(visitor Visitor) {
	visitor.Enter(param)
	if param.Constraint != nil {
		param.Constraint.Walk(visitor)
	}
	visitor.Exit(param)
}

//
// Parameter
//

type ParameterKind string

const (
	ArgParameter      = ParameterKind("arg")
	ReceiverParameter = ParameterKind("receiver")
	VarargParameter   = ParameterKind("vararg")
)

type Parameter struct {
	StartEndPos
	LeadingTrailingComments

	Kind ParameterKind

	Name string // optional. NOTE: "_" is converted to ""
	Type TypeExpression
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
	SingularArgument = ArgumentKind("singular")
	VariadicArgument = ArgumentKind("variadic")
	// Only used by patterns
	SkipPatternArgument = ArgumentKind("skip-pattern")
)

type Argument struct {
	StartEndPos
	LeadingTrailingComments

	Kind ArgumentKind

	Name string // only used by named singular argument

	// NOTE: Expr may be nil for SkipPatternArgument
	Expr Expression
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
		Kind: SingularArgument,
		Expr: expr,
	}
}

func NewNamedArgument(
	name TokenValue,
	assign TokenValue,
	expr Expression,
) *Argument {
	arg := &Argument{
		StartEndPos: NewStartEndPos(name.Loc(), expr.End()),
		Kind:        SingularArgument,
		Name:        name.Val(),
		Expr:        expr,
	}

	arg.LeadingComment = name.TakeLeading()
	// prepend in reverse order
	expr.PrependToLeading(assign.TakeTrailing())
	expr.PrependToLeading(assign.TakeLeading())
	expr.PrependToLeading(name.TakeTrailing())
	arg.TrailingComment = expr.TakeTrailing()

	return arg
}

func NewVariadicArgument(
	expr Expression,
	ellipsis TokenValue,
) *Argument {
	arg := &Argument{
		StartEndPos: NewStartEndPos(expr.Loc(), ellipsis.End()),
		Kind:        VariadicArgument,
		Expr:        expr,
	}

	arg.LeadingComment = expr.TakeLeading()
	expr.AppendToTrailing(ellipsis.TakeLeading())
	arg.TrailingComment = ellipsis.TakeTrailing()

	return arg
}

func NewSkipPatternArgument(
	ellipsis TokenValue,
) *Argument {
	arg := &Argument{
		StartEndPos: NewStartEndPos(ellipsis.Loc(), ellipsis.End()),
		Kind:        SkipPatternArgument,
		Expr:        nil,
	}
	arg.LeadingComment = ellipsis.TakeLeading()
	arg.TrailingComment = ellipsis.TakeTrailing()
	return arg
}
