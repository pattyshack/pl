package parser

import (
	. "github.com/pattyshack/pl/ast"
)

//
// Parameter
//

type ParameterReducerImpl struct{}

var _ ProperParameterDefReducer = &ParameterReducerImpl{}
var _ ParameterDeclReducer = &ParameterReducerImpl{}
var _ ParameterDefReducer = &ParameterReducerImpl{}

func (reducer *ParameterReducerImpl) namedTypedArg(
	kind ParameterKind,
	name *TokenValue,
	typeExpr TypeExpression,
) (
	*Parameter,
	error,
) {
	param := &Parameter{
		StartEndPos: NewStartEndPos(name.Loc(), typeExpr.End()),
		Kind:        kind,
		Name:        name,
		HasEllipsis: false,
		Type:        typeExpr,
	}
	param.LeadingComment = name.TakeLeading()
	param.TrailingComment = typeExpr.TakeTrailing()
	return param, nil
}

func (reducer *ParameterReducerImpl) NamedTypedArgToProperParameterDef(
	name *TokenValue,
	typeExpr TypeExpression,
) (
	*Parameter,
	error,
) {
	return reducer.namedTypedArg(NamedTypedArgParameter, name, typeExpr)
}

func (reducer *ParameterReducerImpl) IgnoreTypedArgToProperParameterDef(
	underscore *TokenValue,
	typeExpr TypeExpression,
) (
	*Parameter,
	error,
) {
	return reducer.namedTypedArg(IgnoreTypedArgParameter, underscore, typeExpr)
}

func (reducer *ParameterReducerImpl) namedTypedVararg(
	kind ParameterKind,
	name *TokenValue,
	ellipsis *TokenValue,
	typeExpr TypeExpression,
) (
	*Parameter,
	error,
) {
	param := &Parameter{
		StartEndPos: NewStartEndPos(name.Loc(), typeExpr.End()),
		Kind:        kind,
		Name:        name,
		HasEllipsis: true,
		Type:        typeExpr,
	}
	param.LeadingComment = name.TakeLeading()
	name.AppendToTrailing(ellipsis.TakeLeading())
	typeExpr.PrependToLeading(ellipsis.TakeTrailing())
	param.TrailingComment = typeExpr.TakeTrailing()
	return param, nil
}

func (reducer *ParameterReducerImpl) NamedTypedVarargToProperParameterDef(
	name *TokenValue,
	ellipsis *TokenValue,
	typeExpr TypeExpression,
) (
	*Parameter,
	error,
) {
	return reducer.namedTypedVararg(
		NamedTypedVarargParameter,
		name,
		ellipsis,
		typeExpr)
}

func (reducer *ParameterReducerImpl) IgnoreTypedVarargToProperParameterDef(
	underscore *TokenValue,
	ellipsis *TokenValue,
	typeExpr TypeExpression,
) (
	*Parameter,
	error,
) {
	return reducer.namedTypedVararg(
		IgnoreTypedVarargParameter,
		underscore,
		ellipsis,
		typeExpr)
}

func (reducer *ParameterReducerImpl) namedInferredVararg(
	kind ParameterKind,
	name *TokenValue,
	ellipsis *TokenValue,
) (
	*Parameter,
	error,
) {
	param := &Parameter{
		StartEndPos: NewStartEndPos(name.Loc(), ellipsis.End()),
		Kind:        kind,
		Name:        name,
		HasEllipsis: true,
		Type:        nil,
	}
	param.LeadingComment = name.TakeLeading()
	name.AppendToTrailing(ellipsis.TakeLeading())
	param.TrailingComment = ellipsis.TakeTrailing()
	return param, nil
}

func (reducer *ParameterReducerImpl) NamedInferredVarargToProperParameterDef(
	name *TokenValue,
	ellipsis *TokenValue,
) (
	*Parameter,
	error,
) {
	return reducer.namedInferredVararg(
		NamedInferredVarargParameter,
		name,
		ellipsis)
}

func (reducer *ParameterReducerImpl) IgnoreInferredVarargToProperParameterDef(
	underscore *TokenValue,
	ellipsis *TokenValue,
) (
	*Parameter,
	error,
) {
	return reducer.namedInferredVararg(
		IgnoreInferredVarargParameter,
		underscore,
		ellipsis)
}

func (reducer *ParameterReducerImpl) UnnamedTypedArgToParameterDecl(
	typeExpr TypeExpression,
) (
	*Parameter,
	error,
) {
	param := &Parameter{
		StartEndPos: NewStartEndPos(typeExpr.Loc(), typeExpr.End()),
		Kind:        UnnamedTypedArgParameter,
		Name:        nil,
		HasEllipsis: false,
		Type:        typeExpr,
	}
	param.LeadingComment = typeExpr.TakeLeading()
	param.TrailingComment = typeExpr.TakeTrailing()
	return param, nil
}

