package reducer

import (
	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser/lr"
)

func (reducer *Reducer) NilToReturnType() (
	ast.TypeExpression,
	error,
) {
	return nil, nil
}

func (reducer *Reducer) toFuncSignature(
	funcKW *lr.TokenValue,
	receiver *ast.Parameter, // optional
	nameToken *lr.TokenValue, // optional
	genericParameters *ast.GenericParameterList, // optional
	parameters *ast.ParameterList,
	returnType ast.TypeExpression,
) *ast.FuncSignature {
	leading := funcKW.TakeLeading()

	if receiver != nil {
		leading.Append(receiver.TakeLeading())
		leading.Append(receiver.TakeTrailing())
	}

	name := ""
	if nameToken != nil {
		name = nameToken.Value
		leading.Append(nameToken.TakeLeading())
		leading.Append(nameToken.TakeTrailing())
	}

	var generics []*ast.GenericParameter
	if genericParameters != nil {
		generics = genericParameters.Elements

		leading.Append(genericParameters.TakeLeading())
		leading.Append(genericParameters.MiddleComment)
		leading.Append(genericParameters.TakeTrailing())
	}

	leading.Append(parameters.TakeLeading())
	leading.Append(parameters.MiddleComment)

	trailing := parameters.TakeTrailing()

	end := parameters.End()
	if returnType != nil {
		end = returnType.End()

		trailing.Append(returnType.TakeLeading())
		trailing.Append(returnType.TakeTrailing())
	}

	sig := &ast.FuncSignature{
		StartEndPos:       ast.NewStartEndPos(funcKW.Loc(), end),
		Receiver:          receiver,
		Name:              name,
		GenericParameters: generics,
		Parameters:        parameters.Elements,
		ReturnType:        returnType,
	}
	sig.LeadingComment = leading
	sig.TrailingComment = trailing

	return sig
}

func (reducer *Reducer) toFuncDefinition(
	funcKW *lr.TokenValue,
	receiver *ast.Parameter,
	nameToken *lr.TokenValue,
	genericParameters *ast.GenericParameterList,
	parameters *ast.ParameterList,
	returnType ast.TypeExpression,
	body ast.Expression,
) *ast.FuncDefinition {
	sig := reducer.toFuncSignature(
		funcKW,
		receiver,
		nameToken,
		genericParameters,
		parameters,
		returnType)

	return &ast.FuncDefinition{
		StartEndPos: ast.NewStartEndPos(sig.Loc(), body.End()),
		LeadingTrailingComments: ast.LeadingTrailingComments{
			LeadingComment:  sig.TakeLeading(),
			TrailingComment: body.TakeTrailing(),
		},
		Signature: *sig,
		Body:      body,
	}
}

func (reducer *Reducer) ToFuncTypeExpr(
	funcKW *lr.TokenValue,
	parameters *ast.ParameterList,
	returnType ast.TypeExpression,
) (
	ast.TypeExpression,
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
	parameters *ast.ParameterList,
	returnType ast.TypeExpression,
) (
	ast.TypeProperty,
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
	genericParameters *ast.GenericParameterList,
	parameters *ast.ParameterList,
	returnType ast.TypeExpression,
	body ast.Expression,
) (
	ast.Definition,
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
	receiver *ast.Parameter,
	rparen *lr.TokenValue,
	name *lr.TokenValue,
	parameters *ast.ParameterList,
	returnType ast.TypeExpression,
	body ast.Expression,
) (
	ast.Definition,
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
	parameters *ast.ParameterList,
	returnType ast.TypeExpression,
	body ast.Expression,
) (
	ast.Expression,
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
