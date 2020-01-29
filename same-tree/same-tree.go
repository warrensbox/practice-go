package main

import (
	"fmt"
	"math/rand"
)

func main() {

	t1 := newOrder(1)
	t2 := newOrder(1)

	tree := isSameTree(t1, t2)
	fmt.Println(tree)

	r1 := newRandom(3, 1)
	r2 := newRandom(3, 1)

	tree = isSameTree(r1, r2)
	fmt.Println(tree)

}

func newOrder(k int) *TreeNode {
	var t *TreeNode

	test := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for _, v := range test {
		t = insert(t, (1+v)*k)
	}
	return t
}

func newRandom(n, k int) *TreeNode {
	var t *TreeNode
	for _, v := range rand.Perm(n) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *TreeNode, v int) *TreeNode {
	if t == nil {
		return &TreeNode{nil, v, nil}
	}
	if v < t.Val {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}

//TreeNode : tree
type TreeNode struct {
	Left  *TreeNode
	Val   int
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return (isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right))
}
