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

func sumOfLeftLeaves(root *TreeNode) int {

	sum := 0
	dfs(root, &sum, false)
	return sum
}

func dfs(curr *TreeNode, sum *int, left bool) {

	if curr == nil {
		return
	}

	dfs(curr.Left, sum, true)
	if curr.Left == nil && curr.Right == nil && left {
		*sum += curr.Val
	}
	dfs(curr.Right, sum, false)
}
