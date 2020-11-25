package main

import (
	"container/list"
	"fmt"
)

//define Node
type Node struct {
	Key         rune
	Value       int
	Left, Right *Node
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
	bts.put('x', 3)
	bts.put('r', 5)
	bts.put('h', 5)
	bts.put('m', 3)
	bts.put('c', 5)
	queue := list.New()
	inOrder(bts.root, queue)

	sm := bts.smallest()
	fmt.Println("smallest", string(sm.Key))

	lg := bts.largest()
	fmt.Println("largest", string(lg.Key))

	// Loop over container list.
	for temp := queue.Front(); temp != nil; temp = temp.Next() {
		fmt.Println(temp.Value)
	}
	// bts.ShowAllLeft()
	// bts.ShowAllRight()
}

func (bts *Tree) put(key rune, value int) {

	root := bts.root

	bts.root = Put(root, key, value)

}

//Put : recursion
func Put(node *Node, key rune, value int) *Node {

	if node == nil {
		var newNode Node
		newNode.Key = key
		newNode.Value = value
		return &newNode
	}
	cmp := compareTo(key, node.Key) // "key" compared to "node's key"
	if cmp < 0 {
		node.Left = Put(node.Left, key, value)
	} else if cmp > 0 {
		node.Right = Put(node.Right, key, value)
	} else if cmp == 0 {
		node.Value = value
	}

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

func inOrder(node *Node, q *list.List) {

	if node == nil {
		return
	}

	inOrder(node.Left, q)
	q.PushBack(string(node.Key))
	inOrder(node.Right, q)
}

func compareTo(node, key rune) int {
	if node > key {
		return 1
	} else if node < key {
		return -1
	}
	return 0
}

//smallest key
func (bts *Tree) smallest() *Node {

	current := bts.root
	sm := &Node{}
	for current != nil {
		sm = current
		current = current.Left
	}

	return sm
}

//largest key
func (bts *Tree) largest() *Node {

	current := bts.root
	lg := &Node{}
	for current != nil {
		lg = current
		current = current.Right
	}

	return lg
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
