package main

import (
	"container/list"
)

func main() {

	list := list.New()

	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(1)

	previous := *list
	head := list.Front()

	for list != nil {

	}
}
