/*
We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once. For example, 2143 is a 4-digit pandigital and is also prime.

What is the largest n-digit pandigital prime that exists?
*/

package algo

import (
	"fmt"
	"math"
	"math/big"
)

func E41() {
	var nPermutations []int
	for i := 9; i > 0; i-- {
		nPermutations = permutations([]int{}, i)
		for j := len(nPermutations) - 1; j > 0; j-- {
			if big.NewInt(int64(nPermutations[j])).ProbablyPrime(20) {
				fmt.Println(nPermutations[j])
				i, j = 0, 0
			}
		}
	}
}
func join(i, j int) int {
	e, tempJ := 0, j
	for tempJ > 0 {
		e++
		tempJ = tempJ / 10
	}
	return i*int(math.Pow10(e)) + j
}

func permutations(alreadyUsed []int, n int) []int {
	digitPermutations := []int{}
	if len(alreadyUsed) == n-1 {
		for i := 1; i <= n; i++ {
			if !Contains(alreadyUsed, i) {
				digitPermutations = append(digitPermutations, i)
			}
		}
	} else {
		for i := 1; i <= n; i++ {
			if !Contains(alreadyUsed, i) {
				newPermutations := permutations(append(alreadyUsed, i), n)
				for j := 0; j < len(newPermutations); j++ {
					digitPermutations = append(digitPermutations, join(i, newPermutations[j]))
				}
			}
		}
	}
	return digitPermutations
}
