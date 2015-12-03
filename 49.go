/*
The arithmetic sequence, 1487, 4817, 8147, in which each of the terms increases by 3330, is unusual in two ways: (i) each of the three terms are prime, and, (ii) each of the 4-digit numbers are permutations of one another.

There are no arithmetic sequences made up of three 1-, 2-, or 3-digit primes, exhibiting this property, but there is one other 4-digit increasing sequence.

What 12-digit number do you form by concatenating the three terms in this sequence?
*/

package main

import (
	"fmt"
	"math/big"
)

func indexOf(s []int, e int) int {
	for i := 0; i < len(s); i++ {
		if s[i] == e {
			return i
		}
	}
	return -1
}

func sameDigits(i int, j int) bool {
	digitsI := []int{}
	for i > 0 {
		digitsI = append(digitsI, i%10)
		i = i / 10
	}
	digitsJ := []int{}
	for j > 0 {
		digitsJ = append(digitsJ, j%10)
		j = j / 10
	}

	var positionK int
	for k := 0; k < len(digitsJ); k++ {
		positionK = indexOf(digitsI, digitsJ[k])
		if positionK == -1 {
			return false
		} else { // Removing the digit used.
			digitsI = append(digitsI[:positionK], digitsI[positionK+1:]...)
		}
	}
	return true
}

func main() {
	primes := []int{}
	for i := 1000; i < 10000; i++ {
		if big.NewInt(int64(i)).ProbablyPrime(20) {
			primes = append(primes, i)
		}
	}

	for i := 0; i < len(primes)-2; i++ {
		for j := i + 1; j < len(primes)-1; j++ {
			if sameDigits(primes[i], primes[j]) {
				for k := j + 1; k < len(primes)-1; k++ {
					if sameDigits(primes[j], primes[k]) && primes[j]-primes[i] == primes[k]-primes[j] {
						fmt.Printf("%d, %d, %d\n", primes[i], primes[j], primes[k])
					}
				}
			}
		}
	}
}
