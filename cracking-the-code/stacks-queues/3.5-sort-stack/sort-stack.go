package main

import "fmt"

type Stack struct {
	stk []int
}

//Sort : sorting the stack
func (s *Stack) Sort() *Stack {

	r := Stack{}

	//length := len(s.stk)
	for len(s.stk) != 0 {
		tmp := s.Pop()
		fmt.Println("tmp", tmp)                 //step 1
		for len(r.stk) != 0 && r.Peek() > tmp { //step 2
			s.Push(r.Pop())
			fmt.Println("exchange")
			fmt.Println(s.stk)
		}
		r.Push(tmp) //step 3
		fmt.Println("r", r.stk)
	}

	return &r
}

//Pop : popping
func (s *Stack) Pop() int {

	n := len(s.stk) - 1
	top := s.stk[n]
	s.stk = s.stk[:n]

	return top
}

//Peek : peeking
func (s *Stack) Peek() int {

	n := len(s.stk) - 1
	top := s.stk[n]

	return top
}

//Push : push into stack
func (s *Stack) Push(item int) {

	s.stk = append(s.stk, item)
}

func main() {

	s := Stack{}

	s.Push(1)
	s.Push(4)
	s.Push(3)
	s.Push(5)

	fmt.Println(s.stk)
	r := s.Sort()
	fmt.Printf("%v", *r)
}
