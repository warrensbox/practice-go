package main

import "fmt"

func main() {
	/*label        0  1  2  3  4 */
	input := []int{5, 7, 4, 9, 4}
	fmt.Println(maxMinute2(input))
}

func maxMinute2(arr []int) int {

	twoPrev := 0
	onePrev := 0
	max := 0
	for i := len(arr) - 1; i >= 0; i-- {
		include := twoPrev + arr[i]
		exclude := onePrev
		max = Max(include, exclude)
		twoPrev = onePrev
		onePrev = max
	}

	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
