package syntax

import (
	"github.com/pattyshack/pl/ast"
)

type ImportClausesCollector struct {
	clauses []*ast.ImportClause
}

func NewImportClausesCollector() *ImportClausesCollector {
	return &ImportClausesCollector{}
}

func (collector *ImportClausesCollector) Clauses() []*ast.ImportClause {
	return collector.clauses
}

func (collector *ImportClausesCollector) Process(node ast.Node) {
	list, ok := node.(*ast.StatementList)
	if !ok {
		return
	}

	for _, stmt := range list.Elements {
		importStmt, ok := stmt.(*ast.ImportStmt)
		if !ok {
			continue
		}

		collector.clauses = append(collector.clauses, importStmt.ImportClauses...)
	}
}
