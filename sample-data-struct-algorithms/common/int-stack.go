package common

import "fmt"

//stack implementation
type IntStack struct {
	stack []int
}

func (s *IntStack) Push(item int) {

	s.stack = append(s.stack, item)

}

func (s *IntStack) Pop() int {

	//nothing ot return
	if len(s.stack) == 0 {
		fmt.Println("Nothing to return")
		return -1
	}

	n := len(s.stack) - 1
	top := s.stack[n]
	s.stack = s.stack[:n] //pop

	return top

}

func (s *IntStack) Peek() int {

	//nothing ot return
	if len(s.stack) == 0 {
		fmt.Println("Nothing to return")
		return -1
	}

	n := len(s.stack) - 1
	top := s.stack[n]

	return top

}

func (s *IntStack) Len() int {

	return len(s.stack)

}
