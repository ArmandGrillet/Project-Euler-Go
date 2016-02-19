/*
It can be seen that the number, 125874, and its double, 251748, contain exactly the same digits, but in a different order.

Find the smallest positive integer, x, such that 2x, 3x, 4x, 5x, and 6x, contain the same digits.
*/

package algo

import (
	"fmt"
	"sort"
)

func E52() {
	i := 10
	found := false
	for !found {
		if length(i) == length(i*6) && containSameDigits(i, 2) {
			fmt.Println(i)
			found = true
		}
		i++
	}
}

func sumOfDigits(number int) int {
	sum := 0
	for number > 0 {
		sum += number % 10
		number = number / 10
	}
	return sum
}

func orderedDigitsOf(number int) []int {
	digits := []int{}
	for number > 0 {
		digits = append(digits, number%10)
		number = number / 10
	}
	sort.Sort(sort.IntSlice(digits))
	return digits
}

func containSameDigits(x, order int) bool {
	if sumOfDigits(x) == sumOfDigits(order*x) && EqualSlices(orderedDigitsOf(x), orderedDigitsOf(order*x)) {
		if order == 6 {
			return true
		} else {
			return containSameDigits(x, order+1)
		}
	}
	return false
}

func length(number int) int {
	length := 0
	for number > 0 {
		length++
		number = number / 10
	}
	return length
}
