package main

import "fmt"

/*
goal: check if bst is valid
sol:
- left leaf has to be less than parent
- right leaf has to be more than parent
- use dfs to check above condition
O(n) time and O(n)O(n) space.
SEE OTHER BETTER EXAMPLE
*/

func main() {

	root := &TreeNode{5, nil, nil}
	obj := Constructor(root)
	obj.Insert(4)
	obj.Insert(7)
	param_2 := obj.Get_root()
	fmt.Println(valid(param_2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func valid(root *TreeNode) bool {

	return validBst(root, 0, 10000)
}

func validBst(root *TreeNode, min, max int) bool {
	// fmt.Println("root.Val", root.Val)
	// fmt.Println("min", min)
	// fmt.Println("max", max)
	// fmt.Println("******")

	if root == nil {
		return true
	} else if root.Val <= min || root.Val >= max {
		return false
	}

	return validBst(root.Left, min, root.Val) && validBst(root.Right, root.Val, max)
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
