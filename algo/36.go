/*
The decimal number, 585 = 10010010012 (binary), is palindromic in both bases.

Find the sum of all numbers, less than one million, which are palindromic in base 10 and base 2.

(Please note that the palindromic number, in either base, may not include leading zeros.)
*/

package algo

import (
	"fmt"
	"strconv"
)

func E36() {
	max := int64(1000000)
	sum := int64(0)

	for i := int64(0); i < max; i++ {
		if strconv.FormatInt(i, 10) == reverse(strconv.FormatInt(i, 10)) && strconv.FormatInt(i, 2) == reverse(strconv.FormatInt(i, 2)) {
			sum += i
		}
	}
	fmt.Println(sum)
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
