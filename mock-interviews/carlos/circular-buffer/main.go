package main

import "errors"

// A ring (or circular) buffer is a data structure that uses a fixed-size
// buffer to store elements in FIFO order. The buffer is seen as if it were
// connected at both ends which explains the naming.
//
// When the buffer is full, instead of resizing and allocating more space,
// the buffer starts from the beginning.
//
// As an example, consider the following buffer where items were written from
// left to right. This means the oldest value is 0 and the newest is 9:
//
// ┌▶[0, 1, 2, 3, 4, 5, 6, 7, 8, 9]─┐
// └────────────────────────────────┘
//
// If another item were to be added to the buffer it would overwrite the oldest,
// in this case 0:
//
// ┌▶[10, 1, 2, 3, 4, 5, 6, 7, 8, 9]─┐
// └─────────────────────────────────┘

// If another item were to be added to the buffer it would overwrite the oldest,
// in this case 1:
//
// ┌▶[10, 11, 2, 3, 4, 5, 6, 7, 8, 9]─┐
// └──────────────────────────────────┘
//
// The task is to implement a ring buffer conforming to the interface below.
type RingBuffer interface {
	// Returns whether the buffer is empty or not.
	isEmpty() bool

	// Returns whether the buffer is full or not.
	isFull() bool

	// Pushes an item into the buffer, if the buffer is full this will overwrite
	// the oldest value.
	push(x interface{})

	// Removes the oldest value inserted in the buffer if any, if the buffer is
	// empty returns an error.
	pop() (interface{}, error)
}

type MyBuffer struct {
	head      int
	tail      int
	cap       int
	itemCount int
	items     []interface{}
}

func NewBuffer(cap int) RingBuffer {
	return &MyBuffer{
		head:      0,
		tail:      0,
		cap:       cap,
		itemCount: 0,
		items:     make([]interface{}, cap),
	}
}

func (b *MyBuffer) isEmpty() bool {
	return b.itemCount == 0
}

func (b *MyBuffer) isFull() bool {
	return b.itemCount == b.cap
}

func (b *MyBuffer) push(x interface{}) {
	b.items[b.tail] = x
	if b.isFull() {
		b.head = (b.head + 1) % b.cap
	} else {
		b.itemCount += 1
	}
	b.tail = (b.tail + 1) % b.cap
}

func (b *MyBuffer) pop() (interface{}, error) {
	if b.isEmpty() {
		return nil, errors.New("empty buffer")
	}

	x := b.items[b.head]
	b.head = (b.head + 1) % b.cap
	b.itemCount -= 1

	return x, nil
}
