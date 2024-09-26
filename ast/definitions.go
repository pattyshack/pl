package ast

//
// FloatingComment
//

type FloatingComment struct {
	IsDef
	StartEndPos
	LeadingTrailingComments
}

var _ Definition = &FloatingComment{}

func (def *FloatingComment) Walk(visitor Visitor) {
	visitor.Enter(def)
	visitor.Exit(def)
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

func (def *PackageDef) Walk(visitor Visitor) {
	visitor.Enter(def)
	def.Body.Walk(visitor)
	visitor.Exit(def)
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
	GenericParameters *GenericParameterList // optional
	BaseType          TypeExpression
	Constraint        TypeExpression // optional
}

var _ Definition = &TypeDef{}

func (def *TypeDef) Walk(visitor Visitor) {
	visitor.Enter(def)
	if def.GenericParameters != nil {
		def.GenericParameters.Walk(visitor)
	}
	def.BaseType.Walk(visitor)
	if def.Constraint != nil {
		def.Constraint.Walk(visitor)
	}
	visitor.Exit(def)
}
