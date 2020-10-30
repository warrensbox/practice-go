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
	for i := 0; i < k; i++ {
		if pointer1 == nil {
			return nil
		}
		pointer1 = pointer1.next
	}

	for pointer1 != nil {
		pointer1 = pointer1.next
		pointer2 = pointer2.next
	}

	fmt.Println(pointer1.item)
	return pointer1
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
