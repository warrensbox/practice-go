package main

import (
	"fmt"
	"math"
)

func main() {

	A := []int{1, 3, 5, 4}
	B := []int{1, 2, 3, 7}

	output := minSwap(A, B)

	fmt.Println(output)
}

func minSwap(A []int, B []int) int {

	// n : natural, s : swapped
	n1 := 0
	s1 := 1

	for i := 1; i < len(A); i++ {

		n2 := math.MaxInt32
		s2 := math.MaxInt32

		if A[i-1] < A[i] && B[i-1] < B[i] {
			n2 = Min(n2, n1)
			s2 = Min(s2, s2+1)
		}

		if A[i-1] < B[i] && B[i-1] < A[i] {
			n2 = Min(n2, s1)
			s2 = Min(s2, n2+1)
		}

		n1 = n2
		s1 = s2
	}

	return Min(n1, s1)
}

func Min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
