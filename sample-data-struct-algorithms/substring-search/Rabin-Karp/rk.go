package main

import "fmt"

func main() {

	txt := "intotherabbtithole"
	search(txt)

}

func search(txt string) int {
	var r RK
	r.pattern = "rabbit"
	r.RabinKarp(r.pattern)

	txtHash := hash(txt, len(r.pattern)) //used to test basic case
	if r.patHash == txtHash {
		fmt.Println("Match found")
		return 0
	}

	for i := r.M; i < len(txt); i++ {
		txtHash = (txtHash + Q - r.RM*int(txt[i-r.M])%Q)
		txtHash = (txtHash*R + int(txt[i])) % Q
		if r.patHash == txtHash {
			fmt.Println("Match found")
			return i - r.M - 1
		}
	}

	fmt.Println("Match NOT found")
	return len(r.pattern)
}

type RK struct {
	pattern string
	patHash int
	M       int
	RM      int
}

const (
	R = 256
	Q = 997
)

func (t *RK) RabinKarp(pat string) {

	t.M = len(pat)
	t.RM = 1

	for i := 1; i <= t.M-1; i++ {
		t.RM = (R * t.RM) % Q
	}
	t.patHash = hash(pat, t.M)
}

func hash(key string, M int) int {
	h := 0
	for j := 0; j < M; j++ {
		h = (R*h + int(key[j])) % Q
	}
	return h
}
