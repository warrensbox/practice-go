package main

import (
	"fmt"

	"github.com/warrensbox/practice-go/sample-data-struct-algorithms/libsample"
)

func main() {

	g := libsample.NewDigraph(8)
	g.Connect(0, 5)
	g.Connect(0, 1)
	g.Connect(2, 0)
	g.Connect(2, 3)
	g.Connect(3, 5)
	g.Connect(3, 2)
	g.Connect(4, 3)
	g.Connect(4, 2)
	g.Connect(5, 4)
	g.Connect(7, 6)

	bag := libsample.NewBag()
	for i := 0; i < 6; i++ {
		bag.Insert(i)
	}

	reachable := libsample.DirectedDFS_Iterable(g, bag)

	for v := 0; v < g.NumofVertices(); v++ {
		if reachable.Marked(v) {
			fmt.Println("v", v)
		}
	}

	reachable = libsample.DirectedDFS(g, 7)
	if reachable.Marked(6) {
		fmt.Println("v", 6)
	}
}
