package main

import (
	"fmt"
)

func main() {

	input := []int{1, 2, 3, 4, 5}

	output := isMonotonic(input)

	fmt.Println(output)
}

func isMonotonic(A []int) bool {

	increasing := true
	decreasing := true

	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			increasing = false
		}
		if A[i] > A[i+1] {
			decreasing = false
		}
	}

	return increasing || decreasing

}
