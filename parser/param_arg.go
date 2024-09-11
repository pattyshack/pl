package parser

import (
	"fmt"
)

//
// Parameter
//

type ParameterKind string

const (
	// NOTE: There is no implicit unnamed inferred arg variant
	NamedTypedArgParameter         = ParameterKind("named-typed-arg")
	NamedInferredArgParameter      = ParameterKind("named-inferred-arg")
	NamedTypedVarargParameter      = ParameterKind("named-typed-vararg")
	NamedInferredVarargParameter   = ParameterKind("named-inferred-vararg")
	UnnamedTypedArgParameter       = ParameterKind("unnamed-typed-arg")
	UnnamedTypedVarargParameter    = ParameterKind("unnamed-typed-vararg")
	UnnamedInferredVarargParameter = ParameterKind("unnamed-inferred-vararg")
	IgnoreTypedArgParameter        = ParameterKind("ignore-typed-arg")
	IgnoreInferredArgParameter     = ParameterKind("ignore-inferred-arg")
	IgnoreTypedVarargParameter     = ParameterKind("ignore-typed-vararg")
	IgnoreInferredVarargParameter  = ParameterKind("ignore-inferred-vararg")
)

type Parameter struct {
	StartEndPos
	LeadingTrailingComments

	Kind ParameterKind

	Name        *TokenValue // optional
	HasEllipsis bool
	Type        TypeExpression // optional
}

var _ Node = &Parameter{}

func (param Parameter) TreeString(indent string, label string) string {
	name := ""
	if param.Name != nil {
		name = param.Name.Value
	}
	result := fmt.Sprintf(
		"%s%s[Parameter: Kind=%v Name=%s HasEllipsis=%v\n",
		indent,
		label,
		param.Kind,
		name,
		param.HasEllipsis)
	if param.Type != nil {
		result += param.Type.TreeString(indent+"  ", "Type=")
		result += "\n" + indent + "]"
	} else {
		result += "]"
	}
	return result
}

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
		StartEndPos: newStartEndPos(name.Loc(), typeExpr.End()),
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
	name TokenValue,
	typeExpr TypeExpression,
) (
	*Parameter,
	error,
) {
	return reducer.namedTypedArg(NamedTypedArgParameter, &name, typeExpr)
}

func (reducer *ParameterReducerImpl) IgnoreTypedArgToProperParameterDef(
	underscore TokenValue,
	typeExpr TypeExpression,
) (
	*Parameter,
	error,
) {
	return reducer.namedTypedArg(IgnoreTypedArgParameter, &underscore, typeExpr)
}

