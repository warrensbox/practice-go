package main

import (
	"fmt"

	"github.com/warrensbox/practice-go/sample-data-struct-algorithms/libsample"
)

func main() {

	g := libsample.NewGraph(9)
	g.Connect(0, 1)
	g.Connect(0, 2)
	g.Connect(0, 6)
	g.Connect(6, 4)
	g.Connect(4, 5)
	g.Connect(5, 3)
	g.Connect(3, 4)
	g.Connect(5, 0)
	g.Connect(8, 7)

	d := Connected(g)
	connected := d.isConnected(0, 6)
	fmt.Println(connected)
	connected = d.isConnected(0, 7)
	fmt.Println(connected)
}

type DFS struct {
	marked []bool
	id     []int
	count  int
}

func Connected(g *libsample.Graph) *DFS {

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

	return &d
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

func (d *DFS) isConnected(s, v int) bool {

	if d.id[s] == d.id[v] {
		return true
	}
	return false
}
