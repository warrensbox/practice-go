package main

/**
 * Definition for a Node.
 */
type Node struct {
	Val  int
	Next *Node
}

func insert(aNode *Node, x int) *Node {

	var node Node
	node.Val = x

	//if node is empty
	if aNode == nil {
		aNode = &node
		node.Next = aNode
		return &node
	}

	//if there is only one node
	if aNode.Next == aNode {
		aNode.Next = &node
		node.Next = aNode
		return aNode
	}

	//when there is more than one node || [3,3,3] =>
	currNode := aNode
	for {
		nextNode := currNode.Next
		if nextNode.Val >= currNode.Val {
			if (x >= currNode.Val && x < nextNode.Val) || nextNode == aNode {
				currNode.Next = &node
				node.Next = nextNode
				break
			}

		} else {
			if x >= currNode.Val || x <= nextNode.Val {
				currNode.Next = &node
				node.Next = nextNode
				break
			}
		}

		currNode = nextNode
	}
	return aNode
}
