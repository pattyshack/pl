package ast

import (
	"fmt"
)

//
// FloatingComment
//

type FloatingComment struct {
	IsDef
	StartEndPos
	LeadingTrailingComments
}

var _ Definition = &FloatingComment{}

func (FloatingComment) TreeString(indent string, label string) string {
	return fmt.Sprintf("%s%s[FloatingComment]", indent, label)
}

//
// PackageDef
//

type PackageDef struct {
	IsDef
	StartEndPos
	LeadingTrailingComments

	Body *StatementsExpr
}

var _ Definition = &PackageDef{}

func (def PackageDef) TreeString(indent string, label string) string {
	result := fmt.Sprintf("%s%s[PackageDef:\n", indent, label)
	result += def.Body.TreeString(indent+"  ", "Body=")
	result += "\n" + indent + "]"
	return result
}

//
// TypeDef
//

type TypeDef struct {
	IsDef
	StartEndPos
	LeadingTrailingComments

	Name              string
	IsAlias           bool
	GenericParameters *GenericParameterList
	BaseType          TypeExpression
	Constraint        TypeExpression // optional
}

var _ Definition = &TypeDef{}

func (def TypeDef) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[TypeDef: Name=%s IsAlias=%v\n",
		indent,
		label,
		def.Name,
		def.IsAlias)
	result += def.GenericParameters.TreeString(
		indent+"  ",
		"GenericParameters=")
	result += "\n" + def.BaseType.TreeString(indent+"  ", "BaseType=")
	if def.Constraint != nil {
		result += "\n" + def.Constraint.TreeString(indent+"  ", "Constraint=")
	}
	result += "\n" + indent + "]"
	return result
}
