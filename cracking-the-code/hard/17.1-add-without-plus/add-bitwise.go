package main

import "fmt"

func main() {

	fmt.Println(sum(3, 4))
}

func sum(a, b int) int {

	for b != 0 {
		sum := a ^ b
		carry := (a & b) << 1
		a = sum
		b = carry
	}

	fmt.Println(a)
	return a
}
