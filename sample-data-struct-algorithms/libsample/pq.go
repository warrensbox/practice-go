package libsample

import (
	"container/heap"
)

// An Item is something we manage in a priority queue.
// type Item struct {
// 	V      int     // The value of the item; arbitrary.
// 	Weight float32 // The priority of the item in the queue.

// 	// The index is needed by update and is maintained by the heap.Interface methods.
// 	index int // The index of the item in the heap.
// }

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*DiEdge

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	//fmt.Println("N", n)
	item := x.(*DiEdge)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) Update(item *DiEdge, v int, weight float32) {
	item.V = v
	item.Priority = weight
	heap.Fix(pq, item.index)
}

// see if pq contains
func (pq *PriorityQueue) Contains(v int) bool {
	arr := *pq
	for i := 0; i < len(arr); i++ {
		if arr[i].V == v {
			return true
		}
	}
	return false
}
