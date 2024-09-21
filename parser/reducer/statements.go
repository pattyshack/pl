package reducer

import (
	"fmt"

	. "github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser/lr"
)

//
// StatementsExpr
//

func (reducer *Reducer) AddImplicitToProperStatementList(
	statements *StatementsExpr,
	newlines lr.TokenCount,
	statement Statement,
) (
	*StatementsExpr,
	error,
) {
	statements.ReduceAdd(&lr.TokenValue{}, statement)
	return statements, nil
}

func (reducer *Reducer) AddExplicitToProperStatementList(
	statements *StatementsExpr,
	semicolon *lr.TokenValue,
	statement Statement,
) (
	*StatementsExpr,
	error,
) {
	statements.ReduceAdd(semicolon, statement)
	return statements, nil
}

func (reducer *Reducer) StatementToProperStatementList(
	statement Statement,
) (
	*StatementsExpr,
	error,
) {
	list := NewStatementsExpr()
	list.Add(statement)
	reducer.StatementsExprs = append(reducer.StatementsExprs, list)
	return list, nil
}

func (reducer *Reducer) ImproperImplicitToStatementList(
	statements *StatementsExpr,
	newlines lr.TokenCount,
) (
	*StatementsExpr,
	error,
) {
	return statements, nil
}

func (reducer *Reducer) ImproperExplicitToStatementList(
	statements *StatementsExpr,
	semicolon *lr.TokenValue,
) (
	*StatementsExpr,
	error,
) {
	statements.ReduceImproper(semicolon)
	return statements, nil
}

func (reducer *Reducer) NilToStatementList() (*StatementsExpr, error) {
	list := NewStatementsExpr()
	reducer.StatementsExprs = append(reducer.StatementsExprs, list)
	return list, nil
}

func (reducer *Reducer) ToStatements(
	lbrace *lr.TokenValue,
	list *StatementsExpr,
	rbrace *lr.TokenValue,
) (
	Expression,
	error,
) {
	list.ReduceMarkers(lbrace, rbrace)
	return list, nil
}

func (reducer *Reducer) LabelledToStatementsExpr(
	labelDecl *lr.TokenValue,
	statementsExprOrParseError Expression,
) (
	Expression,
	error,
) {
	switch expr := statementsExprOrParseError.(type) {
	case *ParseErrorNode:
		return expr, nil
	case *StatementsExpr:
		expr.StartPos = labelDecl.Loc()
		expr.LabelDecl = labelDecl.Value
		expr.PrependToLeading(labelDecl.TakeTrailing())
		expr.PrependToLeading(labelDecl.TakeLeading())
		return expr, nil
	}

	panic(fmt.Sprintf("Unexpected expression: %v", statementsExprOrParseError))
}

func (reducer *Reducer) StatementToTrailingStatement(
	stmt Statement,
) (
	*StatementsExpr,
	error,
) {
	expr := NewStatementsExpr()
	expr.Add(stmt)
	reducer.StatementsExprs = append(reducer.StatementsExprs, expr)
	return expr, nil
}

func (reducer *Reducer) NilToTrailingStatement() (
	*StatementsExpr,
	error,
) {
	expr := NewStatementsExpr()
	reducer.StatementsExprs = append(reducer.StatementsExprs, expr)
	return expr, nil
}

//
// ImportClause
//

func (reducer *Reducer) StringLiteralToImportClause(
	pkg *lr.TokenValue,
) (
	*ImportClause,
	error,
) {
	clause := &ImportClause{
		StartEndPos:             pkg.StartEndPos,
		LeadingTrailingComments: pkg.LeadingTrailingComments,
		Package:                 pkg.Value,
	}

	reducer.ImportClauses = append(reducer.ImportClauses, clause)
	return clause, nil
}

func (reducer *Reducer) aliasImport(
	alias *lr.TokenValue,
	pkg *lr.TokenValue,
) *ImportClause {
	clause := &ImportClause{
		StartEndPos: NewStartEndPos(alias.Loc(), pkg.End()),
		Alias:       alias.Value,
		Package:     pkg.Value,
	}

	leading := alias.TakeLeading()
	leading.Append(alias.TakeTrailing())
	leading.Append(pkg.TakeLeading())

	clause.LeadingComment = leading
	clause.TrailingComment = pkg.TakeTrailing()

	reducer.ImportClauses = append(reducer.ImportClauses, clause)
	return clause
}

func (reducer *Reducer) AliasToImportClause(
	alias *lr.TokenValue,
	pkg *lr.TokenValue,
) (
	*ImportClause,
	error,
) {
	return reducer.aliasImport(alias, pkg), nil
}

func (reducer *Reducer) UnusableImportToImportClause(
	underscore *lr.TokenValue,
	pkg *lr.TokenValue,
) (
	*ImportClause,
	error,
) {
	return reducer.aliasImport(underscore, pkg), nil
}

func (reducer *Reducer) ImportToLocalToImportClause(
	dot *lr.TokenValue,
	pkg *lr.TokenValue,
) (
	*ImportClause,
	error,
) {
	return reducer.aliasImport(dot, pkg), nil
}

//
// ImportStatement
//

func (reducer *Reducer) SingleToImportStatement(
	importKW *lr.TokenValue,
	importClause *ImportClause,
) (
	Statement,
	error,
) {
	stmt := NewImportStatement()
	stmt.Add(importClause)

	stmt.StartPos = importKW.Loc()
	stmt.PrependToLeading(importKW.TakeTrailing())
	stmt.PrependToLeading(importKW.TakeLeading())
	return stmt, nil
}

