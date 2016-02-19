/*
We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once; for example, the 5-digit number, 15234, is 1 through 5 pandigital.

The product 7254 is unusual, as the identity, 39 × 186 = 7254, containing multiplicand, multiplier, and product is 1 through 9 pandigital.

Find the sum of all products whose multiplicand/multiplier/product identity can be written as a 1 through 9 pandigital.

HINT: Some products can be obtained in more than one way so be sure to only include it once in your sum.
*/

package algo

import (
	"fmt"
	"math"
	"sort"
)

func E32() {
	min := 1234
	max := 9876 // X * y = 9876 is the maximum possibility to get as a 1 through 9 pandigital

	var pandigitals []int
	var divisors []int
	var sum int

	for i := max; i > min; i-- {
		divisors = getDivisors(i)
		for j := 0; j < len(divisors); j++ {
			if isPandigital(i, divisors[j]) {
				fmt.Printf("%d * %d = %d\n", divisors[j], i/divisors[j], i)
				sum += i
				break
				pandigitals = append(pandigitals, i)
			}
		}
	}
	fmt.Println(sum)
}

func equalSlices(a, b []int) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func numberDigits(number int) []int {
	digits := []int{}
	for number > 0 {
		digits = append(digits, number%10)
		number /= 10
	}
	return digits
}

func isPandigital(product, multiplier int) bool {
	var digits []int
	digits = append(digits, numberDigits(product)...)
	digits = append(digits, numberDigits(multiplier)...)
	digits = append(digits, numberDigits(product/multiplier)...)
	sort.Ints(digits)

	if equalSlices(digits, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		return true
	} else {
		return false
	}
}

func getDivisors(number int) []int {
	var divisors []int
	for i := 1; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}
