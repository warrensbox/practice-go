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

		fmt.Println("--FOR LOOP START---")

		n2 := math.MaxInt32
		s2 := math.MaxInt32

		fmt.Print("n2", n2)
		fmt.Println(" s2", s2)
		fmt.Print(" A[i-1]", A[i-1])
		fmt.Println(" A[i] ", A[i])
		fmt.Print(" B[i-1] ", B[i-1])
		fmt.Println(" B[i]", B[i])
		if A[i-1] < A[i] && B[i-1] < B[i] {
			fmt.Println("------INLINE-----")
			fmt.Println("n2", n2)
			fmt.Println("s2", s2)
			fmt.Println("n1", n1)
			fmt.Println("s1", s1)
			n2 = Min(n2, n1)
			s2 = Min(s2, s1+1)
		}

		if A[i-1] < B[i] && B[i-1] < A[i] {
			fmt.Println("------DIAGONAL-----")
			n2 = Min(n2, s1)
			s2 = Min(s2, n1+1)
			fmt.Println("n2", n2)
			fmt.Println("s2", s2)
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
