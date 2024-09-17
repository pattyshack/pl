package parser

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

type FuncReducerImpl struct {
	FuncTypeExprs []*FuncSignature

	FuncDefinitions    []*FuncDefinition
	MethodDefinitions  []*FuncDefinition
	AnonymousFuncExprs []*FuncDefinition
}

var _ FuncTypeExprReducer = &FuncReducerImpl{}
var _ MethodSignatureReducer = &FuncReducerImpl{}
var _ NamedFuncDefReducer = &FuncReducerImpl{}
var _ AnonymousFuncExprReducer = &FuncReducerImpl{}

func (reducer *FuncReducerImpl) NilToReturnType() (
	TypeExpression,
	error,
) {
	return nil, nil
}

func (reducer *FuncReducerImpl) toFuncSignature(
	funcKW *TokenValue,
	receiver *Parameter, // optional
	nameToken *TokenValue, // optional
	genericParameters *GenericParameterList, // optional
	parameters *ParameterList,
	returnType TypeExpression,
) *FuncSignature {
	leading := funcKW.TakeLeading()

	var params Node = parameters
	if genericParameters != nil {
		params = genericParameters
	} else {
		genericParameters = NewGenericParameterList()
	}

	name := ""
	if nameToken != nil {
		name = nameToken.Value
		params.PrependToLeading(nameToken.TakeTrailing())

		if receiver != nil {
			receiver.AppendToTrailing(nameToken.TakeLeading())
		} else {
			funcKW.AppendToTrailing(nameToken.TakeLeading())
		}
	}

	if receiver != nil {
		receiver.PrependToLeading(funcKW.TakeTrailing())
	} else {
		params.PrependToLeading(funcKW.TakeTrailing())
	}

	var trailing CommentGroups
	end := parameters.End()
	if returnType != nil {
		trailing = returnType.TakeTrailing()
		end = returnType.End()
	} else {
		trailing = parameters.TakeTrailing()
	}

	sig := &FuncSignature{
		StartEndPos:       NewStartEndPos(funcKW.Loc(), end),
		Receiver:          receiver,
		Name:              name,
		GenericParameters: *genericParameters,
		Parameters:        *parameters,
		ReturnType:        returnType,
	}
	sig.LeadingComment = leading
	sig.TrailingComment = trailing

	return sig
}

func (reducer *FuncReducerImpl) toFuncDefinition(
	funcKW *TokenValue,
	receiver *Parameter,
	nameToken *TokenValue,
	genericParameters *GenericParameterList,
	parameters *ParameterList,
	returnType TypeExpression,
	body Expression,
) *FuncDefinition {
	sig := reducer.toFuncSignature(
		funcKW,
		receiver,
		nameToken,
		genericParameters,
		parameters,
		returnType)

	return &FuncDefinition{
		StartEndPos: NewStartEndPos(sig.Loc(), body.End()),
		LeadingTrailingComments: LeadingTrailingComments{
			LeadingComment:  sig.TakeLeading(),
			TrailingComment: body.TakeTrailing(),
		},
		Signature: *sig,
		Body:      body,
	}
}

func (reducer *FuncReducerImpl) ToFuncTypeExpr(
	funcKW *TokenValue,
	parameters *ParameterList,
	returnType TypeExpression,
) (
	TypeExpression,
	error,
) {
	sig := reducer.toFuncSignature(
		funcKW,
		nil, // receiver
		nil, // name
		nil, // generic parameters
		parameters,
		returnType)
	reducer.FuncTypeExprs = append(reducer.FuncTypeExprs, sig)
	return sig, nil
}

func (reducer *FuncReducerImpl) ToMethodSignature(
	funcKW *TokenValue,
	name *TokenValue,
	parameters *ParameterList,
	returnType TypeExpression,
) (
	TypeProperty,
	error,
) {
	return reducer.toFuncSignature(
		funcKW,
		nil, // receiver
		name,
		nil, // generic parameters
		parameters,
		returnType), nil
}

func (reducer *FuncReducerImpl) FuncDefToNamedFuncDef(
	funcKW *TokenValue,
	name *TokenValue,
	genericParameters *GenericParameterList,
	parameters *ParameterList,
	returnType TypeExpression,
	body Expression,
) (
	Definition,
	error,
) {
	def := reducer.toFuncDefinition(
		funcKW,
		nil, // receiver
		name,
		genericParameters,
		parameters,
		returnType,
		body)

	reducer.FuncDefinitions = append(reducer.FuncDefinitions, def)
	return def, nil
}

func (reducer *FuncReducerImpl) MethodDefToNamedFuncDef(
	funcKW *TokenValue,
	lparen *TokenValue,
	receiver *Parameter,
	rparen *TokenValue,
	name *TokenValue,
	parameters *ParameterList,
	returnType TypeExpression,
	body Expression,
) (
	Definition,
	error,
) {
	funcKW.AppendToTrailing(lparen.TakeLeading())
	receiver.PrependToLeading(lparen.TakeTrailing())
	receiver.AppendToTrailing(rparen.TakeLeading())
	name.PrependToLeading(rparen.TakeTrailing())

	def := reducer.toFuncDefinition(
		funcKW,
		receiver,
		name,
		nil, // generic parameters
		parameters,
		returnType,
		body)
	reducer.MethodDefinitions = append(reducer.MethodDefinitions, def)
	return def, nil
}

func (reducer *FuncReducerImpl) ToAnonymousFuncExpr(
	funcKW *TokenValue,
	parameters *ParameterList,
	returnType TypeExpression,
	body Expression,
) (
	Expression,
	error,
) {
	def := reducer.toFuncDefinition(
		funcKW,
		nil, // receiver
		nil, // name
		nil, // generic parameters
		parameters,
		returnType,
		body)
	reducer.AnonymousFuncExprs = append(reducer.AnonymousFuncExprs, def)
	return def, nil
}
