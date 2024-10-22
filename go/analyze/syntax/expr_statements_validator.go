package syntax

import (
	"github.com/pattyshack/pl/analyze/process"
	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/errors"
	"github.com/pattyshack/pl/util"
)

type ExprStatementsValidator struct {
	processed map[*ast.StatementsExpr]struct{}

	*errors.Emitter
}

func ValidateExprStatements(emitter *errors.Emitter) process.Pass {
	return &ExprStatementsValidator{
		processed: map[*ast.StatementsExpr]struct{}{},
		Emitter:   emitter,
	}
}

func (validator *ExprStatementsValidator) Process(node ast.Node) {
	node.Walk(validator)
}

func (validator *ExprStatementsValidator) Enter(n ast.Node) {
	switch node := n.(type) {
	case *ast.SwitchExpr:
		validator.checkSwitchSelectExpr(node.ConditionBranches)
	case *ast.SelectExpr:
		validator.checkSwitchSelectExpr(node.ConditionBranches)
	case *ast.StatementsExpr:
		_, ok := validator.processed[node]
		if ok {
			return
		}
		validator.checkExprStmts(node.Statements, false)
	case *ast.LoopExpr:
		stmts := []ast.Statement{}
		if node.Init != nil {
			stmts = append(stmts, node.Init)
		}
		if node.Post != nil {
			stmts = append(stmts, node.Post)
		}
		validator.checkExprStmts(stmts, false)
	}
}

func (validator *ExprStatementsValidator) checkSwitchSelectExpr(
	conditionBranches []*ast.ConditionBranchStmt,
) {
	for _, stmt := range conditionBranches {
		validator.checkExprStmts(stmt.Branch.Statements, true)
		validator.processed[stmt.Branch] = struct{}{}
	}
}

func (validator *ExprStatementsValidator) checkExprStmts(
	stmts []ast.Statement,
	allowFallthrough bool,
) {
	for _, node := range stmts {
		invalidStmtType := ""
		switch stmt := node.(type) {
		case *ast.UnsafeStmt: // ok
		case *ast.DirectivesDecl:
			invalidStmtType = "directives declaration"
		case *ast.BlockAddrDeclStmt: // ok
		case *ast.ImportStmt:
			invalidStmtType = "import statement"
		case *ast.FloatingComment:
			// This only happen in manually constructed tree.
			invalidStmtType = "floating comment"
		case *ast.TypeDef:
			// XXX: ok for now. should check to ensure the type is used.
		case *ast.AliasDef:
			// XXX: ok for now. should check to ensure the type is used.
		case *ast.ConditionBranchStmt:
			invalidStmtType = "branch statement"
		case *ast.JumpStmt:
			if stmt.Op == ast.FallthroughOp && !allowFallthrough {
				invalidStmtType = "fallthrough statement"
			}
		case ast.Expression: // ok
		default:
			invalidStmtType = "statement type:\n" + util.TreeString(node, "  ") + "\n"
		}

		if invalidStmtType != "" {
			validator.Emit(
				node.Loc(),
				"unexpected %s. expected expression, assign or jump statement",
				invalidStmtType)
		}
	}
}

func (validator *ExprStatementsValidator) Exit(node ast.Node) {
}
