/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/
package main

import "fmt"

func main() {
	number := 600851475143
	var div, largestDiv int

	for number > 1 {
		div = 2

		// Incrementing the divisor until it divides properly.
		for number%div != 0 {
			div++
		}

		if div > largestDiv {
			largestDiv = div
		}

		number /= div
	}

	fmt.Println(div)
}
