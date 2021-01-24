package main

import (
	"fmt"
)

func main() {

	T1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 3, 2, 1}
	t1 := createMinimalBST(T1)
	fmt.Println(countPathWithSum(t1.root, 6))

}

func countPathWithSum(root *Node, targetSum int) int {
	hashMap := make(map[int]int)
	return countPathWithSumFromRoot(root, targetSum, 0, hashMap)
}

func countPathWithSumFromRoot(node *Node, targetSum, runningSum int, hashMap map[int]int) int {
	if node == nil {
		return 0
	}

	/*count path with sum ending at the current node*/
	runningSum += node.value
	sum := runningSum - targetSum

	totalPaths := hashMap[sum]

	if runningSum == targetSum {
		totalPaths++
	}

	hashMap[runningSum]++
	totalPaths += countPathWithSumFromRoot(node.left, targetSum, runningSum, hashMap)
	totalPaths += countPathWithSumFromRoot(node.right, targetSum, runningSum, hashMap)
	hashMap[runningSum]--

	return totalPaths
}

/* create tree for testing */
type Node struct {
	value int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func createMinimalBST(arr []int) *Tree {
	t := Tree{}
	t.root = minimalBST(arr, 0, len(arr)-1)
	return &t
}

func minimalBST(arr []int, start, end int) *Node {
	if end < start {
		return nil
	}
	n := Node{}
	mid := (start + end) / 2
	n.value = arr[mid]
	n.left = minimalBST(arr, start, mid-1)
	n.right = minimalBST(arr, mid+1, end)
	return &n
}
