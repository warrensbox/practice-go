package main

import "fmt"

func main() {

	list6 := List{}
	list6.Add("a")
	list6.Add("b")
	list6.Add("c")
	list6.Add("c")
	list6.Add("b")
	list6.Add("a")

	list6.IsPalindrome()

}

func (l *List) IsPalindrome() {

	var stack []interface{}
	fast := l.head
	slow := l.head

	for fast != nil && fast.next != nil { //add until the middle of the list
		stack = append(stack, slow.item)
		slow = slow.next
		fast = fast.next.next
	}

	//for odd numbers move one more step
	for fast != nil {
		slow = slow.next
	}

	for slow != nil {
		n := len(stack) - 1
		//peek
		top := stack[n]

		if top != slow.item {
			fmt.Println("Not a palindrome")
			return
		}

		//pop
		stack = stack[:n]
		slow = slow.next
	}

	fmt.Println("Is  palindrome")

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
