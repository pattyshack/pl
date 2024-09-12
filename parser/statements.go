package parser

import (
	"fmt"
)

//
// StatementsExpr
//

type StatementsExpr struct {
	isExpression

	LabelDecl string // optional
	NodeList[Statement]
}

func NewStatementList() *StatementsExpr {
	return &StatementsExpr{}
}

func (expr StatementsExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[StatementsExpr: LabelDecl=%s",
		indent,
		label,
		expr.LabelDecl)
	if len(expr.Elements) == 0 {
		return result + "]"
	}

	result += expr.elementsString(indent + "  ")
	result += "\n" + indent + "]"
	return result
}

type StatementsExprReducerImpl struct{}

var _ ProperStatementListReducer = &StatementsExprReducerImpl{}
var _ StatementListReducer = &StatementsExprReducerImpl{}
var _ StatementsReducer = &StatementsExprReducerImpl{}
var _ StatementsExprReducer = &StatementsExprReducerImpl{}

func (StatementsExprReducerImpl) AddImplicitToProperStatementList(
	statements *StatementsExpr,
	newlines TokenCount,
	statement Statement,
) (
	*StatementsExpr,
	error,
) {
	statements.reduceAdd(TokenValue{}, statement)
	return statements, nil
}

func (StatementsExprReducerImpl) AddExplicitToProperStatementList(
	statements *StatementsExpr,
	semicolon TokenValue,
	statement Statement,
) (
	*StatementsExpr,
	error,
) {
	statements.reduceAdd(semicolon, statement)
	return statements, nil
}

func (StatementsExprReducerImpl) StatementToProperStatementList(
	statement Statement,
) (
	*StatementsExpr,
	error,
) {
	list := NewStatementList()
	list.add(statement)
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
	semicolon TokenValue,
) (
	*StatementsExpr,
	error,
) {
	statements.reduceImproper(semicolon)
	return statements, nil
}

func (StatementsExprReducerImpl) NilToStatementList() (*StatementsExpr, error) {
	return NewStatementList(), nil
}

func (StatementsExprReducerImpl) ToStatements(
	lbrace TokenValue,
	list *StatementsExpr,
	rbrace TokenValue,
) (
	Expression,
	error,
) {
	list.reduceMarkers(lbrace, rbrace)
	return list, nil
}

func (StatementsExprReducerImpl) LabelledToStatementsExpr(
	labelDecl TokenValue,
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

//
// ImportClause
//

type ImportClause struct {
	StartEndPos
	LeadingTrailingComments

	Alias   string // Identifier or underscore or dot or ""
	Package string
}

var _ Node = &ImportClause{}

func (clause ImportClause) TreeString(indent string, label string) string {
	return fmt.Sprintf(
		"%s%s[ImportClause: Alias=%s Package=%s]",
		indent,
		label,
		clause.Alias,
		clause.Package)
}

type ImportClauseReducerImpl struct {
	ImportClauses []*ImportClause
}

var _ ImportClauseReducer = &ImportClauseReducerImpl{}

func (reducer *ImportClauseReducerImpl) StringLiteralToImportClause(
	pkg TokenValue,
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
	alias TokenValue,
	pkg TokenValue,
) *ImportClause {
	clause := &ImportClause{
		StartEndPos: newStartEndPos(alias.Loc(), pkg.End()),
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
	alias TokenValue,
	pkg TokenValue,
) (
	*ImportClause,
	error,
) {
	return reducer.aliasImport(alias, pkg), nil
}

func (reducer *ImportClauseReducerImpl) UnusableImportToImportClause(
	underscore TokenValue,
	pkg TokenValue,
) (
	*ImportClause,
	error,
) {
	return reducer.aliasImport(underscore, pkg), nil
}

func (reducer *ImportClauseReducerImpl) ImportToLocalToImportClause(
	dot TokenValue,
	pkg TokenValue,
) (
	*ImportClause,
	error,
) {
	return reducer.aliasImport(dot, pkg), nil
}

//
// ImportStatement
//

type ImportStatement struct {
	isStatement
	NodeList[*ImportClause]
}

var _ Node = &ImportStatement{}

func NewImportStatement() *ImportStatement {
	return &ImportStatement{
		NodeList: *newNodeList[*ImportClause]("ImportStatement"),
	}
}

type ImportStatementReducerImpl struct{}

var _ ImportStatementReducer = &ImportStatementReducerImpl{}
var _ ProperImportClausesReducer = &ImportStatementReducerImpl{}
var _ ImportClausesReducer = &ImportStatementReducerImpl{}

func (ImportStatementReducerImpl) SingleToImportStatement(
	importKW TokenValue,
	importClause *ImportClause,
) (
	Statement,
	error,
) {
	stmt := NewImportStatement()
	stmt.add(importClause)

	stmt.StartPos = importKW.Loc()
	stmt.PrependToLeading(importKW.TakeTrailing())
	stmt.PrependToLeading(importKW.TakeLeading())
	return stmt, nil
}

func (ImportStatementReducerImpl) MultipleToImportStatement(
	importKW TokenValue,
	lparen TokenValue,
	stmt *ImportStatement,
	rparen TokenValue,
) (
	Statement,
	error,
) {
	stmt.reduceMarkers(lparen, rparen)

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
	stmt.reduceAdd(TokenValue{}, importClause)
	return stmt, nil
}

func (ImportStatementReducerImpl) AddExplicitToProperImportClauses(
	stmt *ImportStatement,
	comma TokenValue,
	importClause *ImportClause,
) (
	*ImportStatement,
	error,
) {
	stmt.reduceAdd(comma, importClause)
	return stmt, nil
}

func (ImportStatementReducerImpl) ImportClauseToProperImportClauses(
	importClause *ImportClause) (
	*ImportStatement,
	error,
) {
	stmt := NewImportStatement()
	stmt.add(importClause)
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
	comma TokenValue,
) (
	*ImportStatement,
	error,
) {
	stmt.reduceImproper(comma)
	return stmt, nil
}
