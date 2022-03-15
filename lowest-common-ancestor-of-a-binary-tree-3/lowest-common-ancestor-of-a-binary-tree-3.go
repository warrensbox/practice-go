package main

import "fmt"

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

func lowestCommonAncestor(p *Node, q *Node) *Node {

	if p == q {
		return p
	}

	ancestorsOfP := map[*Node]bool{}

	for p != nil {
		ancestorsOfP[p] = true
		p = p.Parent
	}

	fmt.Println(ancestorsOfP)

	for q != nil {
		if ancestorsOfP[q] {
			return q
		}
		q = q.Parent
	}

	return nil
}
