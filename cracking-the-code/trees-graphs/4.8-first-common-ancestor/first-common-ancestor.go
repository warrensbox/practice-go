package main

import "fmt"

func main() {

	treeSet := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	t := createMinimalBST(treeSet)
	fmt.Println(LCA(t.root, 1, 3))
}

func LCA(node *Node, p, q int) *Node {

	if node == nil {
		return nil
	}

	if node.value == p || node.value == q {
		return node
	}

	left := LCA(node.left, p, q)
	if left != nil && left.value != p && left.value != q { //already found ancestor
		return left
	}
	right := LCA(node.right, p, q)
	if right != nil && right.value != p && right.value != q { //already found ancestor
		return right
	}

	if left != nil && right != nil { //p & q found i different subtrees
		return node //this is the common ancestor
	} else if node.value == p || node.value == q {
		return node
	} else {
		if node.left == nil { //return non-null value
			return node.right
		} else {
			return node.left
		}
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
