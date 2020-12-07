package main

import (
	"fmt"

	"github.com/warrensbox/practice-go/sample-data-struct-algorithms/libsample"
)

func main() {

	bag := libsample.NewBag()
	bag.Insert(1)
	bag.Insert(2)
	bag.Do(print)

	bag2 := libsample.NewBag()
	bag2.Insert(1)
	bag2.Insert(2)
	bag2.Do(print)
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
func (g *Graph) adj(v int) *libsample.Bag {

	return g.edges[v]
}

//return number of vertices in graph
func (g *Graph) Vertices() int {
	return g.nodes
}

func print(val interface{}) {
	fmt.Println(val)
}
