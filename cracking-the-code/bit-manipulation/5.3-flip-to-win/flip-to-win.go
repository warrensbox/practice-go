package main

import (
	"fmt"
)

func main() {

	flipBit(15)
	//fmt.Println(bin)

}

func flipBit(a int) int {

	fmt.Printf("%08b\n", ^a)

	var K uint32 = uint32(0b1111)
	fmt.Printf("%08b\n", ^K)

	return 1
}
