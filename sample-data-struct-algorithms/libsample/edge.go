package libsample

// Data structure for representing an edge.
type Edge struct {
	v, w   int
	weight float32
}

func NewEdge(v, w int, weight float32) *Edge {
	var edge Edge
	edge.v = v
	edge.w = w
	edge.weight = weight

	return &edge
}

//either endpoint
func (e *Edge) Either() int {
	return e.v
}

//other endpoint
func (e *Edge) Other(vertex int) int {
	if vertex == e.v {
		return e.w
	}
	return e.v
}

//compare edges by weight
func (e *Edge) CompareTo(that Edge) int {

	if e.weight < that.weight {
		return -1
	} else if e.weight > that.weight {
		return 1
	} else {
		return 0
	}

}

//Weight endpoint
func (e *Edge) Weight() float32 {
	return e.weight
}