func (reducer *ParameterReducerImpl) UnnamedInferredVarargToParameterDecl(
	ellipsis *TokenValue,
) (
	*Parameter,
	error,
) {
	param := &Parameter{
		StartEndPos: NewStartEndPos(ellipsis.Loc(), ellipsis.End()),
		Kind:        UnnamedInferredVarargParameter,
		Name:        nil,
		HasEllipsis: true,
		Type:        nil,
	}
	param.LeadingTrailingComments = ellipsis.LeadingTrailingComments
	return param, nil
}

func (reducer *ParameterReducerImpl) UnnamedTypedVarargToParameterDecl(
	ellipsis *TokenValue,
	typeExpr TypeExpression,
) (
	*Parameter,
	error,
) {
	param := &Parameter{
		StartEndPos: NewStartEndPos(ellipsis.Loc(), typeExpr.End()),
		Kind:        UnnamedTypedVarargParameter,
		Name:        nil,
		HasEllipsis: true,
		Type:        typeExpr,
	}
	param.LeadingComment = ellipsis.TakeLeading()
	typeExpr.PrependToLeading(ellipsis.TakeTrailing())
	param.TrailingComment = typeExpr.TakeTrailing()
	return param, nil
}

func (reducer *ParameterReducerImpl) namedInferredArg(
	kind ParameterKind,
	name *TokenValue,
) (
	*Parameter,
	error,
) {
	param := &Parameter{
		StartEndPos: NewStartEndPos(name.Loc(), name.End()),
		Kind:        kind,
		Name:        name,
		HasEllipsis: false,
		Type:        nil,
	}
	param.LeadingComment = name.TakeLeading()
	param.TrailingComment = name.TakeTrailing()
	return param, nil
}

func (reducer *ParameterReducerImpl) NamedInferredArgToParameterDef(
	name *TokenValue,
) (
	*Parameter,
	error,
) {
	return reducer.namedInferredArg(NamedInferredArgParameter, name)
}

func (reducer *ParameterReducerImpl) IgnoreInferredArgToParameterDef(
	underscore *TokenValue,
) (
	*Parameter,
	error,
) {
	return reducer.namedInferredArg(IgnoreInferredArgParameter, underscore)
}

//
// ParameterList
//

type ParameterListReducerImpl struct{}

var _ ProperParameterDeclListReducer = &ParameterListReducerImpl{}
var _ ParameterDeclListReducer = &ParameterListReducerImpl{}
var _ ParameterDeclsReducer = &ParameterListReducerImpl{}
var _ ProperParameterDefListReducer = &ParameterListReducerImpl{}
var _ ParameterDefListReducer = &ParameterListReducerImpl{}
var _ ParameterDefsReducer = &ParameterListReducerImpl{}

func (reducer *ParameterListReducerImpl) AddToProperParameterDeclList(
	list *ParameterList,
	comma *TokenValue,
	parameter *Parameter,
) (
	*ParameterList,
	error,
) {
	list.ReduceAdd(comma, parameter)
	return list, nil
}

func (reducer *ParameterListReducerImpl) ParameterDeclToProperParameterDeclList(
	parameter *Parameter,
) (
	*ParameterList,
	error,
) {
	list := NewParameterList()
	list.Add(parameter)
	return list, nil
}

func (reducer *ParameterListReducerImpl) ImproperToParameterDeclList(
	list *ParameterList,
	comma *TokenValue,
) (
	*ParameterList,
	error,
) {
	list.ReduceImproper(comma)
	return list, nil
}

func (reducer *ParameterListReducerImpl) NilToParameterDeclList() (
	*ParameterList,
	error,
) {
	return NewParameterList(), nil
}

func (reducer *ParameterListReducerImpl) ToParameterDecls(
	lparen *TokenValue,
	list *ParameterList,
	rparen *TokenValue,
) (
	*ParameterList,
	error,
) {
	list.ReduceMarkers(lparen, rparen)
	return list, nil
}

func (reducer *ParameterListReducerImpl) AddToProperParameterDefList(
	list *ParameterList,
	comma *TokenValue,
	parameter *Parameter,
) (
	*ParameterList,
	error,
) {
	list.ReduceAdd(comma, parameter)
	return list, nil
}

func (reducer *ParameterListReducerImpl) ParameterDefToProperParameterDefList(
	parameter *Parameter,
) (
	*ParameterList,
	error,
) {
	list := NewParameterList()
	list.Add(parameter)
	return list, nil
}

func (reducer *ParameterListReducerImpl) ImproperToParameterDefList(
	list *ParameterList,
	comma *TokenValue,
) (
	*ParameterList,
	error,
) {
	list.ReduceImproper(comma)
	return list, nil
}

func (reducer *ParameterListReducerImpl) NilToParameterDefList() (
	*ParameterList,
	error,
) {
	return NewParameterList(), nil
}

