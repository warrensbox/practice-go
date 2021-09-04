package main

import (
	"fmt"
)

func main() {

}

//Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {

	var vals []*TreeNode
	if root == nil {
		return vals
	}
	hashTree := make(map[string]int)
	res := make([]*TreeNode, 0)

	postOrderTravesal(root, hashTree, &res)

	return res
}

func postOrderTravesal(node *TreeNode, hashTree map[string]int, res *[]*TreeNode) string {

	if node == nil {
		return "#"
	}

	leftString := postOrderTravesal(node.Left, hashTree, res)
	rightString := postOrderTravesal(node.Right, hashTree, res)

	treeSerial := fmt.Sprintf("%d,%s,%s", node.Val, leftString, rightString)

	hashTree[treeSerial]++
	if val, ok := hashTree[treeSerial]; ok {
		if val == 2 {
			*res = append(*res, node)
		}
	}

	fmt.Println(treeSerial)
	return treeSerial

}
