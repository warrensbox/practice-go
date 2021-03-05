package main

import "fmt"

func main() {
	//fmt.Println(minus(5, 4))
	//fmt.Println(multiply(4, 6))
	fmt.Println(divide(15, 3))
}

func negate(a int) int {
	neg := 0
	var newSign int
	if a < 0 {
		newSign = 1
	} else {
		newSign = -1
	}
	for a != 0 {
		neg += newSign
		a += newSign
	}

	return neg
}

func negate2(a int) int {
	neg := 0
	var newSign int
	if a < 0 {
		newSign = 1
	} else {
		newSign = -1
	}
	delta := newSign
	for a != 0 {
		differenceSigns := (a+delta > 0) != (a > 0)
		if (a+delta != 0) && differenceSigns {
			delta = newSign
		}
		neg += delta
		a += delta
		delta += delta
	}
	return neg
}

func minus(a, b int) int {
	//fmt.Println(negate(b))
	return a + negate(b)
}

func multiply(a, b int) int {

	if a < b {
		multiply(b, a) //we want the smaller value for "b" - it'll be faster
	}
	mul := 0

	for i := 1; i <= Abs(b); i++ {
		mul += a // if a is negative (we just keep adding negative values)
	}
	fmt.Println(mul)
	if a < 0 {
		return negate(mul)
	}
	return mul
}

func divide(a, b int) int {

	// a = bx
	x := 0
	product := 0
	absA := Abs(a)
	absB := Abs(b)

	for product+absB <= absA {
		product += absB
		x++
	}

	if (a < 0 && b < 0) || (a > 0 && b > 0) {
		return x
	} else {
		return negate(x)
	}
}

func Abs(num int) int {
	if num < 0 {
		return negate(num)
	}
	return num
}
