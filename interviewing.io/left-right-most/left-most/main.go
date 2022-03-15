package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 4, 4, 4, 4, 5}
	// 0  1  2  3  4  5  6  7  8

	fmt.Println(leftmost(arr, 4))
}

func leftmost(arr []int, target int) int {

	lo := 0
	hi := len(arr)

	for lo < hi {

		mid := (lo + hi) / 2

		if arr[mid] > target || arr[mid] == target {
			hi = mid
		} else {
			lo = mid + 1
		}

	}

	return lo
}
