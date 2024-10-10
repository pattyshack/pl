package syntax

import (
	"github.com/pattyshack/gt/lexutil"

	"github.com/pattyshack/pl/ast"
)

type NodeValidator struct {
	*lexutil.ErrorEmitter
}

func ValidateNodes() ast.Pass {
	return &NodeValidator{
		ErrorEmitter: &lexutil.ErrorEmitter{},
	}
}

func (validator *NodeValidator) Process(node ast.Node) {
	node.Walk(validator)
}

func (validator *NodeValidator) Enter(node ast.Node) {
	validatable, ok := node.(ast.Validator)
	if ok {
		validatable.Validate(validator.ErrorEmitter)
	}
}

func (validator *NodeValidator) Exit(node ast.Node) {
}
