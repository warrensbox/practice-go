package main

import (
	"fmt"
)

func main() {

	input := 5
	output := fib(input)
	fmt.Println(output)

	fmt.Println(fib2(input))
}

func fib(N int) int {

	if N == 0 {
		return 0
	}

	if N == 1 {
		return 1
	}

	return fib(N-1) + fib(N-2)
}

func fib2(N int) int {

	//use memoization
	memo := make(map[int]int)
	memo[0] = 0
	memo[1] = 1
	for i := 2; i <= N; i++ {
		memo[i] = memo[i-1] + memo[i-2]
		fmt.Println(memo)
	}

	return memo[N]
}
