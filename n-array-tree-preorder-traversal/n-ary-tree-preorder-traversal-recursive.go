//https://www.youtube.com/watch?v=BHB0B1jFKQc&t=138s

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {

	res := make([]int, 0)

	recursive(root, &res)

	return res

}

func recursive(root *Node, res *[]int) {

	if root == nil {
		return
	}

	*res = append(*res, root.Val)
	for _, val := range root.Children {

		recursive(val, res)
	}
}