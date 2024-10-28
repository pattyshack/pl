package source_passes

import (
	"github.com/pattyshack/pl/analyze/process"
	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/errors"
)

type ExplicitlyTypedDefsValidator struct {
	*errors.Emitter
}

func ValidateExplicitlyTypedDefs(emitter *errors.Emitter) process.Pass {
	return &ExplicitlyTypedDefsValidator{
		Emitter: emitter,
	}
}

func (validator *ExplicitlyTypedDefsValidator) Process(
	list *ast.StatementList,
) {
	process.ParallelWalk(
		list,
		func() ast.Visitor {
			return &explicitlyTypedDefsValidator{
				Emitter: validator.Emitter,
			}
		})
}

type explicitlyTypedDefsValidator struct {
	// push "" when inferred type is allowed, push non-empty string when type
	// must be explicit.
	scopeStack []string
	*errors.Emitter
}

func (validator *explicitlyTypedDefsValidator) push(kind string) {
	validator.scopeStack = append(validator.scopeStack, kind)
}

func (validator *explicitlyTypedDefsValidator) pop() {
	validator.scopeStack = validator.scopeStack[:len(validator.scopeStack)-1]
}

func (validator *explicitlyTypedDefsValidator) current() string {
	if len(validator.scopeStack) > 0 {
		return validator.scopeStack[len(validator.scopeStack)-1]
	}
	return ""
}

func (validator *explicitlyTypedDefsValidator) Enter(n ast.Node) {
	switch node := n.(type) {
	case *ast.PropertiesTypeExpr:
		validator.push("")
	case *ast.FieldDef:
		validator.push("field definition")
	case *ast.FuncSignature:
		if node.Name != "" {
			validator.push("named function signature")
		}
	case *ast.Parameter:
		if node.Kind == ast.ReceiverParameter {
			validator.push("")
		}
	case *ast.InferredTypeExpr:
		kind := validator.current()
		if kind != "" {
			validator.Emit(node.Loc(), "unexpected inferred type in %s", kind)
		}
	}
}

func (validator *explicitlyTypedDefsValidator) Exit(n ast.Node) {
	switch node := n.(type) {
	case *ast.PropertiesTypeExpr:
		validator.pop()
	case *ast.FieldDef:
		validator.pop()
	case *ast.FuncSignature:
		if node.Name != "" {
			validator.pop()
		}
	case *ast.Parameter:
		if node.Kind == ast.ReceiverParameter {
			validator.pop()
		}
	}
}
