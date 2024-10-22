package semantic

import (
	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/build/cfg"
)

type buildConstraintExprEvaluator struct {
	*cfg.Config

	stack []bool
}

func (eval *buildConstraintExprEvaluator) push(val bool) {
	eval.stack = append(eval.stack, val)
}

func (eval *buildConstraintExprEvaluator) pop() bool {
	idx := len(eval.stack) - 1
	top := eval.stack[idx]
	eval.stack = eval.stack[:idx]
	return top
}

func (eval *buildConstraintExprEvaluator) Enter(node ast.Node) {
}

func (eval *buildConstraintExprEvaluator) Exit(node ast.Node) {
	switch expr := node.(type) {
	case *ast.DirectiveValueExpr:
		_, found := eval.BuildTags[expr.Value]
		eval.push(found)
	case *ast.DirectiveGroupExpr: // do nothing
	case *ast.DirectiveNotExpr:
		operand := eval.pop()
		eval.push(!operand)
	case *ast.DirectiveAndExpr:
		right := eval.pop()
		left := eval.pop()
		eval.push(left && right)
	case *ast.DirectiveOrExpr:
		right := eval.pop()
		left := eval.pop()
		eval.push(left || right)
	}
}

type BuildConstraintEvaluator struct {
	*cfg.Config

	satisfy bool
}

func NewBuildConstraintEvaluator(config *cfg.Config) *BuildConstraintEvaluator {
	return &BuildConstraintEvaluator{
		Config:  config,
		satisfy: true,
	}
}

func (pass *BuildConstraintEvaluator) Process(node ast.Node) {
	if pass.Config == nil {
		return
	}

	list, ok := node.(*ast.StatementList)
	if !ok || len(list.Elements) == 0 {
		return
	}

	decl, ok := list.Elements[0].(*ast.DirectivesDecl)
	if !ok || len(decl.Directives) == 0 {
		return
	}

	directive := decl.Directives[0]
	if directive.Name != "build" ||
		directive.SubName != "" ||
		len(directive.Arguments) != 1 {

		return
	}

	eval := &buildConstraintExprEvaluator{
		Config: pass.Config,
	}

	directive.Arguments[0].Walk(eval)
	if len(eval.stack) != 1 {
		panic("This should never happen")
	}

	pass.satisfy = eval.stack[0]
}

func (pass *BuildConstraintEvaluator) SatisfyConstraints() bool {
	return pass.satisfy
}
