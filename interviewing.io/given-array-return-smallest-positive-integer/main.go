package main

import (
	"fmt"
	"sort"
)

/*
A =[ 1,1,2,3,4,6]
A =[ 1,2,3,4]
A =[ -3,-1]


*/

func main() {
	//A := []int{1, 1, 2, 3, 4, 6}
	//A := []int{1, 2, 3}
	//A := []int{-3, -1}
	A := []int{0, 1, 2, 3, 4, 6}
	fmt.Println(solution(A))
}

func solution(A []int) int {

	count := 1
	sort.Ints(A)

	for i := 0; i < len(A); i++ {

		if (i > 0 && A[i] == A[i-1]) || A[i] <= 0 {
			continue
		} else if A[i] == count {
			count++
		} else {
			break
		}
	}

	return count
}
