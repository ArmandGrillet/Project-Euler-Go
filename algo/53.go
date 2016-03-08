/*
There are exactly ten ways of selecting three from five, 12345:

123, 124, 125, 134, 135, 145, 234, 235, 245, and 345

In combinatorics, we use the notation, 5C3 = 10.

In general,

nCr =
n!
_______
r!(n−r)!
,where r ≤ n, n! = n×(n−1)×...×3×2×1, and 0! = 1.
It is not until n = 23, that a value exceeds one-million: 23C10 = 1144066.

How many, not necessarily distinct, values of  nCr, for 1 ≤ n ≤ 100, are greater than one-million?
*/

package algo

import (
	"fmt"
	"math/big"
)

func E53() {
	max := int64(1000000)
	greaterThanMax := 0

	for n := 1; n <= 100; n++ {
		for r := 0; r <= n; r++ {
			coef := big.NewInt(int64(0)).Binomial(int64(n), int64(r))
			if coef.Cmp(big.NewInt(max)) >= 0 {
				greaterThanMax++
			}
		}
	}
	fmt.Println(greaterThanMax)
}
