package parser

import (
	"fmt"

	. "github.com/pattyshack/pl/ast"
)

//
// StatementsExpr
//

type StatementsExprReducerImpl struct{}

var _ ProperStatementListReducer = &StatementsExprReducerImpl{}
var _ StatementListReducer = &StatementsExprReducerImpl{}
var _ StatementsReducer = &StatementsExprReducerImpl{}
var _ StatementsExprReducer = &StatementsExprReducerImpl{}
var _ TrailingSimpleStatementReducer = &StatementsExprReducerImpl{}

func (StatementsExprReducerImpl) AddImplicitToProperStatementList(
	statements *StatementsExpr,
	newlines TokenCount,
	statement Statement,
) (
	*StatementsExpr,
	error,
) {
	statements.ReduceAdd(&TokenValue{}, statement)
	return statements, nil
}

func (StatementsExprReducerImpl) AddExplicitToProperStatementList(
	statements *StatementsExpr,
	semicolon *TokenValue,
	statement Statement,
) (
	*StatementsExpr,
	error,
) {
	statements.ReduceAdd(semicolon, statement)
	return statements, nil
}

func (StatementsExprReducerImpl) StatementToProperStatementList(
	statement Statement,
) (
	*StatementsExpr,
	error,
) {
	list := NewStatementsExpr()
	list.Add(statement)
	return list, nil
}

func (StatementsExprReducerImpl) ImproperImplicitToStatementList(
	statements *StatementsExpr,
	newlines TokenCount,
) (
	*StatementsExpr,
	error,
) {
	return statements, nil
}

func (StatementsExprReducerImpl) ImproperExplicitToStatementList(
	statements *StatementsExpr,
	semicolon *TokenValue,
) (
	*StatementsExpr,
	error,
) {
	statements.ReduceImproper(semicolon)
	return statements, nil
}

func (StatementsExprReducerImpl) NilToStatementList() (*StatementsExpr, error) {
	return NewStatementsExpr(), nil
}

func (StatementsExprReducerImpl) ToStatements(
	lbrace *TokenValue,
	list *StatementsExpr,
	rbrace *TokenValue,
) (
	Expression,
	error,
) {
	list.ReduceMarkers(lbrace, rbrace)
	return list, nil
}

