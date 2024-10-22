package analyze

import (
	"github.com/pattyshack/pl/analyze/process"
	"github.com/pattyshack/pl/analyze/semantic"
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

			// NOTE: directive semantic is checked here since the compiler will make
			// use of these directives.
			semantic.ValidateDirectives(emitter),
		},
		{
			patternsAnalyzer.Transform(),
		},
	}

	process.Process(node, passes, nil)
}

func TargetSyntax(
	isLibrary bool,
	node ast.Node,
	emitter *errors.Emitter,
) []*ast.ImportClause {
	importsCollector := syntax.NewImportClausesCollector(emitter)
	passes := [][]process.Pass{
		{
			importsCollector,
			syntax.ValidatePkgInitBlock(isLibrary, emitter),
		},
	}

	process.Process(node, passes, nil)
	return importsCollector.Clauses()
}
