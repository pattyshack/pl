package parser

import (
	"fmt"

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
	case *ast.StatementList:
		detector.checkTopLevelStmts(node)
	case *ast.SwitchExpr:
		detector.checkSwitchSelectExpr(node.ConditionBranches)
	case *ast.SelectExpr:
		detector.checkSwitchSelectExpr(node.ConditionBranches)
	case *ast.StatementsExpr:
		_, ok := detector.processed[node]
		if ok {
			return
		}
		detector.checkExprStmts(node.Statements, false)
	case *ast.LoopExpr:
		stmts := []ast.Statement{}
		if node.Init != nil {
			stmts = append(stmts, node.Init)
		}
		if node.Post != nil {
			stmts = append(stmts, node.Post)
		}
		detector.checkExprStmts(stmts, false)
	}
}

func (detector *unexpectedStatementsDetector) checkTopLevelStmts(
	stmts *ast.StatementList,
) {
	for _, stmt := range stmts.Elements {
		invalidStmtType := ""
		switch stmt.(type) {
		case *ast.UnsafeStmt: // ok
		case *ast.ImportStmt: // ok
		case *ast.FloatingComment: // ok
		case *ast.TypeDef: // ok
		case *ast.ConditionBranchStmt:
			invalidStmtType = "branch statement"
		case *ast.JumpStmt:
			invalidStmtType = "jump statement"
		case ast.Expression:
			// TODO limit allowable top-level expression
		default:
			invalidStmtType = fmt.Sprintf("statement type (%v)", stmt)
		}

		if invalidStmtType != "" {
			detector.Emit(
				"%s: unexpected %s",
				stmt.Loc(),
				invalidStmtType)
		}
	}
}

func (detector *unexpectedStatementsDetector) checkSwitchSelectExpr(
	conditionBranches []*ast.ConditionBranchStmt,
) {
	for _, stmt := range conditionBranches {
		detector.checkExprStmts(stmt.Branch.Statements, true)
		detector.processed[stmt.Branch] = struct{}{}
	}
}

