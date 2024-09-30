package reducer

import (
	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser/lr"
)

//
// StatementsExpr / StatementList
//

func (reducer *Reducer) AddImplicitToProperStatementList(
	statements *ast.StatementList,
	newlines lr.TokenCount,
	statement ast.Statement,
) (
	*ast.StatementList,
	error,
) {
	statements.Add(statement)
	return statements, nil
}

func (reducer *Reducer) AddExplicitToProperStatementList(
	statements *ast.StatementList,
	semicolon *lr.TokenValue,
	statement ast.Statement,
) (
	*ast.StatementList,
	error,
) {
	statements.ReduceAdd(semicolon, statement)
	return statements, nil
}

func (reducer *Reducer) StatementToProperStatementList(
	statement ast.Statement,
) (
	*ast.StatementList,
	error,
) {
	list := ast.NewStatementList()
	list.Add(statement)
	return list, nil
}

func (reducer *Reducer) ImproperImplicitToStatementList(
	statements *ast.StatementList,
	newlines lr.TokenCount,
) (
	*ast.StatementList,
	error,
) {
	return statements, nil
}

func (reducer *Reducer) ImproperExplicitToStatementList(
	statements *ast.StatementList,
	semicolon *lr.TokenValue,
) (
	*ast.StatementList,
	error,
) {
	statements.ReduceImproper(semicolon)
	return statements, nil
}

func (reducer *Reducer) NilToStatementList() (*ast.StatementList, error) {
	return nil, nil
}

func (reducer *Reducer) ToStatements(
	lbrace *lr.TokenValue,
	list *ast.StatementList,
	rbrace *lr.TokenValue,
) (
	*ast.StatementsExpr,
	error,
) {
	if list == nil {
		list = ast.NewStatementList()
	}
	list.ReduceMarkers(lbrace, rbrace)

	expr := &ast.StatementsExpr{
		StartEndPos:             list.StartEndPos,
		LeadingTrailingComments: list.TakeComments(),
		Statements:              list.Elements,
	}
	expr.LeadingComment.Append(list.MiddleComment)

	return expr, nil
}

func (reducer *Reducer) LabelledToStatementsExpr(
	labelDecl *lr.TokenValue,
	expr *ast.StatementsExpr,
) (
	ast.Expression,
	error,
) {
	expr.StartPos = labelDecl.Loc()
	expr.LabelDecl = labelDecl.Value
	expr.PrependToLeading(labelDecl.TakeTrailing())
	expr.PrependToLeading(labelDecl.TakeLeading())
	return expr, nil
}

//
// ImportClause
//

func (reducer *Reducer) StringLiteralToImportClause(
	pkg *lr.TokenValue,
) (
	*ast.ImportClause,
	error,
) {
	clause := &ast.ImportClause{
		StartEndPos:             pkg.StartEndPos,
		LeadingTrailingComments: pkg.LeadingTrailingComments,
		Package:                 pkg.Value,
	}

	return clause, nil
}

func (reducer *Reducer) aliasImport(
	alias *lr.TokenValue,
	pkg *lr.TokenValue,
) *ast.ImportClause {
	clause := &ast.ImportClause{
		StartEndPos: ast.NewStartEndPos(alias.Loc(), pkg.End()),
		Alias:       alias.Value,
		Package:     pkg.Value,
	}

	leading := alias.TakeLeading()
	leading.Append(alias.TakeTrailing())
	leading.Append(pkg.TakeLeading())

	clause.LeadingComment = leading
	clause.TrailingComment = pkg.TakeTrailing()

	return clause
}

func (reducer *Reducer) AliasToImportClause(
	alias *lr.TokenValue,
	pkg *lr.TokenValue,
) (
	*ast.ImportClause,
	error,
) {
	return reducer.aliasImport(alias, pkg), nil
}

func (reducer *Reducer) UnusableImportToImportClause(
	underscore *lr.TokenValue,
	pkg *lr.TokenValue,
) (
	*ast.ImportClause,
	error,
) {
	return reducer.aliasImport(underscore, pkg), nil
}

func (reducer *Reducer) ImportToLocalToImportClause(
	dot *lr.TokenValue,
	pkg *lr.TokenValue,
) (
	*ast.ImportClause,
	error,
) {
	return reducer.aliasImport(dot, pkg), nil
}

//
// ImportStatement
//

func (reducer *Reducer) SingleToImportStatement(
	importKW *lr.TokenValue,
	importClause *ast.ImportClause,
) (
	ast.Statement,
	error,
) {
	leading := importKW.TakeLeading()
	importClause.PrependToLeading(importKW.TakeTrailing())
	trailing := importClause.TakeTrailing()

	list := ast.NewImportClauseList()
	list.Add(importClause)

	stmt := &ast.ImportStatement{
		StartEndPos:   ast.NewStartEndPos(importKW.Loc(), importClause.End()),
		ImportClauses: []*ast.ImportClause{importClause},
	}
	stmt.LeadingComment = leading
	stmt.TrailingComment = trailing

	return stmt, nil
}

