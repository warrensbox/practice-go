/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isSymmetric(root *TreeNode) bool {
    
    
    return mirror(root,root)
    
}


func mirror(node1 *TreeNode, node2 *TreeNode) bool {
    
    if (node1 == nil && node2 == nil ) {
        return true
    }
    
    if (node1 == nil || node2 == nil ) {
        return false
    }
    
    return (node1.Val == node2.Val) && mirror(node1.Left,node2.Right) && mirror(node1.Right,node2.Left)
}