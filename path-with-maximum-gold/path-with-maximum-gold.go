package main

import "fmt"

func main() {

	grid := [][]int{{0, 6, 0}, {5, 8, 7}, {0, 9, 0}}
	output := getMaximumGold(grid)
	fmt.Println(output)
}

func getMaximumGold(grid [][]int) int {

	max := 0

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] != 0 {
				// fmt.Println("------")
				// fmt.Println("val", grid[row][col])
				dfs(&grid, row, col, 0, &max)
				fmt.Println("CURENT TOTAL", max)
			}
		}
	}

	return max
}

func dfs(grid *[][]int, row, col, cur int, sum *int) {
	if row < 0 || col < 0 || row >= len(*grid) || col >= len((*grid)[row]) || (*grid)[row][col] == 0 {
		*sum = max(*sum, cur)
		return
	}
	tmp := (*grid)[row][col]
	(*grid)[row][col] = 0
	cur += tmp
	// fmt.Println("tmp", tmp)
	// fmt.Println("grid", *grid)
	// fmt.Println("row", row)
	// fmt.Println("col", col)
	dfs(grid, row+1, col, cur, sum)
	dfs(grid, row-1, col, cur, sum)
	dfs(grid, row, col+1, cur, sum)
	dfs(grid, row, col-1, cur, sum)
	(*grid)[row][col] = tmp
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
