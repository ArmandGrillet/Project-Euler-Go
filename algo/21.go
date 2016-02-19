/*
Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.
*/

package algo

import (
	"fmt"
	"math"
)

func E21() {
	max := 10000

	var sum, divisorsSum int
	for i := 0; i < max; i++ {
		divisorsSum = getDivisorsSum(i)
		if divisorsSum < i && getDivisorsSum(divisorsSum) == i {
			sum += i + divisorsSum
		}
	}
	fmt.Println(sum)
}

func getDivisorsSum(number int) (divisorsSum int) {
	for i := 1; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			divisorsSum += i
			if i != number/i && i != 1 { // i != 1 because number / i would be number
				divisorsSum += (number / i)
			}
		}
	}
	return
}
