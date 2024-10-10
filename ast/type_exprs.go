package ast

import (
	"fmt"

	"github.com/pattyshack/gt/lexutil"
)

//
// SliceTypeExpr
//

type SliceTypeExpr struct {
	IsTypeExpr
	StartEndPos
	LeadingTrailingComments

	Value TypeExpression
}

var _ TypeExpression = &SliceTypeExpr{}

func (expr *SliceTypeExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.Value.Walk(visitor)
	visitor.Exit(expr)
}

//
// ArrayTypeExpr
//

type ArrayTypeExpr struct {
	IsTypeExpr
	StartEndPos
	LeadingTrailingComments

	Value TypeExpression
	Size  string // int literal string
}

var _ TypeExpression = &ArrayTypeExpr{}

func (expr *ArrayTypeExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.Value.Walk(visitor)
	visitor.Exit(expr)
}

//
// MapTypeExpr
//

type MapTypeExpr struct {
	IsTypeExpr
	StartEndPos
	LeadingTrailingComments

	Key   TypeExpression
	Value TypeExpression
}

var _ TypeExpression = &MapTypeExpr{}

func (expr *MapTypeExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.Key.Walk(visitor)
	expr.Value.Walk(visitor)
	visitor.Exit(expr)
}

//
// InferredTypeExpr
//

type InferredTypeExpr struct {
	IsTypeExpr
	StartEndPos
	LeadingTrailingComments

	IsImplicit bool
}

func NewImplicitInferredTypeExpr(pos lexutil.Location) TypeExpression {
	return &InferredTypeExpr{
		StartEndPos: NewStartEndPos(pos, pos),
		IsImplicit:  true,
	}
}

func (expr *InferredTypeExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	visitor.Exit(expr)
}

//
// NamedTypeExpr
//

type NamedTypeExpr struct {
	IsTypeExpr
	StartEndPos
	LeadingTrailingComments

	Pkg  string // optional.  "" = local
	Name string

	GenericArguments *TypeExpressionList
}

func (expr *NamedTypeExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.GenericArguments.Walk(visitor)
	visitor.Exit(expr)
}

//
// UnaryTypeExpr
//

type UnaryTypeOp string

const (
	ReturnOnDefaultTypeOp          = UnaryTypeOp("?")
	PanicOnDefaultTypeOp           = UnaryTypeOp("!")
	ReferenceTypeOp                = UnaryTypeOp("&")
	AllPublicMethodsTraitTypeOp    = UnaryTypeOp("~")
	AllPublicPropertiesTraitTypeOp = UnaryTypeOp("~~")
)

type UnaryTypeExpr struct {
	IsTypeExpr
	StartEndPos
	LeadingTrailingComments

	Op      UnaryTypeOp
	Operand TypeExpression
}

var _ TypeExpression = &UnaryTypeExpr{}
var _ Validator = &UnaryTypeExpr{}

func (expr *UnaryTypeExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.Operand.Walk(visitor)
	visitor.Exit(expr)
}

func (expr *UnaryTypeExpr) Validate(emitter *lexutil.ErrorEmitter) {
	switch expr.Op {
	case ReturnOnDefaultTypeOp,
		PanicOnDefaultTypeOp,
		ReferenceTypeOp,
		AllPublicMethodsTraitTypeOp,
		AllPublicPropertiesTraitTypeOp:
		// ok
	default:
		emitter.Emit(
			expr.Loc(),
			"invalid ast construction. unexpected unary type expr op (%s)",
			expr.Op)
	}
}

//
// BinaryTypeExpr
//

type BinaryTypeOp string

const (
	TraitIntersectTypeOp  = BinaryTypeOp("*")
	TraitUnionTypeOp      = BinaryTypeOp("+")
	TraitDifferenceTypeOp = BinaryTypeOp("-")
)

type BinaryTypeExpr struct {
	IsTypeExpr
	StartEndPos
	LeadingTrailingComments

	Left  TypeExpression
	Op    BinaryTypeOp
	Right TypeExpression
}

var _ TypeExpression = &BinaryTypeExpr{}
var _ Validator = &BinaryTypeExpr{}

func (expr *BinaryTypeExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.Left.Walk(visitor)
	expr.Right.Walk(visitor)
	visitor.Exit(expr)
}

