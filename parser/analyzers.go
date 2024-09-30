package parser

import (
	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/util"
)

type unreachableStatementsDetector struct {
	util.ErrorEmitter
}

func detectUnreachableStatements() util.Pass {
	return &unreachableStatementsDetector{}
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
		_, ok := stmt.(*ast.JumpStmt)
		if !ok {
			continue
		}

		if idx < len(stmts.Statements)-1 {
			detector.Emit("%s: unreachable statement", stmts.Statements[idx+1].Loc())
			break
		}
	}
}

func (detector *unreachableStatementsDetector) Exit(node ast.Node) {
}

type unexpectedStatementsDetector struct {
	processed map[*ast.StatementsExpr]struct{}

	util.ErrorEmitter
}

func detectUnexpectedStatements() util.Pass {
	return &unexpectedStatementsDetector{
		processed: map[*ast.StatementsExpr]struct{}{},
	}
}

func (detector *unexpectedStatementsDetector) Process(node ast.Node) {
	node.Walk(detector)
}

func (detector *unexpectedStatementsDetector) Enter(n ast.Node) {
	switch node := n.(type) {
	case *ast.PackageDef:
		detector.checkPkgDef(node.Body)
	case *ast.SwitchExpr:
		detector.checkSwitchSelectExpr(node.ConditionBranches)
	case *ast.SelectExpr:
		detector.checkSwitchSelectExpr(node.ConditionBranches)
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
		case *ast.ParseErrorExpr: // ok
		case *ast.UnsafeStmt: // ok
		case *ast.ImportStmt: // ok
		case *ast.ConditionBranchStmt:
			invalidStmtType = "branch statement"
		case *ast.JumpStmt:
			invalidStmtType = "jump statement"
		case ast.Expression:
			invalidStmtType = "expression statement"
		default:
			invalidStmtType = "statement type"
		}

		if invalidStmtType != "" {
			detector.Emit(
				"%s: unexpected %s. expected import statement",
				stmt.Loc(),
				invalidStmtType)
		}
	}
}

func (detector *unexpectedStatementsDetector) checkSwitchSelectExpr(
	conditionBranches []*ast.ConditionBranchStmt,
) {
	for _, stmt := range conditionBranches {
		detector.checkStmtsExpr(stmt.Branch, true)
		detector.processed[stmt.Branch] = struct{}{}
	}
}

func (detector *unexpectedStatementsDetector) checkStmtsExpr(
	stmts *ast.StatementsExpr,
	allowFallthrough bool,
) {
	for _, node := range stmts.Statements {
		invalidStmtType := ""
		switch stmt := node.(type) {
		case *ast.ParseErrorExpr: // ok
		case *ast.UnsafeStmt: // ok
		case ast.Expression: // ok
		case *ast.JumpStmt:
			if stmt.Op == ast.FallthroughOp && !allowFallthrough {
				invalidStmtType = "fallthrough statement"
			}
		case *ast.ConditionBranchStmt:
			invalidStmtType = "branch statement"
		case *ast.ImportStmt:
			invalidStmtType = "import statement"
		default:
			invalidStmtType = "statement type"
		}

		if invalidStmtType != "" {
			detector.Emit(
				"%s: unexpected %s. expected expression, assign or jump statement",
				node.Loc(),
				invalidStmtType)
		}
	}
}

func (detector *unexpectedStatementsDetector) Exit(node ast.Node) {
}

type unexpectedArgumentsDetector struct {
	util.ErrorEmitter
}

func detectUnexpectedArguments() util.Pass {
	return &unexpectedArgumentsDetector{}
}

func (detector *unexpectedArgumentsDetector) Process(node ast.Node) {
	node.Walk(detector)
}

