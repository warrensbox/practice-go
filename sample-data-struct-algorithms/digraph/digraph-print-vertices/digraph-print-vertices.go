package main

import (
	"fmt"

	"github.com/warrensbox/practice-go/sample-data-struct-algorithms/libsample"
)

func main() {

}

func Print(g *libsample.Graph) {

	for v := 0; v < g.NumofVertices(); v++ {
		w_adj := g.Adjacent(v)
		for w := range w_adj {
			fmt.Printf(%v "->" %v,v,w)
		}
	}
}
