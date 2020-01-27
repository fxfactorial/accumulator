package primality

import (
	"math/big"
)

const (
	factor = 10
)

// IsProbPrime ..
func IsProbPrime(candidate *big.Int) bool {
	return candidate.ProbablyPrime(10)
}
