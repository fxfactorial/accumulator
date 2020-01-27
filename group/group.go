package group

import (
	"math/big"
)

// Named is a named group
type Named byte

const (
	// RSA2048 enum
	RSA2048 Named = iota
)

type rSA2048Elem struct {
	number *big.Int
}

type RSA2048Evaluator interface {
	Op() RSA2048Evaluator
	Identity() RSA2048Evaluator
	Inv() RSA2048Evaluator
	Exp() RSA2048Evaluator
}

func (elem *rSA2048Elem) Op() RSA2048Evaluator {
	return nil
}

func (elem *rSA2048Elem) Identity() RSA2048Evaluator {
	return nil
}

func (elem *rSA2048Elem) Inv() RSA2048Evaluator {
	return nil
}

func (elem *rSA2048Elem) Exp() RSA2048Evaluator {
	return nil
}

// RSA2048MODULUSDECIMAL , taken from [Wikipedia](https://en.wikipedia.org/wiki/RSA_numbers#RSA-2048).
const RSA2048MODULUSDECIMAL = "" +
	"251959084756578934940271832400483985714292821262040320277771378360436620207075955562640185258807" +
	"8440691829064124951508218929855914917618450280848912007284499268739280728777673597141834727026189" +
	"6375014971824691165077613379859095700097330459748808428401797429100642458691817195118746121515172" +
	"6546322822168699875491824224336372590851418654620435767984233871847744479207399342365848238242811" +
	"9816381501067481045166037730605620161967625613384414360383390441495263443219011465754445417842402" +
	"0924616515723350778707749817125772467962926386356373289912154831438167899885040445364023527381951" +
	"378636564391212010397122822120720357"

var (
	rsa2048Modulus, _     = new(big.Int).SetString(RSA2048MODULUSDECIMAL, 10)
	rsa2048ModulusHalf, _ = new(big.Int).SetString(RSA2048MODULUSDECIMAL, 10)
)

func init() {
	rsa2048ModulusHalf.Div(rsa2048ModulusHalf, big.NewInt(2))
}

// ..
func NewRSA2048Elem(t int64) RSA2048Evaluator {
	modVal := new(big.Int).Mod(big.NewInt(t), rsa2048Modulus)
	return &rSA2048Elem{modVal}
	// come back to this
	// if mod.Cmp(rsa2048ModulusHalf) == 1 {
	// 	return &rSA2048Elem{}
	// }
}

type ElementSequenceReader interface {
	Compute() (uint64, uint64)
}
