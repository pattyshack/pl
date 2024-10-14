package syntax

import (
	"github.com/pattyshack/gt/lexutil"

	"github.com/pattyshack/pl/ast"
)

type FieldDefsValidator struct {
	// push true when FieldDef is entered.  push false when PropertiesTypeExpr
	// is entered
	scopeStack []bool
	*lexutil.ErrorEmitter
}

func ValidateFieldDefs(emitter *lexutil.ErrorEmitter) ast.Pass {
	return &FieldDefsValidator{
		ErrorEmitter: emitter,
	}
}

func (validator *FieldDefsValidator) Process(node ast.Node) {
	node.Walk(validator)
}

func (validator *FieldDefsValidator) push(scope bool) {
	validator.scopeStack = append(validator.scopeStack, scope)
}

func (validator *FieldDefsValidator) pop() {
	validator.scopeStack = validator.scopeStack[:len(validator.scopeStack)-1]
}

func (validator *FieldDefsValidator) isInFieldDef() bool {
	if len(validator.scopeStack) > 0 {
		return validator.scopeStack[len(validator.scopeStack)-1]
	}
	return false
}

func (validator *FieldDefsValidator) Enter(node ast.Node) {
	switch node.(type) {
	case *ast.PropertiesTypeExpr:
		validator.push(false)
	case *ast.FieldDef:
		validator.push(true)
	case *ast.InferredTypeExpr:
		if validator.isInFieldDef() {
			validator.Emit(node.Loc(), "unexpected inferred type in field definition")
		}
	}
}

func (validator *FieldDefsValidator) Exit(node ast.Node) {
	switch node.(type) {
	case *ast.PropertiesTypeExpr:
		validator.pop()
	case *ast.FieldDef:
		validator.pop()
	}
}
