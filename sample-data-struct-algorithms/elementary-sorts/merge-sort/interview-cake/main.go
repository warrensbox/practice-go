package main

import (
	"fmt"
)

func main() {

	oriArr := []int{1, 3, 5, 7, 5, 10, 11, 12, 2, 4, 6, 8, 9}

	fmt.Println(mergeSort(oriArr))
}

func mergeSort(oriArr []int) []int {

	//base case : single element array
	if len(oriArr) <= 1 {
		return oriArr
	}

	midIdx := len(oriArr) / 2
	left := make([]int, midIdx)
	right := make([]int, len(oriArr)-midIdx)
	//split the input in half
	copy(left, oriArr[:midIdx])
	fmt.Println("left", left)
	copy(right, oriArr[midIdx:])
	fmt.Println("right", right)

	//os.Exit(0)
	//sort each half
	leftSorted := mergeSort(left)
	rightSorted := mergeSort(right)

	//merge the sorted halves
	return combineSortedArray(leftSorted, rightSorted)
}

func combineSortedArray(arrA, arrB []int) []int {

	idxA, idxB, idxMerged := 0, 0, 0

	mergedArr := make([]int, len(arrA)+len(arrB))

	for idxA < len(arrA) && idxb  B < len(arrB) {

		if arrA[idxA] <= arrB[idxB] {
			mergedArr[idxMerged] = arrA[idxA]
			idxA++
		} else {
			mergedArr[idxMerged] = arrB[idxB]
			idxB++
		}

		idxMerged++
	}

	for idxA < len(arrA) {
		mergedArr[idxMerged] = arrA[idxA]
		idxA++
		idxMerged++
	}

	for idxB < len(arrB) {
		mergedArr[idxMerged] = arrB[idxB]
		idxB++
		idxMerged++
	}
	fmt.Println("merged", mergedArr)
	return mergedArr
}
