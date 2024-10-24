package syntax

import (
	"github.com/pattyshack/pl/analyze/process"
	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/errors"
)

type NodeValidator struct {
	*errors.Emitter
}

func ValidateNodes(emitter *errors.Emitter) process.Pass {
	return &NodeValidator{
		Emitter: emitter,
	}
}

func (validator *NodeValidator) Process(node ast.Node) {
	node.Walk(validator)
}

func (validator *NodeValidator) Enter(node ast.Node) {
	validatable, ok := node.(ast.Validator)
	if ok {
		validatable.Validate(validator.Emitter)
	}
}

func (validator *NodeValidator) Exit(node ast.Node) {
}
