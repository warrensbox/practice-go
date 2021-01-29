package main

import "fmt"

func main() {

	bin := uint32(0b110111001111)
	//bin := uint64(0b00000000000)
	//bin := uint32(0b11111111111)
	//flipBit(bin)
	fmt.Println(flipBit(bin))

}

func flipBit(a uint32) int {

	if ^a == 0 {
		return -1
	}

	currentLength := 0
	previousLength := 0
	maxLength := 1 //We can always have a sequence of at least one 1
	for a != 0 {
		//fmt.Printf("(a & 1): %b\n", (a & 1))
		if (a & 1) == 1 { //Current bit is a 1
			currentLength++
		} else if (a & 1) == 0 { //Current bit is a 0
			//update to 0 (if next bit is 0) or currentLength (if next bit is 1)
			if (a & 2) == 0 { //reset
			} else {
				fmt.Println("going here")
				previousLength = currentLength
			}
			currentLength = 0
		}
		//fmt.Printf("maxLength: %d\n", maxLength)
		maxLength = Max(previousLength+currentLength+1, maxLength)
		//fmt.Printf("maxLength: %d\n", maxLength)
		a = a >> 1
		//fmt.Println("----")
	}
	return maxLength
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
