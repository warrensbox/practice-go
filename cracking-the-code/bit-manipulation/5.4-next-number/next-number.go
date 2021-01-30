package main

import "fmt"

func main() {

	bin := uint32(0b11011001111100)
	fmt.Println(getNext(bin))

}

func getNext(n uint32) uint32 {

	//compute c0 and c1
	c := n
	c0 := 0
	c1 := 0

	for ((c & 1) == 0) && (c != 0) {
		c0++
		c = c >> 1
	}

	for (c & 1) == 1 {
		c1++
		c = c >> 1
	}

	p := c1 + c0

	//flip right most non-traling zero
	n = n | (1 << p)

	fmt.Printf("%b\n", n)

	//clear all bits to the right of p
	a := uint32(1 << p)
	b := a - 1
	mask := ^b
	n = n & mask

	fmt.Printf("%b\n", n)

	//insert (c1-1) ones on the right (c1-1 because remember the 1 we flip earlier)
	a = 1 << (c1 - 1)
	b = a - 1
	n = n | b

	fmt.Printf("%b\n", n)
	return n
}

func getPrev(n uint32) uint32 {

	//compute c0 and c1
	temp := n
	c0 := 0
	c1 := 0

	for (temp & 1) == 1 {
		c1++
		temp = temp >> 1
	}

	if temp == 0 {
		return n
	}

	for ((temp & 1) == 0) && (temp != 0) {
		c0++
		temp = temp >> 1
	}

	p := c1 + c0 //position of the rightmost non-traling one

	//clears from bit p onwards
	a := ^0
	b := uint32(a) << (p + 1)
	n = n & b

	a = 1 << (c1 + 1)
	b = uint32(a) - 1
	c := b << (c0 - 1)

	n = n | c

	fmt.Printf("%b\n", n)
	return n
}
