package reducer

import (
	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser/lr"
)

//
// AssignPattern
//

func (reducer *Reducer) toAssignPattern(
	pattern ast.Expression,
	op *lr.TokenValue,
	value ast.Expression,
) *ast.AssignPattern {
	expr := &ast.AssignPattern{
		StartEndPos: ast.NewStartEndPos(pattern.Loc(), value.End()),
		Kind:        ast.AssignKind(op.Value),
		Pattern:     pattern,
		Value:       value,
	}

	expr.LeadingComment = pattern.TakeLeading()
	pattern.AppendToTrailing(op.TakeLeading())
	value.PrependToLeading(op.TakeTrailing())
	expr.TrailingComment = value.TakeTrailing()

	return expr
}

func (reducer *Reducer) ToAssignStmt(
	pattern ast.Expression,
	assign *lr.TokenValue,
	value ast.Expression,
) (
	ast.Statement,
	error,
) {
	return reducer.toAssignPattern(pattern, assign, value), nil
}

func (reducer *Reducer) ToAssignSelectablePattern(
	exprList *ast.ExpressionList,
	assign *lr.TokenValue,
	value ast.Expression,
) (
	ast.Expression,
	error,
) {
	var expr ast.Expression
	if len(exprList.Elements) == 1 {
		expr = exprList.Elements[0]
	} else {
		args := make([]*ast.Argument, 0, len(exprList.Elements))
		for _, item := range exprList.Elements {
			args = append(args, ast.NewPositionalArgument(item))
		}

		expr = &ast.ImplicitStructExpr{
			StartEndPos:             exprList.StartEnd(),
			LeadingTrailingComments: exprList.TakeComments(),
			Kind:                    ast.ImproperImplicitStruct,
			Arguments:               args,
		}
	}

	return reducer.toAssignPattern(expr, assign, value), nil
}

func (reducer *Reducer) ToCasePatternExpr(
	caseKW *lr.TokenValue,
	switchablePatterns *ast.ExpressionList,
	assign *lr.TokenValue,
	value ast.Expression,
) (
	ast.Expression,
	error,
) {
	pattern := reducer.toAssignPattern(
		&ast.CasePatterns{
			StartEndPos: switchablePatterns.StartEnd(),
			Patterns:    switchablePatterns.Elements,
		},
		assign,
		value)

	pattern.PrependToLeading(caseKW.TakeTrailing())
	pattern.PrependToLeading(caseKW.TakeLeading())
	pattern.StartPos = caseKW.Loc()

	return pattern, nil
}

//
// AddrDeclPattern
//

func (reducer *Reducer) NewInferredToAddrDeclPattern(
	varType *lr.TokenValue,
	pattern ast.Expression,
) (
	ast.Expression,
	error,
) {
	return ast.NewAddrDeclPattern(varType, pattern, nil), nil
}

func (Reducer) NewTypedToAddrDeclPattern(
	varType *lr.TokenValue,
	pattern ast.Expression,
	typeExpr ast.TypeExpression,
) (
	ast.Expression,
	error,
) {
	return ast.NewAddrDeclPattern(varType, pattern, typeExpr), nil
}

//
// AssignToAddrPattern
//

func (Reducer) ToAssignToAddrPattern(
	greater *lr.TokenValue,
	pattern ast.Expression,
) (
	ast.Expression,
	error,
) {
	return ast.NewAssignToAddrPattern(greater, pattern), nil
}

//
// ExpressionList
//

func (Reducer) ToCasePatterns(
	expr ast.Expression,
) (
	*ast.ExpressionList,
	error,
) {
	list := ast.NewExpressionList()
	list.Add(expr)
	return list, nil
}

func (Reducer) SwitchableCasePatternToSwitchableCasePatterns(
	pattern ast.Expression,
) (
	*ast.ExpressionList,
	error,
) {
	list := ast.NewExpressionList()
	list.Add(pattern)
	return list, nil
}

func (Reducer) AddToSwitchableCasePatterns(
	list *ast.ExpressionList,
	comma *lr.TokenValue,
	pattern ast.Expression,
) (
	*ast.ExpressionList,
	error,
) {
	list.ReduceAdd(comma, pattern)
	return list, nil
}

//
// EnumPattern
//

func (Reducer) NamedToEnumPattern(
	dot *lr.TokenValue,
	enumValue *lr.TokenValue,
	pattern ast.Expression,
) (
	ast.Expression,
	error,
) {
	leading := dot.TakeLeading()
	leading.Append(dot.TakeTrailing())
	leading.Append(enumValue.TakeLeading())
	pattern.PrependToLeading(enumValue.TakeTrailing())
	trailing := pattern.TakeTrailing()

	enum := &ast.EnumPattern{
		StartEndPos: ast.NewStartEndPos(dot.Loc(), pattern.End()),
		EnumValue:   enumValue.Value,
		Pattern:     pattern,
	}
	enum.LeadingComment = leading
	enum.TrailingComment = trailing
	return enum, nil
}

func (Reducer) NamedUnitToEnumPattern(
	dot *lr.TokenValue,
	enumValue *lr.TokenValue,
) (
	ast.Expression,
	error,
) {
	leading := dot.TakeLeading()
	leading.Append(dot.TakeTrailing())
	leading.Append(enumValue.TakeLeading())
	trailing := enumValue.TakeTrailing()

	enum := &ast.EnumPattern{
		StartEndPos: ast.NewStartEndPos(dot.Loc(), enumValue.End()),
		EnumValue:   enumValue.Value,
		Pattern:     ast.NewImproperUnit(enumValue.End()),
	}
	enum.LeadingComment = leading
	enum.TrailingComment = trailing
	return enum, nil
}

func (Reducer) UnnamedStructToEnumPattern(
	dot *lr.TokenValue,
	pattern ast.Expression,
) (
	ast.Expression,
	error,
) {
	leading := dot.TakeLeading()
	leading.Append(dot.TakeTrailing())
	trailing := pattern.TakeTrailing()

	enum := &ast.EnumPattern{
		StartEndPos: ast.NewStartEndPos(dot.Loc(), pattern.End()),
		Pattern:     pattern,
	}
	enum.LeadingComment = leading
	enum.TrailingComment = trailing
	return enum, nil
}

//
// BlockAddrDeclStmt item
//

func (Reducer) DeclToBlockAddrDeclItem(
	addr ast.Expression,
	typeExpr ast.TypeExpression,
) (
	ast.Expression,
	error,
) {
	// NOTE: IsVar is set by the block
	return ast.NewAddrDeclPattern(nil, addr, typeExpr), nil
}

func (reducer *Reducer) InferredAssignToBlockAddrDeclItem(
	addr ast.Expression,
	assign *lr.TokenValue,
	value ast.Expression,
) (
	ast.Expression,
	error,
) {
	// NOTE: IsVar is set by the block
	pattern := ast.NewAddrDeclPattern(nil, addr, nil)
	return reducer.toAssignPattern(pattern, assign, value), nil
}

func (reducer *Reducer) TypedAssignToBlockAddrDeclItem(
	addr ast.Expression,
	typeExpr ast.TypeExpression,
	assign *lr.TokenValue,
	value ast.Expression,
) (
	ast.Expression,
	error,
) {
	// NOTE: IsVar is set by the block
	pattern := ast.NewAddrDeclPattern(nil, addr, typeExpr)
	return reducer.toAssignPattern(pattern, assign, value), nil
}
