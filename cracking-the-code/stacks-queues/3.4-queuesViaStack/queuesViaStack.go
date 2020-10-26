package main

import "fmt"

///everytime you push you create a temporary buffer

type Stack struct {
	stk       []int
	stkBuffer []int
}

//enqueue
func (s *Stack) enqueue(item int) {

	s.stk = append(s.stk, item)
	//peekStk := s.peeking()

	s.stkBuffer = nil
	for i := len(s.stk) - 1; i >= 0; i-- {
		s.stkBuffer = append(s.stkBuffer, s.stk[i])
	}

}

func (s *Stack) peeking() int {

	//put error checks
	n := len(s.stk) - 1
	topItem := s.stk[n]

	return topItem
}

func (s *Stack) pop() int {

	//put error checks
	n := len(s.stk) - 1
	topItem := s.stk[n]
	s.stk = s.stk[:n]

	return topItem
}

func (s *Stack) dequeue() int {

	//put error checks
	n := len(s.stkBuffer) - 1
	topItemBuffer := s.stk[n]
	s.stkBuffer = s.stkBuffer[:n]

	s.stk = nil
	for i := len(s.stkBuffer) - 1; i >= 0; i-- {
		s.stk = append(s.stk, s.stkBuffer[i])
	}

	return topItemBuffer

}

func main() {

	s := Stack{}

	s.enqueue(1)
	s.enqueue(2)
	s.enqueue(3)
	s.enqueue(4)
	fmt.Println(s.stk)
	fmt.Println(s.stkBuffer)
	s.dequeue()
	fmt.Println(s.stk)
	fmt.Println(s.stkBuffer)
}
