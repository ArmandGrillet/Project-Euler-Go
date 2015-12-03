/*
A unit fraction contains 1 in the numerator. The decimal representation of the unit fractions with denominators 2 to 10 are given:

1/2	= 	0.5
1/3	= 	0.(3)
1/4	= 	0.25
1/5	= 	0.2
1/6	= 	0.1(6)
1/7	= 	0.(142857)
1/8	= 	0.125
1/9	= 	0.(1)
1/10	= 	0.1
Where 0.1(6) means 0.166666..., and has a 1-digit recurring cycle. It can be seen that 1/7 has a 6-digit recurring cycle.

Find the value of d < 1000 for which 1/d contains the longest recurring cycle in its decimal fraction part.
*/

package main

import "fmt"

const (
	min = 2
	max = 999
)

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	remainders := []int{}
	cycleLength := 0
	var container, remainder int
	for i := max; i >= min && cycleLength < i; i-- {
		remainders = []int{}
		remainder = 1 // 1 % i is always equal to 1
		for remainder != 0 && !contains(remainders, remainder) {
			remainders = append(remainders, remainder)
			remainder = (remainder * 10) % i
		}
		if len(remainders) > cycleLength {
			cycleLength = len(remainders)
			container = i
		}
	}
	fmt.Println(container)
}
