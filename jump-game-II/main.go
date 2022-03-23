package main

func main() {

}
func jump(nums []int) int {

	dp := make([]int, len(nums))
	// from := make([]int,len(nums))

	for i, _ := range dp {
		dp[i] = 10000
	}
	dp[0] = 0
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j]+j >= i {
				dp[i] = Min(dp[j]+1, dp[i])
			}
		}
	}

	return dp[len(dp)-1]

}

func Min(x, y int) int {

	if x < y {
		return x
	}

	return y
}
