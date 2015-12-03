/*
The fraction 49/98 is a curious fraction, as an inexperienced mathematician in attempting to simplify it may incorrectly believe that 49/98 = 4/8, which is correct, is obtained by cancelling the 9s.

We shall consider fractions like, 30/50 = 3/5, to be trivial examples.

There are exactly four non-trivial examples of this type of fraction, less than one in value, and containing two digits in the numerator and denominator.

If the product of these four fractions is given in its lowest common terms, find the value of the denominator.
*/

package main

import "fmt"

const (
	max = 100
)

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getCommonDigit(i, j int) int {
	digitsI := []int{}
	for i > 0 {
		digitsI = append(digitsI, i%10)
		i /= 10
	}

	for j > 0 {
		if contains(digitsI, j%10) {
			return j % 10
		}
		j /= 10
	}

	return -1
}

func removeDigitFromTwoDigitsNumber(twoDigitsNumber, digit int) int {
	if twoDigitsNumber%10 == digit {
		return twoDigitsNumber / 10
	} else {
		return twoDigitsNumber % 10
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var commonDigit, newI, newJ int
	nominatorProduct, denominatorProduct := 1, 1
	for i := 10; i < max; i++ {
		for j := i + 1; j < max; j++ {
			commonDigit = getCommonDigit(i, j)
			if commonDigit > 0 { // Non-trivial
				newI, newJ = removeDigitFromTwoDigitsNumber(i, commonDigit), removeDigitFromTwoDigitsNumber(j, commonDigit)
				if newI != 0 && newJ != 0 {
					if float64(i)/float64(j) == float64(newI)/float64(newJ) {
						fmt.Printf("%d / %d = %d / %d\n", i, j, newI, newJ)
						nominatorProduct *= newI
						denominatorProduct *= newJ
					}
				}
			}
		}
	}

	gcd := gcd(nominatorProduct, denominatorProduct)
	fmt.Printf("%d / %d\n", nominatorProduct/gcd, denominatorProduct/gcd)
}
