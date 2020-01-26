package hash

type blake2B struct {
}

type Reader interface {
}

type Named byte

const (
	Blake2B Named = iota
)

func New(kind Named) Reader {
	return &blake2B{}
}
