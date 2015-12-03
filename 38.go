/*
Take the number 192 and multiply it by each of 1, 2, and 3:

192 × 1 = 192
192 × 2 = 384
192 × 3 = 576
By concatenating each product we get the 1 to 9 pandigital, 192384576. We will call 192384576 the concatenated product of 192 and (1,2,3)

The same can be achieved by starting with 9 and multiplying by 1, 2, 3, 4, and 5, giving the pandigital, 918273645, which is the concatenated product of 9 and (1,2,3,4,5).

What is the largest 1 to 9 pandigital 9-digit number that can be formed as the concatenated product of an integer with (1,2, ... , n) where n > 1?
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func digits(number int) []int {
	digits := []int{}
	for number > 0 {
		digits = append(digits, number%10)
		number /= 10
	}
	return digits
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func isPandigital(number int) bool {
	digits := digits(number)
	if len(digits) == 9 {
		for i := 1; i < 10; i++ {
			if !contains(digits, i) {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func concatenMultiplication(number int, multiplication int) int {
	digits := []string{}
	for i := 1; i <= multiplication; i++ {
		digits = append(digits, strconv.Itoa(number*i))
	}

	concatenation, _ := strconv.Atoi(strings.Join(digits, ""))
	return concatenation
}

func main() {
	max := 9999 // *2 = 999 919 998
	biggestPandigital := 0
	var numberMultiplied int
	for i := max; i > 0; i-- {
		numberMultiplied = i
		for j := 1; len(digits(numberMultiplied)) <= 9; j++ {
			numberMultiplied = concatenMultiplication(i, j)
			if isPandigital(numberMultiplied) && numberMultiplied > biggestPandigital {
				fmt.Println(numberMultiplied)
				biggestPandigital = numberMultiplied
			}
		}
	}
}
