package syntax

import (
	"github.com/pattyshack/pl/ast"

	"github.com/pattyshack/gt/lexutil"
)

type ImportClausesCollector struct {
	clauses []*ast.ImportClause
	*lexutil.ErrorEmitter
}

func NewImportClausesCollector(
	emitter *lexutil.ErrorEmitter,
) *ImportClausesCollector {
	return &ImportClausesCollector{
		ErrorEmitter: emitter,
	}
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

		for _, clause := range importStmt.ImportClauses {
			err := clause.PackageID.Validate()
			if err != nil {
				collector.EmitErrors(lexutil.LocationError{Loc: clause.Loc(), Err: err})
				continue
			}

			collector.clauses = append(collector.clauses, clause)
		}
	}
}
