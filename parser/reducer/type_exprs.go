package reducer

import (
	"github.com/pattyshack/gt/lexutil"

	. "github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser"
)

//
// SliceTypeExpr
//

func (reducer *Reducer) ToSliceTypeExpr(
	lbracket *parser.TokenValue,
	value TypeExpression,
	rbracket *parser.TokenValue,
) (
	TypeExpression,
	error,
) {
	slice := &SliceTypeExpr{
		StartEndPos: NewStartEndPos(lbracket.Loc(), rbracket.End()),
		Value:       value,
	}

	slice.LeadingComment = lbracket.TakeLeading()
	value.PrependToLeading(lbracket.TakeTrailing())
	value.AppendToTrailing(rbracket.TakeLeading())
	slice.TrailingComment = rbracket.TakeTrailing()

	reducer.SliceTypeExprs = append(reducer.SliceTypeExprs, slice)
	return slice, nil
}

//
// ArrayTypeExpr
//

func (reducer *Reducer) ToArrayTypeExpr(
	lbracket *parser.TokenValue,
	value TypeExpression,
	comma *parser.TokenValue,
	size *parser.TokenValue,
	rbracket *parser.TokenValue,
) (
	TypeExpression,
	error,
) {
	array := &ArrayTypeExpr{
		StartEndPos: NewStartEndPos(lbracket.Loc(), rbracket.End()),
		Value:       value,
		Size:        size,
	}

	array.LeadingComment = lbracket.TakeLeading()
	value.PrependToLeading(lbracket.TakeTrailing())
	value.AppendToTrailing(comma.TakeLeading())
	array.Size.PrependToLeading(comma.TakeTrailing())
	array.Size.AppendToTrailing(rbracket.TakeLeading())
	array.TrailingComment = rbracket.TakeTrailing()

	reducer.ArrayTypeExprs = append(reducer.ArrayTypeExprs, array)
	return array, nil
}

//
// MapTypeExpr
//

func (reducer *Reducer) ToMapTypeExpr(
	lbracket *parser.TokenValue,
	key TypeExpression,
	semicolon *parser.TokenValue,
	value TypeExpression,
	rbracket *parser.TokenValue,
) (
	TypeExpression,
	error,
) {
	dict := &MapTypeExpr{
		StartEndPos: NewStartEndPos(lbracket.Loc(), rbracket.End()),
		Key:         key,
		Value:       value,
	}

	dict.LeadingComment = lbracket.TakeLeading()
	key.PrependToLeading(lbracket.TakeTrailing())
	key.AppendToTrailing(semicolon.TakeLeading())
	value.PrependToLeading(semicolon.TakeTrailing())
	value.AppendToTrailing(rbracket.TakeLeading())
	dict.TrailingComment = rbracket.TakeTrailing()

	reducer.MapTypeExprs = append(reducer.MapTypeExprs, dict)
	return dict, nil
}

//
// InferredTypeExpr
//

func (reducer *Reducer) DotToInferredTypeExpr(
	dot *parser.TokenValue,
) (
	TypeExpression,
	error,
) {
	expr := &InferredTypeExpr{
		StartEndPos:  NewStartEndPos(dot.Loc(), dot.End()),
		InferMutable: false,
	}
	expr.LeadingComment = dot.TakeLeading()
	expr.TrailingComment = dot.TakeTrailing()
	reducer.InferredTypeExprs = append(reducer.InferredTypeExprs, expr)
	return expr, nil
}

func (reducer *Reducer) UnderscoreToInferredTypeExpr(
	underscore *parser.TokenValue,
) (
	TypeExpression,
	error,
) {
	expr := &InferredTypeExpr{
		StartEndPos:  NewStartEndPos(underscore.Loc(), underscore.End()),
		InferMutable: true,
	}
	expr.LeadingComment = underscore.TakeLeading()
	expr.TrailingComment = underscore.TakeTrailing()
	reducer.InferredTypeExprs = append(reducer.InferredTypeExprs, expr)
	return expr, nil
}

//
// NamedTypeExpr
//

func (reducer *Reducer) toNamedTypeExpr(
	name *parser.TokenValue,
	genericArguments *GenericArgumentList,
) *NamedTypeExpr {
	var endPos lexutil.Location
	var trailing CommentGroups
	if genericArguments == nil {
		genericArguments = &GenericArgumentList{}
		endPos = name.End()
		trailing = name.TakeTrailing()
	} else {
		endPos = genericArguments.End()
		trailing = genericArguments.TakeTrailing()
	}
	named := &NamedTypeExpr{
		StartEndPos: NewStartEndPos(name.Loc(), endPos),
		LeadingTrailingComments: LeadingTrailingComments{
			TrailingComment: trailing,
		},
		Name:             name,
		GenericArguments: *genericArguments,
	}

	reducer.NamedTypeExprs = append(reducer.NamedTypeExprs, named)
	return named
}

func (reducer *Reducer) LocalToNamedTypeExpr(
	name *parser.TokenValue,
	genericArguments *GenericArgumentList,
) (
	TypeExpression,
	error,
) {
	named := reducer.toNamedTypeExpr(name, genericArguments)
	named.LeadingComment = name.TakeLeading()
	return named, nil
}

