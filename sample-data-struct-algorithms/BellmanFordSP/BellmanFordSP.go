package main

import (
	"container/list"
	"fmt"
	"math"

	"github.com/warrensbox/practice-go/sample-data-struct-algorithms/libsample"
)

func main() {

	g := libsample.NewEdgeWeightedDiGraph(8)
	e := libsample.NewDiEdge(4, 5, 0.35)
	g.AddEdge(e)
	e = libsample.NewDiEdge(5, 4, 0.35)
	g.AddEdge(e)
	e = libsample.NewDiEdge(4, 7, 0.37)
	g.AddEdge(e)
	e = libsample.NewDiEdge(5, 7, 0.28)
	g.AddEdge(e)
	e = libsample.NewDiEdge(7, 5, 0.28)
	g.AddEdge(e)
	e = libsample.NewDiEdge(5, 1, 0.32)
	g.AddEdge(e)
	e = libsample.NewDiEdge(0, 4, 0.38)
	g.AddEdge(e)
	e = libsample.NewDiEdge(0, 2, 0.26)
	g.AddEdge(e)
	e = libsample.NewDiEdge(7, 3, 0.39)
	g.AddEdge(e)
	e = libsample.NewDiEdge(1, 3, 0.29)
	g.AddEdge(e)
	e = libsample.NewDiEdge(2, 7, 0.34)
	g.AddEdge(e)
	e = libsample.NewDiEdge(6, 2, 0.40)
	g.AddEdge(e)
	e = libsample.NewDiEdge(3, 6, 0.52)
	g.AddEdge(e)
	e = libsample.NewDiEdge(6, 0, 0.58)
	g.AddEdge(e)
	e = libsample.NewDiEdge(6, 4, 0.93)
	g.AddEdge(e)
	bf := BellmanForbf(g, 0)

	fmt.Print(bf.edgeTo)
}

type BF struct {
	edgeTo  []libsample.DiEdge
	distTo  []float32
	onQueue []bool
	queue   *list.List
	cost    int
}

func BellmanForbf(g *libsample.EdgeWeightedDiGraph, s int) *BF {

	var bf BF
	bf.distTo = make([]float32, g.NumofVertices())
	bf.edgeTo = make([]libsample.DiEdge, g.NumofVertices())
	bf.onQueue = make([]bool, g.NumofVertices())
	bf.queue = list.New()
	bf.cost = 0
	for v := 0; v < g.NumofVertices(); v++ {
		bf.distTo[v] = math.MaxFloat32
	}
	bf.distTo[s] = 0.0
	bf.queue.PushBack(s)
	bf.onQueue[s] = true

	for bf.queue.Len() > 0 {
		elem := bf.queue.Front()
		v := elem.Value.(int)
		bf.queue.Remove(elem)
		bf.onQueue[v] = true
		bf.relax(g, v)
	}
	return &bf
}

func (bf *BF) relax(g *libsample.EdgeWeightedDiGraph, v int) {

	for e := range g.Adjacent(v) {
		w := e.To()

		if bf.distTo[w] > bf.distTo[v]+e.Weight() {

			bf.distTo[w] = bf.distTo[v] + e.Weight()
			bf.edgeTo[w] = e

			if !bf.onQueue[w] {
				bf.queue.PushBack(w)
				bf.onQueue[w] = true
			}
		}
	}

}

//this function is not working as it should
func (bf *BF) findNegativeCycle() bool {
	NumV := len(bf.edgeTo)
	g := libsample.NewEdgeWeightedDiGraph(NumV)
	for v := 0; v < NumV; v++ {
		if &bf.edgeTo[v] != nil {
			g.AddEdge(&bf.edgeTo[v])
		}
	}

	cf := DirectedCycle(g)
	return cf.hasCycle()
}

type DFS struct {
	marked  []bool
	edgeTo  []int
	cycle   []int  //STACK - vertices on a cycle (if one exists)
	onStack []bool //vertices on function-call stack
}

func DirectedCycle(g *libsample.EdgeWeightedDiGraph) *DFS {

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

func (d *DFS) dfs(g *libsample.EdgeWeightedDiGraph, v int) {

	d.onStack[v] = true
	d.marked[v] = true

	for adjV := range g.Adjacent(v) {
		w := adjV.To()
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
