// Find the sum of all the primes below two million.

package main

import (
	"fmt"
	"math/big"
)

const (
	max = 2000000
)

func sumPrimes(start int, end int, ch chan<- int) {
	sum := 0
	for i := start; i < end; i++ {
		if big.NewInt(int64(i)).ProbablyPrime(20) {
			sum += i
		}
	}
	ch <- sum
}

func main() {
	goroutines := 80 // We'll use 80 goroutines to do the job.
	cycle := int(max / goroutines)
	sum := 0

	ch := make(chan int)
	for i := 0; i < max; i += cycle {
		if i+cycle < max {
			go sumPrimes(i, i+cycle, ch)
		} else {
			go sumPrimes(i, max, ch)
		}

	}

	for i := 0; i < goroutines; i++ {
		sum += <-ch
	}

	fmt.Println(sum)
}
