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

	//sort each array
	sort.Ints(arr1)
	sort.Ints(arr2)
	a, b := 0, 0
	for a < len(arr1) && b < len(arr2) {
		val := Abs(arr1[a] - arr2[b])
		if val < diff {
			diff = val
		}

		if arr1[a] < arr2[b] {
			a++
		} else {
			b++
		}
	}

	return diff
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
