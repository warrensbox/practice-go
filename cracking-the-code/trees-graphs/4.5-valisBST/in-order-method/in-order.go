package main

import "fmt"

func main() {

	treeSet := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	t := createMinimalBST(treeSet)
	o := Order{
		make([]int, len(treeSet)),
		0,
	}
	fmt.Println(validBST(o, t))
}

func validBST(o Order, t *Tree) bool {
	o.inOrder(t.root)
	for i := 0; i < len(o.arr)-1; i++ {
		if o.arr[i] > o.arr[i+1] {
			return false
		}
	}
	return true
}

type Order struct {
	arr   []int
	index int
}

func (o *Order) inOrder(node *Node) {
	if node == nil {
		return
	}
	o.inOrder(node.left)
	o.arr[o.index] = node.value
	o.index = o.index + 1
	o.inOrder(node.right)
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
