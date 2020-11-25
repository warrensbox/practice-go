package main

import "fmt"

//define Node
type Node struct {
	Key         rune
	Value       int
	Left, Right *Node
}

type Tree struct {
	root *Node
}

func main() {
	bts := Tree{}
	bts.pute('s', 1)
	bts.pute('e', 5)
	bts.pute('a', 5)
	bts.pute('x', 3)
	bts.ShowAllLeft()
	//fmt.Println(bts.root.Value)
}

func (bts *Tree) pute(key rune, value int) {

	root := bts.root

	bts.root = Put(root, key, value)

}

//Put : recursion
func Put(node *Node, key rune, value int) *Node {

	fmt.Println("key", string(key))
	fmt.Println("val", value)
	if node == nil {
		fmt.Println("new node")
		var newNode Node
		newNode.Key = key
		newNode.Value = value
		return &newNode
	}
	cmp := compareTo(key, node.Key) // "key" compared to "node's key"
	if cmp < 0 {
		fmt.Println("going left")
		node.Left = Put(node.Left, key, value)
	} else if cmp > 0 {
		fmt.Println("going right")
		node.Right = Put(node.Left, key, value)
	} else if cmp == 0 {
		node.Value = value
	}

	fmt.Println("Reached here")
	return node
}

func (bts *Tree) get(key rune) int {

	node := bts.root
	for node != nil {
		cmp := compareTo(key, node.Key) // "key" compared to "node's key"
		if cmp < 0 {
			node = node.Left
		} else if cmp > 0 {
			node = node.Right
		} else if cmp == 0 {
			return node.Value
		}
	}
	return -1
}

func compareTo(node, key rune) int {
	if node > key {
		return 1
	} else if node < key {
		return -1
	}
	return 0
}

func (bts *Tree) inOrder() {

}

//ShowList : show all list of item
func (bts *Tree) ShowAllLeft() {

	tree := bts.root

	for tree != nil {
		fmt.Printf("%+s ->", string(tree.Key))
		fmt.Sprint(tree.Key)
		tree = tree.Left
	}
	fmt.Println()
}
