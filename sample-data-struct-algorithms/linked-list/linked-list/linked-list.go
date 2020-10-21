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

func (l *List) nthToLastRecussion(head *Node, k int) int {

	if head == nil {
		return 0
	}

	index := l.nthToLastRecussion(head.next, k) + 1

	if index == k {
		fmt.Printf("%s and %v", head.item, k)
	}
	return index
}

func (l *List) nthToLastPointer(k int) *Node {
	//have 2 pointer
	pointer1 := l.head
	pointer2 := l.head

	/*move p1 k nodes into the list*/
	for i := 0; i < k; i++ {
		if pointer1 == nil {
			return nil
		}
		pointer1 = pointer1.next
	}

	/* move them at the same pace*/
	/* when p1 hits the end, p2 will be at the right element */
	for pointer1 != nil {
		pointer1 = pointer1.next
		pointer2 = pointer2.next
	}

	fmt.Println(pointer2)
	return pointer2
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
	list.Add("1")
	list.Add("2")
	list.Add("3")
	list.nthToLastPointer(3)
	//list.RemoveDups()
	list.ShowList()
}
