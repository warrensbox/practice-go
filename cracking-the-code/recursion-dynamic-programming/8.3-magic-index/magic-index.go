package main

import "fmt"

func main() {

	//non repeating sorted array
	unique_sorted := []int{-40, -20, -1, 1, 2, 3, 5, 7, 9, 12, 13}
	fmt.Println(magicFast(unique_sorted, 0, len(unique_sorted)-1))
	notunique_sorted := []int{-10, -5, 2, 2, 2, 3, 4, 7, 9, 12, 13}
	fmt.Println(magicFast2(notunique_sorted, 0, len(notunique_sorted)-1))
}

//using binary type search
func magicFast(arr []int, start, end int) int {

	if end < start {
		return -1
	}
	mid := (start + end) / 2
	if arr[mid] == mid {
		return mid
	} else if arr[mid] > mid {
		return magicFast(arr, start, mid-1)
	} else {
		return magicFast(arr, mid+1, end)
	}

}

func magicFast2(arr []int, start, end int) int {

	if end < start {
		return -1
	}
	midIndex := (start + end) / 2
	midValue := arr[midIndex]
	if midValue == midIndex {
		return midIndex
	}

	fmt.Println("midIndex", midIndex)
	fmt.Println("midValue", midValue)

	//Search left
	leftIndex := Min(midIndex-1, midValue)
	fmt.Println("leftIndex", leftIndex)

	left := magicFast2(arr, start, leftIndex)

	fmt.Println("left", left)
	fmt.Println("----")
	if left >= 0 {
		return left
	}

	fmt.Println("going right")
	//Search right
	rightIndex := Max(midIndex+1, midValue)
	right := magicFast2(arr, rightIndex, end)
	fmt.Println("rightIndex", rightIndex)
	fmt.Println("right", right)
	fmt.Println("----")
	return right

	//fmt.Println("down")
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
