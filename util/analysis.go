package util

import (
	"fmt"
	"sync"

	"github.com/pattyshack/pl/ast"
)

type ErrorEmitter struct {
	Errs []error
}

func (emitter *ErrorEmitter) Errors() []error {
	return emitter.Errs
}

func (emitter *ErrorEmitter) MergeFrom(other *ErrorEmitter) {
	emitter.Errs = append(emitter.Errs, other.Errs...)
}

func (emitter *ErrorEmitter) Emit(format string, args ...interface{}) {
	emitter.Errs = append(
		emitter.Errs,
		fmt.Errorf(format, args...))
}

func (emitter *ErrorEmitter) EmitErrors(errs ...error) {
	emitter.Errs = append(emitter.Errs, errs...)
}

// Analysis or transform pass
type Pass interface {
	Process(ast.Node)

	// Accumulated errors from Process
	Errors() []error
}

func Process(
	node ast.Node,
	passes [][]Pass, // sequence of parallelizable passes
	earlyExitErrThreshold int, // 0 or less = unlimited
) []error {
	errors := []error{}
	for _, parallelPasses := range passes {
		wg := sync.WaitGroup{}
		wg.Add(len(parallelPasses))
		for _, pass := range parallelPasses {
			go func(pass Pass) {
				pass.Process(node)
				wg.Done()
			}(pass)
		}

		wg.Wait()

		for _, pass := range parallelPasses {
			errors = append(errors, pass.Errors()...)
		}

		if earlyExitErrThreshold > 0 && len(errors) >= earlyExitErrThreshold {
			return errors
		}
	}

	return errors
}
