package main

type TreeNode struct {
	item int
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
