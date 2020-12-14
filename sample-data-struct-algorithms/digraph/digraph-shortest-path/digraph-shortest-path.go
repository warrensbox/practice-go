package main

import (
	"container/list"
	"fmt"

	"github.com/warrensbox/practice-go/sample-data-struct-algorithms/libsample"
)

func main() {

	g := libsample.NewDigraph(6)
	g.Connect(0, 2)
	g.Connect(0, 1)
	g.Connect(1, 2)
	g.Connect(2, 4)
	g.Connect(4, 3)
	g.Connect(3, 2)
	g.Connect(3, 5)
	g.Connect(5, 0)

	breadthFirstSearch(g, 0)
}

type BFS struct {
	marked []bool
	edgeTo []int
	distTo []int
	source int
}

//bfs
func breadthFirstSearch(g *libsample.Digraph, source int) {

	b := BFS{
		make([]bool, g.NumofVertices()),
		make([]int, g.NumofVertices()),
		make([]int, g.NumofVertices()),
		source,
	}

	b.marked[source] = true
	b.distTo[source] = 0
	queue := list.New()
	queue.PushBack(source)
	dist := 0
	for queue.Len() > 0 {

		val := queue.Front()
		queue.Remove(val)
		v := val.Value.(int)
		fmt.Println("v :", v)
		dist = b.distTo[v] + 1

		arrV := g.Adjacent(v)
		for adjV := range arrV {
			w := adjV.(int)
			if !b.marked[w] {
				queue.PushBack(w)
				b.marked[w] = true
				b.edgeTo[w] = v
				b.distTo[w] = dist
			}
		}

	}

	fmt.Println(b.distTo)
}
