package main

import "fmt"

func main() {
	arr := []int{-2, -1, 0, 3, 5, 6, 7, 9, 13, 14}
	fmt.Println(sumPairs(arr, 8))
}

func sumPairs(arr []int, sum int) [][]int {
	hash := make(map[int]bool)
	pair := [][]int{}

	for _, val := range arr {
		complement := sum - val
		if hash[complement] {
			pair = append(pair, []int{val, complement})
		} else {
			hash[val] = true
		}
	}

	return pair
}
