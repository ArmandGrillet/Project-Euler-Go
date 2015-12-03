/*
The number, 1406357289, is a 0 to 9 pandigital number because it is made up of each of the digits 0 to 9 in some order, but it also has a rather interesting sub-string divisibility property.

Let d1 be the 1st digit, d2 be the 2nd digit, and so on. In this way, we note the following:

d2d3d4=406 is divisible by 2
d3d4d5=063 is divisible by 3
d4d5d6=635 is divisible by 5
d5d6d7=357 is divisible by 7
d6d7d8=572 is divisible by 11
d7d8d9=728 is divisible by 13
d8d9d10=289 is divisible by 17
Find the sum of all 0 to 9 pandigital numbers with this property.
*/

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func join(i, j int) int {
	e, tempJ := 0, j
	for tempJ > 0 {
		e++
		tempJ = tempJ / 10
	}
	return i*int(math.Pow10(e)) + j
}

func OneToNinePandigitals(alreadyUsed []int) []int {
	digitPermutations := []int{}
	if len(alreadyUsed) == 8 {
		for i := 1; i <= 9; i++ {
			if !contains(alreadyUsed, i) {
				digitPermutations = append(digitPermutations, i)
			}
		}
	} else {
		for i := 1; i <= 9; i++ {
			if !contains(alreadyUsed, i) {
				newPermutations := OneToNinePandigitals(append(alreadyUsed, i))
				for j := 0; j < len(newPermutations); j++ {
					digitPermutations = append(digitPermutations, join(i, newPermutations[j]))
				}
			}
		}
	}
	return digitPermutations
}

func addZeroToPermutations(permutations []int) []int {
	initialLength := len(permutations)
	var permutationAsString, newPermutationAsString string
	var newPermutation int
	var err error
	for i := 0; i < initialLength; i++ {
		permutationAsString = strconv.Itoa(permutations[i])
		for j := 0; j < 9; j++ {
			newPermutationAsString = strings.Replace(permutationAsString, string(permutationAsString[j]), fmt.Sprintf("%s0", string(permutationAsString[j])), 1) // Adding the 0
			newPermutation, err = strconv.Atoi(newPermutationAsString)
			if err != nil {
				panic(err)
			}
			permutations = append(permutations, newPermutation)
		}
	}
	return permutations
}

func hasDivisibleSubstrings(number int) bool {
	digits, dividors := []int{}, []int{2, 3, 5, 7, 11, 13, 17}
	threeNumbers := 0
	var err error

	for number > 0 {
		digits = append(digits, number%10)
		number = number / 10
	}

	// Reversing the digits to have a good order
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	if len(digits) == 10 { // The 0 was the first digit thus it could have been removed, in other cases we remove the first value
		digits = append(digits[:0], digits[1:]...)
	}

	for i := 0; i < len(dividors); i++ {
		threeNumbers, err = strconv.Atoi(fmt.Sprintf("%d%d%d", digits[i], digits[i+1], digits[i+2]))
		if err != nil {
			panic(err)
		}
		if threeNumbers%dividors[i] != 0 {
			return false
		}
	}
	return true
}

func main() {
	pandigitals := OneToNinePandigitals([]int{})
	pandigitals = addZeroToPermutations(pandigitals)
	sum := 0
	for i := 0; i < len(pandigitals); i++ {
		if hasDivisibleSubstrings(pandigitals[i]) {
			fmt.Println(pandigitals[i])
			sum += pandigitals[i]
		}
	}
	fmt.Println(sum)
}
