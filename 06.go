// Find the difference between the sum of the squares of the first 100 natural numbers and the square of the sum.

package main

import "fmt"

const (
	naturalNumbers = 100
)

func main() {
	sum, sumOfSquares := 0, 0
	for i := 0; i <= naturalNumbers; i++ {
		sumOfSquares += i * i
		sum += i
	}
	fmt.Println((sum * sum) - sumOfSquares)
}
