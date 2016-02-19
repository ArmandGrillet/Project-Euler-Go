/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

package algo

import (
	"fmt"
	"os"
	"strconv"
)

func E4() {
	maxNumber := 999

	var productAsString string

	for product := maxNumber * maxNumber; product > 0; product-- {
		productAsString = strconv.Itoa(product)
		if productAsString == Reverse(productAsString) { // Palindrome
			for divisor := maxNumber; divisor > 0; divisor-- {
				if product%divisor == 0 && product/divisor <= maxNumber { // It is the product of two 3-digit numbers.
					fmt.Println(product)
					os.Exit(0)
				}
			}
		}
	}
}
