package main

import "fmt"

func main() {
	fmt.Println(countWays1(3))
	fmt.Println(CountWays2(3))
}

func countWays1(n int) int {
	if n < 0 {
		return 0
	} else if n == 0 {
		return 1
	}
	return countWays1(n-1) + countWays1(n-2) + countWays1(n-3)
}

func CountWays2(n int) int {
	memo := make([]int, n+1)
	memo[0] = 1
	ways := countWays2(n, memo)
	return ways
}

func countWays2(n int, memo []int) int {
	if n < 0 {
		return 0
	} else if n == 0 {
		return 1
	} else if memo[n] > 0 {
		return memo[n]
	} else {
		memo[n] = countWays2(n-1, memo) + countWays2(n-2, memo) + countWays2(n-3, memo)
	}

	return memo[n]
}
