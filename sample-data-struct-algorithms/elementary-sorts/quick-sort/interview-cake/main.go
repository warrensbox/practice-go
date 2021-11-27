package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 8, 2, 3, 6, 5, 7, 9, 4}
	//var arr []int
	//arr = append(arr, 1, 3, 2, 4, 6, 5, 7, 9, 8)
	quickSort(&arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(arr *[]int, startIndex, endIndex int) {

	if startIndex >= endIndex {
		return
	}
	pivot := partition(arr, startIndex, endIndex)
	fmt.Println("pivotTOP", pivot)
	quickSort(arr, startIndex, pivot-1)

	quickSort(arr, pivot+1, endIndex)

	fmt.Println("final", *arr)
}

func partition(arr *[]int, startIndex, endIndex int) int {
	left := startIndex
	right := endIndex - 1
	pivot := (*arr)[endIndex]

	fmt.Println("left", left)
	fmt.Println("right", right)
	fmt.Println("pivot", pivot)

	fmt.Println("HERE", *arr)
	for left <= right {
		//fmt.Println("(*arr)[left]", (*arr)[left])
		fmt.Println("pivot", pivot)
		for left <= endIndex && (*arr)[left] < pivot {
			fmt.Println("(*arr)[left]", (*arr)[left])
			left++
		}

		for right >= startIndex && (*arr)[right] > pivot {
			fmt.Println("(*arr)[right]", (*arr)[right])
			right--
		}

		if left < right {
			exchange(arr, left, right)
		} else {
			exchange(arr, left, endIndex)
		}
	}
	fmt.Println("HERE2", *arr)
	return left
}

func exchange(arr *[]int, j, i int) {
	(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
}
