package process

import (
	"sync"

	"github.com/pattyshack/pl/ast"
)

// Analysis or transform pass
type ProcessPass[T any] interface {
	Process(T)
}

type Pass = ProcessPass[*ast.StatementList]

func Process[T any](
	node T,
	passes [][]ProcessPass[T], // sequence of parallelizable passes
	shouldEarlyExit func() bool, // optional
) {
	for _, parallelPasses := range passes {
		wg := sync.WaitGroup{}
		wg.Add(len(parallelPasses))
		for _, pass := range parallelPasses {
			go func(pass ProcessPass[T]) {
				pass.Process(node)
				wg.Done()
			}(pass)
		}

		wg.Wait()

		if shouldEarlyExit != nil && shouldEarlyExit() {
			return
		}
	}
}

func ParallelWalk(
	list *ast.StatementList,
	newVisitor func() ast.Visitor,
) []ast.Visitor {
	wg := sync.WaitGroup{}
	wg.Add(len(list.Elements))

	visitors := make([]ast.Visitor, 0, len(list.Elements))
	for _, s := range list.Elements {
		v := newVisitor()
		visitors = append(visitors, v)

		go func(stmt ast.Statement, visitor ast.Visitor) {
			stmt.Walk(visitor)
			wg.Done()
		}(s, v)
	}

	wg.Wait()
	return visitors
}
