package main

import (
	"fmt"
	"math"
)

func main() {

	root := &TreeNode{1, nil, nil}
	obj := Constructor(root)
	obj.Insert(2)
	obj.Insert(3)
	obj.Insert(4)
	obj.Insert(5)
	obj.Insert(6)
	obj.Insert(7)

	param_2 := obj.Get_root()
	fmt.Println(param_2)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanced(root *TreeNode) bool {

	if root == nil {
		return true
	}

	isBalanced := bfs(root)

	return isBalanced

}

func bfs(root *TreeNode) bool {

	queue := []*TreeNode{}
	queue = append(queue, root)
	depth := 0
	nodesCount := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := (queue)[0]
			queue = (queue)[1:]
			nodesCount++
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth++
	}

	numNodes := math.Pow(2, float64(depth)) - 1
	if numNodes == float64(nodesCount) {
		return true
	}
	return false

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
