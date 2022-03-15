package main

import (
	"fmt"
)

func main() {

	input := []int{23, 2, 4, 6, 2}
	k := 6
	output := checkSubarraySum(input, k)

	fmt.Println(output)
}

func checkSubarraySum(nums []int, k int) bool {

	sum := make([]int, len(nums))

	sum[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		sum[i] = sum[i-1] + nums[i]
		fmt.Print(sum[i])
	}
	fmt.Println()
	for start := 0; start < len(nums)-1; start++ {
		fmt.Println("start", start)
		for end := start + 1; end < len(nums); end++ {
			fmt.Println("end", end)
			summation := sum[end] - sum[start] + nums[start]
			fmt.Println(summation)

			if summation == k || summation%k == 0 {
				return true
			}
		}
		fmt.Println("-----")
	}

	return false
}
