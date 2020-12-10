package main

import (
	"github.com/warrensbox/practice-go/sample-data-struct-algorithms/libsample"
)

func main() {

	g := libsample.NewGraph(7)
	g.Connect(0, 1)
	g.Connect(0, 2)
	g.Connect(0, 6)
	g.Connect(6, 4)
	g.Connect(4, 5)
	g.Connect(5, 3)
	g.Connect(3, 4)
	g.Connect(5, 0)

}

type DFS struct {
	marked []bool
	id     []int
	count  int
}

func Connected(g *libsample.Graph) {

	d := DFS{
		make([]bool, g.NumofVertices()),
		make([]int, g.NumofVertices()),
		0,
	}

	for v := 0; v < g.NumofVertices(); v++ {
		if !d.marked[v] {
			d.depthFirstSearch(g, v)
			d.count++
		}
	}
}

//recursive dfs
func (d *DFS) depthFirstSearch(g *libsample.Graph, vertices int) {

	d.marked[vertices] = true
	d.id[vertices] = d.count
	arrVertices := g.Adjacent(vertices)
	for adjVertices := range arrVertices {
		w := adjVertices.(int) //cast
		if !d.marked[w] {
			d.depthFirstSearch(g, w)
		}
	}

}
