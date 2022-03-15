package main

import "sort"

func minDifference(nums []int) int {

	n := len(nums)
	if n < 5 {
		return 0
	}

	sort.Ints(nums)

	//3 Large
	A := nums[n-4] - nums[0]

	//2 Large + 1 Small
	B := nums[n-3] - nums[1]

	//1 Large + 2 Small
	C := nums[n-2] - nums[2]

	//3 Small
	D := nums[n-1] - nums[3]

	return Min(A, Min(B, Min(C, D)))
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//READ
