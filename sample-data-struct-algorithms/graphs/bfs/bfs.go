package main

import (
	"container/list"
	"fmt"

	"github.com/warrensbox/practice-go/sample-data-struct-algorithms/libsample"
)

func main() {

	g := libsample.NewGraph(7)
	g.Connect(0, 1)
	g.Connect(0, 2)
	g.Connect(0, 6)
	g.Connect(6, 4)
	g.Connect(4, 5)
	g.Connect(5, 3)
	g.Connect(3, 4)
	g.Connect(5, 0)

	breadthFirstSearch(g, 0)

}

//Data structure for bfs
type BFS struct {
	marked []bool
	edgeTo []int
	source int
}

//recursive dfs
func breadthFirstSearch(g *libsample.Graph, source int) {

	b := BFS{make([]bool, g.NumofVertices()), make([]int, g.NumofVertices()), source}

	b.marked[source] = true //mark source as true
	queue := list.New()     //create a queue
	queue.PushBack(source)  //add vertice into queue (enqueue)

	for queue.Len() > 0 {
		val := queue.Front() //dequeue
		queue.Remove(val)
		v := val.Value.(int) //cast

		fmt.Println(v)
		arrVertices := g.Adjacent(v) //find all adjacent vertices (cast them as well)
		for adjVertices := range arrVertices {
			w := adjVertices.(int) //cast
			if !b.marked[w] {
				queue.PushBack(w)
				b.marked[w] = true
				b.edgeTo[w] = v
			}
		}

	}
}
