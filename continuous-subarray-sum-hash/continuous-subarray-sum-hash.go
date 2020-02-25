package main

import (
	"fmt"
)

func main() {

	input := []int{23, 2, 4, 6, 7}
	k := 7
	output := checkSubarraySum(input, k)

	fmt.Println(output)
}

func checkSubarraySum(nums []int, k int) bool {
	if k < 0 {
		k = -k
	}
	m := make(map[int]int) // map[sum](1st pos)
	sum := 0
	m[0] = -1
	for i := 0; i < len(nums); i++ {
		var remainder int
		sum += nums[i]
		if k == 0 {
			remainder = sum
		} else {
			remainder = sum % k
		}
		if _, ok := m[remainder]; !ok {
			m[remainder] = i
		}
		if pos, ok := m[remainder]; ok && i-pos >= 2 {
			return true
		}
	}
	return false
}
