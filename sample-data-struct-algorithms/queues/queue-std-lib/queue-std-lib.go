package main

import (
	"container/list"
	"fmt"
)

func main() {

	queue := list.New() //instantiate queue here
	recursion(1, queue)

	// Loop over container list.
	for list := queue.Front(); list != nil; list = list.Next() {
		fmt.Printf("%v -> ", list.Value)
	}

	queue2 := list.New()

	queue2.PushBack("Hello ") // Enqueue
	queue2.PushBack("world!")

	for queue2.Len() > 0 {
		e := queue2.Front() // First element
		fmt.Print(e.Value)

		queue2.Remove(e) // Dequeue
	}
}

func recursion(num int, q *list.List) int {

	if num > 10 {
		return num
	}
	q.PushBack(num)
	return recursion(num+1, q)

}
