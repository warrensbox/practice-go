/*
Problem description: Given a binary tree and a sum, determine if the tree has a root-to-leaf path such that adding up all the values along the path equals the given sum.

Note: A leaf is a node with no children.

Example:

Given the below binary tree and sum = 22,

      5
     / \
    4   8
   /   / \
  11  13  4
 /  \      \
7    2      1
return true, as there exist a root-to-leaf path which sum is 22.
*/

package main

import "fmt"

func main() {

	//T1 := []int{7, 11, 2, 4, 16, 12, 6, 5, 9, 13, 18, 8, 3, 4, 11}
	T1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	t1 := createMinimalBST(T1)
	fmt.Println(sum(t1.root, 17))
}

func sum(node *Node, forward int) bool {

	if node == nil {
		return false
	}

	bal := forward - node.value
	fmt.Println(bal)
	if bal == 0 {
		return true
	}

	return sum(node.left, bal) || sum(node.right, bal)
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
