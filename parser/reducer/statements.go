package reducer

import (
	"fmt"

	. "github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser"
)

//
// StatementsExpr
//

func (reducer *Reducer) AddImplicitToProperStatementList(
	statements *StatementsExpr,
	newlines parser.TokenCount,
	statement Statement,
) (
	*StatementsExpr,
	error,
) {
	statements.ReduceAdd(&parser.TokenValue{}, statement)
	return statements, nil
}

func (reducer *Reducer) AddExplicitToProperStatementList(
	statements *StatementsExpr,
	semicolon *parser.TokenValue,
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
	return list, nil
}

func (reducer *Reducer) ImproperImplicitToStatementList(
	statements *StatementsExpr,
	newlines parser.TokenCount,
) (
	*StatementsExpr,
	error,
) {
	return statements, nil
}

func (reducer *Reducer) ImproperExplicitToStatementList(
	statements *StatementsExpr,
	semicolon *parser.TokenValue,
) (
	*StatementsExpr,
	error,
) {
	statements.ReduceImproper(semicolon)
	return statements, nil
}

func (reducer *Reducer) NilToStatementList() (*StatementsExpr, error) {
	return NewStatementsExpr(), nil
}

func (reducer *Reducer) ToStatements(
	lbrace *parser.TokenValue,
	list *StatementsExpr,
	rbrace *parser.TokenValue,
) (
	Expression,
	error,
) {
	list.ReduceMarkers(lbrace, rbrace)
	return list, nil
}

func (reducer *Reducer) LabelledToStatementsExpr(
	labelDecl *parser.TokenValue,
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

func (reducer *Reducer) SimpleStatementToTrailingSimpleStatement(
	stmt Statement,
) (
	*StatementsExpr,
	error,
) {
	expr := NewStatementsExpr()
	expr.Add(stmt)
	return expr, nil
}

func (reducer *Reducer) NilToTrailingSimpleStatement() (
	*StatementsExpr,
	error,
) {
	return NewStatementsExpr(), nil
}

//
// ImportClause
//

func (reducer *Reducer) StringLiteralToImportClause(
	pkg *parser.TokenValue,
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
	alias *parser.TokenValue,
	pkg *parser.TokenValue,
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
	alias *parser.TokenValue,
	pkg *parser.TokenValue,
) (
	*ImportClause,
	error,
) {
	return reducer.aliasImport(alias, pkg), nil
}

func (reducer *Reducer) UnusableImportToImportClause(
	underscore *parser.TokenValue,
	pkg *parser.TokenValue,
) (
	*ImportClause,
	error,
) {
	return reducer.aliasImport(underscore, pkg), nil
}

func (reducer *Reducer) ImportToLocalToImportClause(
	dot *parser.TokenValue,
	pkg *parser.TokenValue,
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
	importKW *parser.TokenValue,
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
	importKW *parser.TokenValue,
	lparen *parser.TokenValue,
	stmt *ImportStatement,
	rparen *parser.TokenValue,
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
	newlines parser.TokenCount,
	importClause *ImportClause,
) (
	*ImportStatement,
	error,
) {
	stmt.ReduceAdd(&parser.TokenValue{}, importClause)
	return stmt, nil
}

func (reducer *Reducer) AddExplicitToProperImportClauses(
	stmt *ImportStatement,
	comma *parser.TokenValue,
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
	newlines parser.TokenCount,
) (
	*ImportStatement,
	error,
) {
	return stmt, nil
}

func (reducer *Reducer) ExplicitToImportClauses(
	stmt *ImportStatement,
	comma *parser.TokenValue,
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
	op *parser.TokenValue,
) (
	Statement,
	error,
) {
	return NewJumpStatement(op, nil, nil), nil
}

func (reducer *Reducer) UnlabeledValuedToJumpStatement(
	op *parser.TokenValue,
	value Expression,
) (
	Statement,
	error,
) {
	return NewJumpStatement(op, nil, value), nil
}

func (reducer *Reducer) LabeledNoValueToJumpStatement(
	op *parser.TokenValue,
	label *parser.TokenValue,
) (
	Statement,
	error,
) {
	return NewJumpStatement(op, label, nil), nil
}

func (reducer *Reducer) LabeledValuedToJumpStatement(
	op *parser.TokenValue,
	label *parser.TokenValue,
	value Expression,
) (
	Statement,
	error,
) {
	return NewJumpStatement(op, label, value), nil
}

func (reducer *Reducer) FallthroughToJumpStatement(
	op *parser.TokenValue,
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
	unsafe *parser.TokenValue,
	less *parser.TokenValue,
	language *parser.TokenValue,
	greater *parser.TokenValue,
	verbatimSource *parser.TokenValue,
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
	caseKW *parser.TokenValue,
	casePatterns *ExpressionList,
	colon *parser.TokenValue,
	body *StatementsExpr,
) (
	Statement,
	error,
) {
	end := colon.End()
	if body == nil {
		body = NewStatementsExpr()
	} else {
		end = body.End()
	}

	leading := caseKW.TakeLeading()
	casePatterns.PrependToLeading(caseKW.TakeTrailing())
	casePatterns.AppendToTrailing(colon.TakeLeading())
	body.PrependToLeading(colon.TakeTrailing())

	stmt := &BranchStatement{
		StartEndPos:  NewStartEndPos(caseKW.Loc(), end),
		CasePatterns: *casePatterns,
		Body:         *body,
	}
	stmt.LeadingComment = leading

	return stmt, nil
}

func (reducer *Reducer) DefaultBranchToBranchStatement(
	defaultKW *parser.TokenValue,
	colon *parser.TokenValue,
	body *StatementsExpr,
) (
	Statement,
	error,
) {
	casePatterns := NewExpressionList()

	end := colon.End()
	if body == nil {
		body = NewStatementsExpr()
	} else {
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
		Body:         *body,
	}
	stmt.LeadingComment = leading

	return stmt, nil
}
