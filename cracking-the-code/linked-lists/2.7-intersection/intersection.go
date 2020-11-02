package main

import (
	"fmt"
	"math"
)

func main() {

}

func intersection(list1, list2 *List) {

	if list1 == nil || list2 == nil {
		fmt.Println("No intersection")
		return
	}

	tailSize1 := getTailAndSize(list1)
	tailSize2 := getTailAndSize(list2)

	if tailSize1.tail != tailSize2.tail {
		fmt.Println("No intersection")
		return
	}

	shorter := &List{}
	longer := &List{}

	//set the pointer to the start of each linked list
	if tailSize1.size < tailSize2.size {
		shorter = list1
		longer = list2
	} else {
		shorter = list2
		longer = list1
	}

	longerNodeList := getKthNode(longer, math.Abs(float64(tailSize1.size)-float64(tailSize2.size)))
	shorterNodeList := shorter.head

	for shorterNodeList != longerNodeList {
		shorterNodeList = shorterNodeList.next
		longerNodeList = longerNodeList.next
	}

	fmt.Printf("They intersect at %v", shorterNodeList) //longerNodeList works too
	return
}

func getKthNode(list *List, k float64) *Node {

	current := list.head
	for current != nil && k > 0 {
		current = current.next
		k--
	}

	return current
}

func getTailAndSize(list *List) *TailSize {

	var tailSize TailSize

	size := 1
	current := list.head
	for current.next != nil {
		size++
		current = current.next
	}

	tailSize.tail = current
	tailSize.size = size

	return &tailSize
}

type TailSize struct {
	tail *Node
	size int
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
