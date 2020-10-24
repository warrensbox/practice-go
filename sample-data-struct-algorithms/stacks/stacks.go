package main

import (
	"fmt"
)

func main() {

	stack := current{}
	fmt.Println(stack.push(3))
	fmt.Println("ori", stack.stackOri)
	fmt.Println(stack.push(2))
	fmt.Println("ori", stack.stackOri)
	fmt.Println(stack.push(4))
	fmt.Println("ori", stack.stackOri)
	fmt.Println(stack.push(1))
	fmt.Println("ori", stack.stackOri)
	fmt.Println(stack.push(5))
	fmt.Println("ori", stack.stackOri)
	fmt.Println(stack.pop()) //5
	fmt.Println("ori", stack.stackOri)
	fmt.Println(stack.pop()) //1
	fmt.Println("ori", stack.stackOri)
	fmt.Println(stack.pop()) //4
	fmt.Println("ori", stack.stackOri)
	fmt.Println(stack.pop()) //2
	fmt.Println("ori", stack.stackOri)
	fmt.Println("Last pop")
	fmt.Println(stack.pop()) //3
	fmt.Println("ori", stack.stackOri)
	fmt.Println(stack.pop()) //nothing
	fmt.Println("ori", stack.stackOri)
}

type current struct {
	min      int
	stackOri []int
	stackMin []int
}

func (c *current) push(item int) int {

	min := item

	if len(c.stackOri) == 0 {
		fmt.Println("Stack is empty")
		c.stackMin = append(c.stackMin, item)
		c.stackOri = append(c.stackOri, item)
		return min
	}
	//peek
	if len(c.stackMin) > 0 {
		n := len(c.stackMin) - 1
		min = c.stackMin[n]
	}

	if item < min {
		c.stackMin = append(c.stackMin, item)
		min = item
	}
	c.stackOri = append(c.stackOri, item)

	return min
}

func (c *current) pop() int {

	//nothing ot return
	if len(c.stackOri) == 0 {
		fmt.Println("Nothing to return")
		return -1
	}

	//peek Orignal
	m := len(c.stackOri) - 1
	ori := c.stackOri[m]

	//peek Min
	n := len(c.stackMin) - 1
	min := c.stackMin[n]

	if min == ori && len(c.stackMin) > 0 {
		c.stackOri = c.stackOri[:m]
		c.stackMin = c.stackMin[:n]
	}

	c.stackOri = c.stackOri[:m]
	return min
}
