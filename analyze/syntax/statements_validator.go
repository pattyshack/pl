package syntax

import (
	"github.com/pattyshack/gt/lexutil"

	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/util"
)

type StatementsValidator struct {
	processed map[*ast.StatementsExpr]struct{}

	lexutil.ErrorEmitter
}

func ValidateStatements() ast.Pass {
	return &StatementsValidator{
		processed: map[*ast.StatementsExpr]struct{}{},
	}
}

func (validator *StatementsValidator) Process(node ast.Node) {
	node.Walk(validator)
}

func (validator *StatementsValidator) Enter(n ast.Node) {
	switch node := n.(type) {
	case *ast.StatementList:
		validator.checkTopLevelStmts(node)
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

func (validator *StatementsValidator) checkTopLevelStmts(
	stmts *ast.StatementList,
) {
	for _, stmt := range stmts.Elements {
		invalidStmtType := ""
		switch expr := stmt.(type) {
		case *ast.UnsafeStmt: // ok
		case *ast.ImportStmt: // ok
		case *ast.FloatingComment: // ok
		case *ast.TypeDef: // ok
		case *ast.AliasDef: // ok
		case *ast.BlockAddrDeclStmt: // ok
		case *ast.ConditionBranchStmt:
			invalidStmtType = "branch statement"
		case *ast.JumpStmt:
			invalidStmtType = "jump statement"
		case ast.Expression:
			validator.checkTopLevelExpr(expr)
		default:
			invalidStmtType = "statement type:\n" + util.TreeString(stmt, "  ") + "\n"
		}

		if invalidStmtType != "" {
			validator.Emit(stmt.Loc(), "unexpected %s", invalidStmtType)
		}
	}
}

func (validator *StatementsValidator) checkTopLevelExpr(
	node ast.Expression,
) {
	invalid := ""
	additional := ""
	switch expr := node.(type) {

	// invalid top level expressions

	case *ast.LiteralExpr:
		invalid = "literal"
	case *ast.NamedExpr:
		invalid = "name"
	case *ast.AccessExpr:
		invalid = "access"
	case *ast.UnaryExpr:
		invalid = "unary"
	case *ast.BinaryExpr:
		invalid = "binary"
	case *ast.ImplicitStructExpr:
		invalid = "implicit struct"
	case *ast.CallExpr:
		invalid = "call"
	case *ast.IndexExpr:
		invalid = "index"
	case *ast.AsExpr:
		invalid = "as"
	case *ast.InitializeExpr:
		invalid = "initialize"
	case *ast.IfExpr:
		invalid = "if"
	case *ast.SwitchExpr:
		invalid = "switch"
	case *ast.SelectExpr:
		invalid = "select"
	case *ast.LoopExpr:
		invalid = "loop"

	// invalid top-level patterns (errors are emitted by patternAnalyzer).

	case *ast.CasePatterns:
	case *ast.AssignToAddrPattern:
	case *ast.EnumPattern:

		// conditionally valid top level expressions

	case *ast.FuncDefinition:
		if expr.Signature.Name == "" {
			invalid = "anonymous func"
		}

	case *ast.AssignPattern:
		_, ok := expr.Pattern.(*ast.AddrDeclPattern)
		if !ok {
			invalid = "assign"
			additional = ", left hand side must be proper variable declaration"
		}

		// valid top level expressions

	case *ast.StatementsExpr:
	case *ast.ParseErrorExpr:
	case *ast.AddrDeclPattern:

	default:
		invalid = "statement type:\n" + util.TreeString(node, "  ") + "\n"
	}

	if invalid != "" {
		validator.Emit(
			node.Loc(),
			"unexpected source level %s expression%s",
			invalid,
			additional)
	}
}

func (validator *StatementsValidator) checkSwitchSelectExpr(
	conditionBranches []*ast.ConditionBranchStmt,
) {
	for _, stmt := range conditionBranches {
		validator.checkExprStmts(stmt.Branch.Statements, true)
		validator.processed[stmt.Branch] = struct{}{}
	}
}

func (validator *StatementsValidator) checkExprStmts(
	stmts []ast.Statement,
	allowFallthrough bool,
) {
	for _, node := range stmts {
		invalidStmtType := ""
		switch stmt := node.(type) {
		case *ast.UnsafeStmt: // ok
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

func (validator *StatementsValidator) Exit(node ast.Node) {
}
