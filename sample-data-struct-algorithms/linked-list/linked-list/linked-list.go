package main

import (
	"fmt"
	"math"
)

type Node struct {
	item interface{}
	next *Node
}

type List struct {
	head *Node
}

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

func (l *List) countNumberOfNodes() int {
	list := l.head
	count := 0
	for list != nil {
		list = list.next
		count++
	}
	return count
}

func (l *List) deleteMiddleNode() {

	list := l.head
	count := l.countNumberOfNodes()

	fmt.Println(count)

	mid := count / 2

	for mid > 1 {
		list = list.next
		mid--
	}

	list.next = list.next.next
}

func (l *List) nthToLastRecussion(head *Node, k int) int {

	if head == nil {
		return 0
	}

	index := l.nthToLastRecussion(head.next, k) + 1

	if index == k {
		fmt.Printf("%s and %v", head.item, k)
	}
	return index
}

func (l *List) nthToLastPointer(k int) *Node {
	//have 2 pointer
	pointer1 := l.head
	pointer2 := l.head

	/*move p1 k nodes into the list*/
	for i := 0; i < k; i++ {
		if pointer1 == nil {
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

	fmt.Println(pointer2)
	return pointer2
}

func (l *List) RemoveDups() {

	dict := make(map[interface{}]bool)

	previous := &Node{}
	list := l.head

	for list != nil {
		if _, ok := dict[list.item]; ok {
			previous.next = list.next
		} else {
			dict[list.item] = true
			previous = list
		}
		list = list.next
	}
}

func (l *List) ShowList() {

	list := l.head

	for list != nil {
		fmt.Printf("%+v ->", list.item)
		list = list.next
	}
}

func (l *List) partition(x interface{}) *Node {

	node := l.head
	head := &Node{}
	tail := &Node{}

	for node != nil {

		fmt.Println(node)
		next := node.next
		if (node.item).(int) < x.(int) {
			fmt.Println("less than", node.item)
			// 	/* insert node at head */
			node.next = head
			head = node
		} else {
			fmt.Println("more than", node.item)
			// 	/* insert node at tail */
			tail.next = node
			tail = node
		}

		node = next
	}

	tail.next = nil

	fmt.Println(head)
	ShowListNumerate(head)
	return nil
}

func (l *List) partitionY(x interface{}) *Node {

	node := l.head
	beforeStart := &Node{}
	beforeEnd := &Node{}
	afterStart := &Node{}
	afterEnd := &Node{}

	for node != nil {
		next := node.next
		node.next = nil
		if (node.item).(int) < x.(int) {
			if beforeStart == nil {
				beforeStart = node
				beforeEnd = beforeStart
			} else {
				beforeEnd.next = node
				beforeEnd = node
			}
		} else {
			if afterStart == nil {
				afterStart = node
				afterEnd = afterStart
			} else {
				afterEnd.next = node
				afterEnd = node
			}
		}
		node = next
	}

	if beforeStart == nil {
		ShowListNumerate(afterStart)
		return afterStart
	}

	beforeEnd.next = afterStart
	ShowListNumerate(beforeStart)
	return nil
}

func ShowListNumerate(node *Node) {

	list := node

	for list != nil {
		fmt.Printf("%+v ->", list.item)
		list = list.next
	}
}

func SumListBackward(list1 *List, list2 *List) {

	temp1 := 0
	listOne := list1.head
	list3 := List{}
	for listOne != nil {
		fmt.Printf("%+v ->", listOne.item)
		data := (listOne.item).(int)
		current := temp1 * 10
		temp1 = current + data
		listOne = listOne.next
	}
	fmt.Println()

	temp2 := 0
	listTwo := list2.head
	for listTwo != nil {
		fmt.Printf("%+v ->", listTwo.item)
		data := (listTwo.item).(int)
		current := temp2 * 10
		temp2 = current + data
		listTwo = listTwo.next
	}

	total := reverseNumber(reverseNumber(temp1) + reverseNumber(temp2))

	fmt.Println("total", total)

	x := total
	for x > 0 {
		z := x / 10
		y := x % 10
		fmt.Println(y)
		list3.Add(y)
		x = z
	}
	list3.ShowList()

}

func reverseNumber(x int) int {
	temp3 := 0

	for x > 0 {
		z := x / 10
		y := x % 10
		fmt.Println(y)
		temp3 = (temp3 * 10) + y
		x = z
	}

	return temp3
}

func SumListForward(list1 *List, list2 *List) {

	temp1 := 0
	listOne := list1.head
	list3 := List{}
	for listOne != nil {
		fmt.Printf("%+v ->", listOne.item)
		data := (listOne.item).(int)
		current := temp1 * 10
		temp1 = current + data
		listOne = listOne.next
	}

	fmt.Println()

	temp2 := 0
	listTwo := list2.head
	for listTwo != nil {
		fmt.Printf("%+v ->", listTwo.item)
		data := (listTwo.item).(int)
		current := temp2 * 10
		temp2 = current + data
		listTwo = listTwo.next
	}

	total := reverseNumber(temp1 + temp2)

	fmt.Println("total", total)

	x := total
	for x > 0 {
		z := x / 10
		y := x % 10
		fmt.Println(y)
		list3.Add(y)
		x = z
	}
	list3.ShowList()

}

func (l *List) isPalindrome() {
	fast := l.head
	slow := l.head

	var stack []interface{}

	for fast != nil && fast.next != nil {
		stack = append(stack, slow.item)
		slow = slow.next
		fast = fast.next.next
	}

	if fast == nil {
		slow = slow.next
	}

	top := stack[len(stack)-1]
	for slow != nil {

		if len(stack) > 0 {
			n := len(stack) - 1 // Top element
			top = stack[n]
			stack = stack[:n] // Pop
		}

		if slow.item != top {
			fmt.Println("not palin")
			return
		}

		slow = slow.next

	}

	fmt.Println("true")

}

type TailSize struct {
	tail *Node
	size int
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

func main() {

	//list := List{}
	//list.ShowList()
	// list.Add(1)
	// list.Add(2)
	// list.Add(3)
	// list.Add(4)
	// list.Add(5)
	// list.Add(6)
	// list.Add(8)
	// list.Add(9)
	// list.Add(10)
	// list.Add(5)
	// list.Add(3)
	//list.nthToLastPointer(3)
	//list.RemoveDups()
	//list.deleteMiddleNode()
	// list.partitionY(5)
	// list.ShowList()

	// list1 := List{}
	// list1.Add(7)
	// list1.Add(1)
	// list1.Add(6)
	// //list1.ShowList()

	// fmt.Println()
	// list2 := List{}
	// list2.Add(5)
	// list2.Add(9)
	// list2.Add(2)
	// //list2.ShowList()
	// fmt.Println()

	//SumListForward(&list1, &list2)
	//SumListBackward(&list1, &list2)

	// list4 := List{}
	// list4.Add(6)
	// list4.Add(1)
	// list4.Add(7)
	// //list1.ShowList()

	// fmt.Println()
	// list5 := List{}
	// list5.Add(2)
	// list5.Add(9)
	// list5.Add(5)

	// SumListForward(&list4, &list5)

	list6 := List{}
	list6.Add("a")
	list6.Add("b")
	list6.Add("c")
	list6.Add("c")
	list6.Add("b")
	list6.Add("a")

	list6.isPalindrome()

}
