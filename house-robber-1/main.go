package main

func rob(nums []int) int {

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
