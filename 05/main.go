/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	const maxDiv = 20
	i, div := maxDiv, maxDiv-1

	for {
		if i%div == 0 {
			if div > 1 {
				div--
			} else {
				fmt.Println(i)
				os.Exit(0)
			}
		} else {
			i += maxDiv
			div = maxDiv - 1
		}
	}
}
