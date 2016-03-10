/*
A googol (10100) is a massive number: one followed by one-hundred zeros; 100100 is almost unimaginably large: one followed by two-hundred zeros.
Despite their size, the sum of the digits in each number is only 1.

Considering natural numbers of the form, ab, where a, b < 100, what is the maximum digital sum?
*/

package algo

import (
	"fmt"
	"math/big"
)

func E56() {
	max := BigFromInt(100)

	one := BigFromInt(1) // Used to increment a and b
	maxSum := 0
	exp := BigFromInt(0) // Value of a ** b

	for a := BigFromInt(0); a.Cmp(max) == -1; a.Add(a, one) {
		for b := BigFromInt(0); b.Cmp(max) == -1; b.Add(b, one) {
			exp.Exp(a, b, nil)
			if digitalSum(exp) > maxSum {
				maxSum = digitalSum(exp)
			}
		}
	}

	fmt.Println(maxSum)
}

func digitalSum(x *big.Int) int {
	digitalSum := 0
	for _, c := range x.String() {
		digitalSum += int(c - '0')
	}
	return digitalSum
}
