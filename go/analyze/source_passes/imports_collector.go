package source_passes

import (
	"github.com/pattyshack/gt/parseutil"

	"github.com/pattyshack/pl/ast"
)

type ImportClausesCollector struct {
	clauses []*ast.ImportClause
	*parseutil.Emitter
}

func NewImportClausesCollector(
	emitter *parseutil.Emitter,
) *ImportClausesCollector {
	return &ImportClausesCollector{
		Emitter: emitter,
	}
}

func (collector *ImportClausesCollector) Clauses() []*ast.ImportClause {
	return collector.clauses
}

func (collector *ImportClausesCollector) Process(list *ast.StatementList) {
	for _, stmt := range list.Elements {
		importStmt, ok := stmt.(*ast.ImportStmt)
		if !ok {
			continue
		}

		for _, clause := range importStmt.ImportClauses {
			err := clause.PackageID.Validate()
			if err != nil {
				collector.EmitErrors(parseutil.LocationError{Loc: clause.Loc(), Err: err})
				continue
			}

			collector.clauses = append(collector.clauses, clause)
		}
	}
}
