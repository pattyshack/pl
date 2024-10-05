package reducer

import (
	"github.com/pattyshack/gt/lexutil"

	"github.com/pattyshack/pl/parser/lr"
)

type Reducer struct {
	*lexutil.ErrorEmitter
}

var _ lr.Reducer = &Reducer{}

func NewReducer(emitter *lexutil.ErrorEmitter) *Reducer {
	return &Reducer{
		ErrorEmitter: emitter,
	}
}
