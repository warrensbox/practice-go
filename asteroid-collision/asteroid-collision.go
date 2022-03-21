package main

import (
	"fmt"
)

func main() {

	// S := "a##c"
	// T := "#a#c"

	//S := []int{10, 2, -5}

	S := []int{-5, 5, -3, -12}

	output := asteroidCollision(S)
	fmt.Println(output)
}

func asteroidCollision(asteroids []int) []int {

	var lstack Stack
	var rstack Stack
	var queue []int

	for _, val := range asteroids {
		lstack.push(val)
	}

	fmt.Println(lstack)
	fmt.Println(rstack)

	for !lstack.empty() {
		top := lstack.top() //also pop
		fmt.Println("top", top)
		//continue
		if top > 0 {
			if rstack.empty() || rstack.peek() > 0 {
				rstack.push(top)
			} else if rstack.peek() < 0 {
				topInner := rstack.top()
				if -topInner > top {
					lstack.push(topInner)
				} else if -topInner < top {
					lstack.push(top)
				}
			}
		} else {
			if !lstack.empty() && lstack.peek() > 0 {
				//fmt.Println(lstack.peek())
				topInner := lstack.top() //also pops
				if topInner < -top {
					lstack.push(top)
				} else if topInner > -top {
					lstack.push(topInner)
				}
				//fmt.Println("lstack2",lstack)
			} else {
				rstack.push(top)
				//fmt.Println("rstack2",rstack)
			}
		}

	}

	for !rstack.empty() {
		top := rstack.top()
		queue = append(queue, top)
	}
	return queue
}

type Stack struct {
	stack []int
}

func (this *Stack) init() {
	this.stack = []int{}
}

func (this *Stack) peek() int {
	return this.stack[len(this.stack)-1]
}

func (this *Stack) pop() {
	if len(this.stack) > 0 {
		this.stack = this.stack[:len(this.stack)-1]
	}
}

func (this *Stack) top() int {
	top := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return top
}

func (this *Stack) empty() bool {
	return len(this.stack) == 0
}

func (this *Stack) push(a int) {
	this.stack = append(this.stack, a)
}
