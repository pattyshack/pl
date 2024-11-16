package analyze

import (
	"github.com/pattyshack/gt/parseutil"

	"github.com/pattyshack/pl/analyze/process"
	"github.com/pattyshack/pl/analyze/source_passes"
	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/build/cfg"
)

func Source(
	config *cfg.Config,
	source *ast.StatementList,
	emitter *parseutil.Emitter,
) (
	[]*ast.ImportClause,
	bool,
) {
	patternsAnalyzer := source_passes.NewPatternsAnalyzer(emitter)
	importsCollector := source_passes.NewImportClausesCollector(emitter)
	evaluator := source_passes.NewBuildConstraintEvaluator(config)

	passes := [][]process.Pass{
		{
			source_passes.ValidateNodes(emitter),
			source_passes.ValidateFuncSignatures(emitter),
			source_passes.ValidatePackageStatements(emitter),
			source_passes.ValidateExprStatements(emitter),
			source_passes.ValidateImplicitStructs(emitter),
			source_passes.ValidateExplicitlyTypedDefs(emitter),
			source_passes.RejectUnexpectedAnonymousMethodTypes(emitter),
			patternsAnalyzer.Validate(),
			importsCollector,

			source_passes.ValidateDirectives(emitter),
			evaluator,
		},
		{
			patternsAnalyzer.Transform(),
		},
	}

	process.Process(source, passes, nil)

	return importsCollector.Clauses(), evaluator.SatisfyConstraints()
}
