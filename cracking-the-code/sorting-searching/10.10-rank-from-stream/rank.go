package main

import "fmt"

func main() {
	bts := Tree{}
	bts.put(5)
	bts.put(1)
	bts.put(4)
	bts.put(5)
	bts.put(1)
	bts.put(4)

	rank := bts.rank(5)
	fmt.Println("rank:", rank)
}

type Node struct {
	Key         int
	Left, Right *Node
	Count       int
}

type Tree struct {
	root *Node
}

type Queue []Node

func (bts *Tree) put(key int) {
	root := bts.root
	bts.root = Put(root, key)
}

func Put(node *Node, key int) *Node {

	if node == nil {
		var newNode Node
		newNode.Key = key
		newNode.Count = 1
		return &newNode
	}

	cmp := compareTo(key, node.Key)
	if cmp <= 0 {
		node.Left = Put(node.Left, key)
	} else if cmp > 0 {
		node.Right = Put(node.Right, key)
	} else if cmp == 0 {
		node.Key = key
	}

	node.Count = 1 + Size(node.Left) + Size(node.Right)
	return node
}

func (bts *Tree) rank(key int) int {
	root := bts.root
	return Rank(root, key)
}

func Rank(node *Node, key int) int {

	if node == nil {
		return 0
	}
	cmp := compareTo(key, node.Key)
	if cmp < 0 {
		return Rank(node.Left, key)
	} else if cmp > 0 {
		return 1 + Size(node.Left) + Rank(node.Right, key)
	} else if cmp == 0 {
		return Size(node.Left)
	}

	return -1
}

func compareTo(node, key int) int {
	if node > key {
		return 1
	} else if node < key {
		return -1
	}
	return 0
}

func Size(node *Node) int {
	if node == nil {
		return 0
	}
	return node.Count
}
