package main

import "fmt"

func main() {

	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 5, 3, 2, 1}
	output := validateStackSequences(pushed, popped)

	fmt.Println(output)
}

func validateStackSequences(pushed []int, popped []int) bool {

	var stack []int
	j := 0
	for _, val := range pushed {

		stack = push(stack, val)

		for len(stack) != 0 && j < len(pushed) && peek(stack) == popped[j] {

			stack = pop(stack)
			j++
		}
	}

	fmt.Println(j)
	fmt.Println(len(pushed))

	return false

}

func push(stack []int, charAt int) []int {

	stack = append(stack, charAt) // Push

	return stack
}

func peek(stack []int) int {

	n := len(stack) - 1 // Top element
	val := stack[n]

	return val
}

func pop(stack []int) []int {

	n := len(stack) - 1 // Top element

	stack = stack[:n] // Pop

	return stack
}
