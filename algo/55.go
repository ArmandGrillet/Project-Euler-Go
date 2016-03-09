/*
If we take 47, reverse and add, 47 + 74 = 121, which is palindromic.

Not all numbers produce palindromes so quickly. For example,

349 + 943 = 1292,
1292 + 2921 = 4213
4213 + 3124 = 7337

That is, 349 took three iterations to arrive at a palindrome.

Although no one has proved it yet, it is thought that some numbers, like 196, never produce a palindrome.
A number that never forms a palindrome through the reverse and add process is called a Lychrel number.
Due to the theoretical nature of these numbers, and for the purpose of this problem, we shall assume that a number is Lychrel until proven otherwise.
In addition you are given that for every number below ten-thousand, it will either (i) become a palindrome in less than fifty iterations, or, (ii) no one, with all the computing power that exists, has managed so far to map it to a palindrome.
In fact, 10677 is the first number to be shown to require over fifty iterations before producing a palindrome: 4668731596684224866951378664 (53 iterations, 28-digits).

Surprisingly, there are palindromic numbers that are themselves Lychrel numbers; the first example is 4994.

How many Lychrel numbers are there below ten-thousand?

NOTE: Wording was modified slightly on 24 April 2007 to emphasise the theoretical nature of Lychrel numbers.
*/

package algo

import (
	"fmt"
	"math/big"
	"strings"
)

func E55() {
	max := int64(10000)

	sum := 0
	for i := int64(10); i < max; i++ {
		if isLychrel(i) {
			sum++
		}
	}
	fmt.Println(sum)
}

func isLychrel(x int64) bool {
	y := big.NewInt(x)
	reverse := big.NewInt(int64(0))

	for iteration := 0; iteration < 50; iteration++ {
		reverse.SetString(reverseBig(y), 0)
		y.Add(y, reverse)

		if y.String() == Reverse(y.String()) {
			return false
		}
	}
	return true
}

func reverseBig(s *big.Int) string {
	r := []rune(s.String())
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	// 0123 => 123
	return strings.TrimLeft(string(r), "0")
}
