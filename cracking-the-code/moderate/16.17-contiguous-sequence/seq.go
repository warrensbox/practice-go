package main

import "fmt"

func main() {

	arr := []int{5, -9, 6, -2, 3}
	fmt.Println(getMaxSum(arr))
}

func getMaxSum(arr []int) int {

	sum := 0
	maxSum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if sum > maxSum {
			maxSum = sum
		} else if sum < 0 {
			sum = 0
		}
	}

	return maxSum
}
