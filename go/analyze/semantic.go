package analyze

import (
	"github.com/pattyshack/pl/analyze/process"
	"github.com/pattyshack/pl/analyze/semantic_passes"
	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/errors"
)

func Semantic(pkg *ast.StatementList, emitter *errors.Emitter) {
	passes := [][]process.Pass{
		{
			semantic_passes.TransformTypeDefBaseType(emitter),
		},
	}

	process.Process(pkg, passes, nil)
}
