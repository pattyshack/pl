package analyze

import (
	"github.com/pattyshack/pl/analyze/process"
	"github.com/pattyshack/pl/analyze/syntax"
	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/errors"
)

func SourceSyntax(node ast.Node, emitter *errors.Emitter) {
	patternsAnalyzer := syntax.NewPatternsAnalyzer(emitter)

	passes := [][]process.Pass{
		{
			syntax.ValidateNodes(emitter),
			syntax.ValidateFuncSignatures(emitter),
			syntax.ValidatePackageStatements(emitter),
			syntax.ValidateExprStatements(emitter),
			syntax.ValidateImplicitStructs(emitter),
			syntax.ValidateExplicitlyTypedDefs(emitter),
			patternsAnalyzer.Validate(),
		},
		{
			patternsAnalyzer.Transform(),
		},
	}

	process.Process(node, passes, nil)
}

func PackageSyntax(
	node ast.Node,
	emitter *errors.Emitter,
) []*ast.ImportClause {
	importsCollector := syntax.NewImportClausesCollector(emitter)
	passes := [][]process.Pass{
		{
			importsCollector,
			syntax.ValidatePkgInitBlock(emitter),
		},
	}

	process.Process(node, passes, nil)
	return importsCollector.Clauses()
}
