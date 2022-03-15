package main

import (
	"fmt"
	"sort"
)

func main() {

	input := []string{"a", "b", "ba", "bca", "bda", "bdca"}
	output := longestStrChain(input)

	fmt.Println(output)

}

func longestStrChain(words []string) int {

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	dict := make(map[string]int)
	result := 0
	for _, val := range words {
		fmt.Println("----")
		fmt.Println("words", val)

		for i := 0; i < len(val); i++ {
			prev := val[0:i] + val[i+1:]
			// fmt.Println("i", i)
			// fmt.Println("prev", prev)
			// fmt.Println("dict", dict)
			// fmt.Println("val[0:i]", val[0:i])
			// fmt.Println("val[i+1:]", val[i+1:])
			// fmt.Println("dict[val]", dict[val])

			dict[val] = max(dict[val], dict[prev]+1)
			// fmt.Println("dict[prev]", dict[prev])
		}

		result = max(result, dict[val])
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
