/*
It was proposed by Christian Goldbach that every odd composite number can be written as the sum of a prime and twice a square.

9 = 7 + 2×1**2
15 = 7 + 2×2**2
21 = 3 + 2×3**2
25 = 7 + 2×3**2
27 = 19 + 2×2**2
33 = 31 + 2×1**2

It turns out that the conjecture was false.

What is the smallest odd composite that cannot be written as the sum of a prime and twice a square?
*/

package main

import (
	"fmt"
	"math/big"
)

func isOddComposite(i int) bool {
	if i%2 != 0 && !big.NewInt(int64(i)).ProbablyPrime(20) { // Not divisible by 2 but non-prime
		return true
	} else {
		return false
	}
}

func respectGoldbach(number int) bool {
	for i := 1; 2*(i*i) < number; i++ {
		if big.NewInt(int64(number - 2*(i*i))).ProbablyPrime(20) {
			return true
		}
	}
	return false
}

func main() {
	i, smallestGoldbachError := 2, 0
	for smallestGoldbachError == 0 {
		if isOddComposite(i) && !respectGoldbach(i) {
			smallestGoldbachError = i
		}
		i++
	}
	fmt.Printf("%d doesn't respect Goldbach\n", smallestGoldbachError)
}