func (reducer *Reducer) MultipleToImportStatement(
	importKW *lr.TokenValue,
	lparen *lr.TokenValue,
	clauses *ast.ImportClauseList,
	rparen *lr.TokenValue,
) (
	ast.Statement,
	error,
) {
	clauses.ReduceMarkers(lparen, rparen)

	leading := importKW.TakeLeading()
	leading.Append(importKW.TakeTrailing())
	leading.Append(clauses.TakeLeading())
	trailing := clauses.TakeTrailing()

	stmt := &ast.ImportStatement{
		StartEndPos:   ast.NewStartEndPos(importKW.Loc(), clauses.End()),
		ImportClauses: clauses.Elements,
	}
	stmt.LeadingComment = leading
	stmt.TrailingComment = trailing

	return stmt, nil
}

func (reducer *Reducer) AddImplicitToProperImportClauses(
	clauses *ast.ImportClauseList,
	newlines lr.TokenCount,
	importClause *ast.ImportClause,
) (
	*ast.ImportClauseList,
	error,
) {
	clauses.Add(importClause)
	return clauses, nil
}

func (reducer *Reducer) AddExplicitToProperImportClauses(
	clauses *ast.ImportClauseList,
	comma *lr.TokenValue,
	importClause *ast.ImportClause,
) (
	*ast.ImportClauseList,
	error,
) {
	clauses.ReduceAdd(comma, importClause)
	return clauses, nil
}

func (reducer *Reducer) ImportClauseToProperImportClauses(
	importClause *ast.ImportClause,
) (
	*ast.ImportClauseList,
	error,
) {
	clauses := ast.NewImportClauseList()
	clauses.Add(importClause)
	return clauses, nil
}

func (reducer *Reducer) ImplicitToImportClauses(
	clauses *ast.ImportClauseList,
	newlines lr.TokenCount,
) (
	*ast.ImportClauseList,
	error,
) {
	return clauses, nil
}

func (reducer *Reducer) ExplicitToImportClauses(
	clauses *ast.ImportClauseList,
	comma *lr.TokenValue,
) (
	*ast.ImportClauseList,
	error,
) {
	clauses.ReduceImproper(comma)
	return clauses, nil
}

//
// JumpStatement
//

func (reducer *Reducer) UnlabeledNoValueToJumpStatement(
	op *lr.TokenValue,
) (
	ast.Statement,
	error,
) {
	return ast.NewJumpStatement(op, nil, nil), nil
}

func (reducer *Reducer) UnlabeledValuedToJumpStatement(
	op *lr.TokenValue,
	value ast.Expression,
) (
	ast.Statement,
	error,
) {
	return ast.NewJumpStatement(op, nil, value), nil
}

func (reducer *Reducer) LabeledNoValueToJumpStatement(
	op *lr.TokenValue,
	label *lr.TokenValue,
) (
	ast.Statement,
	error,
) {
	return ast.NewJumpStatement(op, label, nil), nil
}

func (reducer *Reducer) LabeledValuedToJumpStatement(
	op *lr.TokenValue,
	label *lr.TokenValue,
	value ast.Expression,
) (
	ast.Statement,
	error,
) {
	return ast.NewJumpStatement(op, label, value), nil
}

func (reducer *Reducer) FallthroughToJumpStatement(
	op *lr.TokenValue,
) (
	ast.Statement,
	error,
) {
	return ast.NewJumpStatement(op, nil, nil), nil
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
	*ast.UnsafeStatement,
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
	stmt := &ast.UnsafeStatement{
		StartEndPos:    ast.NewStartEndPos(unsafe.Loc(), verbatimSource.End()),
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
	casePatterns *ast.ExpressionList,
	colon *lr.TokenValue,
	trailingStatement ast.Statement,
) (
	ast.Statement,
	error,
) {
	body := &ast.StatementsExpr{
		StartEndPos: colon.StartEnd(),
	}

	if trailingStatement != nil {
		body.EndPos = trailingStatement.End()
		body.Statements = append(body.Statements, trailingStatement)
	}

	leading := caseKW.TakeLeading()

	casePatterns.PrependToLeading(caseKW.TakeTrailing())
	casePatterns.AppendToTrailing(colon.TakeLeading())
	casePatterns.AppendToTrailing(colon.TakeTrailing())

	stmt := &ast.BranchStatement{
		StartEndPos:  ast.NewStartEndPos(caseKW.Loc(), body.End()),
		CasePatterns: casePatterns,
		Body:         body,
	}
	stmt.LeadingComment = leading

	return stmt, nil
}

func (reducer *Reducer) DefaultBranchToBranchStatement(
	defaultKW *lr.TokenValue,
	colon *lr.TokenValue,
	trailingStatement ast.Statement,
) (
	ast.Statement,
	error,
) {
	body := &ast.StatementsExpr{
		StartEndPos: colon.StartEnd(),
	}

	if trailingStatement != nil {
		body.EndPos = trailingStatement.End()
		body.Statements = append(body.Statements, trailingStatement)
	}

	leading := defaultKW.TakeLeading()

	body.LeadingComment = defaultKW.TakeTrailing()
	body.LeadingComment.Append(colon.TakeLeading())
	body.LeadingComment.Append(colon.TakeTrailing())

	stmt := &ast.BranchStatement{
		StartEndPos:  ast.NewStartEndPos(defaultKW.Loc(), body.End()),
		IsDefault:    true,
		CasePatterns: nil,
		Body:         body,
	}
	stmt.LeadingComment = leading

	return stmt, nil
}
