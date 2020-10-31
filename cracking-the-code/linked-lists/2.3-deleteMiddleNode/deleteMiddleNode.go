package main

import "fmt"

func main() {

	list := List{}

	list.Add("A")
	list.Add("B")
	list.Add("C")
	list.Add("D")
	list.Add("E")
	list.Add("F")
	list.Add("G")

	list.DeleteMiddleNode()

	list.ShowList()

}

//DeleteMiddleNode delete the middle node
func (l *List) DeleteMiddleNode() {

	list := l.head
	count := l.Len()
	mid := count / 2

	for mid > 1 {
		list = list.next
		mid--
	}

	list.next = list.next.next

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
