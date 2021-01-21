package main

import "fmt"

func main() {

	arr := []int{1, 2, 4, 5, 6, 7, 8, 9, 10, 12, 13, 14, 15}
	t := createMinimalBST(arr)
	height := getHeight(t.root)
	fmt.Println(height)
}

func getHeight(node *Node) int {

	if node == nil {
		return -1
	}

	return Max(getHeight(node.left), getHeight(node.right)) + 1
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
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
