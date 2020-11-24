package main

import (
	"fmt"
)

type Array struct {
	pq []int
}

func main() {

	a := Array{}
	a.pq = []int{0}

	a.insert(16)

	a.insert(14)
	a.insert(13)
	a.insert(12)
	a.insert(9)
	a.insert(15)
	a.insert(17)
	fmt.Println(a.pq)

	max := a.delMax()
	fmt.Println("max", max)
	fmt.Println(a.pq)
}

func (a *Array) delMax() int {

	n := len(a.pq) - 1
	max := a.pq[1]   //retrieve max from top
	a.exchange(1, n) // exchange top item with last item
	a.pq = a.pq[:n]  //remove last item
	a.sink(1)        // restrore heap property

	return max
}

//AA13F019
func (a *Array) insert(v int) {

	n := len(a.pq) - 1
	n++ //added new element
	a.pq = append(a.pq, v)
	a.swim(n)
}

func (a *Array) swim(k int) {

	for k > 1 && a.pq[k/2] < a.pq[k] {
		a.exchange(k/2, k)
		k = k / 2
	}
}

func (a *Array) sink(k int) {

	n := len(a.pq) - 1
	for 2*k <= n {
		j := 2 * k
		if j < n && a.pq[j] < a.pq[j+1] {
			j++ // moves to the left (between child nodes, this is bigger)
		}
		if a.pq[k] > a.pq[j] {
			break //if parent is bigger than child, then break the swim
		}
		a.exchange(k, j)
		k = j
	}
}

func (a *Array) exchange(i, min int) {
	a.pq[i], a.pq[min] = a.pq[min], a.pq[i]
}
