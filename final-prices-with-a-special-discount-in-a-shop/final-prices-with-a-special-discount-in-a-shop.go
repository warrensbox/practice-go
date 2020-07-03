package main

import "fmt"

func main() {

	prices := []int{8, 4, 6, 2, 3}

	output := finalPrices(prices)

	fmt.Println(output)

}

func finalPrices(prices []int) []int {

	stack := []int{}

	for i, _ := range prices {

		for len(stack) > 0 && prices[peek(stack)] >= prices[i] {
			prices[peek(stack)] = prices[peek(stack)] - prices[i]
			fmt.Println("peek", peek(stack))

			stack = pop(stack)
		}

		stack = push(stack, i)
	}

	return prices
}

func push(stack []int, charAt int) []int {

	stack = append(stack, charAt) // Push

	return stack
}

func peek(stack []int) int {

	n := len(stack) - 1 // Top element

	return stack[n]
}

func pop(stack []int) []int {

	n := len(stack) - 1 // Top element

	stack = stack[:n] // Pop
	return stack
}
