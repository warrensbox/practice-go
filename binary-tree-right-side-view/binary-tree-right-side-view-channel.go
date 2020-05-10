

//https://www.youtube.com/watch?v=NjdOhYKjFrU
//https://www.youtube.com/watch?v=jCqIr_tBLKs

func rightSideView(root *TreeNode) []int {
	ret := make([]int, 0)

	q := make(chan TreeNode, 100)

	if root != nil {
		q <- *root
	}

	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			cur := <-q
			if cur.Left != nil {
				q <- *(cur.Left)
			}
			if cur.Right != nil {
				q <- *(cur.Right)
			}

			if i == size-1 {
				ret = append(ret, cur.Val)
			}
		}
	}
	return ret
}
