package main

import "fmt"

type Node struct {
	item interface{}
	next *Node
}

type Stack struct {
	n    int
	head *Node
}

func main() {

	stack := Stack{}

	stack.Push(3)
	stack.Push(4)
	stack.Pop()
	stack.Push(5)

	fmt.Println(stack.Size())
	stack.Traverse()
}

//Push - push to list
func (s *Stack) Push(item interface{}) {
	var newItem Node
	newItem.item = item
	current := s.head

	if current == nil {
		s.head = &newItem
		s.n++
	} else {
		newItem.next = s.head
		s.head = &newItem
		s.n++
	}

}

//Traverse - go thru stack
func (s *Stack) Traverse() {

	stack := s.head

	for stack != nil {
		fmt.Printf("%+v ->", stack.item)
		stack = stack.next
	}
}

func (s *Stack) Pop() interface{} {

	if s.IsEmpty() {
		return nil
	}

	first := s.head

	item := first.item
	fmt.Println("Pop", item)
	s.head = first.next
	s.n--

	return item
}

//IsEmpty - check is stack is empty
func (s *Stack) IsEmpty() bool {
	return s.head == nil
}

//Size - check is stack is empty
func (s *Stack) Size() int {
	return s.n
}
