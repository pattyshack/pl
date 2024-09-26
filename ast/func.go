package ast

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
	GenericParameters *GenericParameterList

	Parameters *ParameterList
	ReturnType TypeExpression // Optional
}

var _ TypeExpression = &FuncSignature{}
var _ TypeProperty = &FuncSignature{}

func (sig *FuncSignature) Walk(visitor Visitor) {
	visitor.Enter(sig)
	if sig.Receiver != nil {
		sig.Receiver.Walk(visitor)
	}
	if sig.GenericParameters != nil {
		sig.GenericParameters.Walk(visitor)
	}
	sig.Parameters.Walk(visitor)
	if sig.ReturnType != nil {
		sig.ReturnType.Walk(visitor)
	}
	visitor.Exit(sig)
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

func (def *FuncDefinition) Walk(visitor Visitor) {
	visitor.Enter(def)
	def.Signature.Walk(visitor)
	def.Body.Walk(visitor)
	visitor.Exit(def)
}
