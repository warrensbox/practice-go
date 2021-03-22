package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	if root == nil {
		return []int{}
	}

	data := make(map[*TreeNode]*TreeNode)

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Left != nil {
			data[node.Left] = node
		}

		if node.Right != nil {
			data[node.Right] = node
		}

		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	var ans []int
	visited := make(map[*TreeNode]bool)

	var queue []*TreeNode

	queue = append(queue, target)
	var depth int
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if depth == K {
				ans = append(ans, node.Val)
				continue
			}

			parent := data[node]

			if parent != nil && !visited[parent] {
				queue = append(queue, parent)
			}

			if node.Left != nil && !visited[node.Left] {
				queue = append(queue, node.Left)
			}

			if node.Right != nil && !visited[node.Right] {
				queue = append(queue, node.Right)
			}

			visited[node] = true
		}
		depth++
	}

	return ans
}
