/*
2**15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2**1000?
*/

package main

import (
	"fmt"
	"math/big"
	"strconv"
)

const (
	number = 2
	power  = 1000
)

func main() {
	number := new(big.Int).Exp(big.NewInt(number), big.NewInt(power), nil)

	sum := 0
	for _, digit := range number.String() {
		digitAsInt, err := strconv.Atoi(string(digit))
		if err != nil {
			panic(err)
		}
		sum += digitAsInt
	}
	fmt.Println(sum)
}
