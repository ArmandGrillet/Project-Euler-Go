// Find the sum of all the primes below two million.
package main

import (
	"fmt"
	"math"
)

func main() {
	const max = 2000000
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

func sumPrimes(start int, end int, ch chan<- int) {
	sum := 0
	for i := start; i < end; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	ch <- sum
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	s := int(math.Sqrt(float64(n)))
	for i := 3; i <= s; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
