package main

import (
	"fmt"
	"math"
)

func main() {

	price := []int{1, 3, 4, 7, 3}
	fmt.Println(cutRod(price, 3))
	fmt.Println(cutRod2(price, 3))
}

func cutRod(price []int, n int) int {
	if n <= 0 {
		return 0
	}
	max := math.MinInt16

	for i := 0; i < n; i++ {
		max = Max(max, price[i]+cutRod(price, n-i-1))

	}

	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//memoization
func cutRod2(price []int, n int) int {

	val := make([]int, n+1)
	val[0] = 0

	for i := 1; i <= n; i++ {
		max := math.MinInt16
		for j := 0; j < i; j++ {
			max = Max(max, price[j]+val[i-j-1])
		}
		val[i] = max
	}

	return val[n]
}
