package reducer

import (
	"github.com/pattyshack/gt/lexutil"

	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser/lr"
)

//
// SliceTypeExpr
//

func (reducer *Reducer) ToSliceTypeExpr(
	lbracket *lr.TokenValue,
	value ast.TypeExpression,
	rbracket *lr.TokenValue,
) (
	ast.TypeExpression,
	error,
) {
	slice := &ast.SliceTypeExpr{
		StartEndPos: lexutil.NewStartEndPos(lbracket.Loc(), rbracket.End()),
		Value:       value,
	}

	slice.LeadingComment = lbracket.TakeLeading()
	value.PrependToLeading(lbracket.TakeTrailing())
	value.AppendToTrailing(rbracket.TakeLeading())
	slice.TrailingComment = rbracket.TakeTrailing()

	return slice, nil
}

//
// ArrayTypeExpr
//

func (reducer *Reducer) ToArrayTypeExpr(
	lbracket *lr.TokenValue,
	value ast.TypeExpression,
	comma *lr.TokenValue,
	size *lr.TokenValue,
	rbracket *lr.TokenValue,
) (
	ast.TypeExpression,
	error,
) {
	array := &ast.ArrayTypeExpr{
		StartEndPos: lexutil.NewStartEndPos(lbracket.Loc(), rbracket.End()),
		Value:       value,
		Size:        size.Value,
	}

	array.LeadingComment = lbracket.TakeLeading()

	value.PrependToLeading(lbracket.TakeTrailing())
	value.AppendToTrailing(comma.TakeLeading())
	value.AppendToTrailing(comma.TakeTrailing())
	value.AppendToTrailing(size.TakeLeading())

	trailing := size.TakeTrailing()
	trailing.Append(rbracket.TakeLeading())
	trailing.Append(rbracket.TakeTrailing())
	array.TrailingComment = trailing

	return array, nil
}

//
// MapTypeExpr
//

func (reducer *Reducer) ToMapTypeExpr(
	lbracket *lr.TokenValue,
	key ast.TypeExpression,
	semicolon *lr.TokenValue,
	value ast.TypeExpression,
	rbracket *lr.TokenValue,
) (
	ast.TypeExpression,
	error,
) {
	dict := &ast.MapTypeExpr{
		StartEndPos: lexutil.NewStartEndPos(lbracket.Loc(), rbracket.End()),
		Key:         key,
		Value:       value,
	}

	dict.LeadingComment = lbracket.TakeLeading()
	key.PrependToLeading(lbracket.TakeTrailing())
	key.AppendToTrailing(semicolon.TakeLeading())
	value.PrependToLeading(semicolon.TakeTrailing())
	value.AppendToTrailing(rbracket.TakeLeading())
	dict.TrailingComment = rbracket.TakeTrailing()

	return dict, nil
}

//
// InferredTypeExpr
//

func (reducer *Reducer) ToInferredTypeExpr(
	underscore *lr.TokenValue,
) (
	ast.TypeExpression,
	error,
) {
	return &ast.InferredTypeExpr{
		StartEndPos:             underscore.StartEnd(),
		LeadingTrailingComments: underscore.TakeComments(),
	}, nil
}

//
// NamedTypeExpr
//

func (reducer *Reducer) toNamedTypeExpr(
	name *lr.TokenValue,
	genericArguments *ast.TypeExpressionList,
) *ast.NamedTypeExpr {
	leading := name.TakeLeading()
	var end ast.Node = name

	var parameters []ast.TypeExpression
	if genericArguments != nil {
		end = genericArguments
		parameters = genericArguments.Elements
	}

	named := &ast.NamedTypeExpr{
		StartEndPos: lexutil.NewStartEndPos(name.Loc(), end.End()),
		Name:        name.Value,
		Parameters:  parameters,
	}
	named.LeadingComment = leading
	named.TrailingComment = end.TakeTrailing()

	return named
}

func (reducer *Reducer) LocalToNamedTypeExpr(
	name *lr.TokenValue,
	genericArguments *ast.TypeExpressionList,
) (
	ast.TypeExpression,
	error,
) {
	named := reducer.toNamedTypeExpr(name, genericArguments)
	named.LeadingComment = name.TakeLeading()
	return named, nil
}

