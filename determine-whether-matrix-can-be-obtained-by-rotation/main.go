package main

func findRotation(mat [][]int, target [][]int) bool {

	check := make([]bool, 4)

	for i := range check {
		check[i] = true
	}
	n := len(mat)

	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			if mat[r][c] != target[r][c] {
				check[3] = false
			}
			if mat[r][c] != target[c][n-r-1] {
				check[0] = false
			}
			if mat[r][c] != target[n-r-1][n-c-1] {
				check[1] = false
			}
			if mat[r][c] != target[n-c-1][r] {
				check[2] = false
			}
		}
	}

	return check[0] || check[1] || check[2] || check[3]
}
