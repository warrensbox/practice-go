package main

import (
	"fmt"
)

func main() {

	input := []int{5, 1, 10, 6, 7}
	k := 6
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
		fmt.Println("sum", sum)
		if pos, ok := m[remainder]; ok && i-pos >= 2 {
			fmt.Println("m", m)
			return true
		}
	}
	return false
}
