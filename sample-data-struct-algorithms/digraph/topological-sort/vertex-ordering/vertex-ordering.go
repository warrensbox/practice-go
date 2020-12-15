package main

import (
	"container/list"
	"fmt"

	"github.com/warrensbox/practice-go/sample-data-struct-algorithms/libsample"
)

func main() {

	g := libsample.NewDigraph(7)
	g.Connect(0, 1)
	g.Connect(0, 2)
	g.Connect(0, 5)
	g.Connect(3, 6)
	g.Connect(3, 5)
	g.Connect(3, 4)
	g.Connect(5, 2)
	g.Connect(6, 4)
	g.Connect(6, 0)
	g.Connect(3, 2)
	g.Connect(1, 4)

	d := DepthFirstSearch(g)

	fmt.Println(d.reversePost)
}

type Post string

type DFS struct {
	marked      []bool
	reversePost []int      //reserve post order
	pre         *list.List //pre order
	post        *list.List //post order
}

func DepthFirstSearch(g *libsample.Digraph) *DFS {

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

func (d *DFS) dfs(g *libsample.Digraph, v int) {

	d.pre.PushBack(v)

	d.marked[v] = true
	arrV := g.Adjacent(v)
	for adjV := range arrV {
		w := adjV.(int)
		if !d.marked[w] {
			d.dfs(g, w)

		}
	}
	d.post.PushBack(v)
	d.reversePost = append(d.reversePost, v)
}
