package main

import (
	"fmt"
)

func main() {

	fmt.Println(sumForLoop(4))
	fmt.Println(sumRecursion(4))
}

func sumForLoop(n int) int {
	i := 1
	prev := 0
	for i <= n {
		prev = prev + i
		i++
	}
	return prev
}

func sumRecursion(n int) int {
	if n == 0 {
		return 0
	}
	return n + sumRecursion(n-1)
}
