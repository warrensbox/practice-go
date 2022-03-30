package main

func minPathSum(grid [][]int) int {

	//create dp
	dp := make([][]int, len(grid))
	//initialize dp
	for i, _ := range grid {
		dp[i] = make([]int, len(grid[0]))
	}

	for r := len(grid) - 1; r >= 0; r-- {
		for c := len(grid[0]) - 1; c >= 0; c-- {

			if r == len(grid)-1 && c != len(grid[0])-1 {

				dp[r][c] = dp[r][c+1] + grid[r][c]

			} else if c == len(grid[0])-1 && r != len(grid)-1 {

				dp[r][c] = dp[r+1][c] + grid[r][c]

			} else if r != len(grid)-1 && c != len(grid[0])-1 {

				dp[r][c] = Min(dp[r+1][c], dp[r][c+1]) + grid[r][c]

			} else {
				dp[r][c] = grid[r][c]
			}
		}
	}

	return dp[0][0]
}

func Min(x, y int) int {

	if x < y {
		return x
	}

	return y
}
