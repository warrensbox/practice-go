package main

import (
	"fmt"

	"github.com/warrensbox/practice-go/sample-data-struct-algorithms/common"
	"github.com/warrensbox/practice-go/sample-data-struct-algorithms/libsample"
)

func main() {
	re := recognize("ACD")
	fmt.Println(re)
}

func recognize(txt string) bool {
	//Does NFA recognize txt?
	g := NFA_RE("((A*B|AC)D)")
	pc := libsample.NewBag() //create new bag
	dfs := libsample.DirectedDFS(g.G, 0)

	for v := 0; v < g.G.NumofVertices(); v++ {
		if dfs.Marked(v) {
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
		dfs = libsample.DirectedDFS_Iterable(g.G, match)
		for v := 0; v < g.G.NumofVertices(); v++ {
			if dfs.Marked(v) {
				pc.Insert(v)
			}
		}
	}
	for vertices := range pc.Vertices() {
		v := vertices.(int)
		if v == g.m {
			return true
		}
	}
	return false
}

type NFA struct {
	G  *libsample.Digraph // epsilon transitions
	m  int                // number of states
	re []byte             //match transitions
}

func NFA_RE(regex string) *NFA {

	//Create the NFA for the given regular expression
	n := NFA{}
	ops := common.IntStack{} //new stack
	n.re = []byte(regex)     //convert regex to char array
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

		if i < n.m-1 {
			if n.re[i+1] == '*' {
				n.G.Connect(lp, i+1)
				n.G.Connect(i+1, lp)
			}
		}
		if n.re[i] == '(' || n.re[i] == '*' || n.re[i] == ')' {
			n.G.Connect(i, i+1)
		}
	}

	return &n
}
