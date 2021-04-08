package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}
func countNodes(root *TreeNode) int {

	if root == nil { //base case
		return 0
	}

	leftLevel := 1
	leftNode := root.Left
	for leftNode != nil {
		leftNode = leftNode.Left
		leftLevel++
	}

	rightLevel := 1
	rightNode := root.Right
	for rightNode != nil {
		rightNode = rightNode.Right
		rightLevel++
	}

	if leftLevel == rightLevel { //perfect binary tree
		return int(math.Pow(2, float64(leftLevel))) - 1
	}

	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
