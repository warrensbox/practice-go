package main

import "fmt"

func main() {
	source := "abc"
	target := "abcbc"

	output := shortestWay(source, target)

	fmt.Println(output)
}

func shortestWay(source string, target string) int {

	count := 0

	for t := 0; t < len(target); {
		found := false

		for s := 0; s < len(source) && t < len(target); s++ {
			//fmt.Println("source[s]", string(source[s]))
			//fmt.Println("target[t]", string(target[t]))
			if source[s] == target[t] {
				found = true
				t++
			}
		}
		if !found {
			return -1
		}
		count++
	}

	return count
}
