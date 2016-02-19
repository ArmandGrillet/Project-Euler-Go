/*
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a**2 + b**2 = c**2
For example, 32 + 42 = 9 + 16 = 25 = 52.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/

package algo

import (
	"fmt"
	"math"
	"os"
)

func E9() {
	sum := 1000

	a := 0
	var b, c int
	var tempC float64
	for a+(a+1)+(a+2) < sum { // Pythagorean triplet.
		a++
		b = a + 1
		for a+b+(b+1) < sum {
			tempC = math.Sqrt(float64(a*a + b*b))
			if tempC == float64(int(tempC)) { // The square of c is a natural number.
				c = int(tempC)
				if a+b+c == 1000 {
					fmt.Println(a * b * c)
					os.Exit(0)
				}
			}
			b++
		}
	}
}
