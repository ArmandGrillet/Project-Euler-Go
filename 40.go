/*
An irrational decimal fraction is created by concatenating the positive integers:

0.123456789101112131415161718192021...

It can be seen that the 12th digit of the fractional part is 1.

If dn represents the nth digit of the fractional part, find the value of the following expression.

d1 × d10 × d100 × d1000 × d10000 × d100000 × d1000000
*/

package main

import (
	"fmt"
)

func digits(number int) []int {
	digits := []int{}
	for number > 0 {
		digits = append(digits, number%10)
		number = number / 10
	}
	return digits
}

func main() {
	i, position, oldPosition := 0, 0, 0
	product := 1
	wantedDigits, wantedDigitIndex := [7]int{1, 10, 100, 1000, 10000, 100000, 1000000}, 0
	for wantedDigitIndex < len(wantedDigits) {
		i++
		position += len(digits(i))
		if position >= wantedDigits[wantedDigitIndex] && oldPosition < wantedDigits[wantedDigitIndex] {
			product *= digits(i)[position-wantedDigits[wantedDigitIndex]]
			wantedDigitIndex++
		}
		oldPosition++
	}
	fmt.Println(product)
}
