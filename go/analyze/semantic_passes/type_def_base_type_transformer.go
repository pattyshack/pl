package semantic_passes

import (
	"github.com/pattyshack/pl/analyze/process"
	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/errors"
)

// If type def's base type is not a properties type expr, wrap the base type.
// Free-floating top level method definitions are moved into the appropriate
// type def.
type TypeDefBaseTypeTransformer struct {
	*errors.Emitter
}

func TransformTypeDefBaseType(emitter *errors.Emitter) process.Pass {
	return &TypeDefBaseTypeTransformer{
		Emitter: emitter,
	}
}

func (transformer *TypeDefBaseTypeTransformer) Process(
	list *ast.StatementList,
) {
	list.Walk(transformer)

	typeDefs := map[string]*ast.PropertiesTypeExpr{}
	for _, stmt := range list.Elements {
		def, ok := stmt.(*ast.TypeDef)
		if ok {
			typeDefs[def.Name] = def.BaseType.(*ast.PropertiesTypeExpr)
		}
	}

	stmts := []ast.Statement{}
	for _, stmt := range list.Elements {
		switch def := stmt.(type) {
		case *ast.FuncDefinition:
			if !def.Signature.IsMethod() {
				stmts = append(stmts, stmt)
				continue
			}

			receiverType := def.Signature.Parameters.Elements[0].Type
			ref, ok := receiverType.(*ast.RefTypeExpr)
			if ok {
				receiverType = ref.Value
			}

			switch t := receiverType.(type) {
			case *ast.InferredTypeExpr:
				transformer.Emit(
					receiverType.Loc(),
					"unable to associate method to type, receiver type must be named")
			case *ast.NamedTypeExpr:
				prop, ok := typeDefs[t.Name]
				if ok {
					prop.Properties = append(prop.Properties, def)
				} else {
					transformer.Emit(receiverType.Loc(), "undefined: %s", t.Name)
				}
			}
		default:
			stmts = append(stmts, stmt)
		}
	}

	list.Elements = stmts
}

func (transformer *TypeDefBaseTypeTransformer) Enter(n ast.Node) {
	switch def := n.(type) {
	case *ast.TypeDef:
		_, ok := def.BaseType.(*ast.PropertiesTypeExpr)
		if ok {
			return
		}

		def.BaseType = &ast.PropertiesTypeExpr{
			StartEndPos: def.BaseType.StartEnd(),
			Kind:        ast.StructKind,
			IsImplicit:  true,
			Properties: []ast.TypeProperty{
				&ast.FieldDef{
					Type: def.BaseType,
				},
			},
		}
	}
}

func (transformer *TypeDefBaseTypeTransformer) Exit(n ast.Node) {
}
