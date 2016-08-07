/*
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/

package main

import "fmt"

func main() {
	const max = 1000
	muls := [2]int{3, 5}

	var sum int
	for i := 0; i < max; i++ {
		for _, mul := range muls {
			if i%mul == 0 {
				sum += i
				break
			}
		}
	}

	fmt.Println(sum)
}
