package main

import (
	"fmt"
)

func main() {

	arr := []rune{'a', 'b', 'c', 'd', 'e', 'f'}
	g := NewDigraph(256, arr)
	g.Connect('d', 'a')
	g.Connect('b', 'f')
	g.Connect('d', 'b')
	g.Connect('a', 'f')
	g.Connect('c', 'd')
	//g.Connect('f', 'c')

	d := BuildOrderDFS(g)

	if d.cycle {
		fmt.Println("Imposible")
	} else {
		for _, k := range d.reversePost {
			fmt.Println(string(k))
		}
	}
}

type DFS struct {
	marked      []bool
	edgeTo      []rune
	onStack     []bool
	reversePost []rune //reserve post order
	cycle       bool
}

func BuildOrderDFS(g *Digraph) *DFS {

	d := DFS{}
	d.marked = make([]bool, g.NumofVertices())
	d.edgeTo = make([]rune, g.NumofVertices())
	d.onStack = make([]bool, g.NumofVertices())

	arr := []rune{'a', 'b', 'c', 'd', 'e', 'f'}
	for _, v := range arr {
		if !d.marked[v] {
			d.dfs(g, v)
		}
	}
	return &d
}

func (d *DFS) dfs(g *Digraph, v rune) {

	d.marked[v] = true
	d.onStack[v] = true

	for w := range g.Adjacent(v) {
		if d.cycle {
			return
		} else if !d.marked[w] {
			d.edgeTo[w] = v
			d.dfs(g, w)
		} else if d.onStack[w] {
			d.cycle = true
		}
	}
	d.onStack[v] = false
	d.reversePost = append(d.reversePost, v)
}

// Data structure for representing a graph.
type Digraph struct {
	nodes int
	edges []*Bag
}

//create new undirected Digraph
func NewDigraph(vertices int, arr []rune) *Digraph {
	var digraph Digraph
	digraph.nodes = vertices
	digraph.edges = make([]*Bag, vertices)
	for _, v := range arr {
		digraph.edges[v] = NewBag() // arr[0] = bag{1,2,3}; arr[1] = bag{0,5};
	}
	return &digraph
}

//connect two vertices (may loopback)
func (g *Digraph) Connect(v, w rune) {
	g.edges[v].Insert(w)
}

//iterator for vertices adjacent to v
func (g *Digraph) Adjacent(v rune) map[rune]int {
	return g.edges[v].Vertices()
}

//return number of vertices in Digraph
func (g *Digraph) NumofVertices() int {
	return g.nodes
}

// Bag data structure (multiset).
type Bag struct {
	data map[rune]int
}

// Creates a new empty bag.
func NewBag() *Bag {
	return &Bag{make(map[rune]int)}
}

// Inserts an element into the bag.
func (b *Bag) Insert(val rune) {
	b.data[val]++
}

func (b *Bag) Vertices() map[rune]int {
	return b.data
}
