package main

import (
	"fmt"

	"github.com/warrensbox/practice-go/sample-data-struct-algorithms/libsample"
)

func main() {

}

func recognize(txt string) {
	//Does NFA recognize txt?
	g := NFA_RE("((A*B|AC)D)")
	pc := libsample.NewBag() //create new bag
	dfs := DepthFirstSearch(g.G, 0)

	for v := 0; v < g.G.NumofVertices(); v++ {
		if dfs.marked[v] {
			pc.Insert(v)
		}
	}
	for i := 0; i < len(txt); i++ {
		match := libsample.NewBag()
		for vertices := range pc.Vertices() {
			v := vertices.(int)
			if v == g.m {
				continue
			}
			if g.re[v] == txt[i] || g.re[v] == '.' {
				match.Insert(v + 1)
			}
		}
		pc = libsample.NewBag() //create new bag
		dfs = DepthFirstSearch(g.G, sta)
	}

}

type DFS struct {
	marked []bool
	edgeTo []int
}

func DepthFirstSearch(g *libsample.Digraph, v int) *DFS {

	d := DFS{}
	d.marked = make([]bool, g.NumofVertices())
	d.edgeTo = make([]int, g.NumofVertices())

	// for v := 0; v < g.NumofVertices(); v++ {
	// 	if !d.marked[v] {
	d.dfs(g, v)
	// 	}
	// }

	return &d
}

func (d *DFS) dfs(g *libsample.Digraph, v int) {
	d.marked[v] = true

	arrV := g.Adjacent(v)
	for adjV := range arrV {
		w := adjV.(int)
		if !d.marked[w] {
			d.edgeTo[w] = v
			d.dfs(g, w)
		}
	}
}

type NFA struct {
	G  *libsample.Digraph // epsilon transitions
	m  int                // number of states
	re []byte             //match transitions
}

func NFA_RE(regex string) *NFA {

	//Create the NFA for the given regular expression
	n := NFA{}
	ops := Stack{}       //new stack
	n.re = []byte(regex) //convert regex to char array
	n.m = len(regex)
	n.G = libsample.NewDigraph(n.m + 1)

	for i := 0; i < n.m; i++ {

		lp := i
		if n.re[i] == '(' || n.re[i] == '|' {
			ops.Push(i)
		} else if n.re[i] == ')' {
			or := ops.Pop()
			if n.re[or] == '|' {
				lp = ops.Pop()
				n.G.Connect(lp, or+1)
				n.G.Connect(or, i)
			} else {
				lp = or
			}
		}
		if i < n.m && n.re[i+1] == '*' {
			n.G.Connect(lp, i+1)
			n.G.Connect(i+1, lp)
		}
		if n.re[i] == '(' || n.re[i] == '*' || n.re[i] == ')' {
			n.G.Connect(i, i+1)
		}
	}

	//n.G.Adjacent()

	return n
}

//stack implementation
type Stack struct {
	stack []int
}

func (s *Stack) Push(item int) {

	s.stack = append(s.stack, item)

}

func (s *Stack) Pop() int {

	//nothing ot return
	if len(s.stack) == 0 {
		fmt.Println("Nothing to return")
		return -1
	}

	n := len(s.stack) - 1
	top := s.stack[n]
	s.stack = s.stack[:n] //pop

	return top

}

func (s *Stack) Peek() int {

	//nothing ot return
	if len(s.stack) == 0 {
		fmt.Println("Nothing to return")
		return -1
	}

	n := len(s.stack) - 1
	top := s.stack[n]

	return top

}

func (s *Stack) Len() int {

	return len(s.stack)

}
