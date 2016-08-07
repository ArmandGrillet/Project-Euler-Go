// What is the 10 001st prime number?

package main

import (
	"fmt"
	"math"
)

func main() {
	const goal = 10001
	var i, primes int

	for primes < goal {
		i++
		if isPrime(i) {
			primes++
		}
	}

	fmt.Println(i)
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	s := int(math.Sqrt(float64(n)))
	for i := 3; i <= s; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
