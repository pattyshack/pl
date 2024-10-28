package source_passes

import (
	"github.com/pattyshack/pl/analyze/process"
	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/errors"
)

// All errors detected by this pass can only occur due to manual ast
// construction.  lr parsed tree should be free of these errors.
//
// Improper implicit structs are valid in jump statements, statements
// expr, assign patterns, loop's init/post statements (and iterator condition),
// and top level statements.  In index expr and colon implicit struct,
// unit improper implicit structs are valid arguments.
//
// Colon implicit structs are only allowed in callable expr.
type ImplicitStructsValidator struct {
	*errors.Emitter
}

func ValidateImplicitStructs(emitter *errors.Emitter) process.Pass {
	return &ImplicitStructsValidator{
		Emitter: emitter,
	}
}

func (validator *ImplicitStructsValidator) Process(list *ast.StatementList) {
	process.ParallelWalk(
		list,
		func(stmt ast.Statement) ast.Visitor {
			return &implicitStructsValidator{
				root:               stmt,
				validImproperColon: map[*ast.ImplicitStructExpr]struct{}{},
				Emitter:            validator.Emitter,
			}
		})
}

type implicitStructsValidator struct {
	root               ast.Node
	validImproperColon map[*ast.ImplicitStructExpr]struct{}
	*errors.Emitter
}

func (validator *implicitStructsValidator) Enter(n ast.Node) {
	if n == validator.root {
		validator.allowImproperStruct(n.(ast.Statement))
	}

	switch node := n.(type) {
	case *ast.StatementsExpr:
		for _, stmt := range node.Statements {
			validator.allowImproperStruct(stmt)
		}
	case *ast.LoopExpr:
		if node.Init != nil {
			validator.allowImproperStruct(node.Init)
		}
		if node.Post != nil {
			validator.allowImproperStruct(node.Post)
		}
	case *ast.AssignPattern:
		validator.allowImproperStruct(node.Pattern)
		validator.allowImproperStruct(node.Value)
	case *ast.JumpStmt:
		validator.allowImproperStruct(node.Value)
	case *ast.CallExpr:
		validator.allowColonStructArguments(node.Arguments)
	case *ast.InitializeExpr:
		validator.allowColonStructArguments(node.Arguments)
	case *ast.IndexExpr:
		for _, arg := range node.IndexArgs {
			validator.allowUnitImproperStruct(arg)
		}
	case *ast.ImplicitStructExpr:
		if node.Kind == ast.ProperImplicitStruct {
			validator.allowColonStructArguments(node.Arguments)
			return
		}

		if node.Kind == ast.ColonImplicitStruct {
			for _, arg := range node.Arguments {
				if arg.Kind == ast.SingularArgument {
					validator.allowUnitImproperStruct(arg.Expr)
				}
			}
		}

		_, valid := validator.validImproperColon[node]
		if !valid {
			validator.Emit(
				node.Loc(),
				"invalid ast construction. unexpected %s implicit struct",
				node.Kind)
		}
	case *ast.EnumPattern:
		validator.allowUnitImproperStruct(node.Pattern)
	}
}

func (validator *implicitStructsValidator) allowColonStructArguments(
	args []*ast.Argument,
) {
	for _, arg := range args {
		if arg.Kind == ast.SingularArgument {
			exprStruct, ok := arg.Expr.(*ast.ImplicitStructExpr)
			if ok && exprStruct.Kind == ast.ColonImplicitStruct {
				validator.validImproperColon[exprStruct] = struct{}{}
			}
		}
	}
}

func (validator *implicitStructsValidator) allowUnitImproperStruct(
	expr ast.Expression,
) {
	exprStruct, ok := expr.(*ast.ImplicitStructExpr)
	if ok &&
		exprStruct.Kind == ast.ImproperImplicitStruct &&
		len(exprStruct.Arguments) == 0 {

		validator.validImproperColon[exprStruct] = struct{}{}
	}
}

func (validator *implicitStructsValidator) allowImproperStruct(
	node ast.Node,
) {
	exprStruct, ok := node.(*ast.ImplicitStructExpr)
	if ok && exprStruct.Kind == ast.ImproperImplicitStruct {
		validator.validImproperColon[exprStruct] = struct{}{}
	}
}

func (validator *implicitStructsValidator) Exit(node ast.Node) {
}
