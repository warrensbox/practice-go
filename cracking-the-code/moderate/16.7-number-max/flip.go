package main

import "fmt"

func main() {

	fmt.Println(getMaxNaive(24, 15))

}

// flips a 1 -> 0 AND 0 -> 1
func flip(bit int) int {
	//fmt.Printf("%b\n", bit)
	//fmt.Printf("%b\n", 1^bit)
	return 1 ^ bit
}

func sign(a int) int {
	// fmt.Println(a)
	// fmt.Printf("%b\n", a)
	// fmt.Println((a >> 31))
	// fmt.Printf("%b\n", a>>31)
	// fmt.Println(0x1)
	return flip((a >> 31) & 0x1)
}

func getMaxNaive(a, b int) int {
	k := sign(a - b)
	//fmt.Printf("%b\n", k)
	q := flip(k)
	return a*k + b*q
}
