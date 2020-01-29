package main

import (
	"fmt"
)

func main() {

	input := 5
	output := climbStairs(input)
	fmt.Println(output)
}

func climbStairs(n int) int {

	fmt.Println("n", n)

	if n == 0 || n == 1 || n == 2 {
		return n
	}

	mem := make([]int, n+1)
	mem[1] = 1
	mem[2] = 2
	for i := 3; i <= n; i++ {
		fmt.Println("i", i)
		fmt.Println("mem[i-1]", mem[i-1])
		fmt.Println("mem[i-2]", mem[i-2])
		mem[i] = mem[i-1] + mem[i-2]
	}

	return mem[n]
}
