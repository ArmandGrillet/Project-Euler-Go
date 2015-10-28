/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

package main

import (
	"fmt"
	"os"
)

const (
	maxDivisor = 20
)

func main() {
	number, divisor := maxDivisor, maxDivisor-1 // Number is a multiple of maxDivisor thus we know that it is divisible by maxDivisor.

	for {
		if number%divisor == 0 {
			if divisor > 1 {
				divisor--
			} else { // The current divisor is 1, the number is divisible by all of the numbers from divisor to 1.
				fmt.Println(number)
				os.Exit(0)
			}
		} else { // The number cannot be divided by divided, we're moving on to the next number.
			number += maxDivisor
			divisor = maxDivisor - 1
		}
	}
}
