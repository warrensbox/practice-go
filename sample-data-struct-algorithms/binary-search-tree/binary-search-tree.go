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
	Count       int //added later
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

	bts.delete('a')

	rootsize := bts.size()
	fmt.Println("rootsize:", rootsize)

	rank := bts.rank('e')
	fmt.Println("rank:", rank)

	floor := bts.floor('d')
	fmt.Println("floor", string(floor))

	queue := list.New()
	inOrder(bts.root, queue)

	sm := bts.smallest()
	fmt.Println("smallest", string(sm.Key))

	lg := bts.largest()
	fmt.Println("largest", string(lg.Key))

	// Loop over container list.
	for temp := queue.Front(); temp != nil; temp = temp.Next() {
		fmt.Println("Key", temp.Value)
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
		newNode.Count = 1
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
	node.Count = 1 + Size(node.Left) + Size(node.Right)
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

func (bts *Tree) rank(key rune) int {
	root := bts.root
	return Rank(root, key)
}

func Rank(node *Node, key rune) int {

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

func (bts *Tree) deleteMin() {
	root := bts.root
	root = DeleteMin(root)
}

func DeleteMin(node *Node) *Node {

	if node.Left == nil {
		return node.Right
	}
	node.Left = DeleteMin(node.Left)
	node.Count = 1 + Size(node.Left) + Size(node.Right)
	return node
}

func (bts *Tree) delete(key rune) {
	root := bts.root
	root = Delete(root, key)
}

func Delete(node *Node, key rune) *Node {

	if node == nil {
		return nil
	}
	cmp := compareTo(key, node.Key)
	if cmp < 0 {
		node.Left = Delete(node.Left, key) //serching for the key
	} else if cmp > 0 {
		node.Right = Delete(node.Right, key) //serching for the key
	} else {
		if node.Right == nil { //no right child
			return node.Left
		}
		if node.Left == nil { //no left child
			return node.Right
		}

		//replace with succesor
		t := node
		node = Smallest(t.Right)
		node.Left = t.Left
	}
	node.Count = Size(node.Left) + Size(node.Right)
	return node
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

func Smallest(node *Node) *Node {

	current := node
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

func (bts *Tree) floor(key rune) rune {

	root := bts.root
	node := Floor(root, key)
	if node == nil {
		return rune(0)
	}
	return node.Key
}

func Floor(node *Node, key rune) *Node {

	if node == nil {
		return nil
	}
	cmp := compareTo(key, node.Key)
	if cmp == 0 {
		return node
	}
	if cmp < 0 {
		return Floor(node.Left, key)
	}
	t := Floor(node.Right, key)
	if t != nil {
		return t
	}
	return node
}

func (bts *Tree) size() int {
	root := bts.root
	return Size(root)
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
