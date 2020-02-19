package main

import (
	"fmt"
)

func main() {

	input := []int{4, 5, 8, 2}

	output := productExceptSelf(input)
	fmt.Println("output", output)
}

func productExceptSelf(nums []int) []int {

	N := len(nums)
	output := make([]int, N)

	output[0] = 1

	for i := 1; i < N; i++ {
		number := nums[i-1]
		left := output[i-1]
		output[i] = number * left
	}

	R := 1
	for i := N - 1; i >= 0; i-- {
		out := output[i]
		output[i] = out * R
		R = R * nums[i]

	}

	fmt.Println(output)

	return output
}

func productExceptSelf2(nums []int) []int {

	N := len(nums)

	leftProduct := make([]int, N)
	rightProduct := make([]int, N)
	output := make([]int, N)

	leftProduct[0] = 1
	rightProduct[N-1] = 1

	for i := 1; i < N; i++ {
		// fmt.Println("i", i)
		// fmt.Println("nums", nums[i-1])
		// fmt.Println("leftProduct", leftProduct[i-1])
		number := nums[i-1]
		left := leftProduct[i-1]
		leftProduct[i] = number * left
	}

	for i := N - 2; i >= 0; i-- {
		fmt.Println("i", i)
		fmt.Println("nums", nums[i+1])
		fmt.Println("rightProduct", rightProduct[i+1])
		number := nums[i+1]
		right := rightProduct[i+1]
		rightProduct[i] = number * right
	}

	for i := 0; i < N; i++ {
		right := rightProduct[i]
		fmt.Println("right", right)
		left := leftProduct[i]
		fmt.Println("left", left)
		output[i] = right * left
		fmt.Println(output)
	}

	fmt.Println(output)

	return output
}
