package analyze

import (
	"github.com/pattyshack/gt/parseutil"

	"github.com/pattyshack/pl/analyze/process"
	"github.com/pattyshack/pl/analyze/semantic_passes"
	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/types"
)

func Semantic(
	catalog *types.Catalog,
	pkg *ast.StatementList,
	emitter *parseutil.Emitter,
) types.PackageInterface {
	passes := [][]process.Pass{
		{
			semantic_passes.PreSemanticAnalysisTransformation(emitter),
		},
	}

	process.Process(pkg, passes, nil)

	return types.PackageInterface{}
}
