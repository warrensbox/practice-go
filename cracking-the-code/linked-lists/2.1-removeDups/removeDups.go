package main

import "fmt"

func main() {

	list := List{}

	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(1)
	list.Add(3)
	list.RemoveDups()
	list.ShowList()
}

func (l *List) RemoveDups() {

	hash := make(map[interface{}]bool)
	previous := &Node{}
	list := l.head

	for list != nil {

		if _, ok := hash[list.item]; ok {
			previous.next = list.next
		} else {
			hash[list.item] = true
			previous = list
		}

		list = list.next
	}

}

/* HELPER CODE - NOT USED FOR GRADING */
//Node struct
type Node struct {
	item interface{}
	next *Node
}

//List struct
type List struct {
	head *Node
}

//Add : add item to list
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

//ShowList : show all list of item
func (l *List) ShowList() {

	list := l.head

	for list != nil {
		fmt.Printf("%+v ->", list.item)
		list = list.next
	}

}
