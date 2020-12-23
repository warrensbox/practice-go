package main

import (
	"container/list"
	"fmt"
	"reflect"
	"sort"

	"github.com/warrensbox/practice-go/sample-data-struct-algorithms/libsample"
)

var mst *list.List

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

	KruskalMST(g)

	Iterable()
}

func KruskalMST(ewg *libsample.EdgeWeightedGraph) {

	edges := []libsample.Edge{}
	mst = list.New()

	edgeArr := ewg.Edges()
	for e := range edgeArr {
		edges = append(edges, e)
	}

	sort.Sort(ByWeight(edges))
	var unionFind Constructor
	unionFind.length = ewg.NumofVertices()
	uf := Initialize(&unionFind)

	for len(edges) > 0 && mst.Len() < ewg.NumofVertices()-1 {

		e := edges[0]     //get the min edge
		edges = edges[1:] //rm the smallest
		v := e.Either()
		w := e.Other(v)
		if uf.connected(v, w) {
			continue // if connected, ignore
		}
		uf.union(v, w)  //merge components
		mst.PushBack(e) //edge edge to mst
	}

}

func Iterable() {

	for list := mst.Front(); list != nil; list = list.Next() {

		vD := reflect.ValueOf(list.Value)
		weight := vD.FieldByName("weight")

		pV := reflect.ValueOf(list.Value)
		v := pV.FieldByName("v")

		pW := reflect.ValueOf(list.Value)
		w := pW.FieldByName("w")
		fmt.Printf("%v - %v  weight - %v\n", v, w, weight)
	}
}

type ByWeight []libsample.Edge

func (a ByWeight) Len() int           { return len(a) }
func (a ByWeight) Less(i, j int) bool { return a[i].Weight() < a[j].Weight() }
func (a ByWeight) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func Initialize(attr *Constructor) *Constructor {

	attr.arr = make([]int, attr.length)
	for i := 0; i < attr.length; i++ {
		attr.arr[i] = i
	}

	fmt.Println(attr)
	return attr
}

type Constructor struct {
	length int
	arr    []int
}

func (attr *Constructor) connected(p, q int) bool {

	if attr.arr[p] == attr.arr[q] {
		return true
	}
	return false
}

func (attr *Constructor) union(p, q int) {

	pid := attr.arr[p]
	qid := attr.arr[q]

	for i := 0; i < attr.length; i++ {
		if attr.arr[i] == pid {
			attr.arr[i] = qid
		}
	}

}
