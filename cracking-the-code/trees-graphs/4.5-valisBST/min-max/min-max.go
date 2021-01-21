package main

import (
	"fmt"
)

func main() {

	treeSet := []int{1, 2, 3, 4, 5, 16, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	t := createMinimalBST(treeSet)
	o := Order{}
	fmt.Println(validBST(o, t))
}

func validBST(o Order, t *Tree) bool {

	o.min = 10000000
	o.max = 10000000
	return o.inOrder(t.root, o.min, o.max)
}

type Order struct {
	previous int
	min      int
	max      int
}

func (o *Order) inOrder(node *Node, min, max int) bool {
	if node == nil {
		return true
	}
	fmt.Println("node.value", node.value)
	fmt.Println("min", min)
	fmt.Println("max", max)
	fmt.Println("-----")
	if (min != 10000000 && node.value <= min) || (max != 10000000 && node.value > max) {
		return false
	}
	//movess pointer to right, max remains
	if o.inOrder(node.left, min, node.value) == false || o.inOrder(node.right, node.value, max) == false {
		return false
	}

	return true
}

/* create tree for testing */
type Node struct {
	value int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func createMinimalBST(arr []int) *Tree {
	t := Tree{}
	t.root = minimalBST(arr, 0, len(arr)-1)
	return &t
}

func minimalBST(arr []int, start, end int) *Node {
	if end < start {
		return nil
	}
	n := Node{}
	mid := (start + end) / 2
	n.value = arr[mid]
	n.left = minimalBST(arr, start, mid-1)
	n.right = minimalBST(arr, mid+1, end)
	return &n
}
