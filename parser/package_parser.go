package parser

import (
	"github.com/pattyshack/pl/ast"
)

type parseSourceResult struct {
	Defs        *ast.DefinitionList
	ParseErrors []error
	Err         error
}

type packageParser struct {
}

func ParsePackage(
	packageSources []string,
	options ParserOptions,
) (
	*ast.DefinitionList,
	[]error,
	error,
) {
	sourceChan := make(chan parseSourceResult, len(packageSources))
	for _, src := range packageSources {
		go func(fileName string) {
			list, parseErrors, err := ParseSource(fileName, options)
			sourceChan <- parseSourceResult{
				Defs:        list,
				ParseErrors: parseErrors,
				Err:         err,
			}
		}(src)
	}

	var parseErrors []error
	var list *ast.DefinitionList
	for i := 0; i < len(packageSources); i++ {
		result := <-sourceChan
		if result.Err != nil {
			return nil, nil, result.Err
		}

		if i == 0 {
			list = result.Defs
			parseErrors = result.ParseErrors
		} else {
			list.Elements = append(list.Elements, result.Defs.Elements...)
			parseErrors = append(parseErrors, result.ParseErrors...)
		}
	}

	return list, parseErrors, nil
}
