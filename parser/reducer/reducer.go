package reducer

import (
	"github.com/pattyshack/pl/parser/lr"
	"github.com/pattyshack/pl/util"
)

type Reducer struct {
	// TODO: rm
	*util.ErrorEmitter
}

var _ lr.Reducer = &Reducer{}

func NewReducer(emitter *util.ErrorEmitter) *Reducer {
	return &Reducer{
		ErrorEmitter: emitter,
	}
}
