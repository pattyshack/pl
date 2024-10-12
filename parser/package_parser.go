package parser

import (
	"github.com/pattyshack/pl/analyze"
	"github.com/pattyshack/pl/ast"
	"sync"
)

type parseSourceResult struct {
	Stmts       *ast.StatementList
	ParseErrors []error
	Err         error
}

type packageParser struct {
}

func ParsePackage(
	packageSources []string,
	options ParserOptions,
) (
	*ast.StatementList,
	[]*ast.ImportClause,
	[]error,
	error,
) {
	if len(packageSources) == 0 {
		return nil, nil, nil, nil
	}

	results := make([]parseSourceResult, len(packageSources), len(packageSources))

	wg := &sync.WaitGroup{}
	wg.Add(len(packageSources))

	for idx, src := range packageSources {
		go func(idx int, fileName string) {
			defer wg.Done()
			list, parseErrors, err := ParseSource(fileName, options)
			results[idx] = parseSourceResult{
				Stmts:       list,
				ParseErrors: parseErrors,
				Err:         err,
			}
		}(idx, src)
	}

	wg.Wait()

	var list *ast.StatementList
	var parseErrors []error

	for idx, result := range results {
		if result.Err != nil {
			return nil, nil, nil, result.Err
		}

		if idx == 0 {
			list = result.Stmts
			parseErrors = result.ParseErrors
		} else {
			list.Elements = append(list.Elements, result.Stmts.Elements...)
			parseErrors = append(parseErrors, result.ParseErrors...)
		}
	}

	var importClauses []*ast.ImportClause
	if options.AnalyzeSyntax {
		var errs []error
		importClauses, errs = analyze.PackageSyntax(list)
		parseErrors = append(parseErrors, errs...)
	}

	return list, importClauses, parseErrors, nil
}
