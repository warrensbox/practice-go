package main

import (
	"fmt"
)

/*
goal : find location of closing parenthesis
sol:
- use stack
- trim white space
- push index of ( into stack (i)
- when you find a ) pop stack, get the index => map(popped) = curr (i)
*/

func main() {

	str := "Sometimes (when I nest them (my parentheticals) too much (like this (and this))) they get confusing."

	fmt.Println(getClosingBracket(str, 10))

}

func getClosingBracket(str string, position int) int {

	var stack []int
	hash := make(map[int]int)
	for currIndex, val := range str {

		if val == '(' {
			stack = append(stack, currIndex)
		} else if val == ')' {

			if len(stack) > 0 {
				//top
				top := stack[len(stack)-1]
				//pop
				stack = stack[:len(stack)-1]
				hash[top] = currIndex
			}

		}
	}

	return hash[position]
}
