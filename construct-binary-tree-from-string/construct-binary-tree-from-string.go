package main

import (
	"fmt"
	"strconv"
)

/*
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func str2tree(s string) *TreeNode {

	if len(s) == 0 {
		return nil
	}

	n := len(s)
	i := 0
	pos := 1
	var stack []*TreeNode
	for i < n {
		if s[i] >= '0' && s[i] <= '9' {
			fmt.Println("top", string(s[i]))
			tmp := string(s[i])
			for i+1 < n && s[i+1] >= '0' && s[i+1] <= '9' {
				fmt.Println("top2", string(s[i+1]))
				tmp += string(s[i+1])
				i++
			}
			t, _ := strconv.Atoi(string(tmp))
			t = (t * pos)
			fmt.Println("t", t)
			node := &TreeNode{Val: t}
			pos = 1
			if len(stack) > 0 {
				fmt.Println("entering loop")
				r := stack[len(stack)-1]
				fmt.Println("r.val", r.Val)
				if r.Left != nil {
					fmt.Println("Left is NOT nil")
					r.Right = node
				} else {
					fmt.Println("Right is NOT nil")
					r.Left = node
				}
			}
			fmt.Println("node.Val", node.Val)
			fmt.Println("stack", stack)
			stack = append(stack, node)
		} else if s[i] == ')' {
			fmt.Println("2nd", string(s[i]))
			fmt.Println("stack", stack)
			fmt.Println("popping", stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		} else if s[i] == '-' {
			fmt.Println("3rd", string(s[i]))
			pos = -1
		}
		i++
	}

	return stack[0]
}
