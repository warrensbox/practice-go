package main

import (
	"fmt"
	"sort"
)

func main() {

	s := "aacaagtttacaagc"
	fmt.Println(lrs(s))

}

func lrs(s string) string {

	n := len(s)

	suffixes := make([]string, n)

	for i := 0; i < n; i++ {
		suffixes[i] = s[i:n]
	}

	sort.Strings(suffixes)

	lrstr := ""
	for i := 0; i < n-1; i++ {
		length := lcp(suffixes[i], suffixes[i+1])
		if length > len(lrstr) {
			lrstr = suffixes[i][0:length]
		}
	}

	return lrstr
}

func lcp(w1, w2 string) int {

	count := 0
	len := min(w1, w2)
	for i := 0; i < len; i++ {
		if w1[:i+1] == w2[:i+1] {
			count++
		}
	}
	return count
}

func min(w1, w2 string) int {
	if len(w2) < len(w1) {
		return len(w2)
	}
	return len(w1)
}
