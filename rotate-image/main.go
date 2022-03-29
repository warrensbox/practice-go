package main

func rotate(matrix [][]int) {

	left := 0
	right := len(matrix) - 1
	top := 0
	bottom := len(matrix) - 1

	for left < right {
		for i := 0; i < right-left; i++ {

			tmp := matrix[top][left+i]
			matrix[top][left+i] = matrix[bottom-i][left]
			matrix[bottom-i][left] = matrix[bottom][right-i]
			matrix[bottom][right-i] = matrix[top+i][right]
			matrix[top+i][right] = tmp
		}
		left++
		right--
		bottom--
		top++
	}

}
