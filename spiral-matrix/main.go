package main

import "fmt"

func spiralOrder(matrix [][]int) []int {

	colStart := 0
	colEnd := len(matrix[0]) - 1
	rowStart := 0
	rowEnd := len(matrix) - 1

	//total := len(matrix[0]) * len(matrix)
	res := []int{}
	dir := 0
	for colStart <= colEnd && rowStart <= rowEnd {

		switch dir {
		case 0: //go right
			for col := colStart; col <= colEnd; col++ {
				fmt.Println(matrix[rowStart][col])
				res = append(res, matrix[rowStart][col])
			}
			rowStart++
		case 1: //go down
			for row := rowStart; row <= rowEnd; row++ {
				fmt.Println(matrix[row][colEnd])
				res = append(res, matrix[row][colEnd])
			}
			colEnd--
		case 2: //go left
			for col := colEnd; col >= colStart; col-- {
				fmt.Println(matrix[rowEnd][col])
				res = append(res, matrix[rowEnd][col])
			}
			rowEnd--
		case 3: //go up
			for row := rowEnd; row >= rowStart; row-- {
				fmt.Println(matrix[row][colStart])
				res = append(res, matrix[row][colStart])
			}
			colStart++
		}
		dir = (dir + 1) % 4
	}

	return res

}
