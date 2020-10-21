package main

import (
	"fmt"
)

type Node struct {
	item interface{}
	next *Node
}

type List struct {
	head *Node
}

func (l *List) Add(item interface{}) {

	var node Node
	node.item = item

	if l.head == nil {
		l.head = &node
	} else {
		list := l.head
		for list.next != nil {
			list = list.next
		}
		list.next = &node
	}
}

func (l *List) nthToLast(k int) {
	//have 2 pointer
	pointer1 := l.head
	pointer2 := l.head

	/*move p1 k nodes into the list*/
	for i := 0; i < k; i++ {
		if pointer1 == nil {
			return
		}
		pointer1 = pointer1.next
	}

	for pointer1 != nil {
		pointer1 = pointer1.next
		pointer2 = pointer2.next
	}

	fmt.Println(pointer2)
}

func (l *List) RemoveDups() {

	dict := make(map[interface{}]bool)

	previous := &Node{}
	list := l.head

	for list != nil {
		if _, ok := dict[list.item]; ok {
			previous.next = list.next
		} else {
			dict[list.item] = true
			previous = list
		}
		list = list.next
	}
}

func (l *List) ShowList() {

	list := l.head

	for list != nil {
		fmt.Printf("%+v ->", list.item)
		list = list.next
	}
}

func main() {

	list := List{}
	list.ShowList()
	list.Add("1")
	list.Add("2")
	list.Add("3")
	list.Add("4")
	list.Add("5")
	list.Add("1")
	list.Add("2")
	list.Add("3")
	list.RemoveDups()
	list.ShowList()
}
