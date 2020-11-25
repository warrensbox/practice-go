package main

import "fmt"

func main() {

	queue := Queue{}

	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)
	queue.Dequeue()
	queue.Dequeue()
	queue.Traverse()
}

type Node struct {
	item interface{}
	next *Node
}

type Queue struct {
	n    int
	head *Node
	tail *Node
}

//Enqueue - Enqueue to list
func (q *Queue) Enqueue(item interface{}) {

	var newNode Node
	newNode.item = item
	oldLast := q.tail //get current tail "X"
	q.tail = &newNode //set new node to tail

	if q.IsEmpty() {
		q.head = q.tail
	} else {
		oldLast.next = q.tail // set "X" to point to the new tail
	}
}

//Enqueue - Enqueue to list
func (q *Queue) Dequeue() interface{} {

	if q.IsEmpty() {
		return nil
	}

	first := q.head

	item := first.item
	fmt.Println("Dequeue", item)
	q.head = first.next
	q.n--

	if q.IsEmpty() {
		q.tail = nil
	}

	return item
}

//IsEmpty - check is stack is empty
func (q *Queue) IsEmpty() bool {
	return q.head == nil
}

//Traverse - go thru stack
func (q *Queue) Traverse() {

	queue := q.head

	for queue != nil {
		fmt.Printf("%+v ->", queue.item)
		queue = queue.next
	}
}
