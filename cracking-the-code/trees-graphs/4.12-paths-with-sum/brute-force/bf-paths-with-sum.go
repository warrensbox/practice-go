package main

import "fmt"

func main() {

	T1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 3, 2, 1}
	t1 := createMinimalBST(T1)
	fmt.Println(countPathWithSum(t1.root, 6))

}

func countPathWithSum(root *Node, targetSum int) int {

	if root == nil {
		return 0
	}

	pathsFromRoot := countPathWithSumFromRoot(root, targetSum, 0)

	pathsOnLeft := countPathWithSum(root.left, targetSum)
	pathsOnRight := countPathWithSum(root.right, targetSum)

	return pathsFromRoot + pathsOnLeft + pathsOnRight
}

func countPathWithSumFromRoot(node *Node, targetSum int, currentSum int) int {
	if node == nil {
		return 0
	}

	currentSum += node.value

	totalPaths := 0
	if currentSum == targetSum {
		totalPaths++
	}

	totalPaths += countPathWithSumFromRoot(node.left, targetSum, currentSum)
	totalPaths += countPathWithSumFromRoot(node.right, targetSum, currentSum)
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
