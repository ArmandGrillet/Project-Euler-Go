/*
Surprisingly there are only three numbers that can be written as the sum of fourth powers of their digits:

1634 = 14 + 64 + 34 + 44
8208 = 84 + 24 + 04 + 84
9474 = 94 + 44 + 74 + 44
As 1 = 14 is not a sum it is not included.

The sum of these numbers is 1634 + 8208 + 9474 = 19316.

Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.
*/

package main

import (
	"fmt"
	"math"
)

const (
	power = 5
	min   = 10000
	max   = 236196 // (9^5) * 5
)

func sumOfPowers(number int, pow int) int {
	sum := 0
	for number > 0 {
		sum += int(math.Pow(float64(number%10), float64(pow)))
		number /= 10
	}
	return sum
}

func main() {
	sum := 0
	for i := min; i <= max; i++ {
		if i == sumOfPowers(i, power) {
			sum += i
		}
	}
	fmt.Println(sum)
}
