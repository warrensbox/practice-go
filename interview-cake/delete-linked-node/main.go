package main

import "fmt"

/*
goal : remove a node from linked list
sol :
- iterate though list
- assign previous to next

 head 2->3->4->5


 prev current (you know where you came from)

 if match {
	 prev.next.next = list.next
 }

 prev = current
 current = current.next

*/

type LinkedList struct {
	Head *List
}
type List struct {
	Val  int
	Next *List
}

func New() *LinkedList {
	var newList LinkedList
	return &newList
}

func (list *LinkedList) Add(item int) {

	newItem := &List{
		Val: item,
	}

	if list.Head == nil {
		list.Head = newItem //if head is nil, set head to new item
	} else {
		currentNode := list.Head
		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}
		currentNode.Next = newItem
	}

}

func (list *LinkedList) ShowAll() {

	if list.Head == nil {
		fmt.Println("Nothing in list")
	} else {
		currentNode := list.Head
		for currentNode != nil {
			fmt.Println(currentNode.Val)
			currentNode = currentNode.Next
		}
	}
}

func (list *LinkedList) Delete(item int) {

	if list.Head == nil {
		fmt.Println("Nothing to delete")
	} else if list.Head.Val == item {
		list.Head = list.Head.Next
	} else {
		currentNode := list.Head
		prev := &List{}
		for currentNode != nil {

			if currentNode.Val == item {
				prev.Next = currentNode.Next
			} else {
				prev = currentNode
			}

			currentNode = currentNode.Next
		}
	}

}

func main() {

	list := New()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.ShowAll()

	list.Delete(4)
	list.ShowAll()

}
