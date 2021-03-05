package main

import "fmt"

func main() {

	island := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}

	fmt.Println(getSize(island))
}

func getSize(grid [][]int) int {

	//base case
	if len(grid) == 0 {
		return 0
	}

	numberOfIsland := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				numberOfIsland += dfs(grid, i, j)
			}

		}
	}

	return numberOfIsland
}

func dfs(grid [][]int, i, j int) int {

	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 {
		return 0
	}

	grid[i][j] = 0
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
	return 1
}
