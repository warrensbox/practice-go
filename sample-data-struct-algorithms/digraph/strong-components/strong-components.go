package main

import (
	"fmt"

	"github.com/warrensbox/practice-go/sample-data-struct-algorithms/libsample"
)

func main() {

	g := libsample.NewDigraph(13)
	g.Connect(4, 2)
	g.Connect(2, 3)
	g.Connect(3, 2)
	g.Connect(6, 0)
	g.Connect(0, 1)
	g.Connect(2, 3)
	g.Connect(11, 12)
	g.Connect(12, 9)
	g.Connect(9, 10)
	g.Connect(9, 11)
	g.Connect(7, 9)
	g.Connect(10, 12)
	g.Connect(11, 4)
	g.Connect(4, 3)
	g.Connect(3, 5)
	g.Connect(6, 8)
	g.Connect(8, 6)
	g.Connect(5, 4)
	g.Connect(0, 5)
	g.Connect(6, 4)
	g.Connect(6, 9)
	g.Connect(7, 6)

	d := KosarajuSharirSCC(g)
	connected := d.isStronglyConnected(0, 6)
	fmt.Println(connected)
	connected = d.isStronglyConnected(0, 7)
	fmt.Println(connected)

	fmt.Println(d.id)
}

type DFS struct {
	marked      []bool
	id          []int
	reversePost []int //reserve post order
	count       int
}

func KosarajuSharirSCC(g *libsample.Digraph) *DFS {

	var d2 DFS
	d2.marked = make([]bool, g.NumofVertices())
	d2.id = make([]int, g.NumofVertices())
	d2.count = 0

	dR := DepthFirstSearch(g.Reverse())
	fmt.Println("dg", dR.reversePost)

	for pop := len(dR.reversePost) - 1; pop > 0; pop-- {
		v := dR.reversePost[pop]

		if !d2.marked[v] {
			d2.dfs(g, v)
			d2.count++
		}
	}

	return &d2

}

func DepthFirstSearch(g *libsample.Digraph) *DFS {

	var d DFS
	d.marked = make([]bool, g.NumofVertices())
	d.id = make([]int, g.NumofVertices())
	d.count = 0

	for v := 0; v < g.NumofVertices(); v++ {
		if !d.marked[v] {
			d.dfs(g, v)
			d.count++
		}
	}

	return &d
}

//recursive dfs
func (d *DFS) dfs(g *libsample.Digraph, vertices int) {

	d.marked[vertices] = true
	d.id[vertices] = d.count
	arrVertices := g.Adjacent(vertices)
	for adjVertices := range arrVertices {
		w := adjVertices.(int) //cast it
		if !d.marked[w] {
			d.dfs(g, w)
		}
	}
	d.reversePost = append(d.reversePost, vertices)
}

func (d *DFS) isStronglyConnected(s, v int) bool {

	if d.id[s] == d.id[v] {
		return true
	}
	return false
}

func Print(g *libsample.Digraph) {

	for v := 0; v < g.NumofVertices(); v++ {
		w_adj := g.Adjacent(v)
		for w := range w_adj {
			fmt.Printf("%v->%v\n", v, w.(int))
		}
	}
}
