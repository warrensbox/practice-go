package main

import (
	"fmt"
)

func main() {

	x := 5
	y := 5

	output := minKnightMoves(x, y)
	fmt.Println(output)
}

type Move struct {
	X     int
	Y     int
	Jumps int
}

type coordinate struct {
	X int
	Y int
}

func minKnightMoves(x int, y int) int {

	queue := New()
	queue.Enqueue(Move{0, 0, 0})

	seen := make(map[coordinate]bool)

	for queue.Len() > 0 {
		move := queue.Dequeue().(Move)
		coor := coordinate{move.X, move.Y}
		fmt.Println("********")
		fmt.Println("move", move)

		if seen[coor] {
			fmt.Println("SEEEN!!!")
			continue
		} else {
			seen[coor] = true
		}

		if move.X == x && move.Y == y {

			return move.Jumps
		}
		//fmt.Println("LENGTH", queue.Len())
		addNextMoves(queue, move)
	}

	return 3
}

func addNextMoves(queue *Queue, currMove Move) {
	directions := [][]int{[]int{-2, -1}, []int{-2, 1}, []int{-1, 2}, []int{1, 2}, []int{2, 1}, []int{2, -1}, []int{1, -2}, []int{-1, -2}}
	fmt.Println("------")
	for _, pos := range directions {
		fmt.Println("currMove.X ", currMove.X)
		fmt.Println("pos[0]", pos[0])
		fmt.Println(currMove.X + pos[0])

		queue.Enqueue(Move{currMove.X + pos[0], currMove.Y + pos[1], currMove.Jumps + 1})
		fmt.Println("length", queue.Len())
	}
}

///-------------------------------------- utils
type QueueNode struct {
	Value interface{}
	Next  *QueueNode
}

// Queue generic stack implementation using a Singly Linked List
type Queue struct {
	Head   *QueueNode
	Tail   *QueueNode
	Length int
}

// New Create a new queue
// O(1)
func New() *Queue {
	return &Queue{
		Head:   nil,
		Tail:   nil,
		Length: 0,
	}
}

// Dequeue Take the next item off the front of the queue
// O(1)
func (q *Queue) Dequeue() interface{} {
	if q.Head == nil {
		return nil
	}

	front := q.Peek()
	if q.Head == q.Tail {
		q.Head = nil
		q.Tail = nil
	} else {
		q.Head = q.Head.Next
	}

	q.Length--
	return front
}

// Enqueue Put an item on the end of a queue
// O(1)
func (q *Queue) Enqueue(value interface{}) {
	t := &QueueNode{
		Next:  nil,
		Value: value,
	}

	if q.Head == nil {
		q.Head = t
		q.Tail = t
	} else {
		q.Tail.Next = t
		q.Tail = t
	}

	q.Length++
}

// Len Return the number of items in the queue
// O(1)
func (q *Queue) Len() int {
	return q.Length
}

// Peek Return the first item in the queue without removing it
// O(1)
func (q *Queue) Peek() interface{} {
	if q.Head == nil {
		return nil
	}

	return q.Head.Value
}
