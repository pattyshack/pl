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

func (reducer *Reducer) DefToGlobalVarDef(
	pattern ast.Expression,
	assign *lr.TokenValue,
	value ast.Expression,
) (
	ast.Definition,
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
		// TODO: move error check into pattern analyzer
		_, ok := expr.(*ast.EnumPattern)
		if ok {
			reducer.Emit("%s: unexpected case enum pattern", expr.Loc())
		}

	} else {
		args := make([]*ast.Argument, 0, len(exprList.Elements))
		for _, item := range exprList.Elements {
			// TODO: move error check into pattern analyzer
			_, ok := item.(*ast.EnumPattern)
			if ok {
				reducer.Emit("%s: unexpected case enum pattern", item.Loc())
			}
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
// AddressPattern
//

func (reducer *Reducer) NewInferredToAddressPattern(
	varType *lr.TokenValue,
	pattern ast.Expression,
) (
	ast.Expression,
	error,
) {
	return ast.NewAddressPattern(varType, pattern, nil), nil
}

func (Reducer) NewTypedToAddressPattern(
	varType *lr.TokenValue,
	pattern ast.Expression,
	typeExpr ast.TypeExpression,
) (
	ast.Expression,
	error,
) {
	return ast.NewAddressPattern(varType, pattern, typeExpr), nil
}

func (Reducer) ExistingToAddressPattern(
	greater *lr.TokenValue,
	pattern ast.Expression,
) (
	ast.Expression,
	error,
) {
	return ast.NewAddressPattern(greater, pattern, nil), nil
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

func (Reducer) MatchToEnumPattern(
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
		StartEndPos: ast.NewStartEndPos(dot.Loc(), enumValue.End()),
		EnumValue:   enumValue.Value,
		Pattern:     pattern,
	}
	enum.LeadingComment = leading
	enum.TrailingComment = trailing
	return enum, nil
}

func (Reducer) NondataToEnumPattern(
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

	pattern := &ast.EnumPattern{
		StartEndPos: ast.NewStartEndPos(dot.Loc(), enumValue.End()),
		EnumValue:   enumValue.Value,
	}
	pattern.LeadingComment = leading
	pattern.TrailingComment = trailing
	return pattern, nil
}

func (Reducer) AssignToNewAddressToEnumPattern(
	varType *lr.TokenValue,
	dot *lr.TokenValue,
	enumValue *lr.TokenValue,
	implicitStruct ast.Expression,
) (
	ast.Expression,
	error,
) {
	leading := varType.TakeLeading()
	leading.Append(varType.TakeTrailing())
	leading.Append(dot.TakeLeading())
	leading.Append(dot.TakeTrailing())
	leading.Append(enumValue.TakeLeading())

	varPattern := ast.NewAddressPattern(varType, implicitStruct, nil)
	varPattern.PrependToLeading(enumValue.TakeTrailing())

	trailing := implicitStruct.TakeTrailing()

	pattern := &ast.EnumPattern{
		StartEndPos: ast.NewStartEndPos(varType.Loc(), enumValue.End()),
		EnumValue:   enumValue.Value,
		Pattern:     varPattern,
	}
	pattern.LeadingComment = leading
	pattern.TrailingComment = trailing
	return pattern, nil
}

func (Reducer) AssignToExistingAddressToEnumPattern(
	greater *lr.TokenValue,
	dot *lr.TokenValue,
	enumValue *lr.TokenValue,
	implicitStruct ast.Expression,
) (
	ast.Expression,
	error,
) {
	leading := greater.TakeLeading()
	leading.Append(greater.TakeTrailing())
	leading.Append(dot.TakeLeading())
	leading.Append(dot.TakeTrailing())
	leading.Append(enumValue.TakeLeading())

	varPattern := ast.NewAddressPattern(greater, implicitStruct, nil)
	varPattern.PrependToLeading(enumValue.TakeTrailing())

	trailing := implicitStruct.TakeTrailing()

	pattern := &ast.EnumPattern{
		StartEndPos: ast.NewStartEndPos(greater.Loc(), enumValue.End()),
		EnumValue:   enumValue.Value,
		Pattern:     varPattern,
	}
	pattern.LeadingComment = leading
	pattern.TrailingComment = trailing
	return pattern, nil
}
