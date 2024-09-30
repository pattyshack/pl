package reducer

import (
	"github.com/pattyshack/pl/parser/lr"
	"github.com/pattyshack/pl/util"
)

type Reducer struct {
	util.ErrorEmitter
}

var _ lr.Reducer = &Reducer{}

func NewReducer() *Reducer {
	return &Reducer{}
}
