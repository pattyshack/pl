package reducer

import (
	"github.com/pattyshack/pl/errors"
	"github.com/pattyshack/pl/parser/lr"
)

type Reducer struct {
	*errors.Emitter
}

var _ lr.Reducer = &Reducer{}

func NewReducer(emitter *errors.Emitter) *Reducer {
	return &Reducer{
		Emitter: emitter,
	}
}
