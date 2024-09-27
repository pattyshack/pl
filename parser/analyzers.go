package parser

import (
	"fmt"

	"github.com/pattyshack/pl/ast"
)

type detectUnreachableStatements struct {
	errs []error
}

func (detector *detectUnreachableStatements) Errors() []error {
	return detector.errs
}

func (detector *detectUnreachableStatements) Process(node ast.Node) {
	node.Walk(detector)
}

func (detector *detectUnreachableStatements) Enter(node ast.Node) {
	stmts, ok := node.(*ast.StatementsExpr)
	if !ok {
		return
	}

	for idx, stmt := range stmts.Statements.Elements {
		_, ok := stmt.(*ast.JumpStatement)
		if !ok {
			continue
		}

		if idx < len(stmts.Statements.Elements)-1 {
			detector.errs = append(
				detector.errs,
				fmt.Errorf(
					"unreachable statement: %s",
					stmts.Statements.Elements[idx+1].Loc()))
			break
		}
	}
}

func (detector *detectUnreachableStatements) Exit(node ast.Node) {
}
