package main

import "fmt"

/*
goal : put zero at front
constrains : in place
sol 1 : Iterate through array, keep loop and index pointer
*/

func main() {

	zero := []int{0, 10, 20, 0, 59, 63, 0, 88, 0}
	fmt.Println(Zero(zero))
}

func Zero(zero []int) []int {

	index := 0
	pointer := 0

	for pointer < len(zero) {

		if zero[pointer] == 0 {
			zero[index], zero[pointer] = zero[pointer], zero[index]
			index++
		}
		pointer++
	}
	return zero
}
