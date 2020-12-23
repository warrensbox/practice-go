package main

import "github.com/warrensbox/practice-go/sample-data-struct-algorithms/libsample"

func main() {

	g := libsample.NewEdgeWeightedGraph(8)

	e := libsample.NewEdge(4, 5, 0.35)
	g.AddEdge(e)
	e = libsample.NewEdge(4, 7, 0.37)
	g.AddEdge(e)
	e = libsample.NewEdge(5, 7, 0.28)
	g.AddEdge(e)

	g.Edges()

}
