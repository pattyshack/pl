package ast

//
// FieldDef
//

type FieldDefKind string

const (
	NamedFieldDef   = FieldDefKind("named")
	UnnamedFieldDef = FieldDefKind("unnamed")
	PaddingFieldDef = FieldDefKind("padding")
)

type FieldDef struct {
	IsTypeProp
	StartEndPos
	LeadingTrailingComments

	Kind      FieldDefKind
	IsDefault bool

	Name string // Could be identifier, underscore, or ""
	Type TypeExpression
}

var _ TypeProperty = &FieldDef{}

func (def *FieldDef) Walk(visitor Visitor) {
	visitor.Enter(def)
	def.Type.Walk(visitor)
	visitor.Exit(def)
}
