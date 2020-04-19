package main

import (
	"fmt"
	"math"
)

func main() {

}

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

func countNodes(root *TreeNode) int {

	if root == nil { // if the tree is empty
		return 0
	}

	d := computeDepth(root)
	fmt.Println("depth", d)
	if d == 0 { // if the tree contains 1 node
		return 1
	}

	// Last level nodes are enumerated from 0 to 2**d - 1 (left -> right).
	// Perform binary search to check how many nodes exist.
	left := 1
	rightPow := math.Pow(2, float64(d)) - 1
	right := int(rightPow)
	fmt.Println("RRRRR", right)
	for left <= right {
		pivot := left + (right-left)/2
		fmt.Println("PIVOT", pivot)
		if exist(pivot, d, root) {
			fmt.Println("EXIST")
			left = pivot + 1
			fmt.Println("left", left)
			fmt.Println("right", right)
		} else {
			fmt.Println("DONT EXIST")
			right = pivot - 1
			fmt.Println("left", left)
			fmt.Println("right", right)
		}
	}

	containsPow := math.Pow(2, float64(d)) - 1
	contains := int(containsPow)

	return contains + left
}

func computeDepth(root *TreeNode) int {

	d := 0
	for root.Left != nil {
		root = root.Left
		d++
	}
	return d

}

// Last level nodes are enumerated from 0 to 2**d - 1 (left -> right).
// Return True if last level node idx exists.
// Binary search with O(d) complexity.
func exist(idx int, d int, node *TreeNode) bool {
	fmt.Println("----")
	left := 0
	rightPow := math.Pow(2, float64(d)) - 1
	right := int(rightPow)
	fmt.Println("righRt", right)
	for i := 0; i < d; i++ {
		fmt.Println("+++++")
		pivot := left + (right-left)/2
		fmt.Println("pivot", pivot)
		fmt.Println("right", right)
		fmt.Println("left", left)
		fmt.Println("idx", idx)
		if idx <= pivot {
			fmt.Println("==pivot")
			node = node.Left
			right = pivot
		} else {
			fmt.Println("!=pivot")
			node = node.Right
			left = pivot + 1
		}

		//fmt.Println("right",right)

	}

	return node != nil
}