func (reducer *Reducer) MultipleToImportStatement(
	importKW *lr.TokenValue,
	lparen *lr.TokenValue,
	stmt *ImportStatement,
	rparen *lr.TokenValue,
) (
	Statement,
	error,
) {
	stmt.ReduceMarkers(lparen, rparen)

	stmt.StartPos = importKW.Loc()
	stmt.PrependToLeading(importKW.TakeTrailing())
	stmt.PrependToLeading(importKW.TakeLeading())
	return stmt, nil
}

func (reducer *Reducer) AddImplicitToProperImportClauses(
	stmt *ImportStatement,
	newlines lr.TokenCount,
	importClause *ImportClause,
) (
	*ImportStatement,
	error,
) {
	stmt.ReduceAdd(&lr.TokenValue{}, importClause)
	return stmt, nil
}

func (reducer *Reducer) AddExplicitToProperImportClauses(
	stmt *ImportStatement,
	comma *lr.TokenValue,
	importClause *ImportClause,
) (
	*ImportStatement,
	error,
) {
	stmt.ReduceAdd(comma, importClause)
	return stmt, nil
}

func (reducer *Reducer) ImportClauseToProperImportClauses(
	importClause *ImportClause) (
	*ImportStatement,
	error,
) {
	stmt := NewImportStatement()
	stmt.Add(importClause)
	return stmt, nil
}

func (reducer *Reducer) ImplicitToImportClauses(
	stmt *ImportStatement,
	newlines lr.TokenCount,
) (
	*ImportStatement,
	error,
) {
	return stmt, nil
}

func (reducer *Reducer) ExplicitToImportClauses(
	stmt *ImportStatement,
	comma *lr.TokenValue,
) (
	*ImportStatement,
	error,
) {
	stmt.ReduceImproper(comma)
	return stmt, nil
}

//
// JumpStatement
//

func (reducer *Reducer) UnlabeledNoValueToJumpStatement(
	op *lr.TokenValue,
) (
	Statement,
	error,
) {
	return NewJumpStatement(op, nil, nil), nil
}

func (reducer *Reducer) UnlabeledValuedToJumpStatement(
	op *lr.TokenValue,
	value Expression,
) (
	Statement,
	error,
) {
	return NewJumpStatement(op, nil, value), nil
}

func (reducer *Reducer) LabeledNoValueToJumpStatement(
	op *lr.TokenValue,
	label *lr.TokenValue,
) (
	Statement,
	error,
) {
	return NewJumpStatement(op, label, nil), nil
}

func (reducer *Reducer) LabeledValuedToJumpStatement(
	op *lr.TokenValue,
	label *lr.TokenValue,
	value Expression,
) (
	Statement,
	error,
) {
	return NewJumpStatement(op, label, value), nil
}

func (reducer *Reducer) FallthroughToJumpStatement(
	op *lr.TokenValue,
) (
	Statement,
	error,
) {
	return NewJumpStatement(op, nil, nil), nil
}

//
// UnsafeStatement
//

func (reducer *Reducer) ToUnsafeStatement(
	unsafe *lr.TokenValue,
	less *lr.TokenValue,
	language *lr.TokenValue,
	greater *lr.TokenValue,
	verbatimSource *lr.TokenValue,
) (
	*UnsafeStatement,
	error,
) {
	leading := unsafe.TakeLeading()
	leading.Append(unsafe.TakeTrailing())
	leading.Append(less.TakeLeading())
	leading.Append(less.TakeTrailing())
	leading.Append(language.TakeLeading())
	leading.Append(language.TakeTrailing())
	leading.Append(greater.TakeLeading())
	leading.Append(greater.TakeTrailing())
	leading.Append(verbatimSource.TakeLeading())
	stmt := &UnsafeStatement{
		StartEndPos:    NewStartEndPos(unsafe.Loc(), verbatimSource.End()),
		Language:       language.Value,
		VerbatimSource: verbatimSource.Value,
	}
	stmt.LeadingComment = leading
	stmt.TrailingComment = verbatimSource.TakeTrailing()
	return stmt, nil
}

//
// BranchStatement
//

func (reducer *Reducer) CaseBranchToBranchStatement(
	caseKW *lr.TokenValue,
	casePatterns *ExpressionList,
	colon *lr.TokenValue,
	body *StatementsExpr,
) (
	Statement,
	error,
) {
	end := colon.End()
	if len(body.Elements) > 0 {
		// body.End() is invalid for "nil trailing statement"
		end = body.End()
	}

	leading := caseKW.TakeLeading()
	casePatterns.PrependToLeading(caseKW.TakeTrailing())
	casePatterns.AppendToTrailing(colon.TakeLeading())
	body.PrependToLeading(colon.TakeTrailing())

	stmt := &BranchStatement{
		StartEndPos:  NewStartEndPos(caseKW.Loc(), end),
		CasePatterns: *casePatterns,
		Body:         body,
	}
	stmt.LeadingComment = leading

	return stmt, nil
}

func (reducer *Reducer) DefaultBranchToBranchStatement(
	defaultKW *lr.TokenValue,
	colon *lr.TokenValue,
	body *StatementsExpr,
) (
	Statement,
	error,
) {
	casePatterns := NewExpressionList()

	end := colon.End()
	if len(body.Elements) > 0 {
		// body.End() is invalid for "nil trailing statement"
		end = body.End()
	}

	leading := defaultKW.TakeLeading()
	casePatterns.PrependToLeading(defaultKW.TakeTrailing())
	casePatterns.AppendToTrailing(colon.TakeLeading())
	body.PrependToLeading(colon.TakeTrailing())

	stmt := &BranchStatement{
		StartEndPos:  NewStartEndPos(defaultKW.Loc(), end),
		IsDefault:    true,
		CasePatterns: *casePatterns,
		Body:         body,
	}
	stmt.LeadingComment = leading

	return stmt, nil
}
