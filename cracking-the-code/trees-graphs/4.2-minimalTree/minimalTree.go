package main

type Node struct {
	value int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func main() {

	arr := []int{1, 2, 4, 5, 6, 7, 8, 9, 10, 12, 13, 14, 15}
	createMinimalBST(arr)
}

func createMinimalBST(arr []int) *Node {

	return minimalBST(arr, 0, len(arr)-1)
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
