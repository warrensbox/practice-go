package main

import (
	"fmt"
)

func main() {

	inputArr := []int{1, 2, 1, 1, 5}
	targetVal := 1
	output := removeElement(inputArr, targetVal)
	fmt.Println(output)
}

func removeElement(nums []int, val int) int {

	lengthArr := len(nums)
	i := 0
	count := 0
	for i < lengthArr {

		if nums[i] == val {
			nums[i] = nums[lengthArr-1]
			fmt.Println("nums[i]", nums)
			lengthArr--

		} else {
			fmt.Println("Not same")
			i++
			count++
		}
	}

	fmt.Println("lengthArr", lengthArr)
	fmt.Println("count", count)
	return lengthArr

}
