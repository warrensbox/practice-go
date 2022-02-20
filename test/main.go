package main

import (
	"fmt"
	"sort"
)

func main() {

	//words := []string{"a", "b", "ab", "ac", "abc", "abd"}
	//words := []string{"a", "b", "ac", "ap", "ba", "app", "ban", "appl", "bana", "apple", "apply", "banan"}
	//words := []string{"a", "b", "ba", "bca", "bda", "bdca"}
	words := []string{"a", "b"}
	fmt.Println(longesr(words))

}

func longesr(words []string) string {

	// sort.Slice(words, func(i, j int) bool {
	// 	return len(words[i]) < len(words[j])
	// })

	sort.Strings(words)
	fmt.Println("word", words)
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

		fmt.Println(hash)
	}

	return res
}
