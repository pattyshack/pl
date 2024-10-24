package syntax

import (
	"github.com/pattyshack/pl/analyze/process"
	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/errors"
	"github.com/pattyshack/pl/util"
)

type PackageStatementsValidator struct {
	*errors.Emitter
}

func ValidatePackageStatements(emitter *errors.Emitter) process.Pass {
	return &PackageStatementsValidator{
		Emitter: emitter,
	}
}

func (validator *PackageStatementsValidator) Process(node ast.Node) {
	if node == nil {
		return
	}

	stmts, ok := node.(*ast.StatementList)
	if !ok {
		return
	}

	for _, stmt := range stmts.Elements {
		validator.checkStmt(stmt)
	}
}

func (validator *PackageStatementsValidator) checkStmt(stmt ast.Statement) {
	invalidType := ""
	switch expr := stmt.(type) {
	case *ast.UnsafeStmt: // ok
	case *ast.DirectivesDecl: // ok
	case *ast.ImportStmt: // ok
	case *ast.FloatingComment: // ok
	case *ast.TypeDef: // ok
	case *ast.AliasDef: // ok
	case *ast.BlockAddrDeclStmt: // ok
	case *ast.ConditionBranchStmt:
		invalidType = "branch statement"
	case *ast.JumpStmt:
		invalidType = "jump statement"
	case ast.Expression:
		validator.checkExpr(expr)
	default:
		invalidType = "statement type:\n" + util.TreeString(stmt, "  ") + "\n"
	}

	if invalidType != "" {
		validator.Emit(stmt.Loc(), "unexpected package level %s", invalidType)
	}
}

func (validator *PackageStatementsValidator) checkExpr(node ast.Expression) {
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
	case *ast.ParameterizedExpr:
		invalid = "parameterized"
	case *ast.CallExpr:
		invalid = "call"
	case *ast.IndexExpr:
		invalid = "index"
	case *ast.AsExpr:
		invalid = "as"
	case *ast.MakeExpr:
		invalid = "make"
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
		if expr.Label != "" {
			validator.Emit(
				node.Loc(),
				"package level statement block must be unlabelled")
		}
	case *ast.ParseErrorExpr:
	case *ast.AddrDeclPattern:

	default:
		invalid = "expression type:\n" + util.TreeString(node, "  ") + "\n"
	}

	if invalid != "" {
		validator.Emit(
			node.Loc(),
			"unexpected package level %s expression%s",
			invalid,
			additional)
	}
}
