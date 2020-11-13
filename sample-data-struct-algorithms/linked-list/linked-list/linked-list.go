package main

import (
	"fmt"
)

func main() {

	//examples
	//AddToLinkedList()
	//RemoveFromLinkedList()
	//FindInLinkedList()
	//ReverseLinkedList()
	//AddToFrontLinkedList()
	DeleteFromFront()
}

//DeleteFromFront : delete from front of list
func DeleteFromFront() {

	list := List{}

	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)

	list.DeleteFront()

	list.ShowList()
}

//AddToLinkedList : add item to list
func AddToLinkedList() {

	list := List{}

	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)

	list.ShowList()
}

//RemoveFromLinkedList : remove item to list
func RemoveFromLinkedList() {

	list := List{}

	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)

	list.ShowList()

	list.Remove(2)
	fmt.Println()
	list.ShowList()
}

//FindInLinkedList : find item to list
func FindInLinkedList() {

	list := List{}

	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)

	list.Find(2)
	list.Find(20)
}

func ReverseLinkedList() {

	list := List{}

	list.Add("A")
	list.Add("B")
	list.Add("C")
	list.Add("D")
	list.Add("E")

	list.ShowList()
	list.Reverse()
	list.ShowList()

}

func AddToFrontLinkedList() {

	list := List{}

	list.Add("A")
	list.Add("B")
	list.Add("C")
	list.Add("D")
	list.Add("E")

	list.ShowList()

	list.Front("Z")

	list.ShowList()

}

func (l *List) Reverse() {

	current := l.head
	next := &Node{}
	prev := &Node{}

	for current != nil {
		next = current.next //like a temporary variable
		current.next = prev
		prev = current
		current = next
	}

	l.head = prev

}

func (l *List) Front(item interface{}) {

	//current := l.head
	var node Node
	node.item = item
	fmt.Println(node)
	if l.head == nil {
		l.head = &node
	} else {
		node.next = l.head
		l.head = &node
	}

}

func (l *List) DeleteFront() {

	if l.head == nil {
		l.head = nil
	} else {
		if l.head.next != nil {
			l.head = l.head.next
		} else {
			l.head = nil
		}
	}

}

/* HELPER CODE -- NOT USED FOR GRADING */
//Node struct
type Node struct {
	item interface{}
	next *Node
}

//List struct
type List struct {
	head *Node
}

//Remove : Remove item to list
func (l *List) Remove(item interface{}) {

	var node Node
	node.item = item
	previous := &Node{}
	list := l.head

	for list != nil {

		if list.item == item {
			previous.next = list.next
		} else {
			previous = list
		}
		list = list.next
	}

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

//Find : find item to list
func (l *List) Find(item interface{}) {

	list := l.head

	for list != nil {

		if list.item == item {
			fmt.Println("Found")
			return
		}

		list = list.next
	}

	fmt.Println("Not Found")
}

//ShowList : show all list of item
func (l *List) ShowList() {

	list := l.head

	for list != nil {
		fmt.Printf("%+v ->", list.item)
		list = list.next
	}
	fmt.Println()
}

//Len : Length Of List
func (l *List) Len() int {
	list := l.head
	count := 0

	for list != nil {
		count++
		list = list.next
	}
	return count
}

///END OF LIBRARY
