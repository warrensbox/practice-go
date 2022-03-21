package main

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {

	//base case
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	length := len(inorder)
	idx := indexOf(inorder, postorder[length-1])
	root := &TreeNode{
		Val:   postorder[length-1],
		Left:  buildTree(inorder[:idx], postorder[:idx]),
		Right: buildTree(inorder[idx+1:], postorder[idx:length-1]),
	}

	return root
}

func indexOf(arr []int, val int) int {
	for pos, v := range arr {
		if v == val {
			return pos
		}
	}
	return -1
}
