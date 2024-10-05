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

	InferMutable bool // "_" infers mutable ref. "." infers immutabel ref.
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

	GenericArguments *TypeExpressionList // optional
}

func (expr *NamedTypeExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	if expr.GenericArguments != nil {
		expr.GenericArguments.Walk(visitor)
	}
	visitor.Exit(expr)
}

//
// UnaryTypeExpr
//

type UnaryTypeOp string

type UnaryTypeExpr struct {
	IsTypeExpr
	StartEndPos
	LeadingTrailingComments

	Op      UnaryTypeOp
	Operand TypeExpression
}

func (expr *UnaryTypeExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.Operand.Walk(visitor)
	visitor.Exit(expr)
}

//
// BinaryTypeExpr
//

type BinaryTypeOp string

type BinaryTypeExpr struct {
	IsTypeExpr
	StartEndPos
	LeadingTrailingComments

	Left  TypeExpression
	Op    BinaryTypeOp
	Right TypeExpression
}

func (expr *BinaryTypeExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.Left.Walk(visitor)
	expr.Right.Walk(visitor)
	visitor.Exit(expr)
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

func (expr *PropertiesTypeExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	for _, prop := range expr.Properties {
		prop.Walk(visitor)
	}
	visitor.Exit(expr)
}

func (expr *PropertiesTypeExpr) Validate(emitter *lexutil.ErrorEmitter) {
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

			if prop.Kind == PaddingFieldDef && expr.Kind != StructKind {
				emitter.Emit(p.Loc(), "unexpected field padding")
			}
		case *FuncSignature:
			if expr.Kind != TraitKind {
				emitter.Emit(p.Loc(), "unexpected method signature in %s", expr.Kind)
			}
		default:
			panic(fmt.Sprintf("unexpected property type: %v", p))
		}
	}
}

//
// FuncSignature
//

type FuncSignatureKind string

type FuncSignature struct {
	IsTypeExpr
	IsTypeProp
	StartEndPos
	LeadingTrailingComments

	// Optional. Only used by named signature
	Name string
	// Optional.  Only used by named signature
	GenericParameters *GenericParameterList

	Parameters *ParameterList
	ReturnType TypeExpression // Optional
}

var _ TypeExpression = &FuncSignature{}
var _ TypeProperty = &FuncSignature{}

func (sig *FuncSignature) Walk(visitor Visitor) {
	visitor.Enter(sig)
	if sig.GenericParameters != nil {
		sig.GenericParameters.Walk(visitor)
	}
	sig.Parameters.Walk(visitor)
	if sig.ReturnType != nil {
		sig.ReturnType.Walk(visitor)
	}
	visitor.Exit(sig)
}
