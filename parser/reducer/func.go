package reducer

import (
	. "github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser/lr"
)

func (reducer *Reducer) NilToReturnType() (
	TypeExpression,
	error,
) {
	return nil, nil
}

func (reducer *Reducer) toFuncSignature(
	funcKW *lr.TokenValue,
	receiver *Parameter, // optional
	nameToken *lr.TokenValue, // optional
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

func (reducer *Reducer) toFuncDefinition(
	funcKW *lr.TokenValue,
	receiver *Parameter,
	nameToken *lr.TokenValue,
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

func (reducer *Reducer) ToFuncTypeExpr(
	funcKW *lr.TokenValue,
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

func (reducer *Reducer) ToMethodSignature(
	funcKW *lr.TokenValue,
	name *lr.TokenValue,
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

func (reducer *Reducer) FuncDefToNamedFuncDef(
	funcKW *lr.TokenValue,
	name *lr.TokenValue,
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

func (reducer *Reducer) MethodDefToNamedFuncDef(
	funcKW *lr.TokenValue,
	lparen *lr.TokenValue,
	receiver *Parameter,
	rparen *lr.TokenValue,
	name *lr.TokenValue,
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

func (reducer *Reducer) ToAnonymousFuncExpr(
	funcKW *lr.TokenValue,
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
