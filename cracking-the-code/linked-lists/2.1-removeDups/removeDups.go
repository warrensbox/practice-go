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
