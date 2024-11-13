package reducer

import (
	"strings"

	"github.com/pattyshack/gt/lexutil"

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

//
// ImportClause
//

func (reducer *Reducer) stringLiteralToPackageID(
	pkg *lr.TokenValue,
) ast.PackageID {
	if strings.HasPrefix(pkg.Value, "```") ||
		strings.HasPrefix(pkg.Value, `"""`) {

		reducer.Emit(
			pkg.Loc(),
			"cannot specify package id using multi-line string literal")

		return ast.NewPackageID(pkg.Value[3 : len(pkg.Value)-3])
	}

	return ast.NewPackageID(pkg.Value[1 : len(pkg.Value)-1])
}

func (reducer *Reducer) StringLiteralToImportClause(
	pkg *lr.TokenValue,
) (
	*ast.ImportClause,
	error,
) {
	clause := &ast.ImportClause{
		StartEndPos:             pkg.StartEndPos,
		LeadingTrailingComments: pkg.LeadingTrailingComments,
		PackageID:               reducer.stringLiteralToPackageID(pkg),
	}

	return clause, nil
}

func (reducer *Reducer) aliasImport(
	alias *lr.TokenValue,
	pkg *lr.TokenValue,
) *ast.ImportClause {
	clause := &ast.ImportClause{
		StartEndPos: lexutil.NewStartEndPos(alias.Loc(), pkg.End()),
		Alias:       alias.Value,
		PackageID:   reducer.stringLiteralToPackageID(pkg),
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
// ImportStmt
//

func (reducer *Reducer) SingleToImportStmt(
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

	stmt := &ast.ImportStmt{
		StartEndPos:   lexutil.NewStartEndPos(importKW.Loc(), importClause.End()),
		ImportClauses: []*ast.ImportClause{importClause},
	}
	stmt.LeadingComment = leading
	stmt.TrailingComment = trailing

	return stmt, nil
}

func (reducer *Reducer) MultipleToImportStmt(
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

	stmt := &ast.ImportStmt{
		StartEndPos:   lexutil.NewStartEndPos(importKW.Loc(), clauses.End()),
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
// JumpStmt
//

func (reducer *Reducer) UnlabeledNoValueToJumpStmt(
	op *lr.TokenValue,
) (
	ast.Statement,
	error,
) {
	return ast.NewJumpStmt(op, nil, nil, ast.NewImproperUnit(op.End())), nil
}

func (reducer *Reducer) UnlabeledValuedToJumpStmt(
	op *lr.TokenValue,
	value ast.Expression,
) (
	ast.Statement,
	error,
) {
	return ast.NewJumpStmt(op, nil, nil, value), nil
}

func (reducer *Reducer) LabeledNoValueToJumpStmt(
	op *lr.TokenValue,
	at *lr.TokenValue,
	label *lr.TokenValue,
) (
	ast.Statement,
	error,
) {
	return ast.NewJumpStmt(
		op,
		at, label,
		ast.NewImproperUnit(label.End())), nil
}

func (reducer *Reducer) LabeledValuedToJumpStmt(
	op *lr.TokenValue,
	at *lr.TokenValue,
	label *lr.TokenValue,
	value ast.Expression,
) (
	ast.Statement,
	error,
) {
	return ast.NewJumpStmt(op, at, label, value), nil
}

func (reducer *Reducer) FallthroughToJumpStmt(
	op *lr.TokenValue,
) (
	ast.Statement,
	error,
) {
	return ast.NewJumpStmt(op, nil, nil, ast.NewImproperUnit(op.End())), nil
}

//
// UnsafeStmt
//

func (reducer *Reducer) ToUnsafeStmt(
	unsafe *lr.TokenValue,
	less *lr.TokenValue,
	language *lr.TokenValue,
	greater *lr.TokenValue,
	verbatimSource *lr.TokenValue,
) (
	*ast.UnsafeStmt,
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
	stmt := &ast.UnsafeStmt{
		StartEndPos:    lexutil.NewStartEndPos(unsafe.Loc(), verbatimSource.End()),
		Language:       language.Value,
		VerbatimSource: verbatimSource.Value,
	}
	stmt.LeadingComment = leading
	stmt.TrailingComment = verbatimSource.TakeTrailing()
	return stmt, nil
}

//
// BranchStmt
//

func (reducer *Reducer) CaseBranchToBranchStmt(
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

	casePatterns.Elements[0].PrependToLeading(caseKW.TakeTrailing())
	lastPattern := casePatterns.Elements[len(casePatterns.Elements)-1]
	lastPattern.AppendToTrailing(colon.TakeLeading())
	lastPattern.AppendToTrailing(colon.TakeTrailing())

	cond := &ast.CasePatterns{
		StartEndPos: casePatterns.StartEnd(),
		Patterns:    casePatterns.Elements,
	}

	stmt := &ast.ConditionBranchStmt{
		StartEndPos: lexutil.NewStartEndPos(caseKW.Loc(), body.End()),
		Condition:   cond,
		Branch:      body,
	}
	stmt.LeadingComment = leading

	return stmt, nil
}

func (reducer *Reducer) DefaultBranchToBranchStmt(
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

	stmt := &ast.ConditionBranchStmt{
		StartEndPos:     lexutil.NewStartEndPos(defaultKW.Loc(), body.End()),
		IsDefaultBranch: true,
		Condition:       nil,
		Branch:          body,
	}
	stmt.LeadingComment = leading

	return stmt, nil
}

//
// BlockAddrDeclStmt
//

func (Reducer) ToBlockAddrDeclStmt(
	varType *lr.TokenValue,
	dollar *lr.TokenValue,
	lparen *lr.TokenValue,
	list *ast.ExpressionList,
	rparen *lr.TokenValue,
) (
	ast.Statement,
	error,
) {
	isVar := varType.Val() == "var"

	list.ReduceMarkers(lparen, rparen)

	leading := varType.TakeLeading()
	leading.Append(varType.TakeTrailing())
	leading.Append(dollar.TakeLeading())
	leading.Append(dollar.TakeTrailing())
	leading.Append(list.TakeLeading())
	leading.Append(list.MiddleComment)

	trailing := list.TakeTrailing()
	for _, item := range list.Elements {
		pattern := item

		assign, ok := item.(*ast.AssignPattern)
		if ok {
			pattern = assign.Pattern
		}

		decl := pattern.(*ast.AddrDeclPattern)
		decl.IsVar = isVar
	}

	block := &ast.BlockAddrDeclStmt{
		StartEndPos: lexutil.NewStartEndPos(varType.Loc(), rparen.End()),
		IsVar:       isVar,
		Patterns:    list.Elements,
	}
	block.LeadingComment = leading
	block.TrailingComment = trailing

	return block, nil
}

func (Reducer) ImproperImplicitToBlockAddrDeclList(
	list *ast.ExpressionList,
	newlines lr.TokenCount,
) (
	*ast.ExpressionList,
	error,
) {
	return list, nil
}

func (Reducer) ImproperExplicitToBlockAddrDeclList(
	list *ast.ExpressionList,
	comma *lr.TokenValue,
) (
	*ast.ExpressionList,
	error,
) {
	list.ReduceImproper(comma)
	return list, nil
}

func (Reducer) NilToBlockAddrDeclList() (*ast.ExpressionList, error) {
	return ast.NewExpressionList(), nil
}

func (Reducer) BlockAddrDeclItemToProperBlockAddrDeclList(
	item ast.Expression,
) (
	*ast.ExpressionList,
	error,
) {
	list := ast.NewExpressionList()
	list.Add(item)
	return list, nil
}

func (Reducer) AddImplicitToProperBlockAddrDeclList(
	list *ast.ExpressionList,
	newlines lr.TokenCount,
	item ast.Expression,
) (
	*ast.ExpressionList,
	error,
) {
	list.Add(item)
	return list, nil
}

func (Reducer) AddExplicitToProperBlockAddrDeclList(
	list *ast.ExpressionList,
	comma *lr.TokenValue,
	item ast.Expression,
) (
	*ast.ExpressionList,
	error,
) {
	list.ReduceAdd(comma, item)
	return list, nil
}
