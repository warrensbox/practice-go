package main

import (
	"fmt"
	"runtime"
)

func main() {

	//Given nums = [2, 7, 11, 15], target = 9,

	// primes := []int{2, 11, 7, 4, 5, 15}
	// target := 9

	// output := twoSum2(primes, target)
	// fmt.Println(output)

	goarch := runtime.GOARCH
	goos := runtime.GOOS

	fmt.Println(goarch)
	fmt.Println(goos)
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

func twoSum2(nums []int, target int) []int {

	dict := make(map[int]int)

	var arr []int

	for index, val := range nums {

		if key, ok := dict[val]; ok {
			arr = append(arr, index)
			arr = append(arr, key)

		} else {
			remains := target - val
			dict[remains] = index
		}

	}

	return arr
}
