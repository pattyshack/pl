package source_passes

import (
	"github.com/pattyshack/gt/parseutil"

	"github.com/pattyshack/pl/analyze/process"
	"github.com/pattyshack/pl/ast"
)

type FuncSignaturesValidator struct {
	*parseutil.Emitter
}

func ValidateFuncSignatures(emitter *parseutil.Emitter) process.Pass {
	return &FuncSignaturesValidator{
		Emitter: emitter,
	}
}

func (validator *FuncSignaturesValidator) Process(list *ast.StatementList) {
	process.ParallelWalk(
		list,
		func(stmt ast.Statement) ast.Visitor {
			return &funcSignaturesValidator{
				root:      stmt,
				processed: map[*ast.FuncSignature]struct{}{},
				Emitter:   validator.Emitter,
			}
		})
}

type funcSignaturesValidator struct {
	root      ast.Node
	processed map[*ast.FuncSignature]struct{}

	*parseutil.Emitter
}

func (validator *funcSignaturesValidator) Enter(node ast.Node) {
	if node == validator.root {
		validator.processStatements(node.(ast.Statement))
	}

	switch expr := node.(type) {
	case *ast.StatementsExpr:
		validator.processStatements(expr.Statements...)
	case *ast.PropertiesTypeExpr:
		validator.processPropertiesTypeExpr(expr)
	case *ast.FuncSignature:
		_, ok := validator.processed[expr]
		if ok {
			return
		}

		validator.processAnonymousSignature(expr)
	}
}

func (validator *funcSignaturesValidator) processStatements(
	stmts ...ast.Statement,
) {
	for _, stmt := range stmts {
		def, ok := stmt.(*ast.FuncDefinition)
		if !ok {
			continue
		}

		validator.processed[def.Signature] = struct{}{}
		if def.Signature.Name != "" {
			validator.processNamedSignature(def.Signature, true)
		} else {
			validator.processAnonymousSignature(def.Signature)
		}
	}
}

func (validator *funcSignaturesValidator) processPropertiesTypeExpr(
	typeExpr *ast.PropertiesTypeExpr,
) {
	for _, property := range typeExpr.Properties {
		// NOTE: valid anonymous signatures are wrapped by FieldDef
		switch prop := property.(type) {
		case *ast.FuncSignature:
			validator.processed[prop] = struct{}{}
			validator.processNamedSignature(prop, false)
		case *ast.FuncDefinition:
			validator.processed[prop.Signature] = struct{}{}
			validator.processNamedSignature(prop.Signature, true)
		}
	}
}

func (validator *funcSignaturesValidator) processNamedSignature(
	sig *ast.FuncSignature,
	allowReceiver bool,
) {
	if sig.Name == "" {
		validator.Emit(sig.Loc(), "unexpected anonymous function signature")
	}

	validator.processParameters(sig, allowReceiver)
}

func (validator *funcSignaturesValidator) processAnonymousSignature(
	sig *ast.FuncSignature,
) {
	if sig.Name != "" {
		validator.Emit(sig.Loc(), "unexpected named function signature")
	} else if sig.GenericParameters != nil {
		validator.Emit(
			sig.Loc(),
			"invalid manual ast construction. "+
				"generic parameters set in anonymous function signature")
	}

	validator.processParameters(sig, false)
}

func (validator *funcSignaturesValidator) processParameters(
	sig *ast.FuncSignature,
	allowReceiver bool,
) {
	for idx, param := range sig.Parameters.Elements {
		if param.Kind == ast.VariadicParameter {
			if idx != len(sig.Parameters.Elements)-1 {
				validator.Emit(
					param.Loc(),
					"%s parameter must be the last parameter in the list",
					param.Kind)
			}
		} else if param.Kind == ast.ReceiverParameter {
			if !allowReceiver {
				validator.Emit(param.Loc(), "unexpect %s parameter", param.Kind)
			} else if idx != 0 {
				validator.Emit(
					param.Loc(),
					"%s parameter must be the first argument in the list",
					param.Kind)
			}
		}
	}
}

func (validator *funcSignaturesValidator) Exit(node ast.Node) {
}
