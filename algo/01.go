/*
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/

package algo

import "fmt"

func E1() {
	max := 1000
	multiples := [...]int{3, 5}

	multiplesFound := []int{}
	sums := make(chan int, len(multiples))
	sum := 0

	for _, multiple := range multiples {
		go multiplesSum(multiple, max, multiplesFound, sums)
		sum += <-sums
		multiplesFound = append(multiplesFound, multiple)
	}

	fmt.Println(sum)
}

func multiplesSum(quantity int, max int, avoid []int, ch chan int) {
	sum := 0
	for i := 0; i < max; i += quantity {
		sum += i
		for _, avoidQuantity := range avoid {
			if i%avoidQuantity == 0 {
				sum -= i
				break
			}
		}
	}
	ch <- sum
}