func (reducer *Reducer) ExternalToNamedTypeExpr(
	pkg *lr.TokenValue,
	dot *lr.TokenValue,
	name *lr.TokenValue,
	genericArguments *ast.TypeExpressionList,
) (
	ast.TypeExpression,
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
// RefTypeExpr
//

func (reducer *Reducer) ToRefTypeExpr(
	and *lr.TokenValue,
	value ast.TypeExpression,
) (
	ast.TypeExpression,
	error,
) {
	expr := &ast.RefTypeExpr{
		StartEndPos: lexutil.NewStartEndPos(and.Loc(), value.End()),
		Value:       value,
	}

	expr.LeadingComment = and.TakeLeading()
	value.PrependToLeading(and.TakeTrailing())
	expr.TrailingComment = value.TakeTrailing()

	return expr, nil
}

//
// DefaultEnumOpTypeExpr
//

func (reducer *Reducer) ToDefaultEnumOpTypeExpr(
	op *lr.TokenValue,
	enum ast.TypeExpression,
) (
	ast.TypeExpression,
	error,
) {
	expr := &ast.DefaultEnumOpTypeExpr{
		StartEndPos: lexutil.NewStartEndPos(op.Loc(), enum.End()),
		Op:          ast.DefaultEnumOp(op.Value),
		Enum:        enum,
	}

	expr.LeadingComment = op.TakeLeading()
	enum.PrependToLeading(op.TakeTrailing())
	expr.TrailingComment = enum.TakeTrailing()

	return expr, nil
}

//
// UnaryTraitOpTypeExpr
//

func (reducer *Reducer) ToUnaryTraitOpTypeExpr(
	op *lr.TokenValue,
	operand ast.TypeExpression,
) (
	ast.TypeExpression,
	error,
) {
	expr := &ast.UnaryTraitOpTypeExpr{
		StartEndPos: lexutil.NewStartEndPos(op.Loc(), operand.End()),
		Op:          ast.UnaryTraitOp(op.Value),
		Base:        operand,
	}

	expr.LeadingComment = op.TakeLeading()
	operand.PrependToLeading(op.TakeTrailing())
	expr.TrailingComment = operand.TakeTrailing()

	return expr, nil
}

//
// BinaryTraitOpTypeExpr
//

func (reducer *Reducer) ToBinaryTraitOpTypeExpr(
	left ast.TypeExpression,
	op *lr.TokenValue,
	right ast.TypeExpression,
) (
	ast.TypeExpression,
	error,
) {
	expr := &ast.BinaryTraitOpTypeExpr{
		StartEndPos: lexutil.NewStartEndPos(left.Loc(), right.End()),
		Left:        left,
		Op:          ast.BinaryTraitOp(op.Value),
		Right:       right,
	}

	expr.LeadingComment = left.TakeLeading()
	left.AppendToTrailing(op.TakeLeading())
	right.PrependToLeading(op.TakeTrailing())
	expr.TrailingComment = right.TakeTrailing()

	return expr, nil
}

//
// Struct / Enum / Trait PropertiesType
//

func (reducer *Reducer) implicitPropertiesTypeExpr(
	kind ast.PropertiesKind,
	lparen *lr.TokenValue,
	properties *ast.TypePropertyList,
	rparen *lr.TokenValue,
) *ast.PropertiesTypeExpr {
	if properties == nil {
		properties = ast.NewTypePropertyList()
	}
	properties.ReduceMarkers(lparen, rparen)

	return &ast.PropertiesTypeExpr{
		StartEndPos:             properties.StartEnd(),
		LeadingTrailingComments: properties.TakeComments(),
		Kind:                    kind,
		IsImplicit:              true,
		Properties:              properties.Elements,
	}
}

func (reducer *Reducer) ToImplicitStructTypeExpr(
	lparen *lr.TokenValue,
	properties *ast.TypePropertyList,
	rparen *lr.TokenValue,
) (
	ast.TypeExpression,
	error,
) {
	expr := reducer.implicitPropertiesTypeExpr(
		ast.StructKind,
		lparen,
		properties,
		rparen)
	return expr, nil
}

func (reducer *Reducer) ToImplicitEnumTypeExpr(
	lparen *lr.TokenValue,
	properties *ast.TypePropertyList,
	rparen *lr.TokenValue,
) (
	ast.TypeExpression,
	error,
) {
	expr := reducer.implicitPropertiesTypeExpr(
		ast.EnumKind,
		lparen,
		properties,
		rparen)
	return expr, nil
}

func (reducer *Reducer) ToPropertiesTypeExpr(
	kw *lr.TokenValue,
	lparen *lr.TokenValue,
	properties *ast.TypePropertyList,
	rparen *lr.TokenValue,
) (
	ast.TypeExpression,
	error,
) {
	if properties == nil {
		properties = ast.NewTypePropertyList()
	}
	properties.ReduceMarkers(lparen, rparen)

	properties.PrependToLeading(kw.TakeTrailing())
	properties.PrependToLeading(kw.TakeLeading())

	return &ast.PropertiesTypeExpr{
		StartEndPos:             lexutil.NewStartEndPos(kw.Loc(), rparen.End()),
		LeadingTrailingComments: properties.TakeComments(),
		Kind:                    ast.PropertiesKind(kw.Value),
		IsImplicit:              false,
		Properties:              properties.Elements,
	}, nil
}
