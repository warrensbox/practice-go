package main

import "fmt"

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {

	//post(root)

	if root == nil {
		return 0
	}

	take, notTake := postOrderTraversal(root)
	return max(take, notTake)
}

func postOrderTraversal(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	takeL, notTakeL := postOrderTraversal(root.Left)       //root.Left node's max amounts under both scenario
	takeR, notTakeR := postOrderTraversal(root.Right)      //root.Right node's max amounts under both scenario
	take := root.Val + notTakeL + notTakeR                 //Take root node case
	notTake := max(takeL, notTakeL) + max(takeR, notTakeR) //NOT take root node case

	fmt.Println("take", take)
	fmt.Println("notTake", notTake)
	fmt.Println("===")
	return take, notTake
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
