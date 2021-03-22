package main

import (
	"fmt"
)

func main() {

	//input := []int{23, 2, 4, 6, 7}
	input := []int{1, 5, 8, 4, 19}
	k := 3
	output := checkSubarraySum(input, k)

	fmt.Println(output)
}

func checkSubarraySum(nums []int, k int) bool {

	hash := make(map[int]int)
	sum := 0
	hash[0] = -1
	// agg := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		// agg += nums[i]
		// fmt.Println("agg", agg)
		if k != 0 {
			sum = sum % k
		}
		fmt.Println("sum", sum)
		fmt.Println("hash[sum]", hash)
		if position, ok := hash[sum]; ok {
			if i-position >= 2 {
				return true
			}
		} else {
			hash[sum] = i
		}
	}
	return false
}
