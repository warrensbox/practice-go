package main

import "fmt"

func maxAreaOfIsland(grid [][]int) int {

	max := 0
	size := 0

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 1 {
				size = findArea(grid, r, c)
				fmt.Println("size", size)
			}
			max = Max(max, size)
		}
	}

	return max
}

func findArea(grid [][]int, r, c int) int {

	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] == 0 {
		return 0
	}
	size := 1
	if grid[r][c] == 1 {

		grid[r][c] = 0
		size += findArea(grid, r+1, c)
		size += findArea(grid, r-1, c)
		size += findArea(grid, r, c+1)
		size += findArea(grid, r, c-1)
	}

	return size
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
