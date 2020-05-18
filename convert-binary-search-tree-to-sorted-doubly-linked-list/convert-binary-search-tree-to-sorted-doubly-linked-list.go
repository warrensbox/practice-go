package main

func main() {

}

//Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func treeToDoublyList(root *Node) *Node {

	if root == nil {
		return root
	}

	var first, prev *Node

	//function declaration
	var dfs func(node *Node)

	//implement function
	dfs = func(node *Node) {

		if node != nil {
			dfs(node.Left)

			if prev != nil {
				prev.Right = node
				node.Left = prev
			} else {
				first = node
			}

			prev = node

			dfs(node.Right)
		}
	}

	//call function
	dfs(root)

	prev.Right = first
	first.Left = prev

	return first

}