func (detector *unexpectedArgumentsDetector) Enter(n ast.Node) {
	switch node := n.(type) {
	case *ast.IndexExpr:
		if node.Index.Kind != ast.PositionalArgument &&
			node.Index.Kind != ast.ColonExprArgument {

			detector.Emit(
				"%s: unexpected %s argument",
				node.Index.Loc(),
				node.Index.Kind)
		}
	case *ast.CallExpr:
		for _, arg := range node.Arguments {
			if arg.Kind == ast.SkipPatternArgument {
				detector.Emit("%s: unexpected %s argument", arg.Loc(), arg.Kind)
			}
		}
	case *ast.InitializeExpr:
		for _, arg := range node.Arguments {
			if arg.Kind == ast.SkipPatternArgument ||
				arg.Kind == ast.VarargAssignmentArgument {

				detector.Emit("%s: unexpected %s argument", arg.Loc(), arg.Kind)
			}
		}
	case *ast.ImplicitStructExpr:
		// TODO detect invalid skip-pattern (...) usage in non-pattern expressions
		for _, arg := range node.Arguments {
			if arg.Kind == ast.VarargAssignmentArgument {
				detector.Emit("%s: unexpected %s argument", arg.Loc(), arg.Kind)
			}
		}
	}
}

func (detector *unexpectedArgumentsDetector) Exit(node ast.Node) {
}

// This verifies:
//   - default branch (if existing) is the last branch
//   - switch's case don't include assign expr
//   - select's case don't include multiple switch cases, expr is either
//     send/recv expr (and assign pattern don't include case enum pattern; this
//     is check by a different pass)
type unexpectedSelectSwitchBranchesDetector struct {
	util.ErrorEmitter
}

func detectUnexpectedSelectSwitchBranches() util.Pass {
	return &unexpectedSelectSwitchBranchesDetector{}
}

func (detector *unexpectedSelectSwitchBranchesDetector) Process(node ast.Node) {
	node.Walk(detector)
}

func (detector *unexpectedSelectSwitchBranchesDetector) Enter(node ast.Node) {
	switch expr := node.(type) {
	case *ast.SwitchExpr:
		detector.checkBranches(expr.ConditionBranches, true)
	case *ast.SelectExpr:
		detector.checkBranches(expr.ConditionBranches, false)
	}
}

func (detector *unexpectedSelectSwitchBranchesDetector) checkBranches(
	condBranches []*ast.ConditionBranchStmt,
	isSwitch bool,
) {
	last := len(condBranches) - 1
	for idx, branch := range condBranches {
		if branch.IsDefaultBranch {
			if idx != last {
				detector.Emit("%s: default branch is not the last branch", branch.Loc())
			}
			continue
		}

		if isSwitch {
			casePatternExpr, ok := branch.Condition.(*ast.CasePatternExpr)
			if !ok {
				detector.Emit("%s: unexpected expression", branch.Condition.Loc())
			} else {
				for _, pattern := range casePatternExpr.Patterns {
					binary, ok := pattern.(*ast.BinaryExpr)
					if ok && binary.Op == ast.BinaryAssignOp {
						detector.Emit("%s: unexpected assignment pattern", binary.Loc())
					}
				}
			}
		} else {
			_, ok := branch.Condition.(*ast.CasePatternExpr)
			if ok {
				detector.Emit(
					"%s: unexpected pattern list, expecting only one pattern per case",
					branch.Condition.Loc())
			} else {
				expr := branch.Condition

				binary, ok := expr.(*ast.BinaryExpr)
				if ok && binary.Op == ast.BinaryAssignOp {
					expr = binary.Right
				}

				isValid := false
				switch nary := expr.(type) {
				case *ast.UnaryExpr:
					isValid = nary.Op == ast.UnaryRecvOp
				case *ast.BinaryExpr:
					isValid = nary.Op == ast.BinarySendOp
				}

				if !isValid {
					detector.Emit(
						"%s: unexpected expression, expecting send/recv expression",
						expr.Loc())
				}
			}
		}
	}
}

func (detector *unexpectedSelectSwitchBranchesDetector) Exit(node ast.Node) {
}
