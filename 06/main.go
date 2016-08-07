// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
package main

import "fmt"

func main() {
	const max = 100
	var sumSquares, sum int

	for i := 0; i <= max; i++ {
		sumSquares += i * i
		sum += i
	}

	fmt.Println(sumSquares - sum*sum)
}
