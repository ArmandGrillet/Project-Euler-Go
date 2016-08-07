/*
Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, there are exactly 6 routes to the bottom right corner.

How many such routes are there through a 20×20 grid?
*/
package main

import "fmt"

func main() {
	const gridSize = 20

	// Creation of an empty grid
	grid := [gridSize + 1][gridSize + 1]int{}
	for i := 0; i <= gridSize; i++ { // Set the edges to 1 because there is no move left.
		grid[gridSize][i] = 1
		grid[i][gridSize] = 1
	}

	// Calculating the number of routes in the grid
	for i := gridSize - 1; i >= 0; i-- {
		for j := gridSize - 1; j >= 0; j-- {
			grid[i][j] = grid[i+1][j] + grid[i][j+1]
		}
	}

	fmt.Println(grid[0][0])
}
