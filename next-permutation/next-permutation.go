package main

import (
	"fmt"
)

func main() {

	input := []int{1, 2, 1}

	output := nextPermutation(input)
	fmt.Println(output)
}

/**
 * Find the last element which is a peak.
 * In case there are multiple equal peaks, return the first of those.
 * @return first index of last peak
 */
func indexOfLastPeak(nums []int) int {

	for i := len(nums) - 1; i > 0; i-- {

		if nums[i-1] < nums[i] {
			return i
		}
	}
	return 0

}

/** @return last index where the {@code num > threshold} or -1 if not found */
func lastIndexOfGreater(nums []int, threshold int) int {
	for i := len(nums) - 1; i >= 0; i-- {

		if threshold < nums[i] {
			return i
		}
	}
	return -1
}

func nextPermutation(nums []int) []int {

	// pivot is the element just before the non-increasing (weakly decreasing) suffix
	pivot := indexOfLastPeak(nums) - 1

	// paritions nums into [prefix pivot suffix]
	if pivot != -1 {
		nextPrefix := lastIndexOfGreater(nums, nums[pivot]) // in the worst case it's suffix[0]
		// next prefix must exist because pivot < suffix[0], otherwise pivot would be part of suffix
		nums = swap(nums, pivot, nextPrefix) // this minimizes the change in prefix
	}

	nums = reverseSuffix(nums, pivot+1)
	return nums
}

func swap(nums []int, i int, j int) []int {
	tempo := nums[i]
	nums[i] = nums[j]
	nums[j] = tempo

	return nums
}

/** Reverse numbers starting from an index till the end. */
func reverseSuffix(nums []int, start int) []int {

	end := len(nums) - 1

	for start < end {
		start = start + 1
		end = end - 1
		nums = swap(nums, start, end)
	}

	return nums
}
