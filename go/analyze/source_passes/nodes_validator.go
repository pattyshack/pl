package source_passes

import (
	"github.com/pattyshack/gt/parseutil"

	"github.com/pattyshack/pl/analyze/process"
	"github.com/pattyshack/pl/ast"
)

type NodeValidator struct {
	*parseutil.Emitter
}

func ValidateNodes(emitter *parseutil.Emitter) process.Pass {
	return &NodeValidator{
		Emitter: emitter,
	}
}

func (validator *NodeValidator) Process(list *ast.StatementList) {
	process.ParallelWalk(
		list,
		func(ast.Statement) ast.Visitor { return validator })
}

func (validator *NodeValidator) Enter(node ast.Node) {
	validatable, ok := node.(ast.Validator)
	if ok {
		validatable.Validate(validator.Emitter)
	}
}

func (validator *NodeValidator) Exit(node ast.Node) {
}
