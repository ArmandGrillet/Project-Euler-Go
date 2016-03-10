/*
It is possible to show that the square root of two can be expressed as an infinite continued fraction.

√ 2 = 1 + 1/(2 + 1/(2 + 1/(2 + ... ))) = 1.414213...

By expanding this for the first four iterations, we get:

1 + 1/2 = 3/2 = 1.5
1 + 1/(2 + 1/2) = 7/5 = 1.4
1 + 1/(2 + 1/(2 + 1/2)) = 17/12 = 1.41666...
1 + 1/(2 + 1/(2 + 1/(2 + 1/2))) = 41/29 = 1.41379...

The next three expansions are 99/70, 239/169, and 577/408, but the eighth expansion, 1393/985, is the first example where the number of digits in the numerator exceeds the number of digits in the denominator.

In the first one-thousand expansions, how many fractions contain a numerator with more digits than denominator?
*/

package algo

import "fmt"

func E57() {
	max := 1000

	sum := 0
	iterations := 2

	previousNumerator, previousDenominator := BigFromInt(3), BigFromInt(2)
	tempNumerator, tempDenominator := BigFromInt(0), BigFromInt(0)
	numerator, denominator := BigFromInt(7), BigFromInt(5)

	for iterations != max {
		tempNumerator.Set(numerator)
		numerator.Add(numerator, numerator)
		numerator.Add(numerator, previousNumerator)
		previousNumerator.Set(tempNumerator)

		tempDenominator.Set(denominator)
		denominator.Add(denominator, denominator)
		denominator.Add(denominator, previousDenominator)
		previousDenominator.Set(tempDenominator)

		if len(numerator.String()) > len(denominator.String()) {
			sum++
		}

		iterations++
	}

	fmt.Println(sum)
}
