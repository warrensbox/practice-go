package main

import "fmt"

/*
goal : reverse string
sol 1: copy in to new array
sol 2: reverse in place

*/

func main() {
	str := "abcd"
	fmt.Println(reverse(str))
}

func reverse(val string) string {

	str := []byte(val)
	for i := 0; i < len(str)/2; i++ {
		str[i], str[len(str)-i-1] = str[len(str)-i-1], str[i]

	}

	return string(str)
}
