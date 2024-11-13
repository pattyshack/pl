package ast

import (
	"github.com/pattyshack/gt/lexutil"

	"github.com/pattyshack/pl/errors"
)

//
// FieldDef
//

type FieldDefQualifier string

const (
	NoFieldDefQualifier      = FieldDefQualifier("")
	VarFieldDefQualifier     = FieldDefQualifier("var")
	LetFieldDefQualifier     = FieldDefQualifier("let")
	DefaultFieldDefQualifier = FieldDefQualifier("default")
)

type FieldDef struct {
	IsTypeProp
	lexutil.StartEndPos
	LeadingTrailingComments

	Qualifier FieldDefQualifier

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
	lexutil.StartEndPos
	LeadingTrailingComments

	Name string

	Constraint TypeExpression // use any trait if unconstrained
}

var _ Node = &GenericParameter{}

func (param *GenericParameter) Walk(visitor Visitor) {
	visitor.Enter(param)
	param.Constraint.Walk(visitor)
	visitor.Exit(param)
}

//
// Parameter
//

type ParameterKind string

const (
	SingularParameter = ParameterKind("singular")
	ReceiverParameter = ParameterKind("receiver")
	VariadicParameter = ParameterKind("variadic")
)

type Parameter struct {
	lexutil.StartEndPos
	LeadingTrailingComments

	Kind ParameterKind

	Name string // either a real name, or "". NOTE: "_" is converted to ""
	Type TypeExpression
}

var _ Node = &Parameter{}
var _ Validator = &Parameter{}

func (param *Parameter) Walk(visitor Visitor) {
	visitor.Enter(param)
	param.Type.Walk(visitor)
	visitor.Exit(param)
}

func (param *Parameter) Validate(emitter *errors.Emitter) {
	switch param.Kind {
	case SingularParameter, ReceiverParameter, VariadicParameter:
		// ok
	default:
		emitter.Emit(
			param.Loc(),
			"invalid ast construction. unexpected parameter kind (%s)",
			param.Kind)
	}
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
	lexutil.StartEndPos
	LeadingTrailingComments

	Kind ArgumentKind

	Name string // only used by named singular argument

	// NOTE: Expr is nil for SkipPatternArgument
	Expr Expression
}

var _ Node = &Argument{}
var _ Validator = &Argument{}

func (arg *Argument) Walk(visitor Visitor) {
	visitor.Enter(arg)
	if arg.Expr != nil {
		arg.Expr.Walk(visitor)
	}
	visitor.Exit(arg)
}

func (arg *Argument) Validate(emitter *errors.Emitter) {
	switch arg.Kind {
	case SingularArgument, VariadicArgument, SkipPatternArgument:
		// ok
	default:
		emitter.Emit(
			arg.Loc(),
			"invalid ast construction. unexpected argument kind (%s)",
			arg.Kind)
	}

	if arg.Name != "" && arg.Kind != SingularArgument {
		emitter.Emit(
			arg.Loc(),
			"invalid ast construction. name set in %s argument",
			arg.Kind)
	}

	if arg.Expr == nil {
		if arg.Kind != SkipPatternArgument {
			emitter.Emit(
				arg.Loc(),
				"invalid ast construction. expression not set in %s argument",
				arg.Kind)
		}
	} else if arg.Kind == SkipPatternArgument {
		emitter.Emit(
			arg.Loc(),
			"invalid ast construction. expression set in %s argument",
			arg.Kind)
	}
}

func NewPositionalArgument(expr Expression) *Argument {
	return &Argument{
		StartEndPos: lexutil.NewStartEndPos(expr.Loc(), expr.End()),
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
		StartEndPos: lexutil.NewStartEndPos(name.Loc(), expr.End()),
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
		StartEndPos: lexutil.NewStartEndPos(expr.Loc(), ellipsis.End()),
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
		StartEndPos: lexutil.NewStartEndPos(ellipsis.Loc(), ellipsis.End()),
		Kind:        SkipPatternArgument,
		Expr:        nil,
	}
	arg.LeadingComment = ellipsis.TakeLeading()
	arg.TrailingComment = ellipsis.TakeTrailing()
	return arg
}
