// Find the maximum total from top to bottom of the triangle below.

package main

import "fmt"

var triangle = [][]int{
	[]int{75},
	[]int{95, 64},
	[]int{17, 47, 82},
	[]int{18, 35, 87, 10},
	[]int{20, 4, 82, 47, 65},
	[]int{19, 1, 23, 75, 03, 34},
	[]int{88, 2, 77, 73, 07, 63, 67},
	[]int{99, 65, 04, 28, 06, 16, 70, 92},
	[]int{41, 41, 26, 56, 83, 40, 80, 70, 33},
	[]int{41, 48, 72, 33, 47, 32, 37, 16, 94, 29},
	[]int{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14},
	[]int{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57},
	[]int{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48},
	[]int{63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31},
	[]int{04, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 04, 23}}

func main() {
	// We're adding step by step the biggest number between two numbers.
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if triangle[i+1][j] > triangle[i+1][j+1] {
				triangle[i][j] += triangle[i+1][j]
			} else {
				triangle[i][j] += triangle[i+1][j+1]
			}
		}
	}
	fmt.Println(triangle[0][0])
}
