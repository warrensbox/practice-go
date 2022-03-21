package main

func main() {
	// var graph = NewNode(1).AddChildren("B", "C", "D")
	// graph.Neighbors[0].AddChildren("E").AddChildren("F")
	// graph.Neighbors[2].AddChildren("G").AddChildren("H")
	// graph.Neighbors[0].Neighbors[1].AddChildren("I").AddChildren("J")
	// graph.Neighbors[2].Neighbors[0].AddChildren("K")

	// cloneGraph(graph)
}

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {

	if node == nil {
		return nil
	}

	visited := make(map[int]*Node)

	var walk func(node *Node) *Node
	walk = func(node *Node) *Node {
		//base case
		if n, ok := visited[node.Val]; ok {
			return n
		}

		cloned := &Node{Val: node.Val, Neighbors: nil}
		visited[node.Val] = cloned
		for _, neigh := range node.Neighbors {
			cn := walk(neigh)
			cloned.Neighbors = append(cloned.Neighbors, cn)
		}

		return cloned
	}

	return walk(node)
}

func NewNode(val int) *Node {
	return &Node{
		Val:       val,
		Neighbors: []*Node{},
	}
}

func (n *Node) AddChildren(names ...int) *Node {
	for _, val := range names {
		child := Node{Val: val}
		n.Neighbors = append(n.Neighbors, &child)
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
