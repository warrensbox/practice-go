package libsample

// Data structure for representing a graph.
type Graph struct {
	nodes int
	edges []*Bag
}

//create new undirected graph
func NewGraph(vertices int) *Graph {
	var graph Graph
	graph.nodes = vertices
	graph.edges = make([]*Bag, vertices)
	for i := 0; i < vertices; i++ {
		graph.edges[i] = NewBag() // arr[0] = bag{1,2,3}; arr[1] = bag{0,5};
	}
	return &graph
}

//connect two vertices (may loopback)
func (g *Graph) Connect(v, w int) {
	g.edges[v].Insert(w)
	if v != w {
		g.edges[w].Insert(v)
	}
}

//iterator for vertices adjacent to v
func (g *Graph) Adjacent(v int) map[interface{}]int {

	return g.edges[v].Vertices()
}

//return number of vertices in graph
func (g *Graph) NumofVertices() int {
	return g.nodes
}
