package main

import "fmt"

func main() {

	//Given nums = [2, 7, 11, 15], target = 9,

	primes := []int{2, 11, 7, 15}
	target := 9

	output := twoSum(primes, target)
	fmt.Println(output)
}

func twoSum(nums []int, target int) []int {

	dict := make(map[int]int)

	var arr []int

	for index, val := range nums {

		remains := target - nums[index]

		_, exist := dict[remains]

		if exist {
			arr = append(arr, dict[remains])
			arr = append(arr, index)

		} else {
			dict[val] = index
		}

	}

	return arr
}
