package main

import (
	"fmt"
)

/* HELPER FUNCTION */
type Iterator interface {
	// Returns the next item or nil if there are no more items
	Next() interface{}
}

type IntIterator struct {
	data []int
}

func NewIntIter(data []int) Iterator {
	return &IntIterator{data: data}
}

func (n *IntIterator) Next() interface{} {
	if len(n.data) > 0 {
		front := n.data[0] // peek queue, get front
		n.data = n.data[1:]
		return front
	}
	return nil
}

/* END HELPER FUNCTION */

func main() {

	n := NewIntIter([]int{1, 2, 3})

	p := New(n)

	fmt.Println(p.Next()) //1
	fmt.Println(p.Peek()) //2
	fmt.Println(p.Peek()) //2
	fmt.Println(p.Next()) //2
	fmt.Println(p.Peek()) //3
	fmt.Println(p.Next()) //3
	fmt.Println(p.Next()) //nil
}

type PeekIterator interface {
	// Peeks at the next value in the iterator without consuming it.
	Peek() interface{}
	Iterator
}

func New(iter Iterator) PeekIterator {
	return &MyPeekIterator{
		iter:   iter,
		cached: false,
		curr:   0,
	}
}

type MyPeekIterator struct {
	iter   Iterator
	cached bool
	curr   interface{}
}

func (m *MyPeekIterator) Next() interface{} {
	if m.cached {
		m.cached = false
		return m.curr
	}
	return m.iter.Next()
}

func (m *MyPeekIterator) Peek() interface{} {
	if !m.cached {
		m.cached = true
		m.curr = m.iter.Next()
	}
	return m.curr
}
