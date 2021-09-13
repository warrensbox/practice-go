package main

import "fmt"

func minAddToMakeValid(s string) int {

	var st []int //save parentheses index which should be delete
	var res string = ""
	for i, ch := range s {
		if ch == '(' {
			st = append(st, i)
		} else if ch == ')' {
			if len(st) > 0 && s[st[len(st)-1]] == '(' {
				st = st[:len(st)-1] //pop
				fmt.Println(st)
			} else { //stack empty or top is ')'
				st = append(st, i)
			}
		}
	}

	fmt.Println(len(st))
	fmt.Println(res)

	return len(st)
}
