/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func constructFromPrePost(pre []int, post []int) *TreeNode {
    root := TreeNode{pre[0],nil,nil} //the parent node is the first node in pre
    
    if len(pre) == 1 {
        return &root
    }
    
    pre = pre[1:] //remove root from the list
    
    hold := pre[0] //pre[0] is the left child
    
    leftChildren := []int{}
    
    for i, val := range post{
        if val == hold {
            leftChildren = post[:i+1]//iterate through post until you see the root. everything before that are left children
            post = post[i+1: ]//everything to the right are the right children
            break
        }
    }
    
    if len(leftChildren) >0{
        root.Left = constructFromPrePost(pre[:len(leftChildren)], leftChildren)
    }
    
    pre = pre[len(leftChildren):] //cut all left children out of pre
    if len(pre) > 0 {
        root.Right = constructFromPrePost(pre, post)
    }
    
    return &root
}