func (reducer *ParameterListReducerImpl) ToParameterDefs(
	lparen *TokenValue,
	list *ParameterList,
	rparen *TokenValue,
) (
	*ParameterList,
	error,
) {
	list.ReduceMarkers(lparen, rparen)
	return list, nil
}

//
// Argument
//

type ArgumentReducerImpl struct {
}

var _ ArgumentReducer = ArgumentReducerImpl{}
var _ FieldVarPatternReducer = ArgumentReducerImpl{}

func (ArgumentReducerImpl) PositionalToArgument(
	expr Expression,
) (
	*Argument,
	error,
) {
	return NewPositionalArgument(expr), nil
}

func (ArgumentReducerImpl) ColonExprToArgument(
	expr *ColonExpr,
) (
	*Argument,
	error,
) {
	return &Argument{
		StartEndPos: NewStartEndPos(expr.Loc(), expr.End()),
		LeadingTrailingComments: LeadingTrailingComments{
			LeadingComment:  expr.TakeLeading(),
			TrailingComment: expr.TakeTrailing(),
		},
		Kind:        ColonExprArgument,
		Expr:        expr,
		HasEllipsis: false,
	}, nil
}

func (ArgumentReducerImpl) NamedAssignmentToArgument(
	name *TokenValue,
	assign *TokenValue,
	expr Expression,
) (
	*Argument,
	error,
) {
	return NewNamedArgument(name, assign, expr), nil
}

func (ArgumentReducerImpl) VarargAssignmentToArgument(
	expr Expression,
	ellipsis *TokenValue,
) (
	*Argument,
	error,
) {
	arg := &Argument{
		StartEndPos: NewStartEndPos(expr.Loc(), ellipsis.End()),
		Kind:        VarargAssignmentArgument,
		Expr:        expr,
		HasEllipsis: true,
	}

	arg.LeadingComment = expr.TakeLeading()
	expr.AppendToTrailing(ellipsis.TakeLeading())
	arg.TrailingComment = ellipsis.TakeTrailing()
	return arg, nil
}

func (ArgumentReducerImpl) SkipPatternToArgument(
	ellipsis *TokenValue,
) (
	*Argument,
	error,
) {
	return NewSkipPatternArgument(ellipsis), nil
}

func (ArgumentReducerImpl) PositionalToFieldVarPattern(
	expr Expression,
) (
	*Argument,
	error,
) {
	return NewPositionalArgument(expr), nil
}

func (ArgumentReducerImpl) NamedAssignmentToFieldVarPattern(
	name *TokenValue,
	assign *TokenValue,
	expr Expression,
) (
	*Argument,
	error,
) {
	return NewNamedArgument(name, assign, expr), nil
}

func (ArgumentReducerImpl) SkipPatternToFieldVarPattern(
	ellipsis *TokenValue,
) (
	*Argument,
	error,
) {
	return NewSkipPatternArgument(ellipsis), nil
}

//
// Argument list
//

type ArgumentListReducerImpl struct{}

var _ ProperArgumentsReducer = &ArgumentListReducerImpl{}
var _ ArgumentsReducer = &ArgumentListReducerImpl{}
var _ ProperFieldVarPatternsReducer = &ArgumentListReducerImpl{}
var _ FieldVarPatternsReducer = &ArgumentListReducerImpl{}

func (reducer *ArgumentListReducerImpl) AddToProperArguments(
	list *ArgumentList,
	comma *TokenValue,
	arg *Argument,
) (
	*ArgumentList,
	error,
) {
	list.ReduceAdd(comma, arg)
	return list, nil
}

func (reducer *ArgumentListReducerImpl) ArgumentToProperArguments(
	arg *Argument,
) (
	*ArgumentList,
	error,
) {
	list := NewArgumentList()
	list.Add(arg)
	return list, nil
}

func (reducer *ArgumentListReducerImpl) ImproperToArguments(
	list *ArgumentList,
	comma *TokenValue,
) (
	*ArgumentList,
	error,
) {
	list.ReduceImproper(comma)
	return list, nil
}

func (reducer *ArgumentListReducerImpl) NilToArguments() (
	*ArgumentList,
	error,
) {
	return NewArgumentList(), nil
}

func (reducer *ArgumentListReducerImpl) FieldVarPatternToProperFieldVarPatterns(
	pattern *Argument,
) (
	*ArgumentList,
	error,
) {
	list := NewArgumentList()
	list.Add(pattern)
	return list, nil
}

func (reducer *ArgumentListReducerImpl) AddToProperFieldVarPatterns(
	list *ArgumentList,
	comma *TokenValue,
	pattern *Argument,
) (
	*ArgumentList,
	error,
) {
	list.ReduceAdd(comma, pattern)
	return list, nil
}

func (reducer *ArgumentListReducerImpl) ImproperToFieldVarPatterns(
	list *ArgumentList,
	comma *TokenValue,
) (
	*ArgumentList,
	error,
) {
	list.ReduceImproper(comma)
	return list, nil
}
