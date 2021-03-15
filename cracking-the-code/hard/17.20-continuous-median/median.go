package main

//SKIPPING
import "fmt"

func main() {
}

/* MAX HEAP */
type MaxHeap struct {
	arr []int
}

func NewMaxHeap() *MaxHeap {
	maxHeap := MaxHeap{}
	maxHeap.arr = make([]int, 0)
	maxHeap.insert(0)
	return &maxHeap
}

func (pq *MaxHeap) insert(v int) {

	n := len(pq.arr) - 1
	n++
	pq.arr = append(pq.arr, v)
	pq.swim(n)
}

func (pq *MaxHeap) deleteMax() int {

	n := len(pq.arr) - 1
	max := pq.arr[1]
	pq.exchange(1, n)
	pq.arr = pq.arr[:n]
	pq.sink(1)
	return max
}

func (pq *MaxHeap) swim(k int) {
	for k > 1 && pq.arr[k/2] < pq.arr[k] {
		pq.exchange(k/2, k)
		k = k / 2
	}
}

func (pq *MaxHeap) sink(k int) {
	fmt.Println("sinking")
	n := len(pq.arr) - 1
	for 2*k <= n {
		i := 2 * k
		if i < n && pq.arr[i] < pq.arr[i+1] {
			i++
		}
		if pq.arr[k] > pq.arr[i] {
			break
		}
		pq.exchange(k, i)
		k = i
	}
}

/* MAX HEAP */
type MinHeap struct {
	arr []int
}

func NewMinHeap() *MinHeap {
	minHeap := MinHeap{}
	minHeap.arr = make([]int, 0)
	minHeap.insert(0)
	return &minHeap
}

func (pq *MinHeap) insert(v int) {

	n := len(pq.arr) - 1
	n++
	pq.arr = append(pq.arr, v)
	pq.swimMin(n)
}

func (pq *MinHeap) deleteMin() int {

	n := len(pq.arr) - 1
	max := pq.arr[1]
	pq.exchangeMin(1, n)
	pq.arr = pq.arr[:n]
	pq.sinkMin(1)
	return max
}

func (pq *MinHeap) swimMin(k int) {
	for k > 1 && pq.arr[k/2] > pq.arr[k] {
		pq.exchangeMin(k/2, k)
		k = k / 2
	}
}

func (pq *MinHeap) sinkMin(k int) {
	fmt.Println("sinking")
	n := len(pq.arr) - 1
	for 2*k <= n {
		i := 2 * k
		if i < n && pq.arr[i] > pq.arr[i+1] {
			i++
		}
		if pq.arr[k] < pq.arr[i] {
			break
		}
		pq.exchangeMin(k, i)
		k = i
	}
}

/* Helper */
func (pq *MaxHeap) exchange(i, j int) {
	pq.arr[i], pq.arr[j] = pq.arr[j], pq.arr[i]
}

func (pq *MinHeap) exchangeMin(i, j int) {
	pq.arr[i], pq.arr[j] = pq.arr[j], pq.arr[i]
}
