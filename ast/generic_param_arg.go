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
