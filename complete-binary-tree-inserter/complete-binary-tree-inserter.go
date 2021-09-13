package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type CBTInserter struct {
	root  *TreeNode
	queue []*TreeNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
	fmt.Println(cbt.queue)
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
	fmt.Println(root.Val)
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

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */
