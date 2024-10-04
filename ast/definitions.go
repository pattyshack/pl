package ast

//
// FloatingComment
//

type FloatingComment struct {
	IsStmt
	StartEndPos
	LeadingTrailingComments
}

var _ Statement = &FloatingComment{}

func (def *FloatingComment) Walk(visitor Visitor) {
	visitor.Enter(def)
	visitor.Exit(def)
}

//
// TypeDef
//

type TypeDef struct {
	IsStmt
	StartEndPos
	LeadingTrailingComments

	Name              string
	IsAlias           bool
	GenericParameters *GenericParameterList // optional
	BaseType          TypeExpression
	Constraint        TypeExpression // optional
}

var _ Statement = &TypeDef{}

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
