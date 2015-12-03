/*
There are exactly ten ways of selecting three from five, 12345:

123, 124, 125, 134, 135, 145, 234, 235, 245, and 345

In combinatorics, we use the notation, 5C3 = 10.

How many, not necessarily distinct, values of  nCr, for 1 ≤ n ≤ 100, are greater than one-million?
*/

package main

import (
	"fmt"
)

func factorial(x int) int {
	if x == 0 {
		return 1
	} else {
		return x * factorial(x-1)
	}
}

func combinator(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
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
