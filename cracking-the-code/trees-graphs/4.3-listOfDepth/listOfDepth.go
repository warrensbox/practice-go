package main

import "container/list"

type Node struct {
	value int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func createLevelLinkedList(t *Tree) []*list.List {

	result := []*list.List{}
	current := list.New()
	root := t.root

	if t.root != nil {
		current.PushBack(root)
	}

	for current.Len() > 0 {
		result = append(result, current)

		parents := current
		current = list.New()

		for parent := parents.Front(); parent != nil; parent = parent.Next() {

			if parent.Value.(*Node).left != nil {
				current.PushBack(parent.Value.(*Node).left)
			}

			if parent.Value.(*Node).right != nil {
				current.PushBack(parent.Value.(*Node).right)
			}
		}
	}

	return result
}
