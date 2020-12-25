package libsample

// Bag data structure (multiset).
type DiEdgeBag struct {
	size int
	data map[DiEdge]int
}

// Creates a new empty bag.
func NewDiEdgeBag() *DiEdgeBag {
	return &DiEdgeBag{0, make(map[DiEdge]int)}
}

// Inserts an element into the bag.
func (b *DiEdgeBag) DiEdgeInsert(val DiEdge) {
	b.data[val]++
	b.size++
}

// Returns the total number of elemens in the bag.
func (b *DiEdgeBag) DiEdgeSize() int {
	return b.size
}

// Counts the number of occurances of an element in the bag.
func (b *DiEdgeBag) DiEdgeCount(val DiEdge) int {
	return b.data[val]
}

// Executes a function for every element in the bag.
func (b *DiEdgeBag) DiEdgeVertices() map[DiEdge]int {

	return b.data
}

// Clears the contents of a bag.
func (b *DiEdgeBag) DiEdgeReset() {
	b.size = 0
	b.data = make(map[DiEdge]int)
}
