package main

import (
	"container/list"
	"fmt"
	"time"
)

func main() {

	queue := list.New()

	queue.PushBack("Hello ")

	queue.PushBack("World ")

	for queue.Len() > 0 {
		e := queue.Front() // First element
		fmt.Print(e.Value)

		queue.Remove(e) // Dequeue
	}

	sec := time.Now()

	fmt.Println(sec.Second())

}

type Animal struct {
	species string
	time    int
}

func enqueue() {

}
