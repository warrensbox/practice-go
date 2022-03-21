package main

import "fmt"

//rotation-point

func main() {

	arr := []int{5, 6, 7, 8, 9, 10, 11, 1, 2, 3, 4}
	//			 0  1  2  3  4   5   6  7  8  9  10
	fmt.Println(findRotationPoint(arr))
}

func findRotationPoint(arr []int) int {

	firstWorld := arr[0]

	left := 0
	right := len(arr) - 1

	for left < right {
		mid := (left + right) / 2
		fmt.Println("mid", mid)
		if arr[mid] >= firstWorld {
			//go right
			left = mid
		} else {
			//go left
			right = mid
		}

		fmt.Println("left", left)
		fmt.Println("right", right)
		if left+1 == right {
			break
		}
	}

	return arr[right]
}
