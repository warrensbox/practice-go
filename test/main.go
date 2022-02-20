package main

import (
	"fmt"
	"math"
)

func main() {

	shares := []int{5, 11, 3, 50, 60, 90}
	fmt.Println(bestProfit(shares, 2))
}

func bestProfit(shares []int, k int) int {

	dp := make([][]int, k+1)

	for i := range dp {
		dp[i] = make([]int, len(shares))
	}

	for t := 1; t < len(dp); t++ {
		maxSoFar := math.MinInt16
		for d := 1; d < len(dp[0]); d++ {
			maxSoFar = Max(maxSoFar, dp[t-1][d-1]-shares[d-1])
			dp[t][d] = Max(dp[t][d-1], maxSoFar+shares[d])
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}

func Max(x, y int) int {

	if x > y {
		return x
	}

	return y
}

/*

   [5 11 3 50 60 90]
0	0  0  0 0  0  0
1	0  6  6 47 57 87
2	0  6  6 53 63 X

x:0 -5+6  = -1
x:1 -11+6 = -5
x:2 -3+6  = 3 (best)
x:3 -50+47= -3
x:4 -60+57= -3


maxSoFar = Max(maxSoFar, profit[t-1][d-1]+price[d-1])


*/
