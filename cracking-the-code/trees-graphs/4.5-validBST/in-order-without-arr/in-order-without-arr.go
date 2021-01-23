package main

import (
	"fmt"
	"math"
)

func main() {

	treeSet := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	t := createMinimalBST(treeSet)
	o := Order{
		math.MinInt16,
	}
	fmt.Println(validBST(o, t))
}

func validBST(o Order, t *Tree) bool {

	return o.inOrder(t.root)
}

type Order struct {
	previous int
}

func (o *Order) inOrder(node *Node) bool {
	if node == nil {
		return true
	}

	if o.inOrder(node.left) == false {
		return false
	}

	if o.previous != math.MinInt16 && node.value <= o.previous {
		return false
	}

	o.previous = node.value

	if o.inOrder(node.right) == false {
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
