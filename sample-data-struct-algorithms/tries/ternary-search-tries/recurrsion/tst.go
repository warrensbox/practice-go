package main

import "fmt"

func main() {

	tst := TST{}
	tst.Put("hello", "world")
	val := tst.Get("hello")
	fmt.Println("val", val)
	tst.Put("hello", "there")
}

type TSTNode struct {
	val              string
	char             byte
	left, mid, right *TSTNode
}

type TST struct {
	root *TSTNode
}

func (t *TST) Put(key, val string) {
	fmt.Println("puttin : ", key)
	t.root = put(t.root, key, val, 0)
}

func put(current *TSTNode, key, val string, d int) *TSTNode {
	c := key[d]
	if current == nil {
		var x TSTNode
		x.char = c
		current = &x
	}

	if c < current.char {
		current.left = put(current.left, key, val, d)
	} else if c > current.char {
		current.right = put(current.right, key, val, d)
	} else if d < len(key)-1 {
		current.mid = put(current.mid, key, val, d+1)
	} else {
		current.val = val
	}
	return current
}

func (t *TST) Get(key string) string {

	fmt.Println("GETTING", key)
	tst := t.root
	tst = get(tst, key, 0)
	return tst.val
}

func get(current *TSTNode, key string, d int) *TSTNode {
	c := key[d]

	if current == nil {
		return nil
	}

	if c < current.char {
		return get(current.left, key, d)
	} else if c > current.char {
		return get(current.right, key, d)
	} else if d < len(key)-1 {
		return get(current.mid, key, d+1)
	} else {
		return current
	}

}
