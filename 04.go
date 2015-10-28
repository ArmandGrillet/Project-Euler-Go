/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	maxNumber = 999
)

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	var productAsString string

	for product := maxNumber * maxNumber; product > 0; product-- {
		productAsString = strconv.Itoa(product)
		if productAsString == reverse(productAsString) { // Palindrome
			for divisor := maxNumber; divisor > 0; divisor-- {
				if product%divisor == 0 && product/divisor <= maxNumber { // It is the product of two 3-digit numbers.
					fmt.Println(product)
					os.Exit(0)
				}
			}
		}
	}
}
