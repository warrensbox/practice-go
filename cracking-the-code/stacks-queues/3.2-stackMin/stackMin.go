package main

import "fmt"

func main() {
	stack := Stack{}
	fmt.Println(stack.Push(3))
	fmt.Println("ori", stack.original)
	fmt.Println(stack.Push(2))
	fmt.Println("ori", stack.original)
	fmt.Println(stack.Push(4))
	fmt.Println("ori", stack.original)
	fmt.Println(stack.Push(1))
	fmt.Println("ori", stack.original)
	fmt.Println(stack.Push(5))
	fmt.Println("ori", stack.original)
	fmt.Println(stack.Pop()) //5
	fmt.Println("ori", stack.original)
	fmt.Println(stack.Pop()) //1
	fmt.Println("ori", stack.original)
	fmt.Println(stack.Pop()) //4
	fmt.Println("ori", stack.original)
	fmt.Println(stack.Pop()) //2
	fmt.Println("ori", stack.original)
	fmt.Println("Last Pop")
	fmt.Println(stack.Pop()) //3
	fmt.Println("ori", stack.original)
	fmt.Println(stack.Pop()) //nothing
	fmt.Println("ori", stack.original)
}

type Stack struct {
	original []int
	minimum  []int
}

func (s *Stack) Push(item int) int {

	if len(s.original) == 0 {
		s.original = append(s.original, item)
		s.minimum = append(s.minimum, item)
		return item
	}

	n := len(s.minimum) - 1
	min := s.minimum[n]

	if item < min {
		s.minimum = append(s.minimum, item)
		min = item
	}

	s.original = append(s.original, item)

	return min
}

func (s *Stack) Pop() int {

	if len(s.original) == 0 {
		// nothing to Pop
		return -1
	}

	m := len(s.original) - 1 //peek ori
	topOri := s.original[m]

	n := len(s.minimum) - 1 //peek min
	topMin := s.minimum[n]

	if topMin == topOri {
		s.minimum = s.minimum[:n] //Pop min
	}

	s.original = s.original[:m] //Pop original

	if len(s.minimum) > 0 {
		n = len(s.minimum) - 1 //get new top min
		topMin = s.minimum[n]
	}

	return topMin
}
