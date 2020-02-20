package main

import "fmt"

type NumMatrix struct {
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	return NumMatrix{matrix: matrix}
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

	sum := 0
	for i := row1; i <= row2; i++ {

		for j := col1; j <= col2; j++ {
			sum += this.matrix[i][j]
		}
	}

	return sum
}
