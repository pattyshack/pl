package parser

import (
	"github.com/pattyshack/gt/lexutil"

	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/util"
)

type nodeValidator struct {
	*lexutil.ErrorEmitter
}

func validateNodes() ast.Pass {
	return &nodeValidator{
		ErrorEmitter: &lexutil.ErrorEmitter{},
	}
}

func (detector *nodeValidator) Process(node ast.Node) {
	node.Walk(detector)
}

func (detector *nodeValidator) Enter(node ast.Node) {
	validator, ok := node.(ast.Validator)
	if ok {
		validator.Validate(detector.ErrorEmitter)
	}
}

func (detector *nodeValidator) Exit(node ast.Node) {
}

type unexpectedStatementsDetector struct {
	processed map[*ast.StatementsExpr]struct{}

	lexutil.ErrorEmitter
}

func detectUnexpectedStatements() ast.Pass {
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
			detector.checkTopLevelExpr(expr)
		default:
			invalidStmtType = "statement type:\n" + util.TreeString(stmt, "  ") + "\n"
		}

		if invalidStmtType != "" {
			detector.Emit(stmt.Loc(), "unexpected %s", invalidStmtType)
		}
	}
}

func (detector *unexpectedStatementsDetector) checkTopLevelExpr(
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
		detector.Emit(
			node.Loc(),
			"unexpected source level %s expression%s",
			invalid,
			additional)
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
			detector.Emit(
				node.Loc(),
				"unexpected %s. expected expression, assign or jump statement",
				invalidStmtType)
		}
	}
}

func (detector *unexpectedStatementsDetector) Exit(node ast.Node) {
}

type unexpectedFuncSignaturesDetector struct {
	processed map[*ast.FuncSignature]struct{}

	lexutil.ErrorEmitter
}

func detectUnexpectedFuncSignatures() ast.Pass {
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
		if def.Signature.Name != "" {
			detector.processNamedSignature(def.Signature, true)
		} else {
			detector.processAnonymousSignature(def.Signature)
		}
	}
}

func (detector *unexpectedFuncSignaturesDetector) processPropertiesTypeExpr(
	typeExpr *ast.PropertiesTypeExpr,
) {
	for _, property := range typeExpr.Properties {
		// NOTE: valid anonymous signatures are wrapped by FieldDef
		switch prop := property.(type) {
		case *ast.FuncSignature:
			detector.processed[prop] = struct{}{}
			detector.processNamedSignature(prop, false)
		case *ast.FuncDefinition:
			detector.processed[prop.Signature] = struct{}{}
			detector.processNamedSignature(prop.Signature, true)
		}
	}
}

func (detector *unexpectedFuncSignaturesDetector) processNamedSignature(
	sig *ast.FuncSignature,
	allowReceiver bool,
) {
	if sig.Name == "" {
		detector.Emit(sig.Loc(), "unexpected anonymous function signature")
	}

	detector.processParameters(sig, allowReceiver)
}

func (detector *unexpectedFuncSignaturesDetector) processAnonymousSignature(
	sig *ast.FuncSignature,
) {
	if sig.Name != "" {
		detector.Emit(sig.Loc(), "unexpected named function signature")
	} else if sig.GenericParameters != nil {
		detector.Emit(
			sig.Loc(),
			"invalid manual ast construction. "+
				"generic parameters set in anonymous function signature")
	}

	detector.processParameters(sig, false)
}

func (detector *unexpectedFuncSignaturesDetector) processParameters(
	sig *ast.FuncSignature,
	allowReceiver bool,
) {
	for idx, param := range sig.Parameters.Elements {
		if param.Kind == ast.VariadicParameter {
			if idx != len(sig.Parameters.Elements)-1 {
				detector.Emit(
					param.Loc(),
					"%s parameter must be the last parameter in the list",
					param.Kind)
			}
		} else if param.Kind == ast.ReceiverParameter {
			if !allowReceiver {
				detector.Emit(param.Loc(), "unexpect %s parameter", param.Kind)
			} else if idx != 0 {
				detector.Emit(
					param.Loc(),
					"%s parameter must be the first argument in the list",
					param.Kind)
			}
		}
	}
}

func (detector *unexpectedFuncSignaturesDetector) Exit(node ast.Node) {
}

// All errors detected by this pass can only occur due to manual ast
// construction.  lr parsed tree should be free of these errors.
//
// Improper implicit structs are valid in jump statements, statements
// expr, assign patterns, loop's init/post statements (and iterator condition),
// and top level statements.  In index expr and colon implicit struct,
// unit improper implicit structs are valid arguments.
//
// Colon implicit structs are only allowed in callable expr.
type unexpectedImplicitStructsDetector struct {
	validImproperColon map[*ast.ImplicitStructExpr]struct{}
	lexutil.ErrorEmitter
}

func detectUnexpectedImplicitStructs() ast.Pass {
	return &unexpectedImplicitStructsDetector{
		validImproperColon: map[*ast.ImplicitStructExpr]struct{}{},
	}
}

