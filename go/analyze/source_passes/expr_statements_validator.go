package source_passes

import (
	"github.com/pattyshack/gt/parseutil"

	"github.com/pattyshack/pl/analyze/process"
	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/util"
)

type ExprStatementsValidator struct {
	*parseutil.Emitter
}

func ValidateExprStatements(emitter *parseutil.Emitter) process.Pass {
	return &ExprStatementsValidator{
		Emitter: emitter,
	}
}

func (validator *ExprStatementsValidator) Process(
	list *ast.StatementList,
) {
	process.ParallelWalk(
		list,
		func(ast.Statement) ast.Visitor {
			return &exprStatementsValidator{
				processed: map[*ast.StatementsExpr]struct{}{},
				Emitter:   validator.Emitter,
			}
		})
}

type exprStatementsValidator struct {
	processed map[*ast.StatementsExpr]struct{}

	*parseutil.Emitter
}

func (validator *exprStatementsValidator) Enter(n ast.Node) {
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
	}
}

func (validator *exprStatementsValidator) checkSwitchSelectExpr(
	conditionBranches []*ast.ConditionBranchStmt,
) {
	for _, stmt := range conditionBranches {
		validator.checkExprStmts(stmt.Branch.Statements, true)
		validator.processed[stmt.Branch] = struct{}{}
	}
}

func (validator *exprStatementsValidator) checkExprStmts(
	stmts []ast.Statement,
	allowFallthrough bool,
) {
	for _, node := range stmts {
		invalidStmtType := ""
		switch stmt := node.(type) {
		case *ast.UnsafeStmt: // ok
		case *ast.DirectivesDecl: // ok
		case *ast.BlockAddrDeclStmt: // ok
		case *ast.ImportStmt:
			invalidStmtType = "import statement"
		case *ast.FloatingComment:
			// This only happen in manually constructed tree.
			invalidStmtType = "floating comment"
		case *ast.TypeDef:
			// XXX: ok for now. should check to ensure the type is used.
		case *ast.AliasDef:
			invalidStmtType = "alias definition"
		case *ast.ConditionBranchStmt:
			invalidStmtType = "branch statement"
		case *ast.JumpStmt:
			if stmt.Op == ast.FallthroughOp && !allowFallthrough {
				invalidStmtType = "fallthrough statement"
			}
		case *ast.FuncDefinition:
			sig := stmt.Signature
			if len(sig.Parameters.Elements) > 0 &&
				sig.Parameters.Elements[0].Kind == ast.ReceiverParameter {

				invalidStmtType = "method definition"
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

func (validator *exprStatementsValidator) Exit(node ast.Node) {
}
