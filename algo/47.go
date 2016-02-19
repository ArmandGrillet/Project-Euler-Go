/*
The first two consecutive numbers to have two distinct prime factors are:

14 = 2 × 7
15 = 3 × 5

The first three consecutive numbers to have three distinct prime factors are:

644 = 2² × 7 × 23
645 = 3 × 5 × 43
646 = 2 × 17 × 19.

Find the first four consecutive integers to have four distinct prime factors. What is the first of these numbers?
*/

package algo

import (
	"fmt"
)

func E47() {
	i, fourConsecutive := 644, 0
	havePrimeFactors := [4]bool{false, false, false, false}
	for fourConsecutive == 0 {
		for j := 0; j < len(havePrimeFactors)-1; j++ {
			havePrimeFactors[j] = havePrimeFactors[j+1]
		}

		if numberOfPrimeFactors(i) == 4 {
			havePrimeFactors[3] = true
			if havePrimeFactors[0] && havePrimeFactors[1] && havePrimeFactors[2] && havePrimeFactors[3] {
				fourConsecutive = i - 3
			} else {
				i++
			}
		} else {
			havePrimeFactors[3] = false
			i += 4
		}
	}

	fmt.Println(fourConsecutive)
}
func numberOfPrimeFactors(number int) int {
	var divisor int
	primeFactors := []int{}

	for number != 1 {
		divisor = 2
		// Incrementing the divisor until it divides
		for number%divisor != 0 {
			divisor++
		}

		// Dividing the number
		number = number / divisor

		if !Contains(primeFactors, divisor) {
			primeFactors = append(primeFactors, divisor)
		}
	}
	return len(primeFactors)
}
