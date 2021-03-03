package main

import "fmt"

func main() {

	fmt.Println(factorial(10))
	fmt.Println(countFactZeros(10))
}

func factorial(x int) int {
	//base case
	if x == 1 {
		return x
	}

	return x * factorial(x-1)
}

func countFactZeros(num int) int {
	count := 0
	if num < 0 {
		return -1
	}
	for i := 5; num/i > 0; i *= 5 {
		count += num / i
	}
	return count
}
