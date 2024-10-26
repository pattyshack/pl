package analyze

import (
	"github.com/pattyshack/pl/analyze/process"
	"github.com/pattyshack/pl/analyze/semantic"
	"github.com/pattyshack/pl/analyze/syntax"
	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/build/cfg"
	"github.com/pattyshack/pl/errors"
)

func Source(
	config *cfg.Config,
	source *ast.StatementList,
	emitter *errors.Emitter,
) (
	[]*ast.ImportClause,
	bool,
) {
	patternsAnalyzer := syntax.NewPatternsAnalyzer(emitter)
	importsCollector := syntax.NewImportClausesCollector(emitter)
	evaluator := semantic.NewBuildConstraintEvaluator(config)

	passes := [][]process.Pass{
		{
			syntax.ValidateNodes(emitter),
			syntax.ValidateFuncSignatures(emitter),
			syntax.ValidatePackageStatements(emitter),
			syntax.ValidateExprStatements(emitter),
			syntax.ValidateImplicitStructs(emitter),
			syntax.ValidateExplicitlyTypedDefs(emitter),
			patternsAnalyzer.Validate(),
			importsCollector,

			semantic.ValidateDirectives(emitter),
			evaluator,
		},
		{
			patternsAnalyzer.Transform(),
		},
	}

	process.Process(source, passes, nil)

	return importsCollector.Clauses(), evaluator.SatisfyConstraints()
}
