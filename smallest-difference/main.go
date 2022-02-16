package main

import (
	"math"
	"sort"
)

func SmallestDifference(array1, array2 []int) []int {
	// Write your code here.
	sort.Ints(array1)
	sort.Ints(array2)
	idx1, idx2 := 0, 0
	memo := make(map[int][]int)
	min := math.MaxInt64
	for idx1 < len(array1) && idx2 < len(array2) {

		smallest := array1[idx1] - array2[idx2]
		memo[abs(smallest)] = []int{array1[idx1], array2[idx2]}
		min = Min(min, abs(smallest))

		if array1[idx1] < array2[idx2] {
			idx1++
		} else {
			idx2++
		}

	}

	return memo[min]
	//return nil
}

func Min(x, y int) int {

	if x < y {
		return x
	}

	return y
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
