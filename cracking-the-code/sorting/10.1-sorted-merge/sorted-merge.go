package main

import (
	"fmt"
)

func main() {

	arrB := []int{2, 6, 7, 8, 10, 14, 15}
	arrA := [14]int{1, 3, 5, 4, 9, 11, 13}

	mergeSort(arrA, arrB, 7, 7)
}

func mergeSort(arrA [14]int, arrB []int, lastindexA, lastindexB int) {

	aIndex := lastindexA - 1
	bIndex := lastindexB - 1
	mergedIndex := aIndex + bIndex + 1

	for bIndex >= 0 {

		if aIndex >= 0 && arrA[aIndex] > arrB[bIndex] {
			arrA[mergedIndex] = arrA[aIndex]
			aIndex--
		} else {
			arrA[mergedIndex] = arrB[bIndex]
			bIndex--
		}
		mergedIndex--
	}

	fmt.Println(arrA)
}
