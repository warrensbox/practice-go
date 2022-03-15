package main

import "fmt"

/*
goal: combination of elements that total target sum


*/

func main() {
	target := 7
	candidates := []int{2, 3, 6, 7}
	results := [][]int{}
	sum(candidates, target, []int{}, &results)

	fmt.Println(results)
}

func sum(candidates []int, target int, stack []int, result *[][]int) {

	if target == 0 {
		// make a deep copy of the current combination
		*result = append(*result, stack)
		fmt.Println("reached end")
		fmt.Println("*result", *result)
		fmt.Println()
		return
	} else if target < 0 {
		// exceed the scope, stop exploration.
		fmt.Println("ops returning")
		return
	}
	//st := []int{}
	for i := 0; i < len(candidates); i++ {
		fmt.Println("target", target)
		fmt.Println("candidates[i]", candidates[i])
		if target >= candidates[i] {
			// add the number into the combination
			stack = append(stack, candidates[i])
			sum(candidates, target-candidates[i], stack, result)
			// backtrack, remove the number from the combination
			stack = stack[:len(stack)-1]
			fmt.Println("stack", stack)
		}
	}
	fmt.Println("stack", stack)
	fmt.Println("ops returning from end")
	return
}
