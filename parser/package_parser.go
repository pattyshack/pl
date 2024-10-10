package parser

import (
	"github.com/pattyshack/pl/analyze"
	"github.com/pattyshack/pl/ast"
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
	[]error,
	error,
) {
	sourceChan := make(chan parseSourceResult, len(packageSources))
	for _, src := range packageSources {
		go func(fileName string) {
			list, parseErrors, err := ParseSource(fileName, options)
			sourceChan <- parseSourceResult{
				Stmts:       list,
				ParseErrors: parseErrors,
				Err:         err,
			}
		}(src)
	}

	var parseErrors []error
	var list *ast.StatementList
	for i := 0; i < len(packageSources); i++ {
		result := <-sourceChan
		if result.Err != nil {
			return nil, nil, result.Err
		}

		if i == 0 {
			list = result.Stmts
			parseErrors = result.ParseErrors
		} else {
			list.Elements = append(list.Elements, result.Stmts.Elements...)
			parseErrors = append(parseErrors, result.ParseErrors...)
		}
	}

	parseErrors = append(parseErrors, analyze.PackageSyntax(list)...)

	return list, parseErrors, nil
}
