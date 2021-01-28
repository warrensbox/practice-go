package main

import "fmt"

func main() {

	j := 6
	i := 2
	var N uint32 = uint32(0b10000011000)
	fmt.Printf("N:\t%08b\n", N)

	var M uint32 = uint32(0b10011)
	fmt.Printf("M:\t%b\n", M)

	var K uint32 = uint32(0b0)
	allOnes := ^K
	left := allOnes << (j + 1)
	fmt.Printf("left:\t%08b\n", left)

	right := uint32((1 << i) - 1)
	fmt.Printf("right:\t%08b\n", right)

	mask := left | right
	fmt.Printf("mask:\t%08b\n", mask)

	n_cleared := N & uint32(mask)
	fmt.Printf("n_cleared:\t%08b\n", n_cleared)

	m_shifted := M << i
	fmt.Printf("m_shifted:\t%08b\n", m_shifted)

	fmt.Printf("FINAL:\t%08b\n", n_cleared|m_shifted)
}
