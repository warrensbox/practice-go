package main

import (
	"fmt"
)

//2nd-largest-binary-node

/*
goal :  find 2nd largest
sol:
- if we have a left subtree but not a right subtree, then the current node is the largest overall. The second largest  element must
  be the largest element in the left subtree
- if we have a right child, but that right child node does not have any children, then the right child must be the largest element
  and current node must be the largest element
- Else, we have a right subtree with more than one element, so the largest and second largest are somewhere in that subtree. So we step right.
*/

func main() {

	root := &TreeNode{5, nil, nil}
	obj := Constructor(root)
	obj.Insert(3)
	obj.Insert(8)
	obj.Insert(1)
	obj.Insert(4)
	obj.Insert(7)
	obj.Insert(9)
	param_2 := obj.Get_root()
	fmt.Println(secondLargest(param_2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findLargest(node *TreeNode) int {
	var largest *TreeNode
	curr := node
	fmt.Println("going here")
	for curr != nil {
		if curr.Right == nil {
			largest = curr
		}
		curr = curr.Right
	}
	return largest.Val
}

func secondLargest(root *TreeNode) int {

	if root == nil || (root.Left == nil && root.Right == nil) {
		return -1
	}

	curr := root
	// fmt.Println("left", curr.Left)
	// fmt.Println("Right", curr.Right)
	for {

		// case: current is largest and has a left subtree
		// 2nd largest is the largest in that subtree
		if curr.Left != nil && curr.Right == nil { // has left subtree
			return findLargest(curr.Left)
		}

		// case: current is parent of largest, and largest has no children,
		// so current is 2nd largest
		if curr.Right != nil && curr.Right.Left == nil && curr.Right.Right == nil {
			fmt.Println("tt")
			return curr.Val
		}

		fmt.Println("end")
		curr = curr.Right
	}

}

/*  BINARY TREE INSERTER */
type CBTInserter struct {
	root  *TreeNode
	queue []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {

	cbt := CBTInserter{root: root}

	if root == nil {
		return cbt
	}

	cbt.queue = append(cbt.queue, root)

	for len(cbt.queue) != 0 {
		root := cbt.queue[0]
		if root.Left == nil {
			return cbt
		} else if root.Right == nil {
			cbt.queue = append(cbt.queue, root.Left)
			return cbt
		}
		cbt.queue = cbt.queue[1:]
		cbt.queue = append(cbt.queue, root.Left, root.Right)

	}
	return cbt
}

func (this *CBTInserter) Insert(val int) int {

	newNode := &TreeNode{Val: val}

	if len(this.queue) == 0 {
		this.root = newNode
		this.queue = append(this.queue, newNode)
		return -1
	}
	root := this.queue[0]
	if root.Left == nil {
		root.Left = newNode
	} else {
		root.Right = newNode
		this.queue = this.queue[1:]
	}
	this.queue = append(this.queue, newNode)
	return root.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}
