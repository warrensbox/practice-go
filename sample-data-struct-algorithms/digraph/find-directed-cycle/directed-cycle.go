package main

import (
	"fmt"

	"github.com/warrensbox/practice-go/sample-data-struct-algorithms/libsample"
)

func main() {

	g := libsample.NewDigraph(6)

	//g := libsample.NewDigraph(7)
	// g.Connect(0, 1)
	// g.Connect(0, 2)
	// g.Connect(0, 5)
	// g.Connect(3, 6)
	// g.Connect(3, 5)
	// g.Connect(3, 4)
	// g.Connect(5, 2)
	// g.Connect(6, 4)
	// g.Connect(6, 0)
	// g.Connect(3, 2)
	// g.Connect(1, 4)
	g.Connect(0, 5)
	g.Connect(5, 4)
	g.Connect(4, 3)
	//g.Connect(3, 6)
	g.Connect(3, 5)

	DirectedCycle(g)

}

type DFS struct {
	marked  []bool
	edgeTo  []int
	cycle   []int  //STACK - vertices on a cycle (if one exists)
	onStack []bool //vertices on function-call stack
}

func DirectedCycle(g *libsample.Digraph) *DFS {

	d := DFS{}
	d.marked = make([]bool, g.NumofVertices())
	d.edgeTo = make([]int, g.NumofVertices())
	d.onStack = make([]bool, g.NumofVertices())

	for v := 0; v < g.NumofVertices(); v++ {
		if !d.marked[v] {
			d.dfs(g, v)
		}
	}

	return &d
}

func (d *DFS) dfs(g *libsample.Digraph, v int) {

	d.onStack[v] = true
	d.marked[v] = true

	arrV := g.Adjacent(v)
	for adjV := range arrV {
		w := adjV.(int)
		if d.hasCycle() { // I do not understand the significance of this line??
			fmt.Println("Has cycle")
			return
		} else if !d.marked[w] {
			d.edgeTo[w] = v
			d.dfs(g, w)
		} else if d.onStack[w] {
			//d.cycle = nil
			for x := v; x != w; x = d.edgeTo[x] {
				d.cycle = append(d.cycle, x)
			}
			d.cycle = append(d.cycle, w)
			d.cycle = append(d.cycle, v)
			fmt.Println(d.cycle)
		}
	}
	d.onStack[v] = false
}

func (d *DFS) hasCycle() bool {
	if len(d.cycle) > 0 {
		return true
	}
	return false
}
