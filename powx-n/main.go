package main

import "fmt"

func main() {
	/*
		Input: x = 2.00000, n = 10
		Output: 1024.00000
	*/
	fmt.Println(myPow(2.0, 10))
}

func myPow(x float64, n int) float64 {

	if n < 0 {
		n = -n
	}

	return pow(x, n)
}

func pow(x float64, n int) float64 {

	if n == 0 {
		return 1.0
	}

	half := pow(x, n/2)

	if n%2 == 0 {
		return half * half
	} else {
		return half * half * x
	}
}
