/*
A permutation is an ordered arrangement of objects.
For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4.
If all of the permutations are listed numerically or alphabetically, we call it lexicographic order.
The lexicographic permutations of 0, 1 and 2 are:

012   021   102   120   201   210

What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
*/

package main

import (
	"fmt"
	"math"
	"sort"
)

const (
	max            = 9
	maxPermutation = 1000000
)

func arrayToInt(digits []int) int {
	var number int
	for index, digit := range digits {
		number += digit * int(math.Pow10(index))
	}
	return number
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func mix(firstNumber, secondNumber int) int {
	if secondNumber == 0 {
		return firstNumber * 10
	} else {
		copySecond := secondNumber

		for copySecond > 0 {
			firstNumber *= 10
			copySecond /= 10
		}

		return firstNumber + secondNumber
	}
}

func permutation(alreadyUsed []int) []int {
	beginning := arrayToInt(alreadyUsed)
	permutations := []int{}

	for i := 0; i <= max; i++ {
		if !contains(alreadyUsed, i) {
			if len(alreadyUsed) == max { // The end, we're creating the final int and permutations will only be one integer.
				permutations = append(permutations, mix(beginning, i))
			} else {
				newPermutations := permutation(append(alreadyUsed, i))
				for _, newPermutation := range newPermutations {
					permutations = append(permutations, newPermutation)
				}
			}
		}
	}
	return permutations
}

func main() {
	permutations := permutation([]int{})
	sort.Ints(permutations)
	fmt.Println(permutations[maxPermutation-1])
}
