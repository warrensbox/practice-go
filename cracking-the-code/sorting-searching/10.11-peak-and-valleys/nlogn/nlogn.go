package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int{5, 3, 1, 2, 4}
	fmt.Println(array)
	arr := sortValleyPeak(array)
	fmt.Println(arr)
}

func sortValleyPeak(array []int) []int {
	sort.Ints(array)
	fmt.Println(array)
	for i := 1; i < len(array); i += 2 {
		array[i], array[i-1] = array[i-1], array[i]
	}

	return array
}
