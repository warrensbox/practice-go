package main

import "fmt"

func main() {
	//minProduct(8, 7)
	fmt.Println(MinProduct(7, 8))
}

func MinProduct(a, b int) int {

	//bigger = a <b ? b : a
	//smaller = a <b ? a : b
	smaller := b
	bigger := a
	if b > a {
		bigger = b
		smaller = a
	}
	memo := make([]int, smaller+1)
	return minProduct(smaller, bigger, memo)
}

func minProduct(smaller, bigger int, memo []int) int {

	fmt.Println(smaller)
	fmt.Println(bigger)
	//base case
	if smaller == 0 {
		return 0
	} else if smaller == 1 {
		return bigger
	} else if memo[smaller] > 0 {
		return memo[smaller]
	}

	//compute half. if uneven, compute other half. If even, double it
	s := smaller >> 1 //divide by 2
	fmt.Printf("%08b\n", s)
	side1 := minProduct(s, bigger, memo) //compute half
	side2 := side1
	if smaller%2 == 1 {
		side2 = minProduct(smaller-s, bigger, memo)
	}

	fmt.Println("side1", side1)
	fmt.Println("side2", side2)

	memo[smaller] = side2 + side1
	return memo[smaller]
}
