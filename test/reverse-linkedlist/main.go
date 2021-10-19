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

func (list *ListNode) Reverse() *List {

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
