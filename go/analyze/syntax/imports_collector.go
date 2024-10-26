package syntax

import (
	"github.com/pattyshack/gt/lexutil"

	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/errors"
)

type ImportClausesCollector struct {
	clauses []*ast.ImportClause
	*errors.Emitter
}

func NewImportClausesCollector(
	emitter *errors.Emitter,
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
				collector.EmitErrors(lexutil.LocationError{Loc: clause.Loc(), Err: err})
				continue
			}

			collector.clauses = append(collector.clauses, clause)
		}
	}
}
