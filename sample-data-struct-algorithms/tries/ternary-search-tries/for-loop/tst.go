package main

import "fmt"

func main() {

	tst := New()
	tst.Put("hello", "world")
	fmt.Println(tst.Get("hello"))
}

type Node struct {
	val              string
	char             byte
	left, mid, right *Node
}

type TST struct {
	length int
	root   *Node
}

func New() *TST {
	tree := &TST{}
	tree.root = nil
	tree.length = 0
	return tree
}

func (this *TST) Put(key, val string) {

	if this.length == 0 { //if there are no tries yet
		this.root = &Node{}
	}

	t := this.root
	KEY := []byte(key)
	for i := 0; i < len(KEY); {
		c := key[i]
		if c > t.char {
			if t.right == nil {
				var node Node
				node.char = c
				t.right = &node
			}
			t = t.right
		} else if c < t.char {
			if t.left == nil {
				var node Node
				node.char = c
				t.left = &node
			}
			t = t.left
		} else {
			i++
			if i < len(KEY) {
				if t.mid == nil {
					c := key[i]
					var node Node
					node.char = c
					t.mid = &node
				}
				t = t.mid
			}
		}
	}
	if t.val == "" {
		this.length++
	}
	t.val = val
}

func (this *TST) Get(key string) string {

	if this.length == 0 { //if there are no tries yet
		return ""
	}

	t := this.root
	KEY := []byte(key)
	for i := 0; i < len(KEY); {
		c := key[i]
		if c > t.char {
			if t.right == nil {
				return ""
			}
			t = t.right
		} else if c < t.char {
			if t.left == nil {
				return ""
			}
			t = t.left
		} else {
			i++
			if i < len(KEY) {
				if t.mid == nil {
					return ""
				}
				t = t.mid
			} else {
				break
			}
		}
	}
	return t.val
}
