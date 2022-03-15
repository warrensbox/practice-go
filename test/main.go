package main

import "fmt"

/*
   Given a reference of a node in a connected undirected graph. Return a deep
   copy of the graph.

   A deep copy is a full clone of the all nodes to a new memory location.
*/

func main() {
	var graph = NewNode("A").AddChildren("B", "C", "D")
	graph.Children[0].AddChildren("E").AddChildren("F")
	graph.Children[2].AddChildren("G").AddChildren("H")
	graph.Children[0].Children[1].AddChildren("I").AddChildren("J")
	graph.Children[2].Children[0].AddChildren("K")

	cloneGraph(graph)
}

type Node struct {
	Name     string
	Children []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	var cloned = &Node{Name: node.Name, Children: nil}
	fmt.Println("cloned", cloned)
	for _, n := range node.Children {
		cn := cloneGraph(n)
		fmt.Println("cloned neighbor", cn)
		cloned.Children = append(cloned.Children, cn)
		fmt.Println("new neighbors", cloned.Children)
	}

	return cloned

}

func NewNode(name string) *Node {
	return &Node{
		Name:     name,
		Children: []*Node{},
	}
}

func (n *Node) AddChildren(names ...string) *Node {
	for _, name := range names {
		child := Node{Name: name}
		n.Children = append(n.Children, &child)
	}
	return n
}

/*
node2 = A

node2.Neighbors = append(neighbors, B)
node2.Neighbors = append(neighbors, C)


node2.neighbors = append(..., B')
node2.neighbors = append(..., C')
*/
