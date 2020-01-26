package accumulator

import (
	"github.com/fxfactorial/accumulator/group"
)

// Reader ..
type Reader interface {
}

type t struct {
}

// New ..
func New(g group.Named) Reader {

	return &t{}
}
