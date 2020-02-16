package main

import (
	"fmt"
)

func main() {

	input := 10
	target := 4

	if target > input {
		fmt.Println("Target is bigger then input")
	}

	firstBadVersion(input, target)

}

func firstBadVersion(input int, target int) {

	left := 1
	right := input

	for left < right {
		mid := left + (right-left)/2
		fmt.Println("mid", mid)

		if mid >= target {
			fmt.Println("True", mid)
			left = mid + 1
		} else {
			fmt.Println("False")
			right = mid
		}
	}
}
