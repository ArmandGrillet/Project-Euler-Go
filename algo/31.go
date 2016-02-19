/*
In England the currency is made up of pound, £, and pence, p, and there are eight coins in general circulation:

1p, 2p, 5p, 10p, 20p, 50p, £1 (100p) and £2 (200p).
It is possible to make £2 in the following way:

1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p
How many different ways can £2 be made using any number of coins?
*/

package algo

import "fmt"

func E31() {
	max := 200
	pences := []int{1, 2, 5, 10, 20, 50, 100, 200}

	var possibilities = make([]int, max+1)
	possibilities[0] = 1 // If there is 0p to pay, you only have one possibility: leave.
	for i := 0; i < len(pences); i++ {
		for j := pences[i]; j < len(possibilities); j++ {
			possibilities[j] += possibilities[j-pences[i]]
		}
	}

	fmt.Println(possibilities[max]) // Number of possibilities to pay 200p
}
