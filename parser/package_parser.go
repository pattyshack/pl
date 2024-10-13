package parser

import (
	"github.com/pattyshack/gt/lexutil"

	"github.com/pattyshack/pl/analyze"
	"github.com/pattyshack/pl/ast"
	"sync"
)

type packageParser struct {
}

func ParsePackage(
	packageSources []string,
	emitter *lexutil.ErrorEmitter,
	options ParserOptions,
) (
	*ast.StatementList,
	[]*ast.ImportClause,
	bool,
) {
	if len(packageSources) == 0 {
		return nil, nil, false
	}

	results := make(
		[]*ast.StatementList,
		len(packageSources),
		len(packageSources))

	pkgErrEmitter := &lexutil.ErrorEmitter{}

	wg := &sync.WaitGroup{}
	wg.Add(len(packageSources))

	for idx, src := range packageSources {
		go func(idx int, fileName string) {
			defer wg.Done()
			results[idx] = ParseSource(fileName, pkgErrEmitter, options)
		}(idx, src)
	}

	wg.Wait()

	var list *ast.StatementList

	for idx, result := range results {
		if idx == 0 {
			list = result
		} else {
			list.Elements = append(list.Elements, result.Elements...)
		}
	}

	var importClauses []*ast.ImportClause
	if !options.DisableSyntaxAnalysis {
		importClauses = analyze.PackageSyntax(list, pkgErrEmitter)
	}

	emitter.MergeFrom(pkgErrEmitter)
	return list, importClauses, pkgErrEmitter.HasErrors()
}
