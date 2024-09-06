package parser

import (
	"fmt"
)

type TypeArgumentList struct {
	StartPos Location
	EndPos   Location

	// The leading / trailing comments form the generic argument start/end marker
	// tokens $[ ]
	LeadingTrailingComments

	// MiddleComment is only used when the list is empty, but contains a comment,
	// e.g., Foo$[/*comment*/]
	MiddleComment CommentGroups

	TypeArguments []TypeExpression
}

var _ Node = &TypeArgumentList{}

func (list *TypeArgumentList) Loc() Location {
	return list.StartPos
}

func (list *TypeArgumentList) End() Location {
	return list.EndPos
}

func (list *TypeArgumentList) String() string {
	return list.TreeString("", "")
}

func (list *TypeArgumentList) TreeString(indent string, label string) string {
	result := fmt.Sprintf("%s%s[TypeArgumentList:", indent, label)
	if len(list.TypeArguments) == 0 {
		return result + "]"
	}

	for idx, arg := range list.TypeArguments {
		result += "\n" + arg.TreeString(indent+"  ", fmt.Sprintf("Arg %d=", idx))
	}

	result += "\n" + indent + "]"
	return result
}

type TypeArgumentListReducer struct{}

var _ GenericTypeArgumentsReducer = &TypeArgumentListReducer{}
var _ ProperTypeArgumentsReducer = &TypeArgumentListReducer{}
var _ TypeArgumentsReducer = &TypeArgumentListReducer{}

func (reducer *TypeArgumentListReducer) BindingToGenericTypeArguments(
	dollarLbracket TokenValue,
	list *TypeArgumentList,
	rbracket TokenValue,
) (
	*TypeArgumentList,
	error,
) {
	list.StartPos = dollarLbracket.Loc()
	list.EndPos = rbracket.End()

	list.LeadingComment = dollarLbracket.TakeLeading()

	if len(list.TypeArguments) > 0 {
		list.TypeArguments[0].PrependToLeading(dollarLbracket.TakeTrailing())
		list.TypeArguments[len(list.TypeArguments)-1].AppendToTrailing(
			rbracket.TakeLeading())
	} else {
		list.MiddleComment = dollarLbracket.TakeTrailing()
		list.MiddleComment.Append(rbracket.TakeLeading())
	}

	list.TrailingComment = rbracket.TakeTrailing()

	return list, nil
}

func (reducer *TypeArgumentListReducer) NilToGenericTypeArguments() (
	*TypeArgumentList,
	error,
) {
	return nil, nil
}

func (reducer *TypeArgumentListReducer) AddToProperTypeArguments(
	list *TypeArgumentList,
	comma TokenValue,
	arg TypeExpression,
) (
	*TypeArgumentList,
	error,
) {
	list.EndPos = arg.End()

	prevArg := list.TypeArguments[len(list.TypeArguments)-1]
	prevArg.AppendToTrailing(comma.TakeLeading())
	prevArg.AppendToTrailing(comma.TakeTrailing())

	list.TypeArguments = append(list.TypeArguments, arg)
	return list, nil
}

func (reducer *TypeArgumentListReducer) ValueTypeToProperTypeArguments(
	arg TypeExpression,
) (
	*TypeArgumentList,
	error,
) {
	return &TypeArgumentList{
		StartPos:      arg.Loc(),
		EndPos:        arg.End(),
		TypeArguments: []TypeExpression{arg},
	}, nil
}

func (reducer *TypeArgumentListReducer) ImproperToTypeArguments(
	list *TypeArgumentList,
	comma TokenValue,
) (
	*TypeArgumentList,
	error,
) {
	list.EndPos = comma.End()

	lastArg := list.TypeArguments[len(list.TypeArguments)-1]
	lastArg.AppendToTrailing(comma.TakeLeading())
	lastArg.AppendToTrailing(comma.TakeTrailing())

	return list, nil
}

func (reducer *TypeArgumentListReducer) NilToTypeArguments() (
	*TypeArgumentList,
	error,
) {
	return &TypeArgumentList{}, nil
}
