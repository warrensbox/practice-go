package main

import (
	"fmt"
)

func main() {

	stack := Stack{}
	stack.Push(4)
}

type Stack struct {
	stack []interface{}
}

func (s *Stack) Push(item interface{}) {

	s.stack = append(s.stack, item)

}

func (s *Stack) pop() interface{} {

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

func (s *Stack) Peek() interface{} {

	//nothing ot return
	if len(s.stack) == 0 {
		fmt.Println("Nothing to return")
		return -1
	}

	n := len(s.stack) - 1
	top := s.stack[n]

	return top

}

func (s *Stack) Len() int {

	return len(s.stack)

}
