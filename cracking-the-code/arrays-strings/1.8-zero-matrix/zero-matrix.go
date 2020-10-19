package main

import "fmt"

func main() {

	arr := [][]int{
		{1, 2, 3, 4},
		{5, 0, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 0},
	}

	printMatrix(arr)
	fmt.Println("======")
	zeroMatrix(arr)
	printMatrix(arr)

}

func zeroMatrix(matrix [][]int) {

	row := make([]bool, len(matrix))
	col := make([]bool, len(matrix[0]))

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}

	for i := 0; i < len(row); i++ {
		if row[i] {
			nullifyRow(matrix, i)
		}
	}

	for j := 0; j < len(col); j++ {
		if row[j] {
			nullifyCol(matrix, j)
		}
	}

}

func nullifyRow(matrix [][]int, row int) {
	for j := 0; j < len(matrix[row]); j++ {
		matrix[row][j] = 0
	}
}

func nullifyCol(matrix [][]int, col int) {
	for i := 0; i < len(matrix); i++ {
		matrix[i][col] = 0
	}
}

func printMatrix(arr [][]int) {

	//i = row
	//j = column
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			fmt.Print(" ")
			fmt.Print(arr[i][j])
		}
		fmt.Println()
	}

}
