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
		StartEndPos: ast.NewStartEndPos(lbracket.Loc(), rbracket.End()),
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
		StartEndPos: ast.NewStartEndPos(lbracket.Loc(), rbracket.End()),
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
		StartEndPos: ast.NewStartEndPos(lbracket.Loc(), rbracket.End()),
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
	dot *lr.TokenValue,
) (
	ast.TypeExpression,
	error,
) {
	expr := &ast.InferredTypeExpr{
		StartEndPos:  ast.NewStartEndPos(dot.Loc(), dot.End()),
		InferMutable: false,
	}
	expr.LeadingComment = dot.TakeLeading()
	expr.TrailingComment = dot.TakeTrailing()
	reducer.InferredTypeExprs = append(reducer.InferredTypeExprs, expr)
	return expr, nil
}

func (reducer *Reducer) UnderscoreToInferredTypeExpr(
	underscore *lr.TokenValue,
) (
	ast.TypeExpression,
	error,
) {
	expr := &ast.InferredTypeExpr{
		StartEndPos:  ast.NewStartEndPos(underscore.Loc(), underscore.End()),
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
	name *lr.TokenValue,
	genericArguments *ast.GenericArgumentList,
) *ast.NamedTypeExpr {
	var endPos lexutil.Location
	var trailing ast.CommentGroups
	if genericArguments == nil {
		genericArguments = ast.NewGenericArgumentList()
		endPos = name.End()
		trailing = name.TakeTrailing()
	} else {
		endPos = genericArguments.End()
		trailing = genericArguments.TakeTrailing()
	}
	named := &ast.NamedTypeExpr{
		StartEndPos: ast.NewStartEndPos(name.Loc(), endPos),
		LeadingTrailingComments: ast.LeadingTrailingComments{
			TrailingComment: trailing,
		},
		Name:             name,
		GenericArguments: *genericArguments,
	}

	reducer.NamedTypeExprs = append(reducer.NamedTypeExprs, named)
	return named
}

func (reducer *Reducer) LocalToNamedTypeExpr(
	name *lr.TokenValue,
	genericArguments *ast.GenericArgumentList,
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
	genericArguments *ast.GenericArgumentList,
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
// UnaryTypeExpr
//

func (reducer *Reducer) ToPrefixUnaryTypeExpr(
	op *lr.TokenValue,
	operand ast.TypeExpression,
) (
	ast.TypeExpression,
	error,
) {
	expr := &ast.UnaryTypeExpr{
		StartEndPos: ast.NewStartEndPos(op.Loc(), operand.End()),
		Op:          ast.UnaryTypeOp(op.Value),
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
	left ast.TypeExpression,
	op *lr.TokenValue,
	right ast.TypeExpression,
) (
	ast.TypeExpression,
	error,
) {
	expr := &ast.BinaryTypeExpr{
		StartEndPos: ast.NewStartEndPos(left.Loc(), right.End()),
		Left:        left,
		Op:          ast.BinaryTypeOp(op.Value),
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
	kind ast.PropertiesKind,
	lparen *lr.TokenValue,
	properties *ast.TypePropertyList,
	rparen *lr.TokenValue,
) *ast.PropertiesTypeExpr {
	properties.ReduceMarkers(lparen, rparen)
	expr := &ast.PropertiesTypeExpr{
		StartEndPos: ast.NewStartEndPos(lparen.Loc(), rparen.End()),
		Kind:        kind,
		IsImplicit:  true,
		Properties:  *properties,
	}
	expr.LeadingComment = properties.TakeLeading()
	expr.TrailingComment = properties.TakeTrailing()

	return expr
}

func (reducer *Reducer) explicit(
	kind ast.PropertiesKind,
	kw *lr.TokenValue,
	lparen *lr.TokenValue,
	properties *ast.TypePropertyList,
	rparen *lr.TokenValue,
) *ast.PropertiesTypeExpr {
	properties.ReduceMarkers(lparen, rparen)

	properties.PrependToLeading(kw.TakeTrailing())
	trailing := properties.TakeTrailing()

	expr := &ast.PropertiesTypeExpr{
		StartEndPos: ast.NewStartEndPos(kw.Loc(), rparen.End()),
		Kind:        kind,
		IsImplicit:  false,
		Properties:  *properties,
	}
	expr.LeadingComment = kw.TakeLeading()
	expr.TrailingComment = trailing

	return expr
}

func (reducer *Reducer) ToImplicitStructTypeExpr(
	lparen *lr.TokenValue,
	properties *ast.TypePropertyList,
	rparen *lr.TokenValue,
) (
	ast.TypeExpression,
	error,
) {
	expr := reducer.implicit(ast.StructKind, lparen, properties, rparen)
	reducer.StructTypeExprs = append(reducer.StructTypeExprs, expr)
	return expr, nil
}

func (reducer *Reducer) ToExplicitStructTypeExpr(
	structKW *lr.TokenValue,
	lparen *lr.TokenValue,
	properties *ast.TypePropertyList,
	rparen *lr.TokenValue,
) (
	ast.TypeExpression,
	error,
) {
	expr := reducer.explicit(ast.StructKind, structKW, lparen, properties, rparen)
	reducer.StructTypeExprs = append(reducer.StructTypeExprs, expr)
	return expr, nil
}

func (reducer *Reducer) ToTraitTypeExpr(
	trait *lr.TokenValue,
	lparen *lr.TokenValue,
	properties *ast.TypePropertyList,
	rparen *lr.TokenValue,
) (
	ast.TypeExpression,
	error,
) {
	expr := reducer.explicit(ast.TraitKind, trait, lparen, properties, rparen)
	reducer.TraitTypeExprs = append(reducer.TraitTypeExprs, expr)
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
	expr := reducer.implicit(ast.EnumKind, lparen, properties, rparen)
	reducer.EnumTypeExprs = append(reducer.EnumTypeExprs, expr)
	return expr, nil
}

func (reducer *Reducer) ToExplicitEnumTypeExpr(
	enum *lr.TokenValue,
	lparen *lr.TokenValue,
	properties *ast.TypePropertyList,
	rparen *lr.TokenValue,
) (
	ast.TypeExpression,
	error,
) {
	expr := reducer.explicit(ast.EnumKind, enum, lparen, properties, rparen)
	reducer.EnumTypeExprs = append(reducer.EnumTypeExprs, expr)
	return expr, nil
}
