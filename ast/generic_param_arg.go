package ast

import (
	"fmt"
)

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

func (param GenericParameter) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[GenericParameter: Name=%s",
		indent,
		label,
		param.Name)

	if param.Constraint == nil {
		result += "]"
	} else {
		result += "\n" + param.Constraint.TreeString(indent+"  ", "Constraint=")
		result += "\n" + indent + "]"
	}

	return result
}
