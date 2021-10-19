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
	list.Add("H")

	list.nthToLast(6)

}

func (l *List) nthToLast(k int) *Node {

	pointer1 := l.head
	pointer2 := l.head

	/*move p1 to k nodes into the list*/
	for i := 0; i < k-1; i++ {
		if pointer1.next == nil {
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

	fmt.Println(pointer2.item)
	return pointer2
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
