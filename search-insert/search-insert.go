package main

import (
	"fmt"
)

func main() {

	input := []int{1, 3, 5, 6}
	target := 7
	output := searchInsert(input, target)

	fmt.Println(output)
}

func searchInsert(nums []int, target int) int {

	index, i := 0, 0
	for i < len(nums) {
		fmt.Println("nums", nums[i])

		if target-nums[i] == 0 {
			index = i
			break
		} else if target-nums[i] > 0 {
			if i == len(nums)-1 {
				i++
				index = i
				break
			} else {
				i++
			}
		} else if target-nums[i] < 0 {
			index = i
			break
		}
	}

	return index
}
