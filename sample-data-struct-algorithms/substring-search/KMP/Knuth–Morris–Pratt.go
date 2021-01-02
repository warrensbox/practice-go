package main

import "fmt"

const (
	R = 256
)

func main() {

	var p Pattern
	p.pat = "rabbit"
	p.KMP()

	testWord := "aaaaaarabhbitbraabit"
	i := 0
	j := 0

	for i < len(testWord) && j < len(p.pat) {
		fmt.Println(string(testWord[i]))
		j = p.dfa[testWord[i]][j]
		i++
	}
	if j == len(p.pat) {
		fmt.Println("Found pattern")
	} else {
		fmt.Println("NO pattern found")
	}

}

type Pattern struct {
	pat string
	dfa [R][300]int
}

func (this *Pattern) KMP() {

	pattern := this.pat
	fmt.Println("pattern", pattern[0])
	M := len(pattern)
	this.dfa[pattern[0]][0] = 1
	X := 0
	for j := 1; j < M; j++ {
		for c := 0; c < R; c++ {
			this.dfa[c][j] = this.dfa[c][X] //copy mismatch cases
		}
		this.dfa[pattern[j]][j] = j + 1 //set match cases
		X = this.dfa[pattern[j]][X]     //update restart state
	}
}
