package main

import (
	"fmt"
)

func main() {

	//input := []int{23, 2, 4, 6, 7}
	input := []int{90, 5, 6, 4, 7}
	k := 3
	output := checkSubarraySum(input, k)

	fmt.Println(output)
}

func checkSubarraySum(nums []int, k int) bool {
	if k < 0 {
		k = -k
	}
	hash := make(map[int]int)
	sum := 0
	hash[0] = -1

	for i := 0; i < len(nums); i++ {
		var remainder int
		sum += nums[i]
		fmt.Println("sum", sum)
		if k == 0 {
			remainder = sum
		} else {
			remainder = sum % k
		}
		if _, ok := hash[remainder]; !ok {
			hash[remainder] = i
		}
		fmt.Println("hash", hash)
		fmt.Println("=====")
		if position, ok := hash[remainder]; ok && i-position >= 2 {
			return true
		}
	}
	return false
}
