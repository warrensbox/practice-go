package main

import "fmt"

func (queue *Queue) enqueue(node int) {

	//push stack bottomup
	queue.BottomUp.Push(node)
	for len(queue.BottomUp.Arr) > 0 {
		top := queue.BottomUp.Pop()
		queue.TopDown.Push(top)
	}
}

func (queue *Queue) dequeue() int {

	if len(queue.TopDown.Arr) > 0 {
		top := queue.TopDown.Pop()
		return top
	}
	return -1
}

type Queue struct {
	BottomUp Stack
	TopDown  Stack
}

type Stack struct {
	Arr []int
}

func (stack *Stack) Push(item int) {
	stack.Arr = append(stack.Arr, item)
}

func (stack *Stack) Pop() int {
	top := -1
	if len(stack.Arr) > 0 {
		top = stack.Arr[len(stack.Arr)-1]
		stack.Arr = stack.Arr[:len(stack.Arr)-1]
	}

	return top
}

func main() {

	var queue Queue
	queue.enqueue(1)
	queue.enqueue(2)
	queue.enqueue(3)

	fmt.Println(queue.dequeue())
	fmt.Println(queue.dequeue())

}
