package main

import "fmt"

func countSquares(matrix [][]int) int {

	fmt.Println(matrix)

	//define cache arr
	//     cache := make([][]int,len(matrix))

	//     //initialize the cache
	//     for i,_ := range matrix {
	//         cache[i] = make([]int,len(matrix[0]))
	//     }count-square-submatrices-with-all-ones

	res := 0
	// hash := make(map[int]int)
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[0]); c++ {

			if r > 0 && c > 0 && matrix[r][c] > 0 {
				min := Min(matrix[r-1][c], Min(matrix[r][c-1], matrix[r-1][c-1]))
				fmt.Println("min", min)
				matrix[r][c] = min + 1
			}

			res += matrix[r][c]

			// if matrix[r][c] == 1 {
			//     ones++
			// }
			// if cache[r][c] > 1 {
			//     hash[cache[r][c]]++
			// }
		}
	}

	fmt.Println(matrix)
	// fmt.Println(hash)

	return res
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
