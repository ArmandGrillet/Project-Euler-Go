/*
A permutation is an ordered arrangement of objects.
For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4.
If all of the permutations are listed numerically or alphabetically, we call it lexicographic order.
The lexicographic permutations of 0, 1 and 2 are:

012   021   102   120   201   210

What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
*/

package algo

import (
	"fmt"
	"strconv"
)

func E24() {
	goal := 1000000
	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // Digits available

	var permutationGap int // Number of possible permutations for a certain number of digits.
	millionth := ""        // Millionth lexicographic permutation

	for len(digits) > 0 {
		permutationGap = Factorial(len(digits)) / len(digits)
		for i := 0; ; i++ {
			if goal < (i*permutationGap)+permutationGap {
				millionth += strconv.Itoa(digits[i])
				digits = append(digits[:i], digits[i+1:]...)
				goal -= i * permutationGap
				break
			}
		}
	}
	fmt.Println(millionth)
}
