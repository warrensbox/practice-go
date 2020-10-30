package main

import (
	"container/list"
	"fmt"
)

func main() {

	dict := make(map[interface{}]bool)

	previous := list.Element{}
	list := list.New()

	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(1)

	node := list.Front()

	for node != nil {
		if _, ok := dict[node.Value]; ok {
			previous.Next().Value = node.Next()
		} else {
			dict[node.Value] = true
			previous = *node
		}

		//fmt.Println(e.Value)

		//break
		node = node.Next()
	}

	n := list.Front()

	for list.Len() > 0 {
		for n != nil {

			fmt.Println(n.Value)
			//l.queueCat.Remove(e) // Dequeue
			//break
			n = n.Next()
		}
	}

}
