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

func main() {
	const (
		num = int64(2)
		pow = int64(1000)
	)

	res := new(big.Int).Exp(big.NewInt(num), big.NewInt(pow), nil)
	sum := 0
	for _, digit := range res.String() {
		intDigit, err := strconv.Atoi(string(digit))
		if err != nil {
			panic(err)
		}
		sum += intDigit
	}
	fmt.Println(sum)
}
