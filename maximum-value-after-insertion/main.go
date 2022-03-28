package main

import (
	"fmt"
	"strconv"
)

func maxValue(n string, x int) string {

	if n[0] == '-' {

		for i, letter := range n {
			fmt.Println(string(letter))
			if letter-'0' > int32(x) {
				return n[:i] + strconv.Itoa(x) + n[i:]
			}
		}
		return n + strconv.Itoa(x)

	}

	for i, letter := range n {
		if letter-'0' < int32(x) {
			return n[:i] + strconv.Itoa(x) + n[i:]
		}
	}

	return n + strconv.Itoa(x)

}
