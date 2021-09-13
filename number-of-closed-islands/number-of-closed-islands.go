package main

func closedIsland2(grid [][]int) int {

	num := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 0 {
				if dfs(grid, r, c) == 1 {
					num++
				}
			}
		}
	}
	return num
}

func dfs2(grid [][]int, r, c int) int {

	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
		return 0
	}

	if grid[r][c] == 1 {
		return 1
	}

	grid[r][c] = 1

	res := 1

	res = res & dfs(grid, r-1, c)
	res = res & dfs(grid, r+1, c)
	res = res & dfs(grid, r, c-1)
	res = res & dfs(grid, r, c+1)

	return res
}

func closedIsland(grid [][]int) int {

	num := 0
	for r := 1; r < len(grid)-1; r++ {
		for c := 1; c < len(grid[0])-1; c++ {
			if grid[r][c] == 0 {
				if dfs(grid, r, c) == 1 {
					num++
				}
			}
		}
	}
	return num
}

func dfs(grid [][]int, r, c int) int {

	if grid[r][c] == 1 {
		return 1
	}

	if r == 0 || r == len(grid)-1 || c == 0 || c == len(grid[0])-1 {
		return 0
	}

	grid[r][c] = 1

	res := 1

	res = res & dfs(grid, r-1, c)
	res = res & dfs(grid, r+1, c)
	res = res & dfs(grid, r, c-1)
	res = res & dfs(grid, r, c+1)

	return res
}
