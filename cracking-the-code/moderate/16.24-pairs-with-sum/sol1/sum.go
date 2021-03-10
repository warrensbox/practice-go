package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{-2, -1, 0, 3, 5, 6, 7, 9, 13, 14}
	fmt.Println(sumPairs(arr, 8))
}

func sumPairs(arr []int, sum int) [][]int {
	pair := [][]int{}

	sort.Ints(arr)
	first := 0
	last := len(arr) - 1

	for first < last {
		s := arr[first] + arr[last]
		if s == sum {
			pair = append(pair, []int{arr[first], arr[last]})
			first++
			last--
		} else {
			if s < sum {
				first++
			} else {
				last--
			}
		}
	}

	return pair
}
