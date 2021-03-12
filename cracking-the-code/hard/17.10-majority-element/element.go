package main

import "fmt"

func main() {
	arr := []int{3, 1, 7, 1, 1, 7, 7, 3, 7, 7, 7}
	candidate := getCandidate(arr)
	if validate(arr, candidate) {
		fmt.Println(candidate)
	} else {
		fmt.Println(-1)
	}
}

func getCandidate(array []int) int {
	majority := 0
	count := 0
	for _, val := range array {
		if count == 0 {
			majority = val
		}

		if majority == val {
			count++
		} else {
			count--
		}
	}

	return majority
}

func validate(array []int, n int) bool {
	count := 0
	for _, val := range array {
		if val == n {
			count++
		}
	}

	return count > len(array)/2
}
