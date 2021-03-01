package main

import (
	"fmt"
	"math"
)

func main() {
	array := []int{5, 3, 1, 2, 4}
	// fmt.Println(array)
	arr := sortValleyPeak(array)
	fmt.Println(arr)
	//fmt.Println(maxIndex(array, 0, 1, 2))
}

func sortValleyPeak(array []int) []int {

	for i := 1; i < len(array); i += 2 {
		biggestIndex := maxIndex(array, i-1, i, i+1)
		if i != biggestIndex {
			array[i], array[biggestIndex] = array[biggestIndex], array[i]
		}
	}

	return array
}

func maxIndex(array []int, a, b, c int) int {
	max := math.MinInt64
	length := len(array)
	index := a
	if array[a] > max && a < length {
		max = array[a]
		index = a
	}
	if array[b] > max && b < length {
		max = array[b]
		index = b
	}
	if array[c] > max && c < length {
		max = array[c]
		index = c
	}
	return index
}
