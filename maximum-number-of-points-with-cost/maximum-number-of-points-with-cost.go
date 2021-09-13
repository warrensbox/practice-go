package main

import "fmt"

/*
dp[i][j] = dp[i-1][k]+point[i][j]-abs(j-k)


case 1 : j > k
dp[i][j] = dp[i-1][k]+point[i][j]-j+k
         = (dp[i-1][k]+k)+(point[i][j]-j)

case 2 : j < k
dp[i][j] = dp[i-1][k]+point[i][j]-k+j
         = (dp[i-1][k]+j)+(point[i][j]-k)
*/

func maxPoints(points [][]int) int64 {

	dp := make([][]int, len(points))

	for i := 0; i < len(points); i++ {
		dp[i] = make([]int, len(points[0]))
	}

	for j := 0; j < len(points[0]); j++ {
		dp[0][j] = points[0][j]
	}
	//fmt.Println("dp",dp)

	for i := 1; i < len(points); i++ {
		// fmt.Println("i",i)
		left_dp := make([]int, len(points[i]))

		left_dp[0] = dp[i-1][0]
		// fmt.Println("left_dp[0]",left_dp[0])
		// fmt.Println("0000")
		for k := 1; k < len(points[i]); k++ {
			// fmt.Println("left_dp[k-1]",left_dp[k-1])
			// fmt.Println("dp[k]",dp[k])
			// fmt.Println("points[i][k]",points[i][k])
			// fmt.Println("points[i][k-1]",points[i][k-1])
			// left_dp[k] = Max(left_dp[k-1]-points[i][k-1] - 1, dp[k]) + points[i][k];
			// fmt.Println("left_dp",left_dp)
			// fmt.Println("---")
			left_dp[k] = Max(left_dp[k-1], dp[i-1][k]+k)
			// fmt.Println("left_dp",left_dp)
			// fmt.Println("left_dp[k]",left_dp[k])
			// fmt.Println(" dp[i - 1][k]", dp[i - 1][k])
			// fmt.Println("k",k)
			// fmt.Println("---")
		}

		right_dp := make([]int, len(points[0]))
		right_dp[len(right_dp)-1] = dp[i-1][len(right_dp)-1] - len(points[i]) + 1
		//right_dp[len(points[i])-1] = points[i][len(points[i])-1]+dp[len(points[i])-1]
		for k := len(points[i]) - 2; k >= 0; k-- {
			//right_dp[k] = Max(right_dp[k+1]-points[i][k+1] - 1, dp[k]) + points[i][k];
			right_dp[k] = Max(right_dp[k+1], dp[i-1][k]-k)
			// fmt.Println("right_dp",right_dp)
		}

		for j := 0; j < len(points[i]); j++ {
			//dp[k] = Max(left_dp[k] , right_dp[k])
			dp[i][j] = Max(left_dp[j]-j, right_dp[j]+j) + points[i][j]
		}
	}

	fmt.Println("************")
	fmt.Println(dp)
	res := -1
	for j := 0; j < len(points[0]); j++ {
		res = Max(res, dp[len(dp)-1][j])
	}

	return int64(res)

	//return 64

}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
