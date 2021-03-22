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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	sorted []int
}

func Constructor(root *TreeNode) BSTIterator {
	var iterator BSTIterator
	iterator.inOrder(root)
	fmt.Println(iterator.sorted)

	return iterator
}

func (this *BSTIterator) inOrder(curr *TreeNode) {

	if curr == nil {
		fmt.Println("nothing")
		return
	}

	this.inOrder(curr.Left)
	fmt.Println(curr.Val)
	this.sorted = append(this.sorted, curr.Val)
	fmt.Println("going right")
	this.inOrder(curr.Right)

}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {

	n := this.sorted[0]
	this.sorted = this.sorted[1:]
	return n
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.sorted) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
