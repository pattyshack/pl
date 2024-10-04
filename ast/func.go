package ast

type FuncSignatureKind string

const (
	AnonymousFunc = FuncSignatureKind("anonymous")
	NamedFunc     = FuncSignatureKind("named")
)

type FuncSignature struct {
	IsTypeExpr
	IsTypeProp
	StartEndPos
	LeadingTrailingComments

	Kind FuncSignatureKind

	// Optional. Only used by named signature
	Name string
	// Optional.  Only used by named signature
	GenericParameters *GenericParameterList

	Parameters *ParameterList
	ReturnType TypeExpression // Optional
}

var _ TypeExpression = &FuncSignature{}
var _ TypeProperty = &FuncSignature{}

func (sig *FuncSignature) Walk(visitor Visitor) {
	visitor.Enter(sig)
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

	Signature *FuncSignature
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
