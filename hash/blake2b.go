package hash

import (
	h "hash"

	"golang.org/x/crypto/blake2b"
)

type Reader interface {
}

type Named byte

const (
	// Blake2B ..
	Blake2B Named = iota
)

// New ..
func New(kind Named, payload []byte) h.Hash {
	switch kind {
	case Blake2B:
		hash, _ := blake2b.New256(payload)
		return hash
	default:
		return nil
	}
}
