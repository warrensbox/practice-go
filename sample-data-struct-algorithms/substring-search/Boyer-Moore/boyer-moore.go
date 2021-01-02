package main

import "fmt"

func main() {

	var p Pattern
	p.pat = "rabbit"
	p.BM()

	txtWord := "aaaaaarabkbitbrabbit"
	M := len(p.pat)
	N := len(txtWord)
	skip := 0

	for i := 0; i <= N-M; i += skip {
		skip = 0
		for j := M - 1; j >= 0; j-- {
			if p.pat[j] != txtWord[i+j] {
				skip = Max(1, j-p.right[txtWord[i+j]])
				break
			}
		}

		if skip == 0 {
			fmt.Println("Match found")
			return
		}
	}

}

const (
	R = 256
)

type Pattern struct {
	pat   string
	right [R]int
}

func (this *Pattern) BM() {
	M := len(this.pat)
	for c := 0; c < R; c++ { //initialize R (all of it)
		this.right[c] = -1
	}
	for j := 0; j < M; j++ {
		this.right[this.pat[j]] = j
	}
}

func Max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
