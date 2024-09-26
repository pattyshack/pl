package ast

//
// GenericParameter
//

type GenericParameter struct {
	StartEndPos
	LeadingTrailingComments

	Name string

	Constraint TypeExpression // optional
}

var _ Node = &GenericParameter{}

func (param *GenericParameter) Walk(visitor Visitor) {
	visitor.Enter(param)
	if param.Constraint != nil {
		param.Constraint.Walk(visitor)
	}
	visitor.Exit(param)
}
