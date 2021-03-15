package main

import "fmt"

func main() {
	input := []int{0, 0, 4, 0, 0, 6, 0, 0, 3, 0, 8, 0, 2, 0, 5, 2, 0, 3, 0, 0}
	fmt.Println(computeHistogramVolume(input))
}

func computeHistogramVolume(histogram []int) int {

	leftHistogram := make([]int, len(histogram))
	leftMax := histogram[0]
	for i := 0; i < len(histogram); i++ {
		leftMax = Max(leftMax, histogram[i])
		leftHistogram[i] = leftMax
	}

	sum := 0
	rightMax := histogram[len(histogram)-1]
	for i := len(histogram) - 1; i >= 0; i-- {
		rightMax = Max(rightMax, histogram[i])
		secondTallest := Min(rightMax, leftHistogram[i])

		if secondTallest > histogram[i] {
			sum += secondTallest - histogram[i]
		}
	}

	//fmt.Println(sum)
	return sum
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
