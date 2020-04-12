/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }
 
func longestConsecutive(root *TreeNode) int {
    
    var parent *TreeNode
    maxLength := 0
    length := 0
    
    
    dfs(root,parent,length,&maxLength)
        
    return maxLength
}

func dfs(p *TreeNode, parent *TreeNode, length int,maxLength *int) {
 
    if p == nil {
        return
    }
    
    if (parent != nil && p.Val == parent.Val+1) {
        length = length +1
    }else{
        length = 1
    }

    (*maxLength) = max((*maxLength),length)
    fmt.Println(*maxLength)
    dfs(p.Left,p,length,maxLength)
    dfs(p.Right,p,length,maxLength)
    
}

func max(a, b int) int{
    if a > b {
        return a
    }
    
    return b
}