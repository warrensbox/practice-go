package main

import "fmt"

func main() {

	a := uint32(0b101010)
	fmt.Printf("%b\n", swabOddEvenBits(a))

}

func swabOddEvenBits(a uint32) uint32 {

	return (((a & 0xaaaaaaaa) >> 1) | ((a & 0x55555555) << 1))
}
