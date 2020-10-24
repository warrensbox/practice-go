package main

import "fmt"

func main() {

	s := Stack{}
	s.push(5)
	fmt.Println(s.stack)
	s.push(3)
	fmt.Println(s.stack)
	s.push(1)
	fmt.Println(s.stack)
	s.push(10)
	fmt.Println(s.stack)
	s.push(11)
	fmt.Println(s.stack)
	s.push(12)
	fmt.Println(s.stack)
	s.push(14)
	fmt.Println(s.stack)
	s.push(15)
	fmt.Println(s.stack)
	s.push(16)
	fmt.Println(s.stack)
	s.pop()
	fmt.Println(s.stack)
	s.pop()
	fmt.Println(s.stack)
	s.pop()
	fmt.Println(s.stack)
}

type Stack struct {
	stack [][]int
}

func (o *Stack) push(item int) {

	if len(o.stack) == 0 {
		topStack := []int{}
		topStack = append(topStack, item)
		o.stack = append(o.stack, topStack)
		return
	}

	topOuter := len(o.stack) - 1

	if len(o.stack[topOuter]) == 4 {
		fmt.Println("its 4")
		topStack := []int{}
		topStack = append(topStack, item)
		o.stack = append(o.stack, topStack)
	} else {
		o.stack[topOuter] = append(o.stack[topOuter], item)
	}

}

func (o *Stack) pop() {
	if len(o.stack) == 0 {
		fmt.Print("nothing to pop -  empty stack")
		return
	}

	topOuter := len(o.stack) - 1
	o.stack[topOuter] = o.stack[topOuter][:len(o.stack[topOuter])-1]

	if len(o.stack[topOuter]) == 0 {
		fmt.Println("its zero")
		o.stack = o.stack[:topOuter]
	}
}
