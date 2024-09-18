package ast

import (
	"fmt"
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

func (slice SliceTypeExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf("%s%s[SliceTypeExpr:\n", indent, label)
	result += slice.Value.TreeString(indent+"  ", "Value=")
	result += "\n" + indent + "]"
	return result
}

//
// ArrayTypeExpr
//

type ArrayTypeExpr struct {
	IsTypeExpr
	StartEndPos
	LeadingTrailingComments

	Value TypeExpression
	Size  ValuedNode
}

var _ TypeExpression = &ArrayTypeExpr{}

func (array ArrayTypeExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[ArrayTypeExpr: Size=%s\n",
		indent,
		label,
		array.Size.Val())
	result += array.Value.TreeString(indent+"  ", "Value=")
	result += "\n" + indent + "]"
	return result
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

func (dict MapTypeExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf("%s%s[MapTypeExpr:\n", indent, label)
	result += dict.Key.TreeString(indent+"  ", "Key=") + "\n"
	result += dict.Value.TreeString(indent+"  ", "Value=")
	result += "\n" + indent + "]"
	return result
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

func (expr InferredTypeExpr) TreeString(indent string, label string) string {
	return fmt.Sprintf(
		"%s%s[InferredTypeExpr: InferMutable=%v]",
		indent,
		label,
		expr.InferMutable)
}

//
// NamedTypeExpr
//

type NamedTypeExpr struct {
	IsTypeExpr
	StartEndPos
	LeadingTrailingComments

	Pkg string // optional.  "" = local

	Name ValuedNode

	GenericArguments GenericArgumentList
}

func (expr NamedTypeExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[NamedTypeExpr: Pkg=%s Name=%s",
		indent,
		label,
		expr.Pkg,
		expr.Name.Val())

	if len(expr.GenericArguments.Elements) == 0 {
		return result + "]"
	}

	result += "\n" + expr.GenericArguments.TreeString(
		indent+"  ",
		"GenericArguments=")
	result += "\n" + indent + "]"
	return result
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

func (expr UnaryTypeExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[UnaryTypeExpr: Op=(%s)\n",
		indent,
		label,
		expr.Op)
	result += expr.Operand.TreeString(indent+"  ", "Operand=")
	result += "\n" + indent + "]"
	return result
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

func (expr BinaryTypeExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[BinaryTypeExpr: Op=(%s)\n",
		indent,
		label,
		expr.Op)
	result += expr.Left.TreeString(indent+"  ", "Left=") + "\n"
	result += expr.Right.TreeString(indent+"  ", "Right=")
	result += "\n" + indent + "]"
	return result
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

	Properties TypePropertyList
}

var _ TypeExpression = &PropertiesTypeExpr{}

func (expr PropertiesTypeExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[PropertiesTypeExpr: Kind=%s IsImplicit=%v\n",
		indent,
		label,
		expr.Kind,
		expr.IsImplicit)
	result += expr.Properties.TreeString(indent+"  ", "Properties=")
	result += "\n" + indent + "]"
	return result
}