func (detector *unexpectedStatementsDetector) checkExprStmts(
	stmts []ast.Statement,
	allowFallthrough bool,
) {
	for _, node := range stmts {
		invalidStmtType := ""
		switch stmt := node.(type) {
		case *ast.UnsafeStmt: // ok
		case *ast.ImportStmt:
			invalidStmtType = "import statement"
		case *ast.FloatingComment:
			// This only happen in manually constructed tree.
			invalidStmtType = "floating comment"
		case *ast.TypeDef:
			// XXX: ok for now. should check to ensure the type is used.
		case *ast.ConditionBranchStmt:
			invalidStmtType = "branch statement"
		case *ast.JumpStmt:
			if stmt.Op == ast.FallthroughOp && !allowFallthrough {
				invalidStmtType = "fallthrough statement"
			}
		case ast.Expression: // ok
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
	// NOTE: ImplicitStructExpr's arguments are handled by
	// unexpectedPatternsDetector.
	switch node := n.(type) {
	case *ast.CallExpr:
		for idx, arg := range node.Arguments {
			if arg.Kind == ast.SkipPatternArgument {
				detector.Emit("%s: unexpected %s argument", arg.Loc(), arg.Kind)
			}

			if arg.Kind == ast.VariadicArgument &&
				idx != len(node.Arguments)-1 {

				detector.Emit(
					"%s: %s argument must be the last argument in the list",
					arg.Loc(),
					arg.Kind)
			}
		}
	case *ast.InitializeExpr:
		for _, arg := range node.Arguments {
			if arg.Kind == ast.SkipPatternArgument ||
				arg.Kind == ast.VariadicArgument {

				detector.Emit("%s: unexpected %s argument", arg.Loc(), arg.Kind)
			}
		}
	}
}

func (detector *unexpectedArgumentsDetector) Exit(node ast.Node) {
}

// This verifies:
//   - default branch (if existing) is the last branch
type unexpectedDefaultBranchesDetector struct {
	util.ErrorEmitter
}

func detectUnexpectedDefaultBranches() util.Pass {
	return &unexpectedDefaultBranchesDetector{}
}

func (detector *unexpectedDefaultBranchesDetector) Process(node ast.Node) {
	node.Walk(detector)
}

func (detector *unexpectedDefaultBranchesDetector) Enter(node ast.Node) {
	switch expr := node.(type) {
	case *ast.IfExpr:
		detector.checkBranches(expr.ConditionBranches)
	case *ast.SwitchExpr:
		detector.checkBranches(expr.ConditionBranches)
	case *ast.SelectExpr:
		detector.checkBranches(expr.ConditionBranches)
	}
}

func (detector *unexpectedDefaultBranchesDetector) checkBranches(
	condBranches []*ast.ConditionBranchStmt,
) {
	last := len(condBranches) - 1
	for idx, branch := range condBranches {
		if branch.IsDefaultBranch {
			if idx != last {
				detector.Emit("%s: default branch is not the last branch", branch.Loc())
			}
		}
	}
}

func (detector *unexpectedDefaultBranchesDetector) Exit(node ast.Node) {
}

type unexpectedFuncSignaturesDetector struct {
	processed map[*ast.FuncSignature]struct{}

	util.ErrorEmitter
}

func detectUnexpectedFuncSignatures() util.Pass {
	return &unexpectedFuncSignaturesDetector{
		processed: map[*ast.FuncSignature]struct{}{},
	}
}

func (detector *unexpectedFuncSignaturesDetector) Process(node ast.Node) {
	node.Walk(detector)
}

func (detector *unexpectedFuncSignaturesDetector) Enter(node ast.Node) {
	switch expr := node.(type) {
	case *ast.StatementList:
		detector.processStatements(expr.Elements)
	case *ast.StatementsExpr:
		detector.processStatements(expr.Statements)
	case *ast.PropertiesTypeExpr:
		detector.processPropertiesTypeExpr(expr)
	case *ast.FuncSignature:
		_, ok := detector.processed[expr]
		if ok {
			return
		}

		detector.processAnonymousSignature(expr)
	}
}

func (detector *unexpectedFuncSignaturesDetector) processStatements(
	stmts []ast.Statement,
) {
	for _, stmt := range stmts {
		def, ok := stmt.(*ast.FuncDefinition)
		if !ok {
			continue
		}

		detector.processed[def.Signature] = struct{}{}
		detector.processNamedSignature(def.Signature, true)
	}
}

func (detector *unexpectedFuncSignaturesDetector) processPropertiesTypeExpr(
	typeExpr *ast.PropertiesTypeExpr,
) {
	for _, property := range typeExpr.Properties {
		sig, ok := property.(*ast.FuncSignature)
		if !ok {
			continue
		}

		detector.processed[sig] = struct{}{}

		if typeExpr.Kind != ast.TraitKind {
			detector.Emit(
				"%s: unexpected %s function signature in %s type",
				property.Loc(),
				sig.Kind,
				typeExpr.Kind)
			continue
		}

		detector.processNamedSignature(sig, false)
	}
}

func (detector *unexpectedFuncSignaturesDetector) processNamedSignature(
	sig *ast.FuncSignature,
	allowReceiver bool,
) {
	if sig.Kind == ast.AnonymousFunc {
		detector.Emit(
			"%s: unexpected %s function signature",
			sig.Loc(),
			sig.Kind)
	} else {
		// manual ast construction
		if sig.Name == "" {
			detector.Emit(
				"%s: unexpected empty name string in function signature",
				sig.Loc())
		}
	}

	detector.processParameters(sig, allowReceiver)
}

func (detector *unexpectedFuncSignaturesDetector) processAnonymousSignature(
	sig *ast.FuncSignature,
) {
	if sig.Kind != ast.AnonymousFunc {
		detector.Emit(
			"%s: unexpected %s function signature",
			sig.Loc(),
			sig.Kind)
	} else {
		// manual ast construction
		if sig.Name != "" {
			detector.Emit(
				"%s: unexpected name string (%s) in function signature",
				sig.Loc(),
				sig.Name)
		}

		// manual ast construction
		if sig.GenericParameters != nil {
			detector.Emit(
				"%s: unexpected generic parameters in function signature",
				sig.Loc())
		}
	}

	detector.processParameters(sig, false)
}

func (detector *unexpectedFuncSignaturesDetector) processParameters(
	sig *ast.FuncSignature,
	allowReceiver bool,
) {
	for idx, param := range sig.Parameters.Elements {
		if param.Kind == ast.VarargParameter {
			if idx != len(sig.Parameters.Elements)-1 {
				detector.Emit(
					"%s: %s parameter must be the last parameter in the list",
					param.Loc(),
					param.Kind)
			}
		} else if param.Kind == ast.ReceiverParameter {
			if !allowReceiver {
				detector.Emit("%s: unexpect %s parameter", param.Loc(), param.Kind)
			} else if idx != 0 {
				detector.Emit(
					"%s: %s parameter must be the first argument in the list",
					param.Loc(),
					param.Kind)
			}
		}
	}
}

func (detector *unexpectedFuncSignaturesDetector) Exit(node ast.Node) {
}
