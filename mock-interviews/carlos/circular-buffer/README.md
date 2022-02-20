Circular Buffer

 A ring (or circular) buffer is a data structure that uses a fixed-size
 buffer to store elements in FIFO order. The buffer is seen as if it were
 connected at both ends which explains the naming.

 When the buffer is full, instead of resizing and allocating more space,
 the buffer starts from the beginning.

 As an example, consider the following buffer where items were written from
 left to right. This means the oldest value is 0 and the newest is 9:

 ┌▶[0, 1, 2, 3, 4, 5, 6, 7, 8, 9]─┐
 └────────────────────────────────┘

 If another item were to be added to the buffer it would overwrite the oldest,
 in this case 0:

 ┌▶[10, 1, 2, 3, 4, 5, 6, 7, 8, 9]─┐
 └─────────────────────────────────
 If another item were to be added to the buffer it would overwrite the oldest,
 in this case 1:

 ┌▶[10, 11, 2, 3, 4, 5, 6, 7, 8, 9]─┐
 └──────────────────────────────────┘

 The task is to implement a ring buffer conforming to the interface below.
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
