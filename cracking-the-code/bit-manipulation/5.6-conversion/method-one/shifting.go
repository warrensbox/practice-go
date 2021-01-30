package main

import "fmt"

func main() {

	a := uint32(0b101011)
	b := uint32(0b111110)
	fmt.Println(bitSwapRequired(a, b))

}

func bitSwapRequired(a, b uint32) int {
	//SHIFTING METHOD
	count := 0
	for c := a ^ b; c != 0; c = c >> 1 {
		if (c & 1) == 1 {
			count++
		}
	}
	return count
}
