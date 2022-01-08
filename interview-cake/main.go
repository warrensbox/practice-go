// To execute Go code, please declare a func main() in a package "main"
//TESTCODE
package main

import (
	"fmt"
)

func main() {

	fmt.Println(getMoneyAmount(2))
}

/*
goal: decode string according to num[char]
*/
func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = 100000
		}
		dp[i][i] = 0
	}

	fmt.Println(dp)
	fmt.Println("-----")
	for i := 1; i < n; i++ {
		for j := 1; j+i <= n; j++ {
			dp[j][j+i] = min(j+dp[j+1][j+i], dp[j][j+i-1]+j+i)
			fmt.Println(dp)
			for k := j + 1; k < j+i; k++ {
				dp[j][j+i] = min(dp[j][j+i],
					k+max(dp[j][k-1], dp[k+1][j+i]))
			}
		}
	}
	return dp[1][n]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
