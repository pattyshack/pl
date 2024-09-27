package parser

import (
	"fmt"

	"github.com/pattyshack/pl/ast"
)

type reorganizeCaseStatements struct {
	errs []error
}

func (transformer *reorganizeCaseStatements) Errors() []error {
	return transformer.errs
}

func (transformer *reorganizeCaseStatements) Process(node ast.Node) {
	node.Walk(transformer)
}

func (transformer *reorganizeCaseStatements) Enter(node ast.Node) {
	switch expr := node.(type) {
	case *ast.SwitchExpr:
		transformer.groupCaseStatements(expr.Branches)
	case *ast.SelectExpr:
		transformer.groupCaseStatements(expr.Branches)
	}
}

func (reorganizeCaseStatements) Exit(node ast.Node) {
}

func (transformer *reorganizeCaseStatements) groupCaseStatements(
	stmts *ast.StatementsExpr,
) {
	newStatements := ast.NewStatementList()
	var current *ast.BranchStatement
	for _, stmt := range stmts.Statements.Elements {
		branch, ok := stmt.(*ast.BranchStatement)
		if !ok {
			if current != nil {
				current.Body.Statements.Add(stmt)
			} else {
				err := fmt.Errorf(
					"statement does not belong to any branch: %s",
					stmt.Loc())
				transformer.errs = append(transformer.errs, err)
				newStatements.Add(
					ast.NewParseErrorNode(stmt.StartEnd(), err))
			}
			continue
		}

		current = branch
		newStatements.Add(branch)

		// Flatten nested case statements
		for {
			if len(current.Body.Statements.Elements) == 0 {
				break
			}

			if len(current.Body.Statements.Elements) != 1 {
				panic("should never happen")
			}

			nested, ok := current.Body.Statements.Elements[0].(*ast.BranchStatement)
			if !ok {
				break
			}

			current.Body.Statements.Elements = nil
			current = nested
			newStatements.Add(nested)
		}
	}

	stmts.Statements = newStatements
}
