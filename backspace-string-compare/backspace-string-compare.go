package main

import (
	"fmt"
)

func main() {

	// S := "a##c"
	// T := "#a#c"

	S := "ab#c"
	T := "ad#c"

	output := backspaceCompare(S, T)
	fmt.Println(output)
}

func backspaceCompare(S string, T string) bool {

	if build(S) == build(T) {
		return true
	}
	return false
}

func build(str string) string {

	var stack []rune

	for _, char := range str {

		if string(char) != "#" {
			stack = append(stack, char)

		} else if len(stack) > 0 {
			n := len(stack) - 1 // Top element
			stack = stack[:n]   // Pop
		}
	}

	return string(stack)
}
