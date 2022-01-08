package main

import "fmt"

/*
goal : find largest element in stack -  keep track of largest element after popping
- use 2 stacks
- main stack - to keep track of items push into stack
- secondary stack - to keep of largest elements

Main stack = [4,3,1,9,11,12,20,15] <- max(11)
Secondary stack = [4,9,11,12,20] <-
*/
func main() {

	var storage Storage
	storage.Push(4)
	storage.Push(3)
	storage.Push(1)
	storage.Push(9)
	storage.Push(11)
	fmt.Println(storage.GetMax())
	storage.Push(12)
	storage.Push(20)
	fmt.Println(storage.GetMax())
	storage.Push(15)
	fmt.Println(storage.GetMax())
	storage.Pop()
	fmt.Println(storage.GetMax())
	storage.Pop()
	fmt.Println(storage.GetMax())
}

func (storage *Storage) Push(item int) {

	//regular push to main storage
	storage.main.arr = append(storage.main.arr, item)

	if len(storage.secondary.arr) > 0 && storage.secondary.Peek() <= item {
		storage.secondary.Push(item)
	} else if len(storage.secondary.arr) == 0 {
		storage.secondary.Push(item)
	}
}

func (storage *Storage) Pop() int {
	top := -1
	if len(storage.main.arr) > 0 {
		top = storage.main.Pop()
	}
	if len(storage.secondary.arr) > 0 && top == storage.secondary.Peek() {
		storage.secondary.Pop()
	}

	return top
}

func (storage *Storage) GetMax() int { //similiar to peek
	if len(storage.secondary.arr) > 0 {
		return storage.secondary.Peek()
	}
	return -1
}

type Storage struct {
	main      Stack
	secondary Stack
}

type Stack struct {
	arr []int
}

func (stack *Stack) Push(item int) {

	stack.arr = append(stack.arr, item)
}

func (stack *Stack) Pop() int {

	if len(stack.arr) > 0 {
		top := stack.arr[len(stack.arr)-1]
		stack.arr = stack.arr[:len(stack.arr)-1]
		return top
	}

	return -1
}

func (stack *Stack) Peek() int {

	if len(stack.arr) > 0 {
		top := stack.arr[len(stack.arr)-1]
		return top
	}
	return -1
}
