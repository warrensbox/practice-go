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
	memo := make([]int, len(arr))
	return maxMinuteRecur(0, arr, memo)
}

func maxMinuteRecur(index int, arr, memo []int) int {

	if index >= len(arr) {
		return 0
	}

	if memo[index] == 0 {
		/*best with this reservation */
		bestWith := arr[index] + maxMinuteRecur(index+2, arr, memo)
		/*best without this reservation */
		bestWithOut := maxMinuteRecur(index+1, arr, memo)
		/*return best of this subarray, starting from index */
		memo[index] = Max(bestWith, bestWithOut)
	}
	return memo[index]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
