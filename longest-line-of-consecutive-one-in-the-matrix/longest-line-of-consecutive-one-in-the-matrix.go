package main

import "fmt"

func longestLine(mat [][]int) int {

	max := 0

	if len(mat) == 0 {
		return max
	}

	dp := make([][][]int, len(mat))

	//initialize (3d array)
	for r := 0; r < len(mat); r++ {
		dp[r] = make([][]int, len(mat[0]))
		for c := 0; c < len(mat[r]); c++ {
			dp[r][c] = make([]int, 4)
		}
	}

	for r := 0; r < len(mat); r++ {

		for c := 0; c < len(mat[r]); c++ {
			if mat[r][c] == 1 {
				for i := 0; i < 4; i++ {
					dp[r][c][i] = 1 //initialize
				}
				if c > 0 {
					dp[r][c][0] += dp[r][c-1][0] //top
					if r > 0 {
						dp[r][c][1] += dp[r-1][c-1][1] //diaganol
					}
				}
				if r > 0 {
					dp[r][c][2] += dp[r-1][c][2]
					if c < len(mat[r])-1 {
						dp[r][c][3] += dp[r-1][c+1][3]
					}
				}
				fmt.Println(dp[r][c])
				max = Max(max, Max(dp[r][c][0], Max(dp[r][c][1], Max(dp[r][c][2], dp[r][c][3]))))
			}
		}
	}

	// fmt.Println(dp)
	fmt.Println(max)
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
