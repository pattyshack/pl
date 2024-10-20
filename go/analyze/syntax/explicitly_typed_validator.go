package syntax

import (
	"github.com/pattyshack/gt/lexutil"

	"github.com/pattyshack/pl/analyze/process"
	"github.com/pattyshack/pl/ast"
)

type ExplicitlyTypedDefsValidator struct {
	// push true when FieldDef / stand-alone AddrDeclPattern is entered.
	// push false when PropertiesTypeExpr is entered
	scopeStack []bool
	*lexutil.ErrorEmitter
}

func ValidateExplicitlyTypedDefs(emitter *lexutil.ErrorEmitter) process.Pass {
	return &ExplicitlyTypedDefsValidator{
		ErrorEmitter: emitter,
	}
}

func (validator *ExplicitlyTypedDefsValidator) Process(node ast.Node) {
	node.Walk(validator)
}

func (validator *ExplicitlyTypedDefsValidator) push(scope bool) {
	validator.scopeStack = append(validator.scopeStack, scope)
}

func (validator *ExplicitlyTypedDefsValidator) pop() {
	validator.scopeStack = validator.scopeStack[:len(validator.scopeStack)-1]
}

func (validator *ExplicitlyTypedDefsValidator) isInDef() bool {
	if len(validator.scopeStack) > 0 {
		return validator.scopeStack[len(validator.scopeStack)-1]
	}
	return false
}

func (validator *ExplicitlyTypedDefsValidator) Enter(node ast.Node) {
	switch node.(type) {
	case *ast.PropertiesTypeExpr:
		validator.push(false)
	case *ast.FieldDef:
		validator.push(true)
	case *ast.InferredTypeExpr:
		if validator.isInDef() {
			validator.Emit(node.Loc(), "unexpected inferred type in definition")
		}
	}
}

func (validator *ExplicitlyTypedDefsValidator) Exit(node ast.Node) {
	switch node.(type) {
	case *ast.PropertiesTypeExpr:
		validator.pop()
	case *ast.FieldDef:
		validator.pop()
	}
}
