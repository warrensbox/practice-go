package main

/*   Below is the interface for Iterator, which is already defined for you.
 */
type Iterator struct {
}

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
	return false
}
func (this *Iterator) next() int {
	// Returns the next element in the iteration.
	return 1
}

type PeekingIterator struct {
	iter  *Iterator
	moved bool
	curr  int
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{
		iter:  iter,
		moved: false,
		curr:  0,
	}
}

func (this *PeekingIterator) hasNext() bool {
	return this.moved || this.iter.hasNext()
}

func (this *PeekingIterator) next() int {
	if this.moved { //check if we have used peek
		this.moved = false //if we have, set it as false
		return this.curr   //return what we see in peek()
	} else {
		return this.iter.next() //else just return next
	}
}

func (this *PeekingIterator) peek() int {
	if !this.moved { //if peek is never used before
		this.moved = true            //set it to true
		this.curr = this.iter.next() //set the tmp variable (curr) as has next
	}

	return this.curr //return tmp variable
}
