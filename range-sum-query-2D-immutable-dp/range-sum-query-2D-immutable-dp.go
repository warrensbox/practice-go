package main

import "fmt"

type NumMatrix struct {
	empty  bool
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{
			empty: true,
		}
	}

	dp := make([][]int, len(matrix)+1)

	for i := 0; i <= len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0])+1)
	}

	for i := 1; i <= len(matrix); i++ {
		for j := 1; j <= len(matrix[0]); j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1] + matrix[i-1][j-1]
			fmt.Print(dp[i][j])
			fmt.Print(",")
		}
		fmt.Println()
	}

	return NumMatrix{
		empty:  false,
		matrix: dp,
	}
}

func main() {

	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}

	obj := Constructor(matrix)

	param_1 := obj.SumRegion(2, 1, 4, 3)

	fmt.Println(param_1)
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {

	sum := this.matrix[row2+1][col2+1] - this.matrix[row1][col2+1] - this.matrix[row2+1][col1] + this.matrix[row1][col1]

	return sum
}
