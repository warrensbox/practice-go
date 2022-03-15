package main

import "fmt"

func main() {

	fmt.Println(steps(10))

}

func steps(numOfdoors int) []int {

	doors := make([]int, numOfdoors) //0-9

	for step := 0; step < numOfdoors; step++ {
		doors[step] = 1
	}

	for step := 2; step < numOfdoors; {
		for i := step - 1; i < numOfdoors; i += step { //
			bit := doors[i]
			newBit := bit ^ 1
			doors[i] = newBit
		}
		step++
	}

	return doors

}
