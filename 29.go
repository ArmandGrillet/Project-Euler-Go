/*
Consider all integer combinations of ab for 2 ≤ a ≤ 5 and 2 ≤ b ≤ 5:

22=4, 23=8, 24=16, 25=32
32=9, 33=27, 34=81, 35=243
42=16, 43=64, 44=256, 45=1024
52=25, 53=125, 54=625, 55=3125
If they are then placed in numerical order, with any repeats removed, we get the following sequence of 15 distinct terms:

4, 8, 9, 16, 25, 27, 32, 64, 81, 125, 243, 256, 625, 1024, 3125

How many distinct terms are in the sequence generated by ab for 2 ≤ a ≤ 100 and 2 ≤ b ≤ 100?
*/

package main

import (
	"fmt"
	"math/big"
)

const (
	min = 2
	max = 100
)

func contains(s []*big.Int, e *big.Int) bool {
	for _, a := range s {
		if a.Cmp(e) == 0 {
			return true
		}
	}
	return false
}

func main() {
	var distinctTerms []*big.Int
	var combination *big.Int
	for a := int64(min); a <= max; a++ {
		for b := int64(min); b <= max; b++ {
			combination = new(big.Int).Exp(big.NewInt(a), big.NewInt(b), nil)
			if !contains(distinctTerms, combination) {
				distinctTerms = append(distinctTerms, combination)
			}
		}
	}
	fmt.Println(len(distinctTerms))
}