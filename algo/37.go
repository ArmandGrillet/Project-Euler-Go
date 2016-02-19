/*
The number 3797 has an interesting property. Being prime itself, it is possible to continuously remove digits from left to right, and remain prime at each stage: 3797, 797, 97, and 7. Similarly we can work from right to left: 3797, 379, 37, and 3.

Find the sum of the only eleven primes that are both truncatable from left to right and right to left.

NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.
*/

package algo

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

func E37() {
	primes := []int{2, 3, 5, 7}
	truncatablePrimes, reducedNumbers := []int{}, []int{}
	number := 10
	var isTruncatablePrime bool

	for len(truncatablePrimes) < 11 {
		if big.NewInt(int64(number)).ProbablyPrime(20) {
			primes = append(primes, number)
		}

		isTruncatablePrime = true
		reducedNumbers = numbersFromNumber(number)
		for i := 0; i < len(reducedNumbers) && isTruncatablePrime; i++ {
			if !Contains(primes, reducedNumbers[i]) {
				isTruncatablePrime = false
			}
		}
		if isTruncatablePrime {
			fmt.Println(number)
			truncatablePrimes = append(truncatablePrimes, number)
		}

		number++
	}

	sum := 0
	for i := 0; i < len(truncatablePrimes); i++ {
		sum += truncatablePrimes[i]
	}
	fmt.Printf("%d\n", sum)
}

func numbersFromNumber(number int) []int {
	numbers := []int{number}
	numberLength := len(strconv.FormatInt(int64(number), 10))
	numberForDigits, numberLeftToRight, numberRightToLeft := number, number, number

	// Getting each digit
	digits := []int{}
	for numberForDigits > 0 {
		digits = append(digits, numberForDigits%10)
		numberForDigits = numberForDigits / 10
	}

	for i := numberLength - 1; i >= 1; i-- {
		numberLeftToRight = numberLeftToRight - digits[i]*int(math.Pow10(i))
		numbers = append(numbers, numberLeftToRight)
	}

	for i := 0; i < numberLength-1; i++ {
		numberRightToLeft = (numberRightToLeft - digits[i]) / 10
		numbers = append(numbers, numberRightToLeft)
	}

	return numbers
}
