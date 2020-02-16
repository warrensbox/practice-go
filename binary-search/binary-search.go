package main

import (
	"fmt"
)

func main() {

	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 4
	output := binarySearch(input, target)

	fmt.Println(output)
}

func binarySearch(num []int, target int) int {

	lo := 0
	hi := len(num)

	fmt.Println("hi", hi)
	fmt.Println("target", target)

	for lo < hi {
		mid := lo + (hi-lo)/2
		fmt.Println("mid", mid)

		if num[mid] > target || num[mid] == target {
			hi = mid
		} else {
			lo = mid + 1
		}

	}

	if num[hi] == target {
		fmt.Println("Found", num[hi])
		return hi
	} else {
		fmt.Println("Not Found")
		return -1
	}
}
