/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

package main

import "fmt"

const (
	number = 600851475143
)

func main() {
	var divisor, largestDivisor int

	for number != 1 {
		divisor = 2
		// Incrementing the divisor until it divides
		for number%divisor != 0 {
			divisor++
		}

		// Dividing the number
		number = number / divisor

		// Updating the largest divisor
		if divisor > largestDivisor {
			largestDivisor = divisor
		}
	}

	fmt.Println(largestDivisor)
}
