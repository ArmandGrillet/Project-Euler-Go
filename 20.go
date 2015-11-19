/*
n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
*/

package main

import (
	"fmt"
	"math/big"
)

const (
	begin = 1
	end   = 100
)

func main() {
	factorial := big.NewInt(0).MulRange(begin, end)
	sum, modulo := big.NewInt(0), big.NewInt(0)
	for factorial.Sign() > 0 {
		modulo.Mod(factorial, big.NewInt(10))
		sum.Add(sum, modulo)
		factorial.Div(factorial, big.NewInt(10))
	}
	fmt.Println(sum)
}
