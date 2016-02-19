/*
145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.

Find the sum of all numbers which are equal to the sum of the factorial of their digits.

Note: as 1! = 1 and 2! = 2 are not sums they are not included.
*/

package algo

import (
	"fmt"
	"math"
)

func E34() {
	max := 9

	// Getting the max
	max, power := 9, 1
	for sumOfDigitFactorials(max) > max {
		max += 9 * int(math.Pow10(power)) // 9 then 99 then 999
		power++
	}

	sum := 0
	for i := 3; i < max; i++ {
		if sumOfDigitFactorials(i) == i {
			sum += i
		}
	}
	fmt.Println(sum)
}

// Sum of the factorial of their digits
func sumOfDigitFactorials(number int) int {
	digits := []int{}
	for number > 0 {
		digits = append(digits, number%10)
		number = number / 10
	}

	sum := 0
	for _, digit := range digits {
		sum += Factorial(digit)
	}
	return sum
}
