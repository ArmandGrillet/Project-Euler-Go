/*
Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, there are exactly 6 routes to the bottom right corner.

How many such routes are there through a 20×20 grid?
*/
package main

import "fmt"

const (
	gridSize = 20
)

func main() {
	// Creation of an empty grid
	grid := make([][]int, gridSize+1)
	for i := range grid {
		grid[i] = make([]int, gridSize+1)
	}

	// Calculating the number of routes in the grid
	for i := len(grid) - 1; i >= 0; i-- {
		for j := len(grid[i]) - 1; j >= 0; j-- {
			if (i == len(grid)-1) || (j == len(grid)-1) {
				grid[i][j] = 1 // Point at the limit of the grid, we initialize it at 1
			} else {
				grid[i][j] = grid[i+1][j] + grid[i][j+1]
			}
		}
	}

	fmt.Println(grid[0][0])
}
