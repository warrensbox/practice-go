package main

import (
	"container/list"
	"fmt"
	"math"
	"sort"

	"github.com/warrensbox/practice-go/sample-data-struct-algorithms/libsample"
)

func main() {

	g := libsample.NewEdgeWeightedDiGraph(8)
	e := libsample.NewDiEdge(5, 4, 0.35)
	g.AddEdge(e)
	e = libsample.NewDiEdge(4, 7, 0.37)
	g.AddEdge(e)
	e = libsample.NewDiEdge(5, 7, 0.28)
	g.AddEdge(e)
	e = libsample.NewDiEdge(5, 1, 0.32)
	g.AddEdge(e)
	e = libsample.NewDiEdge(4, 0, 0.38)
	g.AddEdge(e)
	e = libsample.NewDiEdge(0, 2, 0.26)
	g.AddEdge(e)
	e = libsample.NewDiEdge(3, 7, 0.39)
	g.AddEdge(e)
	e = libsample.NewDiEdge(1, 3, 0.29)
	g.AddEdge(e)
	e = libsample.NewDiEdge(7, 2, 0.34)
	g.AddEdge(e)
	e = libsample.NewDiEdge(6, 2, 0.40)
	g.AddEdge(e)
	e = libsample.NewDiEdge(3, 6, 0.52)
	g.AddEdge(e)
	e = libsample.NewDiEdge(6, 0, 0.58)
	g.AddEdge(e)
	e = libsample.NewDiEdge(6, 4, 0.93)
	g.AddEdge(e)

	dsp := AcyclicSP(g, 0)

	fmt.Println(dsp.pathTo(0))
}

type DSP struct {
	edgeTo []libsample.DiEdge
	distTo []float32
	minPQ  ByWeight
}

func AcyclicSP(g *libsample.EdgeWeightedDiGraph, s int) *DSP {

	var dsp DSP
	dsp.distTo = make([]float32, g.NumofVertices())
	dsp.edgeTo = make([]libsample.DiEdge, g.NumofVertices())
	dsp.minPQ = make(ByWeight, 0)

	for v := 0; v < g.NumofVertices(); v++ {
		dsp.distTo[v] = math.MaxFloat32
	}

	d := TopologicalOrder(g)
	fmt.Println(d.reversePost)
	size := len(d.reversePost)
	dsp.distTo[d.reversePost[size-1]] = 0.0
	for size > 0 {
		v := d.reversePost[size-1]
		d.reversePost = d.reversePost[:size-1]
		dsp.relax(g, v)
		size = len(d.reversePost)
	}

	fmt.Println(dsp.edgeTo)

	return &dsp

}

func (d *DSP) relax(g *libsample.EdgeWeightedDiGraph, v int) {

	for e := range g.Adjacent(v) {
		w := e.To()
		fmt.Println("w", w)
		fmt.Println("v", v)
		fmt.Println("e.Weight()", e.Weight())
		fmt.Println("d.distTo[w]", d.distTo[w])
		fmt.Println("d.distTo[v]", d.distTo[v])
		if d.distTo[w] > d.distTo[v]+e.Weight() {
			fmt.Println("PUSH", w)
			d.distTo[w] = d.distTo[v] + e.Weight()
			d.edgeTo[w] = e
			if d.PQContains(w) {
				d.PQUpdate(w, d.distTo[w])
			} else {
				d.PQInsert(w, d.distTo[w])
			}

		}

	}

}

func (d *DSP) hasPathTo(v int) bool {
	return d.distTo[v] < math.MaxFloat32
}

func (d *DSP) pathTo(v int) []libsample.DiEdge {
	stack := []libsample.DiEdge{}
	if d.hasPathTo(v) {

		for e := d.edgeTo[v]; e.From() != v; e = d.edgeTo[e.From()] {
			stack = append(stack, e)
		}
	}

	return stack
}

type PriorityQueue struct {
	W      int
	Weight float32
}

type ByWeight []PriorityQueue

func (a ByWeight) Len() int           { return len(a) }
func (a ByWeight) Less(i, j int) bool { return a[i].Weight < a[j].Weight }
func (a ByWeight) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (mst *DSP) PQInsert(v int, weight float32) {
	var pq PriorityQueue
	pq.W = v
	pq.Weight = weight
	mst.minPQ = append(mst.minPQ, pq)
	sort.Sort(ByWeight(mst.minPQ))
	//fmt.Println(mst.minPQ)
}

func (mst *DSP) PQPop() int {

	if len(mst.minPQ) > 1 {
		top := mst.minPQ[0]
		mst.minPQ = mst.minPQ[1:]
		return top.W
	} else if len(mst.minPQ) == 1 {

		last := mst.minPQ[0]
		mst.minPQ = nil
		return last.W
	}

	return -1
}

func (mst *DSP) PQContains(v int) bool {
	for _, t := range mst.minPQ {
		if t.W == v {
			return true
		}
	}

	return false
}

func (mst *DSP) PQUpdate(v int, weight float32) {

	for _, t := range mst.minPQ {
		if t.W == v {
			t.Weight = weight
		}
	}

	sort.Sort(ByWeight(mst.minPQ))
}

type DFS struct {
	marked      []bool
	reversePost []int      //reserve post order
	pre         *list.List //pre order
	post        *list.List //post order
}

func TopologicalOrder(g *libsample.EdgeWeightedDiGraph) *DFS {

	d := DFS{}
	d.marked = make([]bool, g.NumofVertices())
	d.pre = list.New()
	d.post = list.New()

	for v := 0; v < g.NumofVertices(); v++ {
		if !d.marked[v] {
			d.dfs(g, v)
		}
	}

	return &d
}

func (d *DFS) dfs(g *libsample.EdgeWeightedDiGraph, v int) {

	d.pre.PushBack(v)

	d.marked[v] = true
	//arrV := g.Adjacent(v)
	for adjV := range g.Adjacent(v) {
		w := adjV.W
		if !d.marked[w] {
			d.dfs(g, w)

		}
	}
	d.post.PushBack(v)
	d.reversePost = append(d.reversePost, v)
}
