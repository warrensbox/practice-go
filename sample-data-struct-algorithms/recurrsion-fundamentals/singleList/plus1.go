package main

import (
	"fmt"
)

func main() {

	list := List{}

	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)

	getHeight(list)

}

func getHeight(list List) {

	curr := list.head
	num := recurrsionAdd(curr)

	fmt.Println(num)
}

func recurrsionAdd(node *Node) int {

	if node == nil {
		return 0
	}

	fmt.Println("item", node.item)
	return recurrsionAdd(node.next) + 1
}

/* linked list sample */
type Node struct {
	item int
	next *Node
}

//List struct
type List struct {
	head *Node
}

//Add : add item to list
func (l *List) Add(item int) {

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
