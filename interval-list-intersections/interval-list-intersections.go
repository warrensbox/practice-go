package main

import "fmt"

func main() {

	A := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}
	B := [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}

	output := intervalIntersection(A, B)

	fmt.Println(output)

}

//loop throug a and b
//get the max of A[i][0] | B[i][0]
//get the min of A[i][1] | B[i][1]
//high >= low -> add to result
func intervalIntersection(A [][]int, B [][]int) [][]int {

	result := [][]int{}

	i := 0
	j := 0
	for i < len(A) && j < len(B) {

		lo := max(A[i][0], B[j][0])
		hi := min(A[i][1], B[j][1])

		if hi >= lo {
			newVal := []int{lo, hi}
			result = append(result, newVal)
		}

		if A[i][1] < B[j][1] {
			i++
		} else {
			j++
		}

	}

	return result
}

func max(x, y int) int {

	if x > y {
		return x
	}

	return y
}

func min(x, y int) int {

	if x < y {
		return x
	}

	return y
}
