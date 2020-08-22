package main

import (
	"fmt"
)

func main() {

	inputArr := []int{3, 2, 1, 2, 5}
	targetVal := 1
	output := removeElement(inputArr, targetVal)
	fmt.Println(output)
}

func removeElement(nums []int, val int) int {

	lengthArr := len(nums)
	i := 0
	for i < lengthArr {

		if nums[i] == val {
			nums[i] = nums[lengthArr-1]
			fmt.Println("nums[i]", nums[i])
			lengthArr--
		} else {
			fmt.Println("Not same")
			i++
		}
	}

	fmt.Println("lengthArr", lengthArr)
	fmt.Println("nums", nums)
	return lengthArr

}
