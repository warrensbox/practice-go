package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5, 5}

	splitArrSum(arr)
}

func splitArrSum(arr []int) {

	splitPoint := findSplitPoint(arr)

	if splitPoint == -1 || splitPoint == len(arr) {
		fmt.Println("Not Possible")
		return
	}

	fmt.Println(arr[:splitPoint])
	fmt.Println(arr[splitPoint:])
}

func findSplitPoint(arr []int) int {
	leftSum := 0
	for i := 0; i < len(arr); i++ {
		leftSum += arr[i]
	}

	if leftSum%2 != 0 {
		return -1
	}

	rightSum := 0
	for i := len(arr) - 1; i >= 0; i-- {
		rightSum += arr[i]

		leftSum -= arr[i]

		if leftSum == rightSum {
			return i
		}
	}
	return -1
}
