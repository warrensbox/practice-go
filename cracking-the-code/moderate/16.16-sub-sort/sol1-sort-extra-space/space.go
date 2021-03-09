package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19}
	fmt.Println(sortSub(arr))
}

func sortSub(orig []int) (int, int) {

	cpy := make([]int, len(orig))
	copy(cpy, orig)
	sort.Ints(cpy)
	// fmt.Println(cpy)
	// fmt.Println(orig)
	min := len(orig) - 1
	max := 0
	for i := 0; i < len(orig); i++ {
		if orig[i] != cpy[i] {
			min = Min(min, i)
			max = Max(max, i)
		}
	}

	return min, max
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
