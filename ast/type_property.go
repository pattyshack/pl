package ast

import (
	"fmt"
)

//
// FieldDef
//

type FieldDefKind string

const (
	NamedFieldDef              = FieldDefKind("named-field-def")
	UnnamedFieldDef            = FieldDefKind("unnamed-field-def")
	NamedDefaultEnumFieldDef   = FieldDefKind("named-default-enum-default-def")
	UnnamedDefaultEnumFieldDef = FieldDefKind("unnamed-default-enum-default-def")
	PaddingFieldDef            = FieldDefKind("padding-field-def")
)

type FieldDef struct {
	IsTypeProp
	StartEndPos
	LeadingTrailingComments

	Kind FieldDefKind

	Name string // Could be identifier, underscore, or ""
	Type TypeExpression
}

var _ TypeProperty = &FieldDef{}

func (def *FieldDef) Walk(visitor Visitor) {
	visitor.Enter(def)
	def.Type.Walk(visitor)
	visitor.Exit(def)
}

func (def FieldDef) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[FieldDef: Kind=%s Name=%s\n",
		indent,
		label,
		def.Kind,
		def.Name)
	result += def.Type.TreeString(indent+"  ", "Type=")
	result += "\n" + indent + "]"
	return result
}
