/*
The prime 41, can be written as the sum of six consecutive primes:

41 = 2 + 3 + 5 + 7 + 11 + 13
This is the longest sum of consecutive primes that adds to a prime below one-hundred.

The longest sum of consecutive primes below one-thousand that adds to a prime, contains 21 terms, and is equal to 953.

Which prime, below one-million, can be written as the sum of the most consecutive primes?
*/

package algo

import (
	"fmt"
	"math/big"
)

func E50() {
	primes := []int{}
	for i := 2; i < 50000; i++ {
		if big.NewInt(int64(i)).ProbablyPrime(20) {
			primes = append(primes, i)
		}
	}

	sum := 0
	biggestConsecutivePrime := []int{}
	for i := len(primes) - 1; i >= 21; i-- {
		for j := 0; j < i-21; j++ {
			if i-j > len(biggestConsecutivePrime) {
				sum = arraySum(primes[j:i])
				if sum < 1000000 && big.NewInt(int64(sum)).ProbablyPrime(20) {
					fmt.Printf("%v = %d\n", primes[j:i], sum)
					biggestConsecutivePrime = primes[j:i]
				}
			}
		}
	}
}

func arraySum(arr []int) int {
	sum := 0
	for _, a := range arr {
		sum += a
	}
	return sum
}
