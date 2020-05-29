/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func largestValues(root *TreeNode) []int {
    
    output := helper(root)
    
    return output
}


func helper(root *TreeNode) []int {
    list := make([]int, 0)
    
    if root == nil {
        return list
    }
    

    queue := Queue{}
	queue.EnQueue(*root)
    
   for queue.Len() > 0 {
       size := queue.Len()
       //fmt.Println(size)
       min := math.MinInt64
       for i := 0; i < size; i++ {
           node := queue.DeQueue()
           if node.Val > min {
               min = node.Val
           }
           
           
           if node.Left != nil {
				queue.EnQueue(*node.Left)
			}

			if node.Right != nil {
				queue.EnQueue(*node.Right)
			}
          
       }
       list = append(list, min)
       fmt.Println(list)
   }
    
    return list
    
}


type Queue []TreeNode

func (q *Queue) EnQueue(n TreeNode) {
	*q = append(*q, n)
}

func (q *Queue) DeQueue() (n TreeNode) {
	n = (*q)[0]
	*q = (*q)[1:]
	return n
}

func (q *Queue) Len() int {
	return len(*q)
}
