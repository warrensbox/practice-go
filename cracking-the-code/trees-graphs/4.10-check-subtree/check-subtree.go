package main

import "fmt"

func main() {

	T1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	t1 := createMinimalBST(T1)

	T2 := []int{9, 10, 11, 12, 13, 14, 19}
	t2 := createMinimalBST(T2)

	fmt.Println(containsTree(t1.root, t2.root))
}

func containsTree(t1 *Node, t2 *Node) bool {
	if t2 == nil {
		return true //the empty subtree is always a subtree
	}
	return subTree(t1, t2)
}

func subTree(r1 *Node, r2 *Node) bool {

	if r1 == nil {
		return false //big tree empty &&& subtree still not found
	} else if r1.value == r2.value && identicalTree(r1, r2) {
		return true
	}
	return subTree(r1.left, r2) || subTree(r1.right, r2)
}

func identicalTree(r1 *Node, r2 *Node) bool {
	if r1 == nil && r2 == nil {
		return true //nothing left in the subtree
	} else if r1 == nil || r2 == nil {
		return false //exactly tree is empty, therefore trees don't match
	} else if r1.value != r2.value {
		return false //data does not match
	} else {
		return identicalTree(r1.left, r2.left) && identicalTree(r1.right, r2.right)
	}
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
