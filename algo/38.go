/*
Take the number 192 and multiply it by each of 1, 2, and 3:

192 × 1 = 192
192 × 2 = 384
192 × 3 = 576
By concatenating each product we get the 1 to 9 pandigital, 192384576. We will call 192384576 the concatenated product of 192 and (1,2,3)

The same can be achieved by starting with 9 and multiplying by 1, 2, 3, 4, and 5, giving the pandigital, 918273645, which is the concatenated product of 9 and (1,2,3,4,5).

What is the largest 1 to 9 pandigital 9-digit number that can be formed as the concatenated product of an integer with (1,2, ... , n) where n > 1?
*/

package algo

import (
	"fmt"
	"strconv"
	"strings"
)

func E38() {
	max := 9999 // *2 = 999 919 998

	biggestPandigital := 0
	var numberMultiplied int
	for i := max; i > 0; i-- {
		numberMultiplied = i
		for j := 1; len(digits(numberMultiplied)) <= 9; j++ {
			numberMultiplied = concatenMultiplication(i, j)
			if IsPandigital(numberMultiplied) && numberMultiplied > biggestPandigital {
				fmt.Println(numberMultiplied)
				biggestPandigital = numberMultiplied
			}
		}
	}
}

func digits(number int) []int {
	digits := []int{}
	for number > 0 {
		digits = append(digits, number%10)
		number /= 10
	}
	return digits
}

func concatenMultiplication(number int, multiplication int) int {
	digits := []string{}
	for i := 1; i <= multiplication; i++ {
		digits = append(digits, strconv.Itoa(number*i))
	}

	concatenation, _ := strconv.Atoi(strings.Join(digits, ""))
	return concatenation
}
