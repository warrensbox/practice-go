/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	left, right := 0, 0
	walk(root, &left, &right, false, false, x)
	// if any subtree nodes more than half, WIN
	if left > n/2 || right > n/2 {
		return true
	}
	// if x root tree nodes less than or equal to n / 2, WIN
	if right+left+1 <= n/2 {
		return true
	}
	return false
}

func walk(node *TreeNode, leftCount, rightCount *int, isLeft, isRight bool, x int) {
	if node == nil {
		return
	}
	if node.Val == x {
		walk(node.Left, leftCount, rightCount, true, false, x)
		walk(node.Right, leftCount, rightCount, false, true, x)
		return
	}
	if isLeft {
		*leftCount++
	}
	if isRight {
		*rightCount++
	}
	walk(node.Left, leftCount, rightCount, isLeft, isRight, x)
	walk(node.Right, leftCount, rightCount, isLeft, isRight, x)
}