package main

import (
	"fmt"
)

func main() {

	input := []int{1, 2, 3, 5, 4}
	//input := []int{5, 4, 3, 2, 1}
	//input := []int{1, 2, 3, 4, 5}

	output := isMonotonic(input)

	fmt.Println(output)
}

func isMonotonic(A []int) bool {

	increasing := true
	decreasing := true

	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			decreasing = false
		}
		if A[i] < A[i+1] {
			increasing = false
		}
	}

	return increasing || decreasing

}
