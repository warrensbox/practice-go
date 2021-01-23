package main

import (
	"fmt"
)

func main() {

	treeSet := []int{1, 2, 3, 4, 5, 16, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	t := createMinimalBST(treeSet)
	fmt.Println(validBST(t))
}

func validBST(t *Tree) bool {
	root := t.root
	min := 10000000
	max := 10000000
	return inOrder(root, min, max)
}

func inOrder(node *Node, min, max int) bool {
	if node == nil {
		return true
	}

	if (min != 10000000 && node.value <= min) || (max != 10000000 && node.value > max) {
		return false
	}
	//movess pointer to right, max remains
	if inOrder(node.left, min, node.value) == false || inOrder(node.right, node.value, max) == false {
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
