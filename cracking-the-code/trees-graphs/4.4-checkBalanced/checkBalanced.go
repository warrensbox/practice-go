package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	t := createMinimalBST(arr)
	fmt.Println(isBalanced(t))
}

func isBalanced(t *Tree) bool {

	return checkHeight(t.root) != -1
}

func checkHeight(node *Node) int {
	if node == nil {
		return 0
	}

	left := checkHeight(node.left)
	right := checkHeight(node.right)

	if left == -1 || right == -1 || Abs(left-right) > 1 {
		return -1
	}

	return Max(left, right) + 1
}

func Abs(v int) int {
	if v < 0 {
		v *= -1
	}
	return v
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
