package syntax

import (
	"github.com/pattyshack/pl/analyze/process"
	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/errors"
)

type FuncSignaturesValidator struct {
	processed map[*ast.FuncSignature]struct{}

	*errors.Emitter
}

func ValidateFuncSignatures(emitter *errors.Emitter) process.Pass {
	return &FuncSignaturesValidator{
		processed: map[*ast.FuncSignature]struct{}{},
		Emitter:   emitter,
	}
}

func (validator *FuncSignaturesValidator) Process(list *ast.StatementList) {
	list.Walk(validator)
}

func (validator *FuncSignaturesValidator) Enter(node ast.Node) {
	switch expr := node.(type) {
	case *ast.StatementList:
		validator.processStatements(expr.Elements)
	case *ast.StatementsExpr:
		validator.processStatements(expr.Statements)
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

func (validator *FuncSignaturesValidator) processStatements(
	stmts []ast.Statement,
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

func (validator *FuncSignaturesValidator) processPropertiesTypeExpr(
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

func (validator *FuncSignaturesValidator) processNamedSignature(
	sig *ast.FuncSignature,
	allowReceiver bool,
) {
	if sig.Name == "" {
		validator.Emit(sig.Loc(), "unexpected anonymous function signature")
	}

	validator.processParameters(sig, allowReceiver)
}

func (validator *FuncSignaturesValidator) processAnonymousSignature(
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

func (validator *FuncSignaturesValidator) processParameters(
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

func (validator *FuncSignaturesValidator) Exit(node ast.Node) {
}
