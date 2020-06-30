package main

import (
	"fmt"
	"sort"
)

func main() {

	input := []int{8, 1, 2, 2, 3}
	output := smallerNumbersThanCurrent(input)
	fmt.Println(output)
}

func smallerNumbersThanCurrent(nums []int) []int {

	arr := make([]int, len(nums))
	copy(arr, nums)
	dict := make(map[int]int)
	sort.Ints(nums)

	for index, val := range nums {
		if _, ok := dict[val]; !ok {
			dict[val] = index
		}
	}

	for i, val := range arr {
		fmt.Println(val)
		arr[i] = dict[val]
	}

	return arr
}
