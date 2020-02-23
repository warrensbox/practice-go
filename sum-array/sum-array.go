package main

import (
	"fmt"
)

func main() {

	input := []int{23, 2, 4, 6, 7}

	output := sumofarray(input)

	fmt.Println(output)
}

func sumofarray(nums []int) int {

	sum := make([]int, len(nums))

	sum[0] = nums[0]

	for i := 1; i < len(nums); i++ {

		sum[i] = sum[i-1] + nums[i]
		//fmt.Println(sum[i])
	}

	return sum[len(nums)-1]
}
