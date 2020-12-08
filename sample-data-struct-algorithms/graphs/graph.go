package main

import (
	"fmt"

	"github.com/warrensbox/practice-go/sample-data-struct-algorithms/libsample"
)

func main() {

	g := NewGraph(9)
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

//Data structure for dfs
type DFS struct {
	marked []bool
	edgeTo []int
	source int
}

//Initialize DS
func DepthFirstSearch(graph *Graph, source int) *DFS {

	dfs := DFS{make([]bool, graph.NumofVertices()), make([]int, graph.NumofVertices()), source}
	dfs.depthFirstSearch(graph, source)

	return &dfs
}

//recursive dfs
func (d *DFS) depthFirstSearch(g *Graph, vertices int) {

	d.marked[vertices] = true
	arr_vertices := g.adj(vertices)
	for adj_vertices := range arr_vertices {
		w := adj_vertices.(int) //cast
		if !d.marked[w] {
			d.depthFirstSearch(g, w)
			fmt.Println("adding edge", w)
			d.edgeTo[w] = vertices
		}
	}

}

// Data structure for representing a graph.
type Graph struct {
	nodes int
	edges []*libsample.Bag
}

//create new undirected graph
func NewGraph(vertices int) *Graph {
	var graph Graph
	graph.nodes = vertices
	graph.edges = make([]*libsample.Bag, vertices)
	for i := 0; i < vertices; i++ {
		graph.edges[i] = libsample.NewBag() // arr[0] = bag{1,2,3}; arr[1] = bag{0,5};
	}
	return &graph
}

//connect two vertices (may loopback)
func (g *Graph) Connect(v, w int) {
	g.edges[v].Insert(w)
	if v != w {
		g.edges[w].Insert(v)
	}
}

//iterator for vertices adjacent to v
func (g *Graph) adj(v int) map[interface{}]int {

	return g.edges[v].Vertices()
}

//return number of vertices in graph
func (g *Graph) NumofVertices() int {
	return g.nodes
}

func print(val interface{}) {
	fmt.Println(val)
}
