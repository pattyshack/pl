package ast

import (
	"fmt"
)

type FuncSignature struct {
	IsTypeExpr
	IsTypeProp
	StartEndPos
	LeadingTrailingComments

	// Optional. Only used by method definition
	Receiver *Parameter
	// Optional. Only used by trait method signature
	Name string
	// Optional.  Only used by function definition
	GenericParameters GenericParameterList

	Parameters ParameterList
	ReturnType TypeExpression // Optional
}

var _ TypeExpression = &FuncSignature{}
var _ TypeProperty = &FuncSignature{}

func (sig FuncSignature) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[FuncSignature: Name=%s\n",
		indent,
		label,
		sig.Name)
	if sig.Receiver != nil {
		result += sig.Receiver.TreeString(indent+"  ", "Receiver=") + "\n"
	}

	if len(sig.GenericParameters.Elements) > 0 {
		result += sig.GenericParameters.TreeString(
			indent+"  ",
			"GenericParameters=")
		result += "\n"
	}

	result += sig.Parameters.TreeString(indent+"  ", "Parameters=") + "\n"

	if sig.ReturnType == nil {
		result += indent + "  ReturnType=(nil)\n"
	} else {
		result += sig.ReturnType.TreeString(indent+"  ", "ReturnType=") + "\n"
	}

	result += indent + "]"
	return result
}

type FuncDefinition struct {
	IsExpr
	IsDef
	StartEndPos
	LeadingTrailingComments

	Signature FuncSignature
	Body      Expression
}

var _ Expression = &FuncDefinition{}
var _ Definition = &FuncDefinition{}

func (def FuncDefinition) TreeString(indent string, label string) string {
	result := fmt.Sprintf("%s%s[FuncDefinition:\n", indent, label)
	result += def.Signature.TreeString(indent+"  ", "Signature=") + "\n"
	result += def.Body.TreeString(indent+"  ", "Body=")
	result += "\n" + indent + "]"
	return result
}
