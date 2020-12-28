package main

import (
	"fmt"
)

func main() {

	words := []string{
		"she",
		"sells",
		"seashells",
		"by",
		"the",
		"sea",
		"shore",
		"the",
		"shells",
		"she",
		"sells",
		"are",
		"surely",
		"seashells",
	}

	Sort(words)

	fmt.Println(words)
}

func Sort(words []string) {

	last := len(words) - 1
	sort(words, 0, last, 1)
}

func sort(words []string, lo, hi, d int) {

	if hi <= lo {
		return
	}

	lt := lo
	gt := hi
	v := charAt(words[lo], d)
	i := lo + 1
	for i <= gt {
		t := charAt(words[i], d)
		if t < v {

			exch(words, lt, i)
			lt = lt + 1
			i = i + 1
		} else if t > v {

			exch(words, i, gt)
			gt = gt - 1
		} else {
			i++
		}
	}

	sort(words, lo, lt-1, d)
	if v >= 0 {
		sort(words, lt, gt, d+1)
	}
	sort(words, gt+1, hi, d)
}

func exch(word []string, lt, i int) {
	word[lt], word[i] = word[i], word[lt]
}

func charAt(word string, d int) int {
	if d > len(word) {
		return -1
	}
	char := []rune(word[d-1 : d])
	inter := char[0] - 'a'
	return int(inter)
}
