package semantic_passes

import (
	"github.com/pattyshack/pl/analyze/process"
	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/errors"
)

// The following transformations are applied:
//   - If type def's base type is not a properties type expr, wrap the base
//     type with a struct properties type expr.
//   - Free-floating top level method definitions are moved into the appropriate
//     type def.
//   - BlockAddrDeclStmt are flattened/replaced by list of pattern expressions.
type PreSemanticAnalysisTransformer struct {
	*errors.Emitter
}

func PreSemanticAnalysisTransformation(emitter *errors.Emitter) process.Pass {
	return &PreSemanticAnalysisTransformer{
		Emitter: emitter,
	}
}

func (transformer *PreSemanticAnalysisTransformer) Process(
	list *ast.StatementList,
) {
	process.ParallelWalk(list, func() ast.Visitor { return transformer })

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
		case *ast.BlockAddrDeclStmt:
			for _, pattern := range def.Patterns {
				stmts = append(stmts, pattern)
			}
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

func (transformer *PreSemanticAnalysisTransformer) Enter(n ast.Node) {
	switch node := n.(type) {
	case *ast.StatementsExpr:
		node.Statements = transformer.replaceBlockAddrDeclStmt(node.Statements)
	case *ast.TypeDef:
		_, ok := node.BaseType.(*ast.PropertiesTypeExpr)
		if ok {
			return
		}

		node.BaseType = &ast.PropertiesTypeExpr{
			StartEndPos: node.BaseType.StartEnd(),
			Kind:        ast.StructKind,
			IsImplicit:  true,
			Properties: []ast.TypeProperty{
				&ast.FieldDef{
					Type: node.BaseType,
				},
			},
		}
	}
}

func (transformer *PreSemanticAnalysisTransformer) Exit(n ast.Node) {
}

func (transformer *PreSemanticAnalysisTransformer) replaceBlockAddrDeclStmt(
	stmts []ast.Statement,
) []ast.Statement {
	result := make([]ast.Statement, 0, len(stmts))
	for _, stmt := range stmts {
		decl, ok := stmt.(*ast.BlockAddrDeclStmt)
		if !ok {
			result = append(result, stmt)
			continue
		}

		for _, pattern := range decl.Patterns {
			result = append(result, pattern)
		}
	}
	return result
}
