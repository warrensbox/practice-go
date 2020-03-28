package main

import (
	"fmt"
)

func main() {

	input := "L(e)(t(()c)d(e"

	output := minRemoveToMakeValid(input)

	fmt.Println(output)

}

func minRemoveToMakeValid(s string) string {

	var st []int //save parentheses index which should be delete
	var res string = ""
	for i, ch := range s {
		if ch == '(' {
			st = append(st, i)
		} else if ch == ')' {
			if len(st) > 0 && s[st[len(st)-1]] == '(' {
				st = st[:len(st)-1] //pop
			} else { //stack empty or top is ')'
				st = append(st, i)
			}
		}
	}
	fmt.Println("stack", st)

	start := 0
	end := len(s)

	for len(st) > 0 { //this psrt rearranges the sequnces

		start = st[len(st)-1] + 1
		fmt.Println("start", start)
		st = st[:len(st)-1]
		fmt.Println("st", st)
		res = s[start:end] + res
		fmt.Println("res", res)
		end = start - 1
		fmt.Println("end", end)
		start = 0
	}

	res = s[start:end] + res
	return res
}
