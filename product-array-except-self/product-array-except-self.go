package main

import (
	"fmt"
)

func main() {

	input := []int{1, 2, 3, 4}

	output := productExceptSelf2(input)
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

	fmt.Println("OUTPUT", output)

	R := 1
	for i := N - 1; i >= 0; i-- {
		out := output[i]
		output[i] = out * R
		R = R * nums[i]
		fmt.Println("R", R)
		fmt.Println("output", output)
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
		fmt.Println("number", number)
		left := leftProduct[i-1]
		fmt.Println("left", left)
		leftProduct[i] = number * left
		fmt.Println("leftProduct", leftProduct)
	}

	fmt.Println("------")
	for i := N - 2; i >= 0; i-- {

		number := nums[i+1]
		right := rightProduct[i+1]
		rightProduct[i] = number * right
		// fmt.Println("nums", nums)

		fmt.Println("rightProduct", rightProduct)
	}

	fmt.Println("------")

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
