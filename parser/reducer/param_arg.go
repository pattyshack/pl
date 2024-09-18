package reducer

import (
	. "github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser/lr"
)

//
// Parameter
//

func (reducer *Reducer) namedTypedArg(
	kind ParameterKind,
	name *lr.TokenValue,
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

func (reducer *Reducer) NamedTypedArgToProperParameterDef(
	name *lr.TokenValue,
	typeExpr TypeExpression,
) (
	*Parameter,
	error,
) {
	return reducer.namedTypedArg(NamedTypedArgParameter, name, typeExpr)
}

func (reducer *Reducer) IgnoreTypedArgToProperParameterDef(
	underscore *lr.TokenValue,
	typeExpr TypeExpression,
) (
	*Parameter,
	error,
) {
	return reducer.namedTypedArg(IgnoreTypedArgParameter, underscore, typeExpr)
}

func (reducer *Reducer) namedTypedVararg(
	kind ParameterKind,
	name *lr.TokenValue,
	ellipsis *lr.TokenValue,
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

func (reducer *Reducer) NamedTypedVarargToProperParameterDef(
	name *lr.TokenValue,
	ellipsis *lr.TokenValue,
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

func (reducer *Reducer) IgnoreTypedVarargToProperParameterDef(
	underscore *lr.TokenValue,
	ellipsis *lr.TokenValue,
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

func (reducer *Reducer) namedInferredVararg(
	kind ParameterKind,
	name *lr.TokenValue,
	ellipsis *lr.TokenValue,
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

func (reducer *Reducer) NamedInferredVarargToProperParameterDef(
	name *lr.TokenValue,
	ellipsis *lr.TokenValue,
) (
	*Parameter,
	error,
) {
	return reducer.namedInferredVararg(
		NamedInferredVarargParameter,
		name,
		ellipsis)
}

func (reducer *Reducer) IgnoreInferredVarargToProperParameterDef(
	underscore *lr.TokenValue,
	ellipsis *lr.TokenValue,
) (
	*Parameter,
	error,
) {
	return reducer.namedInferredVararg(
		IgnoreInferredVarargParameter,
		underscore,
		ellipsis)
}

func (reducer *Reducer) UnnamedTypedArgToParameterDecl(
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

func (reducer *Reducer) UnnamedInferredVarargToParameterDecl(
	ellipsis *lr.TokenValue,
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

func (reducer *Reducer) UnnamedTypedVarargToParameterDecl(
	ellipsis *lr.TokenValue,
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

func (reducer *Reducer) namedInferredArg(
	kind ParameterKind,
	name *lr.TokenValue,
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

func (reducer *Reducer) NamedInferredArgToParameterDef(
	name *lr.TokenValue,
) (
	*Parameter,
	error,
) {
	return reducer.namedInferredArg(NamedInferredArgParameter, name)
}

func (reducer *Reducer) IgnoreInferredArgToParameterDef(
	underscore *lr.TokenValue,
) (
	*Parameter,
	error,
) {
	return reducer.namedInferredArg(IgnoreInferredArgParameter, underscore)
}

//
// ParameterList
//

func (reducer *Reducer) AddToProperParameterDeclList(
	list *ParameterList,
	comma *lr.TokenValue,
	parameter *Parameter,
) (
	*ParameterList,
	error,
) {
	list.ReduceAdd(comma, parameter)
	return list, nil
}

func (reducer *Reducer) ParameterDeclToProperParameterDeclList(
	parameter *Parameter,
) (
	*ParameterList,
	error,
) {
	list := NewParameterList()
	list.Add(parameter)
	return list, nil
}

func (reducer *Reducer) ImproperToParameterDeclList(
	list *ParameterList,
	comma *lr.TokenValue,
) (
	*ParameterList,
	error,
) {
	list.ReduceImproper(comma)
	return list, nil
}

func (reducer *Reducer) NilToParameterDeclList() (
	*ParameterList,
	error,
) {
	return NewParameterList(), nil
}

func (reducer *Reducer) ToParameterDecls(
	lparen *lr.TokenValue,
	list *ParameterList,
	rparen *lr.TokenValue,
) (
	*ParameterList,
	error,
) {
	list.ReduceMarkers(lparen, rparen)
	return list, nil
}

func (reducer *Reducer) AddToProperParameterDefList(
	list *ParameterList,
	comma *lr.TokenValue,
	parameter *Parameter,
) (
	*ParameterList,
	error,
) {
	list.ReduceAdd(comma, parameter)
	return list, nil
}

func (reducer *Reducer) ParameterDefToProperParameterDefList(
	parameter *Parameter,
) (
	*ParameterList,
	error,
) {
	list := NewParameterList()
	list.Add(parameter)
	return list, nil
}

func (reducer *Reducer) ImproperToParameterDefList(
	list *ParameterList,
	comma *lr.TokenValue,
) (
	*ParameterList,
	error,
) {
	list.ReduceImproper(comma)
	return list, nil
}

func (reducer *Reducer) NilToParameterDefList() (
	*ParameterList,
	error,
) {
	return NewParameterList(), nil
}

func (reducer *Reducer) ToParameterDefs(
	lparen *lr.TokenValue,
	list *ParameterList,
	rparen *lr.TokenValue,
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

func (Reducer) PositionalToArgument(
	expr Expression,
) (
	*Argument,
	error,
) {
	return NewPositionalArgument(expr), nil
}

func (Reducer) ColonExprToArgument(
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

func (Reducer) NamedAssignmentToArgument(
	name *lr.TokenValue,
	assign *lr.TokenValue,
	expr Expression,
) (
	*Argument,
	error,
) {
	return NewNamedArgument(name, assign, expr), nil
}

func (Reducer) VarargAssignmentToArgument(
	expr Expression,
	ellipsis *lr.TokenValue,
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

func (Reducer) SkipPatternToArgument(
	ellipsis *lr.TokenValue,
) (
	*Argument,
	error,
) {
	return NewSkipPatternArgument(ellipsis), nil
}

func (Reducer) PositionalToFieldVarPattern(
	expr Expression,
) (
	*Argument,
	error,
) {
	return NewPositionalArgument(expr), nil
}

func (Reducer) NamedAssignmentToFieldVarPattern(
	name *lr.TokenValue,
	assign *lr.TokenValue,
	expr Expression,
) (
	*Argument,
	error,
) {
	return NewNamedArgument(name, assign, expr), nil
}

func (Reducer) SkipPatternToFieldVarPattern(
	ellipsis *lr.TokenValue,
) (
	*Argument,
	error,
) {
	return NewSkipPatternArgument(ellipsis), nil
}

//
// Argument list
//

func (reducer *Reducer) AddToProperArguments(
	list *ArgumentList,
	comma *lr.TokenValue,
	arg *Argument,
) (
	*ArgumentList,
	error,
) {
	list.ReduceAdd(comma, arg)
	return list, nil
}

func (reducer *Reducer) ArgumentToProperArguments(
	arg *Argument,
) (
	*ArgumentList,
	error,
) {
	list := NewArgumentList()
	list.Add(arg)
	return list, nil
}

func (reducer *Reducer) ImproperToArguments(
	list *ArgumentList,
	comma *lr.TokenValue,
) (
	*ArgumentList,
	error,
) {
	list.ReduceImproper(comma)
	return list, nil
}

func (reducer *Reducer) NilToArguments() (
	*ArgumentList,
	error,
) {
	return NewArgumentList(), nil
}

func (reducer *Reducer) FieldVarPatternToProperFieldVarPatterns(
	pattern *Argument,
) (
	*ArgumentList,
	error,
) {
	list := NewArgumentList()
	list.Add(pattern)
	return list, nil
}

func (reducer *Reducer) AddToProperFieldVarPatterns(
	list *ArgumentList,
	comma *lr.TokenValue,
	pattern *Argument,
) (
	*ArgumentList,
	error,
) {
	list.ReduceAdd(comma, pattern)
	return list, nil
}

func (reducer *Reducer) ImproperToFieldVarPatterns(
	list *ArgumentList,
	comma *lr.TokenValue,
) (
	*ArgumentList,
	error,
) {
	list.ReduceImproper(comma)
	return list, nil
}
