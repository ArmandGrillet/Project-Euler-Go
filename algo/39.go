/*
If p is the perimeter of a right angle triangle with integral length sides, {a,b,c}, there are exactly three solutions for p = 120.

{20,48,52}, {24,45,51}, {30,40,50}

For which value of p ≤ 1000, is the number of solutions maximised?
*/

package algo

import (
	"fmt"
	"math"
)

func E39() {
	p := [1001]int{}
	a, b := 1, 1
	for perimeter(a, a) <= 1000. {
		for perimeter(a, b) <= 1000. {
			if math.Mod(perimeter(a, b), 1.) == 0. {
				fmt.Printf("%d + %d + %d = %d\n", a, b, int(math.Sqrt(float64(a*a+b*b))), int(perimeter(a, b)))
				p[int(perimeter(a, b))]++
			}
			b++
		}
		a++
		b = a
	}

	var max, maxPosition int
	for i := 0; i < len(p); i++ {
		if p[i] > max {
			max = p[i]
			maxPosition = i
		}
	}
	fmt.Println(maxPosition)
}

func perimeter(a int, b int) float64 {
	// Finding c
	c := math.Sqrt(float64(a*a + b*b))
	return float64(a+b) + c
}
