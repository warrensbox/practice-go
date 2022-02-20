package main

import "sort"

func longestWord(words []string) string {

	sort.Strings(words)
	hash := make(map[string]bool)

	res := ""
	for _, word := range words {
		prefix := word[:len(word)-1]
		if (len(word) == 1) || hash[prefix] {
			if len(word) > len(res) {
				res = word
			}
			hash[word] = true
		}
	}

	return res
}
