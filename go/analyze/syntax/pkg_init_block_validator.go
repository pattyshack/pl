package syntax

import (
	"github.com/pattyshack/gt/lexutil"

	"github.com/pattyshack/pl/analyze/process"
	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/errors"
)

type PkgInitBlockValidator struct {
	IsLibrary bool
	*errors.Emitter
}

func ValidatePkgInitBlock(
	isLibrary bool,
	emitter *errors.Emitter,
) process.Pass {
	return &PkgInitBlockValidator{
		IsLibrary: isLibrary,
		Emitter:   emitter,
	}
}

func (validator *PkgInitBlockValidator) Process(node ast.Node) {
	if node == nil {
		return
	}

	stmts, ok := node.(*ast.StatementList)
	if !ok {
		return
	}

	count := 0
	var first lexutil.Location
	for _, stmt := range stmts.Elements {
		_, ok := stmt.(*ast.StatementsExpr)
		if !ok {
			continue
		}

		if validator.IsLibrary {
			count++
			if count == 1 {
				first = stmt.Loc()
			} else {
				validator.Emit(
					stmt.Loc(),
					"duplicate package init statement block (first defined at %s), "+
						"each package can have at most package one init statement block",
					first)
			}
		} else {
			validator.Emit(
				stmt.Loc(),
				"cannot specify init statement block in main / test sources")
		}
	}
}
