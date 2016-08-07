/*
The cube, 41063625 (345**3), can be permuted to produce two other cubes: 56623104 (384**3) and 66430125 (405**3).
In fact, 41063625 is the smallest cube which has exactly three permutations of its digits which are also cube.

Find the smallest cube for which exactly five permutations of its digits are cube.
*/

package algo

import (
	"fmt"
	"math"
)

// On créer un array de array de int avec le nombre de digits entre 0 et 9 pour chaque.

func E62() {
	cubeDigits := [][10]int{}
	i := 1.0
	for {
		cube := int(math.Pow(i, 3.0))
		var digits [10]int
		for cube >= 1 {
			digits[cube%10]++
			cube /= 10
		}

		if iterationsOf(cubeDigits, digits) == 4 {
			fmt.Println(int64(math.Pow(float64(firstIteration(cubeDigits, digits)+1), 3.0)))
			break
		}
		cubeDigits = append(cubeDigits, digits)
		i++
	}
}

func iterationsOf(s [][10]int, e [10]int) int {
	iterations := 0
	for i := 0; i < len(s); i++ {
		if s[i] == e {
			iterations++
		}
	}
	return iterations
}

func firstIteration(s [][10]int, e [10]int) int {
	for i := 0; i < len(s); i++ {
		if s[i] == e {
			return i
		}
	}
	return 0
}
