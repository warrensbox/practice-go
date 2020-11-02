package main

import "fmt"

func main() {

	list1 := List{}
	list1.Add(7)
	list1.Add(1)
	list1.Add(6)
	list1.ShowList()

	list2 := List{}
	list2.Add(5)
	list2.Add(9)
	list2.Add(2)
	list2.ShowList()

	SumListBackward(&list1, &list2)

	list3 := List{}
	list3.Add(6)
	list3.Add(1)
	list3.Add(7)
	list3.ShowList()

	list4 := List{}
	list4.Add(2)
	list4.Add(9)
	list4.Add(5)
	list4.ShowList()

	SumListForward(&list3, &list4)

}

func SumListBackward(list1, list2 *List) {

	temp1 := 0
	listOne := list1.head

	for listOne != nil {
		value := (listOne.item).(int)
		data := temp1 * 10
		temp1 = data + value

		listOne = listOne.next
	}

	temp2 := 0
	listTwo := list2.head

	for listTwo != nil {
		value := (listTwo.item).(int)
		data := temp2 * 10
		temp2 = data + value

		listTwo = listTwo.next
	}

	total := reverse(temp1) + reverse(temp2)

	list3 := List{}
	for total > 0 {
		bal := total % 10
		list3.Add(bal)
		total = total / 10
	}

	list3.ShowList()
}

func SumListForward(list1, list2 *List) {

	temp1 := 0
	listOne := list1.head

	for listOne != nil {
		value := (listOne.item).(int)
		data := temp1 * 10
		temp1 = data + value

		listOne = listOne.next
	}

	temp2 := 0
	listTwo := list2.head

	for listTwo != nil {
		value := (listTwo.item).(int)
		data := temp2 * 10
		temp2 = data + value

		listTwo = listTwo.next
	}

	total := temp1 + temp2
	totalR := reverse(total) //reverse this as it is easier to put in list

	list3 := List{}
	for totalR > 0 {
		bal := totalR % 10
		list3.Add(bal)
		totalR = totalR / 10
	}

	list3.ShowList()
}

func reverse(num int) int {

	total := 0
	for num > 0 {
		bal := num % 10
		total = bal + (total * 10)
		num = num / 10
	}

	return total

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
