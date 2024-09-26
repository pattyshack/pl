package reducer

import (
	"github.com/pattyshack/pl/parser/lr"
)

type Reducer struct {
	ParseErrors []error
}

var _ lr.Reducer = &Reducer{}

func NewReducer() *Reducer {
	return &Reducer{}
}

func (reducer *Reducer) MergeFrom(other *Reducer) {
	reducer.ParseErrors = append(reducer.ParseErrors, other.ParseErrors...)
}
