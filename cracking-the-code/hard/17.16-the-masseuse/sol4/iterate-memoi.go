package main

import "fmt"

func main() {
	/*label        0  1  2  3  4 */
	input := []int{5, 7, 4, 9, 4}
	fmt.Println(maxMinute2(input))
}

func maxMinute2(arr []int) int {
	//allocate 2 extra slots in the array so we dont have to do bounds
	memo := make([]int, len(arr)+2)
	memo[len(arr)] = 0
	memo[len(arr)+1] = 0
	for i := len(arr) - 1; i >= 0; i-- {
		include := arr[i] + memo[i+2]
		exclude := memo[i+1]
		memo[i] = Max(include, exclude)
	}

	return memo[0]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