func (detector *unexpectedImplicitStructsDetector) Process(node ast.Node) {
	node.Walk(detector)
}

func (detector *unexpectedImplicitStructsDetector) Enter(n ast.Node) {
	switch node := n.(type) {
	case *ast.StatementList:
		for _, stmt := range node.Elements {
			detector.allowImproperStruct(stmt)
		}
	case *ast.StatementsExpr:
		for _, stmt := range node.Statements {
			detector.allowImproperStruct(stmt)
		}
	case *ast.LoopExpr:
		if node.Init != nil {
			detector.allowImproperStruct(node.Init)
		}
		if node.Post != nil {
			detector.allowImproperStruct(node.Post)
		}
	case *ast.AssignPattern:
		detector.allowImproperStruct(node.Pattern)
		detector.allowImproperStruct(node.Value)
	case *ast.JumpStmt:
		detector.allowImproperStruct(node.Value)
	case *ast.CallExpr:
		detector.allowColonStructArguments(node.Arguments)
	case *ast.InitializeExpr:
		detector.allowColonStructArguments(node.Arguments)
	case *ast.IndexExpr:
		for _, arg := range node.IndexArgs {
			detector.allowUnitImproperStruct(arg)
		}
	case *ast.ImplicitStructExpr:
		if node.Kind == ast.ProperImplicitStruct {
			detector.allowColonStructArguments(node.Arguments)
			return
		}

		if node.Kind == ast.ColonImplicitStruct {
			for _, arg := range node.Arguments {
				if arg.Kind == ast.SingularArgument {
					detector.allowUnitImproperStruct(arg.Expr)
				}
			}
		}

		_, valid := detector.validImproperColon[node]
		if !valid {
			detector.Emit(
				node.Loc(),
				"invalid ast construction. unexpected %s implicit struct",
				node.Kind)
		}
	case *ast.EnumPattern:
		detector.allowUnitImproperStruct(node.Pattern)
	}
}

func (detector *unexpectedImplicitStructsDetector) allowColonStructArguments(
	args []*ast.Argument,
) {
	for _, arg := range args {
		if arg.Kind == ast.SingularArgument {
			exprStruct, ok := arg.Expr.(*ast.ImplicitStructExpr)
			if ok && exprStruct.Kind == ast.ColonImplicitStruct {
				detector.validImproperColon[exprStruct] = struct{}{}
			}
		}
	}
}

func (detector *unexpectedImplicitStructsDetector) allowUnitImproperStruct(
	expr ast.Expression,
) {
	exprStruct, ok := expr.(*ast.ImplicitStructExpr)
	if ok &&
		exprStruct.Kind == ast.ImproperImplicitStruct &&
		len(exprStruct.Arguments) == 0 {

		detector.validImproperColon[exprStruct] = struct{}{}
	}
}

func (detector *unexpectedImplicitStructsDetector) allowImproperStruct(
	node ast.Node,
) {
	exprStruct, ok := node.(*ast.ImplicitStructExpr)
	if ok && exprStruct.Kind == ast.ImproperImplicitStruct {
		detector.validImproperColon[exprStruct] = struct{}{}
	}
}

func (detector *unexpectedImplicitStructsDetector) Exit(node ast.Node) {
}

type unexpectedFieldDefsDetector struct {
	// push true when FieldDef is entered.  push false when PropertiesTypeExpr
	// is entered
	scopeStack []bool
	lexutil.ErrorEmitter
}

func detectUnexpectedFieldDefs() ast.Pass {
	return &unexpectedFieldDefsDetector{}
}

func (detector *unexpectedFieldDefsDetector) Process(node ast.Node) {
	node.Walk(detector)
}

func (detector *unexpectedFieldDefsDetector) push(scope bool) {
	detector.scopeStack = append(detector.scopeStack, scope)
}

func (detector *unexpectedFieldDefsDetector) pop() {
	detector.scopeStack = detector.scopeStack[:len(detector.scopeStack)-1]
}

func (detector *unexpectedFieldDefsDetector) isInFieldDef() bool {
	if len(detector.scopeStack) > 0 {
		return detector.scopeStack[len(detector.scopeStack)-1]
	}
	return false
}

func (detector *unexpectedFieldDefsDetector) Enter(node ast.Node) {
	switch node.(type) {
	case *ast.PropertiesTypeExpr:
		detector.push(false)
	case *ast.FieldDef:
		detector.push(true)
	case *ast.InferredTypeExpr:
		if detector.isInFieldDef() {
			detector.Emit(node.Loc(), "unexpected inferred type in field definition")
		}
	}
}

func (detector *unexpectedFieldDefsDetector) Exit(node ast.Node) {
	switch node.(type) {
	case *ast.PropertiesTypeExpr:
		detector.pop()
	case *ast.FieldDef:
		detector.pop()
	}
}
