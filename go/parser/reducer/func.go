package reducer

import (
	"github.com/pattyshack/gt/parseutil"

	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser/lr"
)

func (reducer *Reducer) NilToReturnType() (
	ast.TypeExpression,
	error,
) {
	return nil, nil
}

func (reducer *Reducer) AnonymousToFuncSignature(
	funcKW *lr.TokenValue,
	parameters *ast.ParameterList,
	returnType ast.TypeExpression, // optional
) (
	*ast.FuncSignature,
	error,
) {
	leading := funcKW.TakeLeading()
	parameters.PrependToLeading(funcKW.TakeTrailing())

	var end ast.Node = parameters
	if returnType != nil {
		end = returnType
	} else {
		returnType = ast.NewAnyTrait(parameters.End())
	}

	sig := &ast.FuncSignature{
		StartEndPos: parseutil.NewStartEndPos(funcKW.Loc(), end.End()),
		Parameters:  parameters,
		ReturnType:  returnType,
	}
	sig.LeadingComment = leading
	sig.TrailingComment = end.TakeTrailing()

	return sig, nil
}

func (reducer *Reducer) NamedToFuncSignature(
	funcKW *lr.TokenValue,
	name *lr.TokenValue,
	genericParameters *ast.GenericParameterList,
	parameters *ast.ParameterList,
	returnType ast.TypeExpression, // optional
) (
	*ast.FuncSignature,
	error,
) {
	leading := funcKW.TakeLeading()
	leading.Append(funcKW.TakeTrailing())
	leading.Append(name.TakeLeading())

	var next ast.Node = parameters
	if genericParameters != nil {
		next = genericParameters
	} else {
		genericParameters = ast.NewImplicitGenericParameterList(name.End())
	}
	next.PrependToLeading(name.TakeTrailing())

	var end ast.Node = parameters
	if returnType != nil {
		end = returnType
	} else {
		returnType = ast.NewAnyTrait(parameters.End())
	}

	sig := &ast.FuncSignature{
		StartEndPos:       parseutil.NewStartEndPos(funcKW.Loc(), end.End()),
		Name:              name.Value,
		GenericParameters: genericParameters,
		Parameters:        parameters,
		ReturnType:        returnType,
	}
	sig.LeadingComment = leading
	sig.TrailingComment = end.TakeTrailing()

	return sig, nil
}

func (reducer *Reducer) ToFuncDef(
	sig *ast.FuncSignature,
	body *ast.StatementsExpr,
) (
	*ast.FuncDefinition,
	error,
) {
	return &ast.FuncDefinition{
		StartEndPos: parseutil.NewStartEndPos(sig.Loc(), body.End()),
		LeadingTrailingComments: ast.LeadingTrailingComments{
			LeadingComment:  sig.TakeLeading(),
			TrailingComment: body.TakeTrailing(),
		},
		Signature: sig,
		Body:      body,
	}, nil
}
