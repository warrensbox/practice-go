package main

import (
	"fmt"
	"sort"
)

func main() {

	arr := [][]int{{1, 3}, {8, 10}, {2, 6}, {15, 18}}

	fmt.Println(arr)

	output := merge(arr)

	fmt.Println(output)
}

func merge(intervals [][]int) [][]int {

	if len(intervals) == 0 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	fmt.Println(intervals)

	result := [][]int{intervals[0]}
	last := 0
	for i := 1; i < len(intervals); i++ {

		if result[last][1] >= intervals[i][0] {

			result[last][1] = max(result[last][1], intervals[i][1])
		} else {
			result = append(result, intervals[i])
			last++
		}
	}
	return result
}

func max(x, y int) int {

	if x > y {
		return x
	}

	return y
}
