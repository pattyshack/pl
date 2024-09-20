package parser

import (
	"github.com/pattyshack/pl/ast"
	reducerImpl "github.com/pattyshack/pl/parser/reducer"
)

type parseSourceResult struct {
	*reducerImpl.Reducer
	*ast.DefinitionList
	Err error
}

type packageParser struct {
}

func ParsePackage(
	packageSources []string,
	options ParserOptions,
) (
	*reducerImpl.Reducer,
	*ast.DefinitionList,
	error,
) {
	sourceChan := make(chan parseSourceResult, len(packageSources))
	for _, src := range packageSources {
		go func(fileName string) {
			reducer, list, err := ParseSource(fileName, options)
			sourceChan <- parseSourceResult{reducer, list, err}
		}(src)
	}

	var reducer *reducerImpl.Reducer
	var list *ast.DefinitionList
	for i := 0; i < len(packageSources); i++ {
		result := <-sourceChan
		if result.Err != nil {
			return nil, nil, result.Err
		}

		if i == 0 {
			reducer = result.Reducer
			list = result.DefinitionList
		} else {
			reducer.MergeFrom(result.Reducer)
			list.Elements = append(list.Elements, result.DefinitionList.Elements...)
		}
	}

	return reducer, list, nil
}
