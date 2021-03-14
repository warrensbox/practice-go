package main

import (
	"fmt"
	"sort"
)

func main() {

	arr := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(smallestK(arr, 4))
}

func smallestK(arr []int, k int) []int {

	sort.Ints(arr)
	smallest := []int{}
	for i := 0; i < k; i++ {
		smallest = append(smallest, arr[i])
	}

	return smallest
}
