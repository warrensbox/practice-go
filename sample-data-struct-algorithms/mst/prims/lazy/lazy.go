package main

import (
	"container/list"
	"fmt"
	"reflect"
	"sort"

	"github.com/warrensbox/practice-go/sample-data-struct-algorithms/libsample"
)

func main() {

	g := libsample.NewEdgeWeightedGraph(8)

	e := libsample.NewEdge(4, 5, 0.35)
	g.AddEdge(e)
	e = libsample.NewEdge(4, 7, 0.37)
	g.AddEdge(e)
	e = libsample.NewEdge(5, 7, 0.28)
	g.AddEdge(e)
	e = libsample.NewEdge(0, 7, 0.16)
	g.AddEdge(e)
	e = libsample.NewEdge(1, 5, 0.32)
	g.AddEdge(e)
	e = libsample.NewEdge(0, 4, 0.38)
	g.AddEdge(e)

	e = libsample.NewEdge(2, 3, 0.17)
	g.AddEdge(e)
	e = libsample.NewEdge(1, 7, 0.19)
	g.AddEdge(e)
	e = libsample.NewEdge(0, 2, 0.26)
	g.AddEdge(e)
	e = libsample.NewEdge(1, 2, 0.36)
	g.AddEdge(e)
	e = libsample.NewEdge(1, 3, 0.29)
	g.AddEdge(e)
	e = libsample.NewEdge(2, 7, 0.34)
	g.AddEdge(e)

	e = libsample.NewEdge(6, 2, 0.40)
	g.AddEdge(e)
	e = libsample.NewEdge(3, 6, 0.52)
	g.AddEdge(e)
	e = libsample.NewEdge(6, 0, 0.58)
	g.AddEdge(e)
	e = libsample.NewEdge(6, 4, 0.93)
	g.AddEdge(e)

	LazyPrimMST(g)
}

type ByWeight []libsample.Edge

type MST struct {
	marked []bool
	queue  *list.List
	minPQ  ByWeight
}

func (a ByWeight) Len() int           { return len(a) }
func (a ByWeight) Less(i, j int) bool { return a[i].Weight() < a[j].Weight() }
func (a ByWeight) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func LazyPrimMST(ewg *libsample.EdgeWeightedGraph) {

	var mst MST
	mst.marked = make([]bool, ewg.NumofVertices())
	mst.queue = list.New()

	mst.visit(ewg, 0)
	for len(mst.minPQ) > 0 {
		e := mst.PQPop()
		v := e.Either()
		w := e.Other(v)
		if mst.marked[v] && mst.marked[w] {
			continue
		}

		mst.queue.PushBack(e)
		if !mst.marked[v] {
			mst.visit(ewg, v)
		}
		if !mst.marked[w] {
			mst.visit(ewg, w)
		}
	}

	mst.Iterable()
}

func (mst *MST) visit(ewg *libsample.EdgeWeightedGraph, v int) {

	mst.marked[v] = true
	for e := range ewg.Adjacent(v) {
		if !mst.marked[e.Other(v)] {
			mst.PQInsert(e)
		}
	}
}

func (mst *MST) PQInsert(e libsample.Edge) {
	mst.minPQ = append(mst.minPQ, e)
	sort.Sort(ByWeight(mst.minPQ))
	fmt.Println(mst.minPQ)
}

func (mst *MST) PQPop() libsample.Edge {

	var e libsample.Edge

	if len(mst.minPQ) > 1 {
		top := mst.minPQ[0]
		mst.minPQ = mst.minPQ[1:]
		fmt.Println("top", top)
		return top
	} else if len(mst.minPQ) == 1 {

		last := mst.minPQ[0]
		mst.minPQ = nil
		return last
	}

	return e
}

func (mst *MST) Iterable() {

	for list := mst.queue.Front(); list != nil; list = list.Next() {

		vD := reflect.ValueOf(list.Value)
		weight := vD.FieldByName("weight")

		pV := reflect.ValueOf(list.Value)
		v := pV.FieldByName("v")

		pW := reflect.ValueOf(list.Value)
		w := pW.FieldByName("w")
		fmt.Printf("%v - %v  weight - %v\n", v, w, weight)
	}
}
