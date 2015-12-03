/*
A perfect number is a number for which the sum of its proper divisors is exactly equal to the number. For example, the sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.

A number n is called deficient if the sum of its proper divisors is less than n and it is called abundant if this sum exceeds n.

As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that can be written as the sum of two abundant numbers is 24.
By mathematical analysis, it can be shown that all integers greater than 28123 can be written as the sum of two abundant numbers.
However, this upper limit cannot be reduced any further by analysis even though it is known that the greatest number that cannot be expressed as the sum of two abundant numbers is less than this limit.

Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.
*/

package main

import (
	"fmt"
	"math"
)

const (
	max = 28123
)

func isAboundant(number int) bool {
	divisorsSum := 0
	for i := 1; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			divisorsSum += i
			if i != number/i && i != 1 { // i != 1 because number / i would be number
				divisorsSum += (number / i)
			}
		}
	}
	if divisorsSum > number {
		return true
	} else {
		return false
	}
}

func main() {
	aboundants := []int{}
	for i := 0; i <= max; i++ {
		if isAboundant(i) {
			aboundants = append(aboundants, i)
		}
	}

	sumAboundants := map[int]bool{} // Key is between 0 and max, value is being the sum of two aboundant numbers.
	for i := 0; i <= max; i++ {
		sumAboundants[i] = false
	}

	for i := 0; i <= len(aboundants)/2; i++ {
		for j := i; j <= len(aboundants); j++ {
			if aboundants[i]+aboundants[j] <= max {
				sumAboundants[aboundants[i]+aboundants[j]] = true // This key is a sum of two aboundant numbers.
			} else {
				j = len(aboundants)
			}
		}
	}

	sum := 0
	for key, value := range sumAboundants {
		if value == false { // Not a sum of two aboundant numbers, we're adding it to the sum.
			sum += key
		}
	}

	fmt.Println(sum)
}
