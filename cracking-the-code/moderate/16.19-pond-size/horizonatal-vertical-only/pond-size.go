package main

import "fmt"

func main() {

	island := [][]int{
		{0, 1, 1, 0},
		{0, 1, 0, 1},
		{1, 1, 0, 1},
		{0, 1, 0, 1},
	}

	// fmt.Println(island)

	fmt.Println(getSize(island))
}

func getSize(grid [][]int) []int {

	//base case
	if len(grid) == 0 {
		return []int{}
	}

	pondSizes := []int{}
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			var size int
			if grid[r][c] == 0 {
				size = dfs(grid, r, c)
			}
			if size > 0 {
				pondSizes = append(pondSizes, size)
			}

		}
	}

	return pondSizes
}

func dfs(grid [][]int, row, column int) int {

	if row < 0 || row >= len(grid) || column < 0 || column >= len(grid[0]) || grid[row][column] == 1 {
		return 0
	}

	fmt.Println(grid)

	count := 1
	grid[row][column] = 1
	count += dfs(grid, row-1, column)
	count += dfs(grid, row+1, column)
	count += dfs(grid, row, column-1)
	count += dfs(grid, row, column+1)
	return count
	//return dfs(grid, i-1, j, count) + dfs(grid, i+1, j, count) + dfs(grid, i, j-1, count) + dfs(grid, i, j+1, count)
}
