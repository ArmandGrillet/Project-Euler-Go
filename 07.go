/*
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
*/

package main

import (
	"fmt"
	"math/big"
)

const (
	searchedPrimeNumber = 10001
)

func main() {
	counter, number := 0, 0
	for counter < searchedPrimeNumber {
		number++
		if big.NewInt(int64(number)).ProbablyPrime(20) {
			counter++
		}
	}
	fmt.Println(number)
}
