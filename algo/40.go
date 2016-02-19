/*
An irrational decimal fraction is created by concatenating the positive integers:

0.123456789101112131415161718192021...

It can be seen that the 12th digit of the fractional part is 1.

If dn represents the nth digit of the fractional part, find the value of the following expression.

d1 × d10 × d100 × d1000 × d10000 × d100000 × d1000000
*/

package algo

import (
	"fmt"
)

func E40() {
	i, position, oldPosition := 0, 0, 0
	product := 1
	wantedDigits, wantedDigitIndex := [7]int{1, 10, 100, 1000, 10000, 100000, 1000000}, 0
	for wantedDigitIndex < len(wantedDigits) {
		i++
		position += len(Digits(i))
		if position >= wantedDigits[wantedDigitIndex] && oldPosition < wantedDigits[wantedDigitIndex] {
			product *= Digits(i)[position-wantedDigits[wantedDigitIndex]]
			wantedDigitIndex++
		}
		oldPosition++
	}
	fmt.Println(product)
}
