package main

import (
	"fmt"
	"math"
	"sort"

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

	DijkstraSP(g, 0)

}

type DSP struct {
	edgeTo []libsample.DiEdge
	distTo []float32
	minPQ  ByWeight
}

func DijkstraSP(g *libsample.EdgeWeightedDiGraph, s int) {

	var dsp DSP
	dsp.distTo = make([]float32, g.NumofVertices())
	dsp.edgeTo = make([]libsample.DiEdge, g.NumofVertices())
	dsp.minPQ = make(ByWeight, 0)

	for v := 0; v < g.NumofVertices(); v++ {
		dsp.distTo[v] = math.MaxFloat32
	}
	dsp.distTo[s] = 0.0

	dsp.PQInsert(s, 0.0)

	for dsp.minPQ.Len() > 0 {
		vertice := dsp.PQPop()
		dsp.relax(g, vertice)
	}

	fmt.Println(dsp.edgeTo)

}

func (d *DSP) relax(g *libsample.EdgeWeightedDiGraph, v int) {

	for e := range g.Adjacent(v) {
		w := e.To()
		if d.distTo[w] > d.distTo[v]+e.Weight() {
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
