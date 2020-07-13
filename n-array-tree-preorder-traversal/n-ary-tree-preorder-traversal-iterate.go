//https://www.youtube.com/watch?v=BHB0B1jFKQc&t=138s

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

 //IS THIS BFS?
func preorder(root *Node) []int {

	res := make([]int, 0)

	if root == nil {
		return res
	}

	stack := []*Node{}

	for len(stack) > 0 {
		r := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, r.Val)

		tmp := []*Node{}
		for _, v := range r.Children {
			tmp = append([]*Node{v}, tmp...)
		}
		stack = append(stack, tmp...)
	}
	return res
}
