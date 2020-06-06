import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func verticalTraversal(root *TreeNode) [][]int {

	output := helper(root)

	return output
}

func helper(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}

	queue := []XTreeNode{XTreeNode{root, 0}}
	dict := make(map[int][]int)
	dict[0] = append(dict[0], root.Val)

	min := 0 //we need this to adjust the index of the list of array later

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {

			node := (queue)[0] //pop

			queue = (queue)[1:] //assign

			if node.X < min {
				min = node.X
			}

			if node.t.Left != nil {
				nodeLeftVal := node.t.Left.Val
				counterVal := node.X - 1
				//fmt.Println(counterVal)
				dict[counterVal] = append(dict[counterVal], nodeLeftVal)
				queue = append(queue, XTreeNode{node.t.Left, counterVal})
			}

			if node.t.Right != nil {
				nodeRightVal := node.t.Right.Val
				counterVal := node.X + 1
				//fmt.Println(counterVal)
				dict[counterVal] = append(dict[counterVal], nodeRightVal)
				queue = append(queue, XTreeNode{node.t.Right, counterVal})
			}

		}

	}

	list := make([][]int, len(dict)) //create a list of array

	for key, value := range dict {
		list[key-min] = value
	}

	fmt.Println(dict)
	return list
}

type XTreeNode struct {
	t *TreeNode
	X int
}
