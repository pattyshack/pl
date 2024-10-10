package analyze

import (
	"github.com/pattyshack/gt/lexutil"

	"github.com/pattyshack/pl/analyze/syntax"
	"github.com/pattyshack/pl/ast"
)

func SourceSyntax(node ast.Node) []error {
	patternsAnalyzer := syntax.NewPatternsAnalyzer()

	passes := [][]ast.Pass{
		{
			syntax.ValidateNodes(),
			syntax.ValidateFuncSignatures(),
			syntax.ValidatePackageStatements(),
			syntax.ValidateExprStatements(),
			syntax.ValidateImplicitStructs(),
			syntax.ValidateFieldDefs(),
			patternsAnalyzer.Validate(),
		},
		{
			patternsAnalyzer.Transform(),
		},
	}

	return lexutil.Process(node, passes, 0)
}

func PackageSyntax(node ast.Node) []error {
	passes := [][]ast.Pass{
		{
		},
	}

	return lexutil.Process(node, passes, 0)
}