func (reducer *ParameterReducerImpl) namedTypedVararg(
	kind ParameterKind,
	name *TokenValue,
	ellipsis TokenValue,
	typeExpr TypeExpression,
) (
	*Parameter,
	error,
) {
	param := &Parameter{
		StartEndPos: newStartEndPos(name.Loc(), typeExpr.End()),
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
	name TokenValue,
	ellipsis TokenValue,
	typeExpr TypeExpression,
) (
	*Parameter,
	error,
) {
	return reducer.namedTypedVararg(
		NamedTypedVarargParameter,
		&name,
		ellipsis,
		typeExpr)
}

func (reducer *ParameterReducerImpl) IgnoreTypedVarargToProperParameterDef(
	underscore TokenValue,
	ellipsis TokenValue,
	typeExpr TypeExpression,
) (
	*Parameter,
	error,
) {
	return reducer.namedTypedVararg(
		IgnoreTypedVarargParameter,
		&underscore,
		ellipsis,
		typeExpr)
}

func (reducer *ParameterReducerImpl) namedInferredVararg(
	kind ParameterKind,
	name *TokenValue,
	ellipsis TokenValue,
) (
	*Parameter,
	error,
) {
	param := &Parameter{
		StartEndPos: newStartEndPos(name.Loc(), ellipsis.End()),
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
	name TokenValue,
	ellipsis TokenValue,
) (
	*Parameter,
	error,
) {
	return reducer.namedInferredVararg(
		NamedInferredVarargParameter,
		&name,
		ellipsis)
}

func (reducer *ParameterReducerImpl) IgnoreInferredVarargToProperParameterDef(
	underscore TokenValue,
	ellipsis TokenValue,
) (
	*Parameter,
	error,
) {
	return reducer.namedInferredVararg(
		IgnoreInferredVarargParameter,
		&underscore,
		ellipsis)
}

func (reducer *ParameterReducerImpl) UnnamedTypedArgToParameterDecl(
	typeExpr TypeExpression,
) (
	*Parameter,
	error,
) {
	param := &Parameter{
		StartEndPos: newStartEndPos(typeExpr.Loc(), typeExpr.End()),
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
	ellipsis TokenValue,
) (
	*Parameter,
	error,
) {
	param := &Parameter{
		StartEndPos: newStartEndPos(ellipsis.Loc(), ellipsis.End()),
		Kind:        UnnamedInferredVarargParameter,
		Name:        nil,
		HasEllipsis: true,
		Type:        nil,
	}
	param.LeadingTrailingComments = ellipsis.LeadingTrailingComments
	return param, nil
}

func (reducer *ParameterReducerImpl) UnnamedTypedVarargToParameterDecl(
	ellipsis TokenValue,
	typeExpr TypeExpression,
) (
	*Parameter,
	error,
) {
	param := &Parameter{
		StartEndPos: newStartEndPos(ellipsis.Loc(), typeExpr.End()),
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
		StartEndPos: newStartEndPos(name.Loc(), name.End()),
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
	name TokenValue,
) (
	*Parameter,
	error,
) {
	return reducer.namedInferredArg(NamedInferredArgParameter, &name)
}

func (reducer *ParameterReducerImpl) IgnoreInferredArgToParameterDef(
	underscore TokenValue,
) (
	*Parameter,
	error,
) {
	return reducer.namedInferredArg(IgnoreInferredArgParameter, &underscore)
}

//
// ParameterList
//

type ParameterList = NodeList[*Parameter]

func NewParameterList() *ParameterList {
	return newNodeList[*Parameter]("ParameterList")
}

var _ Node = &ParameterList{}

type ParameterListReducerImpl struct{}

var _ ProperParameterDeclListReducer = &ParameterListReducerImpl{}
var _ ParameterDeclListReducer = &ParameterListReducerImpl{}
var _ ParameterDeclsReducer = &ParameterListReducerImpl{}
var _ ProperParameterDefListReducer = &ParameterListReducerImpl{}
var _ ParameterDefListReducer = &ParameterListReducerImpl{}
var _ ParameterDefsReducer = &ParameterListReducerImpl{}

func (reducer *ParameterListReducerImpl) AddToProperParameterDeclList(
	list *ParameterList,
	comma TokenValue,
	parameter *Parameter,
) (
	*ParameterList,
	error,
) {
	list.reduceAdd(comma, parameter)
	return list, nil
}

func (reducer *ParameterListReducerImpl) ParameterDeclToProperParameterDeclList(
	parameter *Parameter,
) (
	*ParameterList,
	error,
) {
	list := NewParameterList()
	list.add(parameter)
	return list, nil
}

func (reducer *ParameterListReducerImpl) ImproperToParameterDeclList(
	list *ParameterList,
	comma TokenValue,
) (
	*ParameterList,
	error,
) {
	list.reduceImproper(comma)
	return list, nil
}

func (reducer *ParameterListReducerImpl) NilToParameterDeclList() (
	*ParameterList,
	error,
) {
	return NewParameterList(), nil
}

func (reducer *ParameterListReducerImpl) ToParameterDecls(
	lparen TokenValue,
	list *ParameterList,
	rparen TokenValue,
) (
	*ParameterList,
	error,
) {
	list.reduceMarkers(lparen, rparen)
	return list, nil
}

func (reducer *ParameterListReducerImpl) AddToProperParameterDefList(
	list *ParameterList,
	comma TokenValue,
	parameter *Parameter,
) (
	*ParameterList,
	error,
) {
	list.reduceAdd(comma, parameter)
	return list, nil
}

func (reducer *ParameterListReducerImpl) ParameterDefToProperParameterDefList(
	parameter *Parameter,
) (
	*ParameterList,
	error,
) {
	list := NewParameterList()
	list.add(parameter)
	return list, nil
}

func (reducer *ParameterListReducerImpl) ImproperToParameterDefList(
	list *ParameterList,
	comma TokenValue,
) (
	*ParameterList,
	error,
) {
	list.reduceImproper(comma)
	return list, nil
}

func (reducer *ParameterListReducerImpl) NilToParameterDefList() (
	*ParameterList,
	error,
) {
	return NewParameterList(), nil
}

func (reducer *ParameterListReducerImpl) ToParameterDefs(
	lparen TokenValue,
	list *ParameterList,
	rparen TokenValue,
) (
	*ParameterList,
	error,
) {
	list.reduceMarkers(lparen, rparen)
	return list, nil
}

//
// Argument
//

type ArgumentKind string

const (
	PositionalArgument       = ArgumentKind("positional")
	NamedAssignmentArgument  = ArgumentKind("named-assigned")
	VarargAssignmentArgument = ArgumentKind("vararg-assigned")
	ColonExprArgument        = ArgumentKind("colon-expr")
	// Only used by patterns
	SkipPatternArgument = ArgumentKind("skip-pattern")
	// Only used by ColonExpr
	IsImplicitUnitArgument = ArgumentKind("implicit-unit")
)

type Argument struct {
	StartEndPos
	LeadingTrailingComments

	Kind ArgumentKind

	// Only set for named assignment
	OptionalName string

	// NOTE: Expr may be nil for SkipPatternArgument or IsImplicitUnit.
	Expr Expression

	HasEllipsis bool
}

var _ Node = &Argument{}

func (arg *Argument) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[Argument: Kind=%v OptionalName=%s HasEllipsis=%v",
		indent,
		label,
		arg.Kind,
		arg.OptionalName,
		arg.HasEllipsis)
	if arg.Expr != nil {
		result += "\n"
		result += arg.Expr.TreeString(indent+"  ", "")
		result += "\n" + indent + "]"
	} else {
		result += "]"
	}

	return result
}

type ArgumentReducerImpl struct {
}

var _ ArgumentReducer = ArgumentReducerImpl{}

func (ArgumentReducerImpl) PositionalToArgument(
	expr Expression,
) (
	*Argument,
	error,
) {
	arg := &Argument{
		StartEndPos: newStartEndPos(expr.Loc(), expr.End()),
		LeadingTrailingComments: LeadingTrailingComments{
			LeadingComment:  expr.TakeLeading(),
			TrailingComment: expr.TakeTrailing(),
		},
		Kind:        PositionalArgument,
		Expr:        expr,
		HasEllipsis: false,
	}

	return arg, nil
}

func (ArgumentReducerImpl) ColonExprToArgument(
	expr *ColonExpr,
) (
	*Argument,
	error,
) {
	return &Argument{
		StartEndPos: newStartEndPos(expr.Loc(), expr.End()),
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
	name TokenValue,
	assign TokenValue,
	expr Expression,
) (
	*Argument,
	error,
) {
	arg := &Argument{
		StartEndPos:  newStartEndPos(name.Loc(), expr.End()),
		Kind:         NamedAssignmentArgument,
		OptionalName: name.Value,
		Expr:         expr,
		HasEllipsis:  false,
	}

	arg.LeadingComment = name.TakeLeading()
	// prepend in reverse order
	expr.PrependToLeading(assign.TakeTrailing())
	expr.PrependToLeading(assign.TakeLeading())
	expr.PrependToLeading(name.TakeTrailing())
	arg.TrailingComment = expr.TakeTrailing()

	return arg, nil
}

func (ArgumentReducerImpl) VarargAssignmentToArgument(
	expr Expression,
	ellipsis TokenValue,
) (
	*Argument,
	error,
) {
	arg := &Argument{
		StartEndPos: newStartEndPos(expr.Loc(), ellipsis.End()),
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
	ellipsis TokenValue,
) (
	*Argument,
	error,
) {
	return &Argument{
		StartEndPos:             newStartEndPos(ellipsis.Loc(), ellipsis.End()),
		LeadingTrailingComments: ellipsis.LeadingTrailingComments,
		Kind:                    SkipPatternArgument,
		Expr:                    nil,
		HasEllipsis:             true,
	}, nil
}

//
// Argument list
//

type ArgumentList = NodeList[*Argument]

func NewArgumentList() *ArgumentList {
	return newNodeList[*Argument]("ArgumentList")
}

type ArgumentListReducerImpl struct{}

var _ ProperArgumentsReducer = &ArgumentListReducerImpl{}
var _ ArgumentsReducer = &ArgumentListReducerImpl{}

func (reducer *ArgumentListReducerImpl) AddToProperArguments(
	list *ArgumentList,
	comma TokenValue,
	arg *Argument,
) (
	*ArgumentList,
	error,
) {
	list.reduceAdd(comma, arg)
	return list, nil
}

func (reducer *ArgumentListReducerImpl) ArgumentToProperArguments(
	arg *Argument,
) (
	*ArgumentList,
	error,
) {
	list := NewArgumentList()
	list.add(arg)
	return list, nil
}

func (reducer *ArgumentListReducerImpl) ImproperToArguments(
	list *ArgumentList,
	comma TokenValue,
) (
	*ArgumentList,
	error,
) {
	list.reduceImproper(comma)
	return list, nil
}

func (reducer *ArgumentListReducerImpl) NilToArguments() (
	*ArgumentList,
	error,
) {
	return NewArgumentList(), nil
}
