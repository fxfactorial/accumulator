package main

import (
	"fmt"

	"github.com/fxfactorial/accumulator"
	"github.com/fxfactorial/accumulator/group"
)

// Te
func main() {
	accumSet := []interface{}{}
	accum := accumulator.New(group.RSA2048)
	fmt.Println(accumSet, accum)
}
