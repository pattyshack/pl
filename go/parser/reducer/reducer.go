package reducer

import (
	"github.com/pattyshack/gt/parseutil"

	"github.com/pattyshack/pl/parser/lr"
)

type Reducer struct {
	*parseutil.Emitter
}

var _ lr.Reducer = &Reducer{}

func NewReducer(emitter *parseutil.Emitter) *Reducer {
	return &Reducer{
		Emitter: emitter,
	}
}
