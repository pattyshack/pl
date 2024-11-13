package ast

import (
	"fmt"

	"github.com/pattyshack/gt/lexutil"

	"github.com/pattyshack/pl/errors"
)

//
// SliceTypeExpr
//

type SliceTypeExpr struct {
	IsTypeExpr
	lexutil.StartEndPos
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
	lexutil.StartEndPos
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
	lexutil.StartEndPos
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
	lexutil.StartEndPos
	LeadingTrailingComments

	IsImplicit bool
}

func NewImplicitInferredTypeExpr(pos lexutil.Location) TypeExpression {
	return &InferredTypeExpr{
		StartEndPos: lexutil.NewStartEndPos(pos, pos),
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

// NOTE: NamedTypeExpr fields should be identical to ParameterizedExpr
type NamedTypeExpr struct {
	IsTypeExpr
	lexutil.StartEndPos
	LeadingTrailingComments

	// NOTE: Pkg is an identifier or empty string for normal ast construction.
	// But for the purpose of type cataloging / checking, Pkg is the fully
	// qualified package path (even for local names)
	Pkg  string // optional.  "" = local.
	Name string

	Parameters []TypeExpression
}

func (expr *NamedTypeExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	for _, arg := range expr.Parameters {
		arg.Walk(visitor)
	}
	visitor.Exit(expr)
}

//
// RefTypeExpr
//

type RefTypeExpr struct {
	IsTypeExpr
	lexutil.StartEndPos
	LeadingTrailingComments

	Value TypeExpression
}

var _ TypeExpression = &RefTypeExpr{}

func (expr *RefTypeExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.Value.Walk(visitor)
	visitor.Exit(expr)
}

//
// EnumOpTypeExpr
//

type DefaultEnumOp string

const (
	ReturnOnDefaultEnumOp = DefaultEnumOp("?")
	PanicOnDefaultEnumOp  = DefaultEnumOp("!")
)

type DefaultEnumOpTypeExpr struct {
	IsTypeExpr
	lexutil.StartEndPos
	LeadingTrailingComments

	Op   DefaultEnumOp
	Enum TypeExpression
}

var _ TypeExpression = &DefaultEnumOpTypeExpr{}
var _ Validator = &DefaultEnumOpTypeExpr{}

func (expr *DefaultEnumOpTypeExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.Enum.Walk(visitor)
	visitor.Exit(expr)
}

func (expr *DefaultEnumOpTypeExpr) Validate(emitter *errors.Emitter) {
	switch expr.Op {
	case ReturnOnDefaultEnumOp,
		PanicOnDefaultEnumOp:
		// ok
	default:
		emitter.Emit(
			expr.Loc(),
			"invalid ast construction. unexpected default enum op (%s)",
			expr.Op)
	}
}

//
// UnaryTraitOpTypeExpr
//

type UnaryTraitOp string

const (
	AllPublicMethodsTraitTraitOp    = UnaryTraitOp("~")
	AllPublicPropertiesTraitTraitOp = UnaryTraitOp("~~")
)

type UnaryTraitOpTypeExpr struct {
	IsTypeExpr
	lexutil.StartEndPos
	LeadingTrailingComments

	Op   UnaryTraitOp
	Base TypeExpression
}

var _ TypeExpression = &UnaryTraitOpTypeExpr{}
var _ Validator = &UnaryTraitOpTypeExpr{}

func (expr *UnaryTraitOpTypeExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.Base.Walk(visitor)
	visitor.Exit(expr)
}

func (expr *UnaryTraitOpTypeExpr) Validate(emitter *errors.Emitter) {
	switch expr.Op {
	case AllPublicMethodsTraitTraitOp,
		AllPublicPropertiesTraitTraitOp:
		// ok
	default:
		emitter.Emit(
			expr.Loc(),
			"invalid ast construction. unexpected unary trait op (%s)",
			expr.Op)
	}
}

//
// BinaryTraitOpTypeExpr
//

type BinaryTraitOp string

const (
	TraitIntersectTraitOp  = BinaryTraitOp("*")
	TraitUnionTraitOp      = BinaryTraitOp("+")
	TraitDifferenceTraitOp = BinaryTraitOp("-")
)

type BinaryTraitOpTypeExpr struct {
	IsTypeExpr
	lexutil.StartEndPos
	LeadingTrailingComments

	Left  TypeExpression
	Op    BinaryTraitOp
	Right TypeExpression
}

var _ TypeExpression = &BinaryTraitOpTypeExpr{}
var _ Validator = &BinaryTraitOpTypeExpr{}

func (expr *BinaryTraitOpTypeExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.Left.Walk(visitor)
	expr.Right.Walk(visitor)
	visitor.Exit(expr)
}

func (expr *BinaryTraitOpTypeExpr) Validate(emitter *errors.Emitter) {
	switch expr.Op {
	case TraitIntersectTraitOp, TraitUnionTraitOp, TraitDifferenceTraitOp: // ok
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
	lexutil.StartEndPos
	LeadingTrailingComments

	Kind PropertiesKind

	IsImplicit bool

	Properties []TypeProperty
}

var _ TypeExpression = &PropertiesTypeExpr{}
var _ Validator = &PropertiesTypeExpr{}

func NewAnyTrait(pos lexutil.Location) *PropertiesTypeExpr {
	return &PropertiesTypeExpr{
		StartEndPos: lexutil.NewStartEndPos(pos, pos),
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

func (expr *PropertiesTypeExpr) Validate(emitter *errors.Emitter) {
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
			switch prop.Qualifier {
			case NoFieldDefQualifier:
				// ok
			case LetFieldDefQualifier, VarFieldDefQualifier:
				if expr.Kind == EnumKind {
					emitter.Emit(p.Loc(), "unexpected let field")
				}
			case DefaultFieldDefQualifier:
				defaultCount++
				if expr.Kind != EnumKind {
					emitter.Emit(p.Loc(), "unexpected default field")
				} else if defaultCount > 1 {
					emitter.Emit(p.Loc(), "more than one default field")
				}
			default:
				emitter.Emit(p.Loc(), "unknown field def qualifier: %s", prop.Qualifier)
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
	lexutil.StartEndPos
	LeadingTrailingComments

	// Required for named named signature. Empty for anonymous signature.
	Name string
	// Required for named signature. nil for anonymous signature.
	// NOTE: named method signature must be an empty list
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

func (sig *FuncSignature) IsMethod() bool {
	return sig.Name != "" &&
		len(sig.Parameters.Elements) > 0 &&
		sig.Parameters.Elements[0].Kind == ReceiverParameter
}

func (sig *FuncSignature) Validate(emitter *errors.Emitter) {
	if sig.Name != "" {
		if sig.GenericParameters == nil {
			emitter.Emit(
				sig.Loc(),
				"invalid ast construction, generic parameters not set for "+
					"named func signature")
		} else if sig.IsMethod() {
			if len(sig.GenericParameters.Elements) > 0 {
				emitter.Emit(
					sig.GenericParameters.Loc(),
					"cannot specify generic parameters for method signature")
			}

			receiverType := sig.Parameters.Elements[0].Type
			ref, ok := receiverType.(*RefTypeExpr)
			if ok {
				receiverType = ref.Value
			}

			switch t := receiverType.(type) {
			case *InferredTypeExpr: // ok
			case *NamedTypeExpr:
				if t.Pkg != "" {
					emitter.Emit(
						receiverType.Loc(),
						"cannot define method for non-locally defined type")
				}
			default:
				emitter.Emit(receiverType.Loc(), "receiver must be simple named type")
			}
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
