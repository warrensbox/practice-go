package main

import "fmt"

func main() {

	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	output := maxArea(input)

	fmt.Println(output)
}

func maxArea(height []int) int {

	maxArea := 0
	left := 0
	right := len(height) - 1

	for left < right {

		maxArea = max(maxArea, min(height[left], height[right])*(right-left))

		if height[left] < height[right] {
			left++
		} else {
			right--
		}

	}

	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
