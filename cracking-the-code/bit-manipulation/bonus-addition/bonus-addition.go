package main

import "fmt"

func main() {

	add(3, 1)
}

func add(a, b byte) byte {

	fmt.Printf("a: %04b\n", a)
	fmt.Printf("b: %04b\n", b)

	carry := a & b
	a = a ^ b

	for carry > 0 {
		b = carry << 1
		carry = (a & b)
		a = a ^ b
	}

	//fmt.Printf("total : %04b\n", total)
	fmt.Printf("total:%d\n", a)
	return a

}
