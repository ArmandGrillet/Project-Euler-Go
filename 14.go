/*
The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:

13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.
*/

package main

import "fmt"

const (
	max = 1000000
)

func chainLength(number int) int {
	if number == 1 {
		return 1
	} else if number%2 == 0 { // Even
		return 1 + chainLength(number/2)
	} else { // Odd
		return 1 + chainLength(3*number+1)
	}
}

func main() {
	var biggestChainNumber, biggestChainLength, currentChainLength int
	for i := 1; i < max; i++ {
		currentChainLength = chainLength(i)
		if currentChainLength > biggestChainLength {
			biggestChainLength = currentChainLength
			biggestChainNumber = i
		}
	}
	fmt.Println(biggestChainNumber)
}