func (StatementsExprReducerImpl) LabelledToStatementsExpr(
	labelDecl *TokenValue,
	statementsExprOrParseError Expression,
) (
	Expression,
	error,
) {
	switch expr := statementsExprOrParseError.(type) {
	case *ParseErrorSymbol:
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

func (StatementsExprReducerImpl) SimpleStatementToTrailingSimpleStatement(
	stmt Statement,
) (
	*StatementsExpr,
	error,
) {
	expr := NewStatementsExpr()
	expr.Add(stmt)
	return expr, nil
}

func (StatementsExprReducerImpl) NilToTrailingSimpleStatement() (
	*StatementsExpr,
	error,
) {
	return NewStatementsExpr(), nil
}

//
// ImportClause
//

type ImportClauseReducerImpl struct {
	ImportClauses []*ImportClause
}

var _ ImportClauseReducer = &ImportClauseReducerImpl{}

func (reducer *ImportClauseReducerImpl) StringLiteralToImportClause(
	pkg *TokenValue,
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

func (reducer *ImportClauseReducerImpl) aliasImport(
	alias *TokenValue,
	pkg *TokenValue,
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

func (reducer *ImportClauseReducerImpl) AliasToImportClause(
	alias *TokenValue,
	pkg *TokenValue,
) (
	*ImportClause,
	error,
) {
	return reducer.aliasImport(alias, pkg), nil
}

func (reducer *ImportClauseReducerImpl) UnusableImportToImportClause(
	underscore *TokenValue,
	pkg *TokenValue,
) (
	*ImportClause,
	error,
) {
	return reducer.aliasImport(underscore, pkg), nil
}

func (reducer *ImportClauseReducerImpl) ImportToLocalToImportClause(
	dot *TokenValue,
	pkg *TokenValue,
) (
	*ImportClause,
	error,
) {
	return reducer.aliasImport(dot, pkg), nil
}

//
// ImportStatement
//

type ImportStatementReducerImpl struct{}

var _ ImportStatementReducer = &ImportStatementReducerImpl{}
var _ ProperImportClausesReducer = &ImportStatementReducerImpl{}
var _ ImportClausesReducer = &ImportStatementReducerImpl{}

func (ImportStatementReducerImpl) SingleToImportStatement(
	importKW *TokenValue,
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

func (ImportStatementReducerImpl) MultipleToImportStatement(
	importKW *TokenValue,
	lparen *TokenValue,
	stmt *ImportStatement,
	rparen *TokenValue,
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

func (ImportStatementReducerImpl) AddImplicitToProperImportClauses(
	stmt *ImportStatement,
	newlines TokenCount,
	importClause *ImportClause,
) (
	*ImportStatement,
	error,
) {
	stmt.ReduceAdd(&TokenValue{}, importClause)
	return stmt, nil
}

func (ImportStatementReducerImpl) AddExplicitToProperImportClauses(
	stmt *ImportStatement,
	comma *TokenValue,
	importClause *ImportClause,
) (
	*ImportStatement,
	error,
) {
	stmt.ReduceAdd(comma, importClause)
	return stmt, nil
}

func (ImportStatementReducerImpl) ImportClauseToProperImportClauses(
	importClause *ImportClause) (
	*ImportStatement,
	error,
) {
	stmt := NewImportStatement()
	stmt.Add(importClause)
	return stmt, nil
}

func (ImportStatementReducerImpl) ImplicitToImportClauses(
	stmt *ImportStatement,
	newlines TokenCount,
) (
	*ImportStatement,
	error,
) {
	return stmt, nil
}

func (ImportStatementReducerImpl) ExplicitToImportClauses(
	stmt *ImportStatement,
	comma *TokenValue,
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

type JumpStatementReducerImpl struct{}

var _ JumpStatementReducer = &JumpStatementReducerImpl{}

func (JumpStatementReducerImpl) UnlabeledNoValueToJumpStatement(
	op *TokenValue,
) (
	Statement,
	error,
) {
	return NewJumpStatement(op, nil, nil), nil
}

func (JumpStatementReducerImpl) UnlabeledValuedToJumpStatement(
	op *TokenValue,
	value Expression,
) (
	Statement,
	error,
) {
	return NewJumpStatement(op, nil, value), nil
}

func (JumpStatementReducerImpl) LabeledNoValueToJumpStatement(
	op *TokenValue,
	label *TokenValue,
) (
	Statement,
	error,
) {
	return NewJumpStatement(op, label, nil), nil
}

func (JumpStatementReducerImpl) LabeledValuedToJumpStatement(
	op *TokenValue,
	label *TokenValue,
	value Expression,
) (
	Statement,
	error,
) {
	return NewJumpStatement(op, label, value), nil
}

func (JumpStatementReducerImpl) FallthroughToJumpStatement(
	op *TokenValue,
) (
	Statement,
	error,
) {
	return NewJumpStatement(op, nil, nil), nil
}

//
// UnsafeStatement
//

type UnsafeStatementReducerImpl struct{}

var _ UnsafeStatementReducer = &UnsafeStatementReducerImpl{}

func (UnsafeStatementReducerImpl) ToUnsafeStatement(
	unsafe *TokenValue,
	less *TokenValue,
	language *TokenValue,
	greater *TokenValue,
	verbatimSource *TokenValue,
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

type BranchStatementReducerImpl struct{}

var _ BranchStatementReducer = &BranchStatementReducerImpl{}

func (BranchStatementReducerImpl) CaseBranchToBranchStatement(
	caseKW *TokenValue,
	casePatterns *ExpressionList,
	colon *TokenValue,
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

func (BranchStatementReducerImpl) DefaultBranchToBranchStatement(
	defaultKW *TokenValue,
	colon *TokenValue,
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
