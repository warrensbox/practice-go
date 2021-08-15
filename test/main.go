// To execute Go code, please declare a func main() in a package "main"

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	str := "3[a2[cd]ef]]"
	//str := "2[abc]3[cd]ef"
	//str := "3[a2[c]]"
	//str := "3[a]2[bc]"
	fmt.Println(decodeString(str))
}

/*
goal: decode string according to num[char]


*/

func decodeString(s string) string {

	stack := []string{}
	curr := ""
	for _, val := range s {
		if string(val) != "]" {
			stack = append(stack, string(val)) //push
		} else {

			for len(stack) > 0 {
				//top
				top := stack[len(stack)-1]
				//pop
				stack = stack[:len(stack)-1]
				if top != "[" {
					curr = top + curr
				} else if top == "[" {
					kStr := ""
					for len(stack) > 0 {
						// peek
						n := len(stack) - 1
						top := stack[n]
						if _, err := strconv.Atoi(top); err == nil {
							kStr = top + kStr
							// pop
							stack = stack[:n]
						} else {
							break
						}
					}

					i, _ := strconv.Atoi(kStr)
					curr = strings.Repeat(curr, i)
					stack = append(stack, curr)
					curr = ""
					break
				}
			}
			stack = append(stack, curr)
		}
	}

	var res strings.Builder
	for _, s := range stack {
		res.WriteString(s)
	}
	return res.String()

}
