package main

/*
a red–black tree is a kind of self-balancing binary search tree.
Each node stores an extra bit representing color,
used to ensure that the tree remains approximately
balanced during insertions and deletions.
*/

import "fmt"

const RED = true
const BLACK = false

//define Node
type Node struct {
	Key         rune
	Value       int
	Left, Right *Node
	Count       int
	Color       bool
}

type Tree struct {
	root *Node
}

type Queue []Node

func main() {

	bts := Tree{}
	bts.put('s', 1)
	bts.put('e', 5)
	bts.put('a', 5)
	bts.put('r', 5)
	bts.put('c', 5)
	bts.put('h', 3)
	bts.put('x', 5)
	bts.put('m', 5)
	bts.put('p', 5)
	bts.put('l', 5)

	bts.ShowAllLeft()
	bts.ShowAllRight()
}

func (bts *Tree) put(key rune, value int) {
	root := bts.root
	bts.root = Put(root, key, value)
}

//Put : recursion
func Put(h *Node, key rune, value int) *Node {

	if h == nil {
		var newNode Node
		newNode.Key = key
		newNode.Value = value
		newNode.Count = 1
		newNode.Color = RED
		return &newNode
	}

	cmp := compareTo(key, h.Key)
	if cmp < 0 {
		h.Left = Put(h.Left, key, value)
	} else if cmp > 0 {
		h.Right = Put(h.Right, key, value)
	} else if cmp == 0 {
		h.Value = value
	}

	//conditions to manipulate
	if isRed(h.Right) && !isRed(h.Left) {
		h = rotateLeft(h)
	}
	if isRed(h.Left) && isRed(h.Left.Left) {
		h = rotateRight(h)
	}
	if isRed(h.Left) && isRed(h.Right) {
		flipColor(h)
	}

	h.Count = Size(h.Left) + Size(h.Right)
	return h
}

func isRed(node *Node) bool {

	if node == nil {
		return false //null links are black
	}

	return node.Color == RED
}

func rotateLeft(h *Node) *Node {

	x := h.Right
	h.Right = x.Left
	x.Left = h
	x.Color = h.Color
	h.Color = RED

	return x
}

//temporarily used
func rotateRight(h *Node) *Node {

	x := h.Left
	h.Left = x.Right
	x.Right = h
	x.Color = h.Color
	h.Color = RED

	return x
}

func flipColor(h *Node) {

	h.Color = RED
	h.Left.Color = BLACK
	h.Right.Color = BLACK
}

func compareTo(node, key rune) int {
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

//ShowList : show all list of item
func (bts *Tree) ShowAllLeft() {

	tree := bts.root

	for tree != nil {
		fmt.Printf("%+s ->", string(tree.Key))
		tree = tree.Left
	}
	fmt.Println()
}

//ShowList : show all list of item
func (bts *Tree) ShowAllRight() {

	tree := bts.root

	for tree != nil {
		fmt.Printf("%+s ->", string(tree.Key))
		tree = tree.Right
	}
	fmt.Println()
}
