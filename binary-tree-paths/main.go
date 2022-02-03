package main

import "strconv"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {

	res := []string{}
	helper(root, &res, "")

	return res
}

func helper(root *TreeNode, res *[]string, val string) {

	if root.Left == nil && root.Right == nil {
		val = val + strconv.Itoa(root.Val)
		*res = append(*res, val)
	}

	if root.Left != nil {
		helper(root.Left, res, val+strconv.Itoa(root.Val)+"->")
	}

	if root.Right != nil {
		helper(root.Right, res, val+strconv.Itoa(root.Val)+"->")
	}

}
