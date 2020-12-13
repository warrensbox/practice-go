package main

import (
	"fmt"

	"github.com/warrensbox/practice-go/sample-data-struct-algorithms/libsample"
)

func main() {

	g := libsample.NewDigraph(9)
	g.Connect(0, 5)
	g.Connect(0, 1)
	g.Connect(2, 0)
	g.Connect(2, 3)
	g.Connect(3, 5)
	g.Connect(3, 2)
	g.Connect(4, 3)
	g.Connect(4, 2)
	g.Connect(5, 4)
	g.Connect(6, 8)
	g.Connect(8, 6)
	g.Connect(6, 0)

	Print(g)
}

func Print(g *libsample.Digraph) {

	for v := 0; v < g.NumofVertices(); v++ {
		w_adj := g.Adjacent(v)
		for w := range w_adj {
			fmt.Printf("%v->%v\n", v, w.(int))
		}
	}
}
