package analyze

import (
	"github.com/pattyshack/gt/lexutil"

	"github.com/pattyshack/pl/analyze/syntax"
	"github.com/pattyshack/pl/ast"
)

func SourceSyntax(node ast.Node, emitter *lexutil.ErrorEmitter) {
	patternsAnalyzer := syntax.NewPatternsAnalyzer(emitter)

	passes := [][]ast.Pass{
		{
			syntax.ValidateNodes(emitter),
			syntax.ValidateFuncSignatures(emitter),
			syntax.ValidatePackageStatements(emitter),
			syntax.ValidateExprStatements(emitter),
			syntax.ValidateImplicitStructs(emitter),
			syntax.ValidateFieldDefs(emitter),
			patternsAnalyzer.Validate(),
		},
		{
			patternsAnalyzer.Transform(),
		},
	}

	lexutil.Process(node, passes, nil)
}

func PackageSyntax(
	node ast.Node,
	emitter *lexutil.ErrorEmitter,
) []*ast.ImportClause {
	importsCollector := syntax.NewImportClausesCollector(emitter)
	passes := [][]ast.Pass{
		{
			importsCollector,
			syntax.ValidatePkgInitBlock(emitter),
		},
	}

	lexutil.Process(node, passes, nil)
	return importsCollector.Clauses()
}
