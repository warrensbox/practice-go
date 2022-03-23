package main

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

func pruneTree(root *TreeNode) *TreeNode {

	var containsOne func(root *TreeNode) bool
	containsOne = func(node *TreeNode) bool {

		if node == nil {
			return false
		}

		leftHasOne := containsOne(node.Left)
		rightHasOne := containsOne(node.Right)

		if !leftHasOne { //if left does not have one
			node.Left = nil
		}

		if !rightHasOne { //if left does not have one
			node.Right = nil
		}

		return node.Val == 1 || leftHasOne || rightHasOne
	}

	if !containsOne(root) { //is does not contain 1
		return nil
	}

	return root

}
