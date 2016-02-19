/*
The Fibonacci sequence is defined by the recurrence relation:

Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.
Hence the first 12 terms will be:

F1 = 1
F2 = 1
F3 = 2
F4 = 3
F5 = 5
F6 = 8
F7 = 13
F8 = 21
F9 = 34
F10 = 55
F11 = 89
F12 = 144
The 12th term, F12, is the first term to contain three digits.

What is the index of the first term in the Fibonacci sequence to contain 1000 digits?
*/

package algo

import (
	"fmt"
	"math/big"
)

func E25() {
	fnMinusTwo, fnMinusOne, fn := big.NewInt(1), big.NewInt(1), big.NewInt(2)
	index := big.NewInt(3)
	var limit big.Int
	limit.Exp(big.NewInt(10), big.NewInt(999), nil)
	for fn.Cmp(&limit) < 0 {
		fnMinusTwo = fnMinusOne
		fnMinusOne = fn
		fn = big.NewInt(0)
		fn.Add(fnMinusOne, fnMinusTwo)
		index.Add(index, big.NewInt(1))
	}
	fmt.Println(index)
}
