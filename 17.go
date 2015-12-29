/*
If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?


NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20 letters. The use of "and" when writing out numbers is in compliance with British usage.
*/

package main

import (
	"fmt"
	"unicode/utf8"
)

const (
	start = 1
	end   = 1000
)

var (
	zeroToNineteen = [20]string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	tens           = [10]string{"", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
)

func countTo1000(num int) int {
	if num < 0 || num > 1000 {
		fmt.Println("Error: number is not between 0 and 1000")
		return 0
	} else {
		counter := 0
		switch {
		case num == 1000:
			return utf8.RuneCountInString("onethousand")
		case num > 99:
			counter = utf8.RuneCountInString(zeroToNineteen[int(num/100)])
			counter += utf8.RuneCountInString("hundred")
			updatedNum := int(num % 100) // updated num = 0 up to 99
			if updatedNum > 0 {
				counter += utf8.RuneCountInString("and")
				counter += countTo1000(updatedNum)
			}
		case num >= 20:
			counter = utf8.RuneCountInString(tens[int(num/10)])
			updatedNum := int(num % 10) // updated num = 0 up to 9
			if updatedNum > 0 {
				counter += countTo1000(updatedNum)
			}
		case num >= 0:
			counter = utf8.RuneCountInString(zeroToNineteen[num])
		default:
			fmt.Println("Error: number is not between 0 and 1000")
		}
		return counter
	}
}

func main() {
	sum := 0
	for i := start; i <= end; i++ {
		sum += countTo1000(i)
	}
	fmt.Println(sum)
}
