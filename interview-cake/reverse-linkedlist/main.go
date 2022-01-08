package main

import "fmt"

func myFunction(arg string) string {

	// Write the body of your function here

	return "running with " + arg
}

func main() {

	// Run your function through some test cases here.
	// Remember: debuggin is half the battle!
	fmt.Println(myFunction("test input"))
}

func (l *LinkedList) Reverse() *List {

	current := l.Head
	prev := &List{}
	nextNode := &List{}
	for current != nil {

		// copy a pointer to the next element
		// before we overwrite currentNode.next
		nextNode = current.Next

		// reverse the 'next' pointer
		current.Next = prev

		// step forward in the list
		prev = current
		current = nextNode
	}

	return prev

}

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
