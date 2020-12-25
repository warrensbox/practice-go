package libsample

type EdgeWeightedDiGraph struct {
	v     int
	e     int
	edges []*DiEdgeBag
}

func NewEdgeWeightedDiGraph(vertices int) *EdgeWeightedDiGraph {
	var ewg EdgeWeightedDiGraph
	ewg.v = vertices
	ewg.e = 0
	ewg.edges = make([]*DiEdgeBag, vertices)
	for i := 0; i < vertices; i++ {
		ewg.edges[i] = NewDiEdgeBag() // arr[0] = bag{1,2,3}; arr[1] = bag{0,5};
	}
	return &ewg
}

//return number of vertices in ewg
func (ewg *EdgeWeightedDiGraph) NumofVertices() int {
	return ewg.v
}

//similiar to connecting two vertices (may loopback)
func (ewg *EdgeWeightedDiGraph) AddEdge(e *DiEdge) {
	ewg.edges[e.From()].DiEdgeInsert(*e)
	ewg.e++
}

//iterator for edges adjacent to v
func (ewg *EdgeWeightedDiGraph) Adjacent(v int) map[DiEdge]int {
	return ewg.edges[v].DiEdgeVertices()
}

//get all edges of a edge-weighted-graph
func (ewg *EdgeWeightedDiGraph) Edges() map[DiEdge]int {

	bag := NewDiEdgeBag()
	for v := 0; v < ewg.NumofVertices(); v++ {
		for e := range ewg.Adjacent(v) {
			bag.DiEdgeInsert(e)
		}
	}
	return bag.data
}
