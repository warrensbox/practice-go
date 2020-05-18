package main

import (
	"fmt"
	"math"
)

func main() {

	A := []int{3, 5, 1, 2, 3}
	B := []int{3, 6, 3, 3, 4}
	// 	[3,5,1,2,3]
	// [3,6,3,3,4]

	output := minDominoRotations(A, B)
	fmt.Println(output)
}

func minDominoRotations(A []int, B []int) int {

	minSwap := min(numberOfSwaps(A[0], A, B), numberOfSwaps(B[0], A, B))
	if minSwap == math.MaxInt64 {
		return -1
	}
	minSwap = min(numberOfSwaps(A[0], B, A), minSwap)
	if minSwap == math.MaxInt64 {
		return -1
	}
	minSwap = min(numberOfSwaps(B[0], B, A), minSwap)

	if minSwap == math.MaxInt64 {
		return -1
	}
	return minSwap

}

func numberOfSwaps(target int, arr1 []int, arr2 []int) int {
	countSwaps := 0

	for i := 0; i < len(arr1); i++ {

		if arr1[i] != target && arr2[i] != target {
			return math.MaxInt64
		} else if arr1[i] != target {
			countSwaps++
		}
	}

	return countSwaps
}

func min(a, b int) int {

	if a < b {
		return a
	}

	return b
}
