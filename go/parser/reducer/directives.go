package reducer

import (
	"github.com/pattyshack/gt/parseutil"

	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser/lr"
)

//
// DirectivesDecl
//

func (Reducer) ToDirectivesDeclaration(
	pound *lr.TokenValue,
	lbracket *lr.TokenValue,
	list *ast.DirectiveList,
	rbracket *lr.TokenValue,
) (
	*ast.DirectivesDecl,
	error,
) {
	list.ReduceMarkers(lbracket, rbracket)

	leading := pound.TakeLeading()
	leading.Append(pound.TakeTrailing())
	leading.Append(list.TakeLeading())

	decl := &ast.DirectivesDecl{
		StartEndPos: parseutil.NewStartEndPos(pound.Loc(), list.End()),
		Directives:  list.Elements,
	}
	decl.LeadingComment = leading
	decl.TrailingComment = list.TakeTrailing()

	return decl, nil
}

//
// DirectiveList
//

func (Reducer) DirectiveToDirectives(
	directive *ast.Directive,
) (
	*ast.DirectiveList,
	error,
) {
	list := ast.NewDirectiveList()
	list.Add(directive)
	return list, nil
}

func (Reducer) AddToDirectives(
	list *ast.DirectiveList,
	semicolon *lr.TokenValue,
	directive *ast.Directive,
) (
	*ast.DirectiveList,
	error,
) {
	list.ReduceAdd(semicolon, directive)
	return list, nil
}

func (Reducer) NilToDirectiveExprs() (*ast.DirectiveExpressionList, error) {
	return nil, nil
}

//
// Directive
//

func (Reducer) NamedToDirective(
	name *lr.TokenValue,
	arguments *ast.DirectiveExpressionList,
) (
	*ast.Directive,
	error,
) {
	leading := name.TakeLeading()
	leading.Append(name.TakeTrailing())

	end := name.End()

	var args []ast.DirectiveExpression
	if arguments != nil {
		args = arguments.Elements
		end = arguments.End()
	}

	directive := &ast.Directive{
		StartEndPos: parseutil.NewStartEndPos(name.Loc(), end),
		Name:        name.Value,
		Arguments:   args,
	}
	directive.LeadingComment = leading

	return directive, nil
}

func (Reducer) CompoundNamedToDirective(
	name *lr.TokenValue,
	colon *lr.TokenValue,
	subName *lr.TokenValue,
	arguments *ast.DirectiveExpressionList,
) (
	*ast.Directive,
	error,
) {
	leading := name.TakeLeading()
	leading.Append(name.TakeTrailing())
	leading.Append(colon.TakeLeading())
	leading.Append(colon.TakeTrailing())
	leading.Append(subName.TakeLeading())
	leading.Append(subName.TakeTrailing())

	end := subName.End()

	var args []ast.DirectiveExpression
	if arguments != nil {
		args = arguments.Elements
		end = arguments.End()
	}

	directive := &ast.Directive{
		StartEndPos: parseutil.NewStartEndPos(name.Loc(), end),
		Name:        name.Value,
		SubName:     subName.Value,
		Arguments:   args,
	}
	directive.LeadingComment = leading

	return directive, nil
}

//
// DirectiveExpressionList
//

func (Reducer) DirectiveExprToProperDirectiveExprs(
	expr ast.DirectiveExpression,
) (
	*ast.DirectiveExpressionList,
	error,
) {
	list := ast.NewDirectiveExpressionList()
	list.Add(expr)
	return list, nil
}

func (Reducer) AddToProperDirectiveExprs(
	list *ast.DirectiveExpressionList,
	comma *lr.TokenValue,
	expr ast.DirectiveExpression,
) (
	*ast.DirectiveExpressionList,
	error,
) {
	list.ReduceAdd(comma, expr)
	return list, nil
}

//
// DirectiveExpression
//

func (Reducer) ToDirectiveValueExpr(
	value *lr.TokenValue,
) (
	ast.DirectiveExpression,
	error,
) {
	return &ast.DirectiveValueExpr{
		StartEndPos:             value.StartEndPos,
		LeadingTrailingComments: value.TakeComments(),
		Value:                   value.Value,
	}, nil
}

func (Reducer) GroupToDirectiveAtomExpr(
	lparen *lr.TokenValue,
	sub ast.DirectiveExpression,
	rparen *lr.TokenValue,
) (
	ast.DirectiveExpression,
	error,
) {
	expr := &ast.DirectiveGroupExpr{
		StartEndPos: parseutil.NewStartEndPos(lparen.Loc(), rparen.End()),
		Expr:        sub,
	}
	expr.LeadingComment = lparen.TakeLeading()
	sub.PrependToLeading(lparen.TakeTrailing())
	sub.AppendToTrailing(rparen.TakeLeading())
	expr.TrailingComment = rparen.TakeTrailing()

	return expr, nil
}

func (Reducer) NotToDirectiveNotExpr(
	not *lr.TokenValue,
	operand ast.DirectiveExpression,
) (
	ast.DirectiveExpression,
	error,
) {
	expr := &ast.DirectiveNotExpr{
		StartEndPos: parseutil.NewStartEndPos(not.Loc(), operand.End()),
		Operand:     operand,
	}
	expr.LeadingComment = not.TakeLeading()
	operand.PrependToLeading(not.TakeTrailing())
	expr.TrailingComment = operand.TakeTrailing()

	return expr, nil
}

func (Reducer) AndToDirectiveAndExpr(
	left ast.DirectiveExpression,
	and *lr.TokenValue,
	right ast.DirectiveExpression,
) (
	ast.DirectiveExpression,
	error,
) {
	expr := &ast.DirectiveAndExpr{
		StartEndPos: parseutil.NewStartEndPos(left.Loc(), right.End()),
		Left:        left,
		Right:       right,
	}
	expr.LeadingComment = left.TakeLeading()
	left.AppendToTrailing(and.TakeLeading())
	right.PrependToLeading(and.TakeTrailing())
	expr.TrailingComment = right.TakeTrailing()

	return expr, nil
}

func (Reducer) OrToDirectiveExpr(
	left ast.DirectiveExpression,
	or *lr.TokenValue,
	right ast.DirectiveExpression,
) (
	ast.DirectiveExpression,
	error,
) {
	expr := &ast.DirectiveOrExpr{
		StartEndPos: parseutil.NewStartEndPos(left.Loc(), right.End()),
		Left:        left,
		Right:       right,
	}
	expr.LeadingComment = left.TakeLeading()
	left.AppendToTrailing(or.TakeLeading())
	right.PrependToLeading(or.TakeTrailing())
	expr.TrailingComment = right.TakeTrailing()

	return expr, nil
}
