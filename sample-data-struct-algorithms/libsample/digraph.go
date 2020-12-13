package libsample

// Data structure for representing a graph.
type Digraph struct {
	nodes int
	edges []*Bag
}

//create new undirected Digraph
func NewDigraph(vertices int) *Digraph {
	var digraph Digraph
	digraph.nodes = vertices
	digraph.edges = make([]*Bag, vertices)
	for i := 0; i < vertices; i++ {
		digraph.edges[i] = NewBag() // arr[0] = bag{1,2,3}; arr[1] = bag{0,5};
	}
	return &digraph
}

//connect two vertices (may loopback)
func (g *Digraph) Connect(v, w int) {
	g.edges[v].Insert(w)
}

//iterator for vertices adjacent to v
func (g *Digraph) Adjacent(v int) map[interface{}]int {

	return g.edges[v].Vertices()
}

//return number of vertices in Digraph
func (g *Digraph) NumofVertices() int {
	return g.nodes
}
