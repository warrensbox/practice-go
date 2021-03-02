package main

import "fmt"

func main() {
	a := 4
	b := 5
	fmt.Printf("a:%v, b:%v\n", a, b)

	a = a - b
	b = a + b
	a = b - a
	fmt.Printf("a:%v, b:%v\n", a, b)
}
