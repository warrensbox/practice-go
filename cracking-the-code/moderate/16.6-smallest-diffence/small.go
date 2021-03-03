package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {

	A := []int{1, 2, 11, 15}
	B := []int{4, 12, 19, 23, 127, 235}
	fmt.Println(findSmallestDiff(A, B))

}

func findSmallestDiff(arr1, arr2 []int) int {
	diff := math.MaxInt64

	mergeArr := append(arr1, arr2...)
	sort.Ints(mergeArr)
	fmt.Println(mergeArr)

	for i := 1; i < len(mergeArr)-1; i++ {
		newDiff := mergeArr[i] - mergeArr[i-1]
		fmt.Println("newDiff", newDiff)
		if newDiff < diff {
			diff = newDiff
		}
	}

	return diff
}
