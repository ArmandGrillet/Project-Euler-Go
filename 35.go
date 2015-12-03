/*
The number, 197, is called a circular prime because all rotations of the digits: 197, 971, and 719, are themselves prime.

There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.

How many circular primes are there below one million?
*/

package main

import (
	"fmt"
	"math"
	"math/big"
)

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getRotations(number int) []int {
	digits := []int{}
	for number > 0 {
		digits = append(digits, number%10)
		number = number / 10
	}

	rotations := []int{}
	for i := 0; i < len(digits); i++ {
		rotationNumber := 0
		for j := len(digits) - 1; j >= 0; j-- {
			position := (j + i) % len(digits)
			rotationNumber += digits[position] * int(math.Pow10(j))
		}
		rotations = append(rotations, rotationNumber)
	}

	return rotations
}

func main() {
	max := 1000000
	circularPrimes, primeRotations := []int{}, []int{}
	var isCircularPrime bool
	for i := 0; i < max; i++ {
		if big.NewInt(int64(i)).ProbablyPrime(20) { // Is prime
			if !contains(circularPrimes, i) { // Isn't already in the circular primes
				isCircularPrime = true
				primeRotations = getRotations(i)
				for i := 0; i < len(primeRotations); i++ {
					if isCircularPrime {
						isCircularPrime = big.NewInt(int64(primeRotations[i])).ProbablyPrime(20)
					}
				}
				if isCircularPrime {
					fmt.Println(i)
					for i := 0; i < len(primeRotations); i++ {
						if !contains(circularPrimes, primeRotations[i]) {
							circularPrimes = append(circularPrimes, primeRotations[i]) // Appending all the primes because they're all circulars
						}
					}
				}
			}
		}
	}
	fmt.Println(len(circularPrimes))
}
