package main

//https://www.youtube.com/watch?v=aaSFzFfOQ0o

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func main(){

}


func delNodes(root *TreeNode, toDelete []int) []*TreeNode {

	del := make(map[int]bool)
	remaining := make([]*TreeNode,0,10)

	for _,val := range toDelete {
		del[val] =true
	}

	if root != nil && !del[root.Val] {
		remaining = append(remaining,root)
	}

	helper(root, del, &remaining)
	return remaining
}


func helper(root *TreeNode, delMap map[int]bool, remaining *[]*TreeNode) *TreeNode{

	if root == nil {
		return nil
	}

	root.Left = helper(root.Left, delMap, remaining)
	root.Right = helper(root.Right, delMap, remaining)
	
	if _, found := delMap[root.Val]; found {

		if root.Left != nil && !delMap[root.Left.Val]{
			*remaining = append(*remaining,root.Left)
		}

		if root.Right != nil && !delMap[root.Right.Val]{
			*remaining = append(*remaining,root.Right)
		}

		return nil
	}

	return root
}