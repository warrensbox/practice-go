package libsample

type EdgeWeightedGraph struct {
	nodes int
	edges []*EdgeBag
}

func NewEdgeWeightedGraph(vertices int) *EdgeWeightedGraph {
	var ewg EdgeWeightedGraph
	ewg.nodes = vertices
	ewg.edges = make([]*EdgeBag, vertices)
	for i := 0; i < vertices; i++ {
		ewg.edges[i] = NewEdgeBag() // arr[0] = bag{1,2,3}; arr[1] = bag{0,5};
	}
	return &ewg
}

//similiar to connecting two vertices (may loopback)
func (ewg *EdgeWeightedGraph) AddEdge(e *Edge) {

	v := e.Either()
	w := e.Other(v)
	ewg.edges[v].EdgeInsert(*e)
	if v != w {
		ewg.edges[w].EdgeInsert(*e)
	}
}

//iterator for edges adjacent to v
func (ewg *EdgeWeightedGraph) Adjacent(v int) map[Edge]int {
	return ewg.edges[v].EdgeVertices()
}

//return number of vertices in ewg
func (ewg *EdgeWeightedGraph) NumofVertices() int {
	return ewg.nodes
}

//get all edges of a edge-weighted-graph
func (ewg *EdgeWeightedGraph) Edges() map[Edge]int {

	bag := NewEdgeBag()
	for v := 0; v < ewg.NumofVertices(); v++ {
		for e := range ewg.Adjacent(v) {
			if e.Other(v) > v {
				bag.EdgeInsert(e)
			}
		}
	}
	return bag.data
}
