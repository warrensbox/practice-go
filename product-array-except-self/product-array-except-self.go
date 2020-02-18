package main

import (
	"fmt"
)

func main() {

	input := []int{5, 7, 7, 8, 8, 10}
	target := 8

	output := searchRange(input, target)
	fmt.Println(output)
}
