package main

import "fmt"

func main() {

	nums := []int{12, 345, 2, 6, 7896}
	output := findNumbers(nums)

	fmt.Println(output)
}

func findNumbers(nums []int) int {
	count := 0

	for _, val := range nums {
		if (val > 9 && val < 100) || (val > 999 && val < 10000) {
			count++
		}
	}

	return count
}
