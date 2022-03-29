package main

func longestIncreasingPath(matrix [][]int) int {

	dp := make([][]int, len(matrix))

	for i, _ := range matrix {
		dp[i] = make([]int, len(matrix[0]))
	}
	var dfs func(r, c int, prev int, matrix [][]int) int

	dfs = func(r, c int, prev int, matrix [][]int) int {

		if r < 0 || c < 0 || c >= len(matrix[0]) || r >= len(matrix) || matrix[r][c] <= prev {
			return 0
		}

		if dp[r][c] > 0 {
			return dp[r][c]
		}

		res := 1

		res = Max(res, 1+dfs(r+1, c, matrix[r][c], matrix))
		res = Max(res, 1+dfs(r-1, c, matrix[r][c], matrix))
		res = Max(res, 1+dfs(r, c+1, matrix[r][c], matrix))
		res = Max(res, 1+dfs(r, c-1, matrix[r][c], matrix))

		dp[r][c] = res
		return res
	}

	max := 0
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[0]); c++ {
			max = Max(max, dfs(r, c, -1, matrix))
		}
	}

	return max
}

func Max(x, y int) int {

	if x > y {
		return x
	}

	return y
}
