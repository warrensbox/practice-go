package main

import (
	"testing"
)

func TestMyBuffer_Push(t *testing.T) {
	buf := NewBuffer(10)
	if _, err := buf.pop(); err == nil {
		t.Error("should error popping empty buffer")
	}

	for i := 0; i < 10; i++ {
		buf.push(i)
	}
	if !buf.isFull() {
		t.Error("buffer should be full")
	}
	items := []int{}
	for !buf.isEmpty() {
		x, _ := buf.pop()
		items = append(items, x.(int))
	}
	assertEqual(t, items, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func TestMyBuffer_PushOverwriting(t *testing.T) {
	buf := NewBuffer(10)

	for i := 0; i < 13; i++ {
		buf.push(i)
	}
	if !buf.isFull() {
		t.Error("buffer should be full")
	}
	items := []int{}
	for !buf.isEmpty() {
		x, _ := buf.pop()
		items = append(items, x.(int))
	}
	assertEqual(t, items, []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
}

func TestMyBuffer_PushPop(t *testing.T) {
	buf := NewBuffer(10)

	for i := 0; i < 10; i++ {
		buf.push(i)
	}
	if !buf.isFull() {
		t.Error("buffer should be full")
	}

	items := []int{}
	for i := 10; i < 20; i++ {
		if i%2 == 0 {
			buf.push(i)
		} else {
			x, _ := buf.pop()
			items = append(items, x.(int))
		}
	}
	assertEqual(t, items, []int{1, 2, 3, 4, 5})

	items = []int{}
	for !buf.isEmpty() {
		x, _ := buf.pop()
		items = append(items, x.(int))
	}
	assertEqual(t, items, []int{6, 7, 8, 9, 10, 12, 14, 16, 18})
}

func assertEqual(t *testing.T, a, b []int) {
	if len(a) != len(b) {
		t.Error("different lengths", len(a), len(b))
	}

	for i, v := range a {
		if b[i] != v {
			t.Error("mismatch at index", i, v, b[i])
		}
	}
}
