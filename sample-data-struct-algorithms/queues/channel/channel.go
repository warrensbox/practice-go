package main

import "fmt"

type Item struct {
	name   string
	number int
}

func main() {

	queue := make(chan *Item, 10)

	item1 := &Item{name: "atkins", number: 1}
	item2 := &Item{name: "hand", number: 2}
	item3 := &Item{name: "soap", number: 3}

	//queue
	queue <- item1
	queue <- item2
	queue <- item3

	close(queue)

	//dequeue
	item := <-queue
	fmt.Println(item)

	item = <-queue
	fmt.Println(item)

	item = <-queue
	fmt.Println(item)
}
