package source_passes

import (
	"github.com/pattyshack/pl/analyze/process"
	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/errors"
)

type AnonymousMethodTypesRejector struct {
	*errors.Emitter
}

func RejectUnexpectedAnonymousMethodTypes(
	emitter *errors.Emitter,
) process.Pass {
	return &AnonymousMethodTypesRejector{
		Emitter: emitter,
	}
}

func (rejector *AnonymousMethodTypesRejector) Process(
	list *ast.StatementList,
) {
	process.ParallelWalk(
		list,
		func(ast.Statement) ast.Visitor {
			return &anonymousMethodTypesRejector{Emitter: rejector.Emitter}
		})
}

type anonymousMethodTypesRejector struct {
	scope int // increment on func signature / field def
	*errors.Emitter
}

func (rejector *anonymousMethodTypesRejector) Enter(n ast.Node) {
	switch node := n.(type) {
	case *ast.FieldDef:
		rejector.scope++
	case *ast.FuncSignature:
		rejector.scope++
	case *ast.PropertiesTypeExpr:
		if rejector.scope == 0 {
			return
		}

		for _, prop := range node.Properties {
			_, ok := prop.(*ast.FuncDefinition)
			if ok {
				rejector.Emit(
					prop.Loc(),
					"unexpected method definition, expecting either simple data type "+
						"or named type")
			}
		}
	}
}

func (rejector *anonymousMethodTypesRejector) Exit(n ast.Node) {
	switch n.(type) {
	case *ast.FieldDef:
		rejector.scope--
	case *ast.FuncSignature:
		rejector.scope--
	}
}
