package main

import (
	"fmt"

	"github.com/warrensbox/practice-go/sample-data-struct-algorithms/libsample"
)

func main() {

	g := libsample.NewGraph(6)
	g.Connect(0, 1)
	g.Connect(0, 2)
	g.Connect(0, 6)
	g.Connect(6, 4)
	g.Connect(4, 5)
	g.Connect(5, 3)
	g.Connect(3, 4)
	g.Connect(5, 0)

	dfs := DepthFirstSearch(g, 0)
	fmt.Println(dfs.pathTo(3))
}

//Data structure for dfs
type DFS struct {
	marked []bool
	edgeTo []int
	source int
}

//Initialize DS
func DepthFirstSearch(graph *libsample.Graph, source int) *DFS {

	dfs := DFS{make([]bool, graph.NumofVertices()), make([]int, graph.NumofVertices()), source}
	dfs.depthFirstSearch(graph, source)

	return &dfs
}

//recursive dfs
func (d *DFS) depthFirstSearch(g *libsample.Graph, vertices int) {

	d.marked[vertices] = true
	arrVertices := g.Adjacent(vertices)
	for adjVertices := range arrVertices {
		w := adjVertices.(int) //cast
		if !d.marked[w] {
			d.depthFirstSearch(g, w)
			fmt.Println("adding edge", w)
			d.edgeTo[w] = vertices
		}
	}

}

//is there a path from s to v
func (d *DFS) hasPathTo(v int) bool {
	return d.marked[v]
}

//path from s to v
func (d *DFS) pathTo(v int) []int {

	path := []int{}
	if !d.hasPathTo(v) {
		return path
	}
	x := v //start at end of path
	for x != d.source {
		path = append(path, x)
		x = d.edgeTo[x]
	}

	return path
}