func (expr *BinaryTypeExpr) Validate(emitter *lexutil.ErrorEmitter) {
	switch expr.Op {
	case TraitIntersectTypeOp, TraitUnionTypeOp, TraitDifferenceTypeOp: // ok
	default:
		emitter.Emit(
			expr.Loc(),
			"invalid ast construction. unexpected binary type expr op (%s)",
			expr.Op)
	}
}

//
// Struct / Enum / Trait PropertiesType
//

type PropertiesKind string

const (
	StructKind = PropertiesKind("struct")
	EnumKind   = PropertiesKind("enum")
	TraitKind  = PropertiesKind("trait")
)

type PropertiesTypeExpr struct {
	IsTypeExpr
	StartEndPos
	LeadingTrailingComments

	Kind PropertiesKind

	IsImplicit bool

	Properties []TypeProperty
}

var _ TypeExpression = &PropertiesTypeExpr{}
var _ Validator = &PropertiesTypeExpr{}

func NewAnyTrait(pos lexutil.Location) *PropertiesTypeExpr {
	return &PropertiesTypeExpr{
		StartEndPos: NewStartEndPos(pos, pos),
		Kind:        TraitKind,
		IsImplicit:  true,
	}
}

func (expr *PropertiesTypeExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	for _, prop := range expr.Properties {
		prop.Walk(visitor)
	}
	visitor.Exit(expr)
}

func (expr *PropertiesTypeExpr) Validate(emitter *lexutil.ErrorEmitter) {
	switch expr.Kind {
	case StructKind, EnumKind, TraitKind: // ok
	default:
		emitter.Emit(
			expr.Loc(),
			"invalid ast construction. unexpected properties type expr kind (%s)",
			expr.Kind)
	}

	defaultCount := 0
	for _, p := range expr.Properties {
		switch prop := p.(type) {
		case *FieldDef:
			if prop.IsDefault {
				defaultCount++
				if expr.Kind != EnumKind {
					emitter.Emit(p.Loc(), "unexpected default field")
				} else if defaultCount > 1 {
					emitter.Emit(p.Loc(), "more than one default field")
				}
			}

			if prop.IsPadding() && expr.Kind != StructKind {
				emitter.Emit(p.Loc(), "unexpected field padding")
			}
		case *FuncSignature:
			if expr.Kind != TraitKind {
				emitter.Emit(p.Loc(), "unexpected method signature in %s", expr.Kind)
			}
		case *FuncDefinition:
			if prop.Signature.Name == "" {
				emitter.Emit(prop.Loc(), "unexpected anonymous function definition")
			} else {
				if len(prop.Signature.Parameters.Elements) == 0 ||
					prop.Signature.Parameters.Elements[0].Kind != ReceiverParameter {

					emitter.Emit(
						prop.Loc(),
						"unexpected method definition, expected receiver parameter")
				}
			}
		default:
			panic(fmt.Sprintf("unexpected property type: %v", p))
		}
	}
}

//
// FuncSignature
//

type FuncSignature struct {
	IsTypeExpr
	IsTypeProp
	StartEndPos
	LeadingTrailingComments

	// Required for named named signature. Empty for anonymous signature.
	Name string
	// Required for named signature. nil for anonymous signature.
	GenericParameters *GenericParameterList

	Parameters *ParameterList
	ReturnType TypeExpression // use any trait if unconstrained
}

var _ TypeExpression = &FuncSignature{}
var _ TypeProperty = &FuncSignature{}
var _ Validator = &FuncSignature{}

func (sig *FuncSignature) Walk(visitor Visitor) {
	visitor.Enter(sig)
	if sig.GenericParameters != nil {
		sig.GenericParameters.Walk(visitor)
	}
	sig.Parameters.Walk(visitor)
	sig.ReturnType.Walk(visitor)
	visitor.Exit(sig)
}

func (sig *FuncSignature) Validate(emitter *lexutil.ErrorEmitter) {
	if sig.Name != "" {
		if sig.GenericParameters == nil {
			emitter.Emit(
				sig.Loc(),
				"invalid ast construction, generic parameters not set for "+
					"named func signature")
		}
	} else {
		if sig.GenericParameters != nil {
			emitter.Emit(
				sig.Loc(),
				"invalid ast construction, generic parameters set for "+
					"anonymous func signature")
		}
	}
}
