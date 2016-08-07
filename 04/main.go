/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	const max = 999
	const min = 99
	mul := max * max

	for i := mul; i > 0; i-- {
		if i == reverse(i) {
			for j := max; j > min; j-- {
				if i%j == 0 && i/j < max {
					fmt.Println(i)
					os.Exit(0)
				}
			}
		}
	}
}

func reverse(i int) int {
	var reversed int
	for i != 0 {
		reversed = reversed*10 + i%10
		i /= 10
	}
	return reversed
}
