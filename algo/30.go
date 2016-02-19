/*
Surprisingly there are only three numbers that can be written as the sum of fourth powers of their digits:

1634 = 1^4 + 6^4 + 3^4 + ^4
8208 = 8^4 + 2^4 + 0^4 + 8^4
9474 = 9^4 + 4^4 + 7^4 + 4^4
As 1 = 1^4 is not a sum it is not included.

The sum of these numbers is 1634 + 8208 + 9474 = 19316.

Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.
*/

package algo

import (
	"fmt"
	"math"
)

const ()

func E30() {
	power := 5
	min := 2
	max := 236196 // (9^5) * 5

	sum := 0
	for i := min; i <= max; i++ {
		if i == sumOfPowers(i, power) {
			fmt.Println(i)
			sum += i
		}
	}
	fmt.Println(sum)
}

func sumOfPowers(number int, pow int) int {
	sum := 0
	for number > 0 {
		sum += int(math.Pow(float64(number%10), float64(pow)))
		number /= 10
	}
	return sum
}