func (reducer *Reducer) ExternalToNamedTypeExpr(
	pkg *parser.TokenValue,
	dot *parser.TokenValue,
	name *parser.TokenValue,
	genericArguments *GenericArgumentList,
) (
	TypeExpression,
	error,
) {
	name.PrependToLeading(dot.TakeTrailing())
	name.PrependToLeading(dot.TakeLeading())
	name.PrependToLeading(pkg.TakeTrailing())

	named := reducer.toNamedTypeExpr(name, genericArguments)
	named.StartPos = pkg.Loc()
	named.Pkg = pkg.Value
	named.LeadingComment = pkg.TakeLeading()
	return named, nil
}

//
// UnaryTypeExpr
//

func (reducer *Reducer) ToPrefixUnaryTypeExpr(
	op *parser.TokenValue,
	operand TypeExpression,
) (
	TypeExpression,
	error,
) {
	expr := &UnaryTypeExpr{
		StartEndPos: NewStartEndPos(op.Loc(), operand.End()),
		Op:          UnaryTypeOp(op.Value),
		Operand:     operand,
	}

	expr.LeadingComment = op.TakeLeading()
	operand.PrependToLeading(op.TakeTrailing())
	expr.TrailingComment = operand.TakeTrailing()

	reducer.UnaryTypeExprs = append(reducer.UnaryTypeExprs, expr)
	return expr, nil
}

//
// BinaryTypeExpr
//

func (reducer *Reducer) ToBinaryTypeExpr(
	left TypeExpression,
	op *parser.TokenValue,
	right TypeExpression,
) (
	TypeExpression,
	error,
) {
	expr := &BinaryTypeExpr{
		StartEndPos: NewStartEndPos(left.Loc(), right.End()),
		Left:        left,
		Op:          BinaryTypeOp(op.Value),
		Right:       right,
	}

	expr.LeadingComment = left.TakeLeading()
	left.AppendToTrailing(op.TakeLeading())
	right.PrependToLeading(op.TakeTrailing())
	expr.TrailingComment = right.TakeTrailing()

	reducer.BinaryTypeExprs = append(reducer.BinaryTypeExprs, expr)
	return expr, nil
}

//
// Struct / Enum / Trait PropertiesType
//

func (reducer *Reducer) implicit(
	kind PropertiesKind,
	lparen *parser.TokenValue,
	properties *TypePropertyList,
	rparen *parser.TokenValue,
) *PropertiesTypeExpr {
	properties.ReduceMarkers(lparen, rparen)
	expr := &PropertiesTypeExpr{
		StartEndPos: NewStartEndPos(lparen.Loc(), rparen.End()),
		Kind:        kind,
		IsImplicit:  true,
		Properties:  *properties,
	}
	expr.LeadingComment = properties.TakeLeading()
	expr.TrailingComment = properties.TakeTrailing()

	return expr
}

func (reducer *Reducer) explicit(
	kind PropertiesKind,
	kw *parser.TokenValue,
	lparen *parser.TokenValue,
	properties *TypePropertyList,
	rparen *parser.TokenValue,
) *PropertiesTypeExpr {
	properties.ReduceMarkers(lparen, rparen)

	properties.PrependToLeading(kw.TakeTrailing())
	trailing := properties.TakeTrailing()

	expr := &PropertiesTypeExpr{
		StartEndPos: NewStartEndPos(kw.Loc(), rparen.End()),
		Kind:        kind,
		IsImplicit:  false,
		Properties:  *properties,
	}
	expr.LeadingComment = kw.TakeLeading()
	expr.TrailingComment = trailing

	return expr
}

func (reducer *Reducer) ToImplicitStructTypeExpr(
	lparen *parser.TokenValue,
	properties *TypePropertyList,
	rparen *parser.TokenValue,
) (
	TypeExpression,
	error,
) {
	expr := reducer.implicit(StructKind, lparen, properties, rparen)
	reducer.StructTypeExprs = append(reducer.StructTypeExprs, expr)
	return expr, nil
}

func (reducer *Reducer) ToExplicitStructTypeExpr(
	structKW *parser.TokenValue,
	lparen *parser.TokenValue,
	properties *TypePropertyList,
	rparen *parser.TokenValue,
) (
	TypeExpression,
	error,
) {
	expr := reducer.explicit(StructKind, structKW, lparen, properties, rparen)
	reducer.StructTypeExprs = append(reducer.StructTypeExprs, expr)
	return expr, nil
}

func (reducer *Reducer) ToTraitTypeExpr(
	trait *parser.TokenValue,
	lparen *parser.TokenValue,
	properties *TypePropertyList,
	rparen *parser.TokenValue,
) (
	TypeExpression,
	error,
) {
	expr := reducer.explicit(TraitKind, trait, lparen, properties, rparen)
	reducer.TraitTypeExprs = append(reducer.TraitTypeExprs, expr)
	return expr, nil
}

func (reducer *Reducer) ToImplicitEnumTypeExpr(
	lparen *parser.TokenValue,
	properties *TypePropertyList,
	rparen *parser.TokenValue,
) (
	TypeExpression,
	error,
) {
	expr := reducer.implicit(EnumKind, lparen, properties, rparen)
	reducer.EnumTypeExprs = append(reducer.EnumTypeExprs, expr)
	return expr, nil
}

func (reducer *Reducer) ToExplicitEnumTypeExpr(
	enum *parser.TokenValue,
	lparen *parser.TokenValue,
	properties *TypePropertyList,
	rparen *parser.TokenValue,
) (
	TypeExpression,
	error,
) {
	expr := reducer.explicit(EnumKind, enum, lparen, properties, rparen)
	reducer.EnumTypeExprs = append(reducer.EnumTypeExprs, expr)
	return expr, nil
}
