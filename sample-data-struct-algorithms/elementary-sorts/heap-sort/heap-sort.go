package main

import (
	"fmt"
)

type Array struct {
	pq []int
}

func main() {

	a := Array{}
	a.pq = []int{0, 12, 23, 34, 54, 65, 21, 76, 11, 2, 3, 4}

	fmt.Println(a.pq)
	a.sort()
	fmt.Println(a.pq)
}

func (a *Array) sort() {

	n := len(a.pq) - 1
	for k := n / 2; k >= 1; k-- { //heap ordered
		a.sink(k, n)
	}
	fmt.Println(a.pq)
	for n >= 1 {
		fmt.Println("n", n)
		a.exchange(1, n)
		n--
		a.sink(1, n)

	}
}

func (a *Array) sink(k, n int) {

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
