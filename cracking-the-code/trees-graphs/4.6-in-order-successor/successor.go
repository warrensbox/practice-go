package main

import (
	"fmt"
)

func main() {
	treeSet := []int{5, 10, 14, 15, 16, 30, 32, 35, 36, 37, 40, 45, 50, 75, 80, 81, 82, 85, 87, 90, 100, 105}
	t := createMinimalBST(treeSet)
	fmt.Println(inOrderSucc(35, t))
}

func inOrderSucc(p int, t *Tree) int {

	if t.root == nil {
		return -1
	}

	curr := t.root
	var store *Node
	for curr.value != p {

		if p <= curr.value {
			store = curr
			curr = curr.left //move to left
		} else {
			curr = curr.right //move to right
		}
	}

	found := curr

	if found == nil {
		return -1 //node not found
	}

	if curr.right != nil {
		return leftMostNode(curr.right)
	}
	//if there is no right subtree
	return store.value
}

func leftMostNode(node *Node) int {

	if node == nil {
		return -1
	}

	for node.left != nil {
		node = node.left
	}
	return node.value
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
