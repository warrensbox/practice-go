package main

import (
	"fmt"
	"strings"
)

func main() {

	input := "/../test/../g"
	output := simplifyPath(input)
	fmt.Println(output)
}

func simplifyPath(path string) string {

	var stack []string

	directory := strings.Split(path, "/")
	fmt.Println(len(directory))
	for _, dir := range directory {
		fmt.Println("cd", dir)
		fmt.Println("stack1", stack)
		if dir == ".." {
			if len(stack) > 0 {
				fmt.Println("pop", dir)
				_, stack = pop(stack)
			}
		} else if dir != "" {
			if dir != "." {
				fmt.Println("push", dir)
				stack = push(stack, dir)
			}
		}
		fmt.Println("stack2", stack)
	}

	res := ""
	for _, s := range stack {

		if string(s) == "" {
			continue
		} else {
			fmt.Println("s", s)
			res = res + "/" + s
		}

		fmt.Println("res", res)
	}

	fmt.Println(res)
	if res == "" {
		res = "/"
	}
	return res

}

func push(stack []string, charAt string) []string {

	stack = append(stack, charAt) // Push

	return stack
}

func pop(stack []string) (string, []string) {
	val := ""
	n := len(stack) - 1 // Top element
	val = stack[n]
	stack = stack[:n] // Pop

	fmt.Println(len(stack))
	return val, stack
}
