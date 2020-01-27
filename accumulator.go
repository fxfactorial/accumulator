package accumulator

import (
	"fmt"

	"github.com/fxfactorial/accumulator/group"
)

// Tracker ..
type Tracker interface {
	Add(group.ElementSequenceReader)
}

type accum struct {
	primeHash, accum uint64
}

func (a *accum) Add(elems group.ElementSequenceReader) {
	a.primeHash, a.accum = elems.Compute()
	fmt.Println(a.primeHash, a.accum)
}

// New ..
func New(g group.Named) Tracker {
	return &accum{}
}
