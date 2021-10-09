package main

func countSquares(matrix [][]int) int {

	res := 0
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[0]); c++ {

			if r > 0 && c > 0 && matrix[r][c] > 0 {
				min := Min(matrix[r-1][c], Min(matrix[r][c-1], matrix[r-1][c-1]))
				matrix[r][c] = min + 1
			}

			res += matrix[r][c]

		}
	}

	return res
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
