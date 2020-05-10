package main

//https://www.youtube.com/watch?v=NjdOhYKjFrU
//https://www.youtube.com/watch?v=jCqIr_tBLKs

func main() {

}

func rightSideView(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	rightView := make([]int, 0)
	queue := Queue{}
	queue.EnQueue(*root)
	for queue.Len() > 0 {

		size := queue.Len()
		for i := 0; i < size; i++ {
			node := queue.DeQueue()

			if i == size-1 {
				rightView = append(rightView, node.Val)
			}

			if node.Left != nil {
				queue.EnQueue(*node.Left)
			}

			if node.Right != nil {
				queue.EnQueue(*node.Right)
			}
		}
	}
	return rightView
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
