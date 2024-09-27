package parser

import (
	"fmt"

	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/util"
)

type unreachableStatementsDetector struct {
	errs []error
}

func detectUnreachableStatements() util.Pass {
	return &unreachableStatementsDetector{}
}

func (detector *unreachableStatementsDetector) Errors() []error {
	return detector.errs
}

func (detector *unreachableStatementsDetector) Process(node ast.Node) {
	node.Walk(detector)
}

func (detector *unreachableStatementsDetector) Enter(node ast.Node) {
	stmts, ok := node.(*ast.StatementsExpr)
	if !ok {
		return
	}

	for idx, stmt := range stmts.Statements {
		_, ok := stmt.(*ast.JumpStatement)
		if !ok {
			continue
		}

		if idx < len(stmts.Statements)-1 {
			detector.errs = append(
				detector.errs,
				fmt.Errorf(
					"unreachable statement: %s",
					stmts.Statements[idx+1].Loc()))
			break
		}
	}
}

func (detector *unreachableStatementsDetector) Exit(node ast.Node) {
}

type unexpectedStatementsDetector struct {
	processed map[*ast.StatementsExpr]struct{}

	errs []error
}

func detectUnexpectedStatements() util.Pass {
	return &unexpectedStatementsDetector{
		processed: map[*ast.StatementsExpr]struct{}{},
	}
}

func (detector *unexpectedStatementsDetector) Errors() []error {
	return detector.errs
}

func (detector *unexpectedStatementsDetector) Process(node ast.Node) {
	node.Walk(detector)
}

func (detector *unexpectedStatementsDetector) Enter(n ast.Node) {
	switch node := n.(type) {
	case *ast.PackageDef:
		detector.checkPkgDef(node.Body)
	case *ast.SwitchExpr:
		detector.checkSwitchSelectExpr(node.Branches)
	case *ast.SelectExpr:
		detector.checkSwitchSelectExpr(node.Branches)
	case *ast.StatementsExpr:
		_, ok := detector.processed[node]
		if ok {
			return
		}
		detector.checkStmtsExpr(node, false)
	}
}

func (detector *unexpectedStatementsDetector) checkPkgDef(
	stmts *ast.StatementsExpr,
) {
	detector.processed[stmts] = struct{}{}

	for _, stmt := range stmts.Statements {
		invalidStmtType := ""
		switch stmt.(type) {
		case *ast.ParseErrorNode: // ok
		case *ast.UnsafeStatement: // ok
		case *ast.ImportStatement: // ok
		case *ast.BranchStatement:
			invalidStmtType = "branch statement"
		case *ast.JumpStatement:
			invalidStmtType = "jump statement"
		case ast.Expression:
			invalidStmtType = "expression statement"
		default:
			invalidStmtType = "statement type"
		}

		if invalidStmtType != "" {
			detector.errs = append(
				detector.errs,
				fmt.Errorf("unexpected %s: %s", invalidStmtType, stmt.Loc()))
		}
	}
}

func (detector *unexpectedStatementsDetector) checkSwitchSelectExpr(
	stmts *ast.StatementsExpr,
) {
	detector.processed[stmts] = struct{}{}

	for _, node := range stmts.Statements {
		invalidStmtType := ""
		switch stmt := node.(type) {
		case *ast.ParseErrorNode: // ok
		case *ast.UnsafeStatement: // ok
		case *ast.BranchStatement:
			detector.checkStmtsExpr(stmt.Body, true)
			detector.processed[stmt.Body] = struct{}{}
		case *ast.ImportStatement:
			invalidStmtType = "import statement"
		case *ast.JumpStatement:
			invalidStmtType = "jump statement"
		case ast.Expression:
			invalidStmtType = "expression statement"
		default:
			invalidStmtType = "statement type"
		}

		if invalidStmtType != "" {
			detector.errs = append(
				detector.errs,
				fmt.Errorf("unexpected %s: %s", invalidStmtType, node.Loc()))
		}
	}
}

func (detector *unexpectedStatementsDetector) checkStmtsExpr(
	stmts *ast.StatementsExpr,
	allowFallthrough bool,
) {
	for _, node := range stmts.Statements {
		invalidStmtType := ""
		switch stmt := node.(type) {
		case *ast.ParseErrorNode: // ok
		case *ast.UnsafeStatement: // ok
		case ast.Expression: // ok
		case *ast.JumpStatement:
			if stmt.Op == ast.FallthroughOp && !allowFallthrough {
				invalidStmtType = "fallthrough statement"
			}
		case *ast.BranchStatement:
			invalidStmtType = "branch statement"
		case *ast.ImportStatement:
			invalidStmtType = "import statement"
		default:
			invalidStmtType = "statement type"
		}

		if invalidStmtType != "" {
			detector.errs = append(
				detector.errs,
				fmt.Errorf("unexpected %s: %s", invalidStmtType, node.Loc()))
		}
	}
}

func (detector *unexpectedStatementsDetector) Exit(node ast.Node) {
}
