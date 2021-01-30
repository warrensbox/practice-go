package main

import "fmt"

func main() {

	a := uint32(0b101011)
	b := uint32(0b111110)
	fmt.Println(bitSwapRequired(a, b))

}

func bitSwapRequired(a, b uint32) int {
	//CLAEARING METHOD
	count := 0
	for c := a ^ b; c != 0; c = c & (c - 1) {
		count++
	}
	return count
}
