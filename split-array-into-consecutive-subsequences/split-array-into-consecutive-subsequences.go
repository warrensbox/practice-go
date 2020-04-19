package main

import "fmt"

//https://www.youtube.com/watch?v=uJ8BAQ8lASE

func main() {

	nums := []int{1, 2, 3, 3, 4, 4, 5, 5}

	//nums := []int{1, 2, 3, 4, 4, 5}

	output := isPossible(nums)

	fmt.Println(output)
}

func isPossible(nums []int) bool {

	freqMap := make(map[int]int)
	hypoteticalMap := make(map[int]int)

	for _, val := range nums {
		freqMap[val]++
	}

	for i := range nums {

		if freqMap[nums[i]] == 0 {
			continue
		}

		fmt.Println(freqMap)

		if hypoteticalMap[nums[i]] > 0 {
			hypoteticalMap[nums[i]]--
			hypoteticalMap[nums[i]+1]++
			freqMap[nums[i]]--
			fmt.Println("hypoteticalMap ll", hypoteticalMap)
		} else if freqMap[nums[i]] > 0 && freqMap[nums[i]+1] > 0 && freqMap[nums[i]+2] > 0 {
			freqMap[nums[i]]--
			freqMap[nums[i]+1]--
			freqMap[nums[i]+2]--
			hypoteticalMap[nums[i]+3]++
			fmt.Println("hypoteticalMap", hypoteticalMap)

		} else {
			return false
		}
	}
	return true
}
