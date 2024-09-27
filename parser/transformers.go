package parser

import (
	"fmt"

	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/util"
)

type caseStatementsReorganizer struct {
	errs []error
}

func reorganizeCaseStatements() util.Pass {
	return &caseStatementsReorganizer{}
}

func (transformer *caseStatementsReorganizer) Errors() []error {
	return transformer.errs
}

func (transformer *caseStatementsReorganizer) Process(node ast.Node) {
	node.Walk(transformer)
}

func (transformer *caseStatementsReorganizer) Enter(node ast.Node) {
	switch expr := node.(type) {
	case *ast.SwitchExpr:
		transformer.groupCaseStatements(expr.Branches)
	case *ast.SelectExpr:
		transformer.groupCaseStatements(expr.Branches)
	}
}

func (caseStatementsReorganizer) Exit(node ast.Node) {
}

func (transformer *caseStatementsReorganizer) groupCaseStatements(
	stmts *ast.StatementsExpr,
) {
	newStatements := []ast.Statement{}
	var current *ast.BranchStatement
	for _, stmt := range stmts.Statements {
		branch, ok := stmt.(*ast.BranchStatement)
		if !ok {
			if current != nil {
				current.Body.Statements = append(current.Body.Statements, stmt)
			} else {
				err := fmt.Errorf(
					"statement does not belong to any branch: %s",
					stmt.Loc())
				transformer.errs = append(transformer.errs, err)
				newStatements = append(
					newStatements,
					ast.NewParseErrorNode(stmt.StartEnd(), err))
			}
			continue
		}

		current = branch
		newStatements = append(newStatements, branch)

		// Flatten nested case statements
		for {
			if len(current.Body.Statements) == 0 {
				break
			}

			if len(current.Body.Statements) != 1 {
				panic("should never happen")
			}

			nested, ok := current.Body.Statements[0].(*ast.BranchStatement)
			if !ok {
				break
			}

			current.Body.Statements = nil
			current = nested
			newStatements = append(newStatements, nested)
		}
	}

	stmts.Statements = newStatements
}
