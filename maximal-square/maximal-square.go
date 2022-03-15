package main

import (
	"fmt"
)

func main() {

	input := [][]byte{{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'}}

	//fmt.Println(input[3][4])

	output := maximalSquare(input)

	fmt.Println(output)
}

func maximalSquare(matrix [][]byte) int {

	if len(matrix) == 0 {
		return 0
	}

	row := len(matrix)
	col := len(matrix[0])
	cache := make([][]int, row)
	max := 0
	for i := range cache {
		cache[i] = make([]int, col)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {

			if i == 0 || j == 0 || matrix[i][j] == '0' {

				cache[i][j] = int(matrix[i][j] - '0')
			} else {
				min := min(cache[i][j-1], min(cache[i-1][j], cache[i-1][j-1]))
				cache[i][j] = min + 1
			}
			// fmt.Print(cache[i][j])

			if cache[i][j] > max {
				max = cache[i][j]
			}

			// for i := 0; i < row; i++ {
			// 	for j := 0; j < col; j++ {
			// 		fmt.Print(cache[i][j])
			// 	}
			// 	fmt.Println()
			// }
			// fmt.Println("-----------")
		}
		fmt.Println("-----------")
	}

	return max * max
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
