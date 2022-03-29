package main

func dailyTemperatures(temperatures []int) []int {

	res := make([]int, len(temperatures))
	stack := []int{}
	for i, val := range temperatures {

		for len(stack) > 0 && val > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1] //the index
			stack = stack[:len(stack)-1]
			diff := i - top //diff of index
			res[top] = diff
		}
		stack = append(stack, i)
	}

	return res
}
