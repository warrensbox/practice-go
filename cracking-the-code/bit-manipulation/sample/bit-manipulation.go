package main

import (
	"fmt"
	"math/bits"
)

func main() {

	bin := byte(0b00010011)
	fmt.Printf("%08b %02x\n", bin, bin)
	x := byte(0x13)
	fmt.Printf("%08b %02x\n", x, x)

	//returns the minimum number of bits required to represent a value
	var a uint = 10
	fmt.Printf("Number:\t%d\n", a)
	fmt.Printf("bits.Len(%d) = %d\n", a, bits.Len(a))

	a++
	fmt.Printf("bits.Len(%d) = %d\n", a, bits.Len(a))

	fmt.Println("-----------------")

	//returns the number of bits set to 1 in a given number:
	var b uint = 31
	fmt.Printf("Number:\t%d\n", b)
	fmt.Printf("bits.OnesCount(%d) = %d\n", b, bits.OnesCount(b))

	a++
	fmt.Printf("bits.OnesCount(%d) = %d\n", b, bits.OnesCount(b))

	fmt.Println("-----------------")

	var c uint8 = 4
	fmt.Printf("Number:\t%4d\n", c)
	fmt.Printf("Leading zeros:\t%d\n", bits.LeadingZeros8(c))
	fmt.Printf("Trailing zeros:\t%d\n", bits.TrailingZeros8(c))
	fmt.Printf("Reversed:\t%b\n", bits.Reverse8(c))

	fmt.Println("-----------------")

	var d uint8 = 10
	d |= 5
	fmt.Printf("%b\n", d)

	fmt.Println("-----------------")

	e := 10
	e |= 5
	fmt.Printf("%b\n", e)

	fmt.Println("-----------------")

	var f byte = 0x0F
	fmt.Printf("%08b\n", f)
	fmt.Printf("%08b\n", ^f)

	fmt.Println("-----------------")

	var g byte = 00000001
	fmt.Printf("%08b\n", g)
	fmt.Printf("%08b\n", ^g)
	shift_left := g << 5
	fmt.Println(shift_left)
	fmt.Printf("%08b\n", shift_left)

	fmt.Println("-------Set Bit----------")
	/* Set Bit */
	var h uint8 = 0x0F
	fmt.Printf("Original\t%08b\n", h)
	fmt.Printf("Set\t%08b\n", SetBit(h, 3))

	fmt.Println("-------Clear Bit----------")
	/* Clear Bit */
	var i uint8 = 0x0F
	fmt.Printf("Original\t%08b\n", i)
	fmt.Printf("Clear\t%08b\n", ClearBit(i, 0))

	fmt.Println("-------Flip Bit----------")
	/* Set Bit */
	var bitwisenot byte = 0x0F
	fmt.Printf("Original\t%08b\n", bitwisenot)
	fmt.Printf("Not\t%08b\n", ^bitwisenot)
	fmt.Printf("Clear\t%08b\n", Flip(bitwisenot, 1))

	fmt.Println("-------ALL Ones----------")
	//allOnes := ^0
	fmt.Printf("%08b\n", 1<<5)

}

//func Set(b, flag Bits) Bits
func SetBit(x, position uint8) uint8 {
	var mask uint8
	mask = 1 << position
	return x | mask
}

//func ClearBit(b, flag Bits) Bits
func ClearBit(x, position uint8) uint8 {
	var mask uint8
	mask = 1 << position
	return x & ^mask
}

//func Flip(b, flag Bits) Bits
func Flip(x, position uint8) uint8 {
	var mask uint8
	mask = 1 << position
	return x ^ mask
}
