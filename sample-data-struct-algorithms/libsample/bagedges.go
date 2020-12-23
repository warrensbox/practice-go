package libsample

// Bag data structure (multiset).
type EdgeBag struct {
	size int
	data map[Edge]int
}

// Creates a new empty bag.
func NewEdgeBag() *EdgeBag {
	return &EdgeBag{0, make(map[Edge]int)}
}

// Inserts an element into the bag.
func (b *EdgeBag) EdgeInsert(val Edge) {
	b.data[val]++
	b.size++
}

// Returns the total number of elemens in the bag.
func (b *EdgeBag) EdgeSize() int {
	return b.size
}

// Counts the number of occurances of an element in the bag.
func (b *EdgeBag) EdgeCount(val Edge) int {
	return b.data[val]
}

// Executes a function for every element in the bag.
func (b *EdgeBag) EdgeVertices() map[Edge]int {

	return b.data
}

// Clears the contents of a bag.
func (b *EdgeBag) EdgeReset() {
	b.size = 0
	b.data = make(map[Edge]int)
}
