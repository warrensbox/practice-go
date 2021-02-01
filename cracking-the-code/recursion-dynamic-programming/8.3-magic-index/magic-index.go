package main

import "fmt"

func main() {

	//non repeating sorted array
	unique_sorted := []int{-40, -20, -1, 1, 2, 3, 5, 7, 9, 12, 13}
	fmt.Println(magicFast(unique_sorted, 0, len(unique_sorted)-1))
	fmt.Println(magicFast2(unique_sorted, 0, len(unique_sorted)-1))
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
	left := magicFast2(arr, start, leftIndex)
	fmt.Println("leftIndex", leftIndex)
	fmt.Println("left", left)
	if left >= 0 {
		return left
	}

	//Search right
	rightIndex := Max(midIndex+1, midValue)
	right := magicFast2(arr, rightIndex, end)
	fmt.Println("rightIndex", rightIndex)
	fmt.Println("right", right)
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
