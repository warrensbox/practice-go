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

	aux := make([]string, len(words))
	last := len(words) - 1
	sort(words, aux, 0, last, 1)
}

func sort(words, aux []string, lo, hi, d int) {

	if hi <= lo {
		return
	}
	count := make([]int, 26+1)
	for i := lo; i <= hi; i++ {
		char := charAt(words[i], d)
		if char != -1 {
			count[char+1]++
		}

	}

	for r := 0; r < 26; r++ {
		count[r+1] += count[r] //compute cummulative
	}

	for i := lo; i <= hi; i++ {
		char := charAt(words[i], d)
		if char != -1 {
			index := count[char]
			aux[index] = words[i]
			count[char]++
		}
	}

	for i := lo; i <= hi; i++ {
		words[i] = aux[i-lo]
	}

	for r := 0; r < 26; r++ {
		sort(words, aux, lo+count[r], lo+count[r+1]-1, d+1)
	}
}

func charAt(word string, d int) int {
	if d > len(word) {
		return -1
	}
	char := []rune(word[d-1 : d])
	inter := char[0] - 'a'
	return int(inter)
}
