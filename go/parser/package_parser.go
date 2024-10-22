package parser

import (
	"sync"

	"github.com/pattyshack/pl/analyze"
	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/errors"
)

type parsedSources struct {
	Definitions  *ast.StatementList
	Dependencies []*ast.ImportClause
}

type ParsedPackage struct {
	Library *ast.StatementList
	Test    *ast.StatementList

	Dependencies []*ast.ImportClause
}

func ParsePackage(
	libSrcs []string,
	testSrcs []string,
	emitter *errors.Emitter,
	options ParserOptions,
) ParsedPackage {
	var lib parsedSources
	var test parsedSources

	wg := &sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		lib = parseSources(true, libSrcs, emitter, options)
	}()
	go func() {
		defer wg.Done()
		test = parseSources(false, testSrcs, emitter, options)
	}()

	wg.Wait()

	deps := lib.Dependencies
	deps = append(deps, test.Dependencies...)

	return ParsedPackage{
		Library:      lib.Definitions,
		Test:         test.Definitions,
		Dependencies: deps,
	}
}

func parseSources(
	isLibrary bool,
	sources []string,
	emitter *errors.Emitter,
	options ParserOptions,
) parsedSources {
	if len(sources) == 0 {
		return parsedSources{}
	}

	results := make(
		[]*ast.StatementList,
		len(sources),
		len(sources))

	wg := &sync.WaitGroup{}
	wg.Add(len(sources))

	for idx, src := range sources {
		go func(idx int, fileName string) {
			defer wg.Done()
			results[idx] = ParseSource(fileName, emitter, options)
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
		importClauses = analyze.TargetSyntax(isLibrary, list, emitter)
	}

	return parsedSources{
		Definitions:  list,
		Dependencies: importClauses,
	}
}
