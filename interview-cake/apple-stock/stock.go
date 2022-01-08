package main

import "fmt"

func main() {

	stock := []int{10, 7, 5, 8, 11, 9}

	smallestSoFar := stock[0]

	max := 0

	for i := 1; i < len(stock); i++ {
		if stock[i] > smallestSoFar {
			max = Max(max, stock[i]-smallestSoFar)
		} else {
			smallestSoFar = stock[i]
		}
	}

	fmt.Println(max)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
