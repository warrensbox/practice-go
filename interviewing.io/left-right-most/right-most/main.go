package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 4, 4, 4, 4, 5, 7}

	fmt.Println(rightmost(arr, 4) - 1)
}

func rightmost(arr []int, target int) int {

	lo := 0
	hi := len(arr) - 1

	for lo <= hi {

		mid := (lo + hi) / 2

		if arr[mid] < target || arr[mid] == target {
			lo = mid + 1
		} else {
			hi = mid
		}

	}

	return lo
}
