package main

import "fmt"

func main() {
	//minProduct(8, 7)
	fmt.Println(minProduct(7, 8))
}

func minProduct(a, b int) int {

	//bigger = a <b ? b : a
	//smaller = a <b ? a : b
	smaller := b
	bigger := a
	if b > a {
		bigger = b
		smaller = a
	}
	return minProductHelper(smaller, bigger)
}

func minProductHelper(smaller, bigger int) int {

	fmt.Println(smaller)
	fmt.Println(bigger)
	//base case
	if smaller == 0 {
		return 0
	} else if smaller == 1 {
		return bigger
	}

	//compute half. if uneven, compute other half. If even, double it
	s := smaller >> 1
	fmt.Printf("%08b\n", s)
	halfProd := minProductHelper(s, bigger)

	if smaller%2 == 0 {
		return halfProd + halfProd
	} else {
		return halfProd + halfProd + bigger
	}
}
