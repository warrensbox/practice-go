package main

import "fmt"

func main() {
	/* label        0   1   2   3   4   5   6   7 */
	//input := []int{30, 15, 60, 75, 45, 15, 15, 45}
	/*label        0  1  2  3  4 */
	input := []int{5, 7, 4, 9, 4}
	fmt.Println(maxMinute(input))
}

func maxMinute(arr []int) int {
	return maxMinuteRecur(0, arr)
}

func maxMinuteRecur(index int, arr []int) int {

	if index >= len(arr) {
		return 0
	}

	fmt.Println("index1", index)

	/*best with this reservation */
	bestWith := arr[index] + maxMinuteRecur(index+2, arr)

	fmt.Println("bestWith", bestWith)
	fmt.Println("index2", index)
	/*best without this reservation */
	bestWithOut := maxMinuteRecur(index+1, arr)
	fmt.Println("bestWithOut", bestWithOut)
	fmt.Println("index3", index)

	fmt.Printf("bestWith: %v, bestWithOut: %v \n", bestWith, bestWithOut)
	fmt.Println("_ _ _ _ _ _ _ _ _ _ _ _ _ _ _")
	fmt.Println("")
	/*return best of this subarray, starting from index */
	return Max(bestWith, bestWithOut)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
