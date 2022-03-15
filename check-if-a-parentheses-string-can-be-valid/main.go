package main

import "fmt"

func main() {

	parent := "))()))"
	locked := "010100"
	fmt.Println(canBeValid(parent, locked))
}

func canBeValid(s string, locked string) bool {

	if len(s)%2 == 1 {
		return false
	}

	unlocked, unpair, open, closed := 0, 0, 0, 0

	for i := 0; i < len(s); i++ {
		if locked[i] == '0' {
			unlocked++
		} else if s[i] == ')' {
			closed++
		} else if s[i] == '(' {
			open++
		}
		unpair = closed - open
		if unpair > unlocked {
			return false
		}
	}

	unlocked, unpair, open, closed = 0, 0, 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if locked[i] == '0' {
			unlocked++
		} else if s[i] == ')' {
			closed++
		} else if s[i] == '(' {
			open++
		}
		unpair = open - closed
		if unpair > unlocked {
			return false
		}
	}

	return true
}
