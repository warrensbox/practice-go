package main

import (
	"fmt"
)

/*
We do a depth-first walk ↴ through our tree, keeping track of the depth as we go. When we find a leaf, we add its depth to a list of depths if we haven't seen that depth already.

Each time we hit a leaf with a new depth, there are two ways that our tree might now be unbalanced:

There are more than 2 different leaf depths
There are exactly 2 leaf depths and they are more than 1 apart.
Why are we doing a depth-first walk and not a breadth-first ↴ one? You could make a case for either. We chose depth-first because it reaches leaves faster, which allows us to short-circuit earlier in some cases.


*/

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

type NodeDepthPair struct {
	TreeNode *TreeNode
	Depth    int
}

func isBalanced(root *TreeNode) bool {

	// a tree with no nodes is superbalanced, since there are no leaves!
	if root == nil {
		return true
	}

	// we short-circuit as soon as we find more than 2
	//depths := make(map[int]int)
	depths := make([]int, 3)
	nodeDepthPair := &NodeDepthPair{root, 0}
	stack := []*NodeDepthPair{nodeDepthPair}
	for len(stack) > 0 {

		nodeDepthPair := stack[len(stack)-1]
		node := nodeDepthPair.TreeNode
		depth := nodeDepthPair.Depth

		// case: we found a leaf
		if node.Left == nil && node.Right == nil {

			// we only care if it's a new depth
			//   1) more than 2 different leaf depths
			//   2) 2 leaf depths that are more than 1 apart
			if !contains(depths, depth) {
				depths = append(depths, depth)

				if len(depths) > 2 || len(depths) == 2 && Abs(depths[0]-depths[1]) > 1 {
					return false
				}
			}

			// case: this isn't a leaf - keep stepping down
		} else {
			if node.Left != nil {
				stack = append(stack, &NodeDepthPair{node.Left, depth + 1})
			}
			if node.Right != nil {
				stack = append(stack, &NodeDepthPair{node.Right, depth + 1})
			}
		}

	}

	return false

}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func contains(arr []int, num int) bool {
	for _, val := range arr {
		return val == num
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
