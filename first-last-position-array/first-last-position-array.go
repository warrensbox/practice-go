package main

import (
	"fmt"
)

func main() {

	input := []int{5, 7, 7, 8, 8, 10}
	target := 8

	output := searchRange(input, target)
	fmt.Println(output)
}

func searchRange(nums []int, target int) [2]int {

	targetRange := [2]int{-1, -1}

	leftIdx := extremeInsertionIndex(nums, target, true)

	// assert that `leftIdx` is within the array bounds and that `target`
	// is actually in `nums`.
	if leftIdx == len(nums) || nums[leftIdx] != target {
		return targetRange
	}

	rightIdx := extremeInsertionIndex(nums, target, false)

	targetRange[0] = leftIdx
	targetRange[1] = rightIdx - 1

	return targetRange
}

//     // returns leftmost (or rightmost) index at which `target` should be
//     // inserted in sorted array `nums` via binary search.
func extremeInsertionIndex(nums []int, target int, left bool) int {

	lo := 0
	hi := len(nums)

	for lo < hi {
		mid := (lo + hi) / 2
		if nums[mid] > target || (left && target == nums[mid]) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
