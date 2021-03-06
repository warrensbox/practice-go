package libsample

// Data structure for representing an edge.
type DiEdge struct {
	V        int
	W        int
	Priority float32
}

func NewDiEdge(v, w int, weight float32) *DiEdge {
	var edge DiEdge
	edge.V = v
	edge.W = w
	edge.Priority = weight

	return &edge
}

//from endpoint
func (e *DiEdge) From() int {
	return e.V
}

//to endpoint
func (e *DiEdge) To() int {
	return e.W
}

//Weight endpoint
func (e *DiEdge) Weight() float32 {
	return e.Priority
}

//compare edges by weight
func (e *DiEdge) CompareTo(that DiEdge) int {

	if e.Priority < that.Priority {
		return -1
	} else if e.Priority > that.Priority {
		return 1
	} else {
		return 0
	}

}
