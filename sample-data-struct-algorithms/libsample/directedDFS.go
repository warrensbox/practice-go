package libsample

// Data structure for representing a DirectedDFS.
type DDFS struct {
	marked []bool
}

func DirectedDFS(g *Digraph, v int) *DDFS {

	d := DDFS{}
	d.marked = make([]bool, g.NumofVertices())
	d.dfs(g, v)

	return &d
}

func DirectedDFS_Iterable(g *Digraph, bag *Bag) *DDFS {

	d := DDFS{}
	d.marked = make([]bool, g.NumofVertices())
	for vertices := range bag.Vertices() {
		s := vertices.(int)
		if !d.marked[s] {
			d.dfs(g, s)
		}
	}

	return &d
}

func (d *DDFS) dfs(g *Digraph, v int) {
	d.marked[v] = true

	arrV := g.Adjacent(v)
	for adjV := range arrV {
		w := adjV.(int)
		if !d.marked[w] {
			d.dfs(g, w)
		}
	}
}

func (d *DDFS) Marked(v int) bool {
	return d.marked[v]
}
