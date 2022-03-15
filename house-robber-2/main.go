package main

func rob(nums []int) int {

	if nums == nil || len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	lastIndex := len(nums)
	skip := helper(nums, 0, lastIndex-2)     //start from index 0 - last-1
	dontSkip := helper(nums, 1, lastIndex-1) //start from index 1 -last
	return Max(skip, dontSkip)
}

func helper(arr []int, start, end int) int {

	nums := arr[start : end+1]
	dp := make([]int, len(nums))
	max := -1

	for i := 0; i < len(nums); i++ {

		if i == 0 {
			dp[i] = nums[i]
		} else if i == 1 {
			dp[i] = Max(dp[i-1], nums[i])
		} else {
			dp[i] = Max(dp[i-1], dp[i-2]+nums[i]) //if I rob you, can must consider the house before you ELSE I don't rob you
		}
		max = Max(max, dp[i])
	}
	return max
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
