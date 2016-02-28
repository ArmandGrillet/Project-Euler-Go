/*
There are exactly ten ways of selecting three from five, 12345:

123, 124, 125, 134, 135, 145, 234, 235, 245, and 345

In combinatorics, we use the notation, 5C3 = 10.

How many, not necessarily distinct, values of  nCr, for 1 ≤ n ≤ 100, are greater than one-million?
*/

package algo

import (
	"fmt"
)

func E53() {
	greaterThanOneMillion := 0
	for n := 23; n <= 100; n++ {
		for r := 1; r <= n; r++ {
			fmt.Printf("%d %d\n", n, r)
			if combinator(n, r) > 1000000 {
				greaterThanOneMillion += (n - r) + 1
				r = n
			}
		}
	}
	fmt.Println(greaterThanOneMillion)
}

func combinator(n, r int) int {
	return Factorial(n) / (Factorial(r) * Factorial(n-r))
}
