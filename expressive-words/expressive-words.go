package main

import "fmt"

func main() {

	S := "heeellooo"
	words := []string{"hello"}

	output := expressiveWords(S, words)

	fmt.Println(output)
}

func expressiveWords(S string, words []string) int {

	res := 0
	for _, val := range words {
		if isSrretchy(S, val) {
			res++
		}
	}
	return res
}

func isSrretchy(source string, word string) bool {
	i := 0
	j := 0

	fmt.Println("source", source)
	fmt.Println("word", word)
	for i < len(source) && j < len(word) {
		fmt.Println("----------")
		sourceChar := source[i]
		wordChar := word[j]

		fmt.Println("characterAt Source", string(sourceChar))
		fmt.Println("characterAt Word", string(wordChar))

		if sourceChar != wordChar {
			fmt.Println("First false")
			return false
		}

		indexAtSource := i
		indexAtWord := j

		fmt.Println("indexAtSource", indexAtSource)
		fmt.Println("indexAtWord", indexAtWord)

		for indexAtSource < len(source) && source[indexAtSource] == sourceChar {
			indexAtSource++
		}
		for indexAtWord < len(word) && word[indexAtWord] == wordChar {
			indexAtWord++
		}

		fmt.Println("indexAtSource", indexAtSource)
		fmt.Println("indexAtWord", indexAtWord)

		lenToPointerSource := indexAtSource - i
		lenToPointerWord := indexAtWord - j

		fmt.Println("lenToPointerWord", lenToPointerWord)
		fmt.Println("lenToPointerSource", lenToPointerSource)

		fmt.Println("Check if lenToPointerWord > lenToPointerSource")
		if lenToPointerWord > lenToPointerSource || (lenToPointerSource < 3 && lenToPointerSource != lenToPointerWord) {
			fmt.Println("Second false")
			return false
		}

		i = indexAtSource
		j = indexAtWord
		fmt.Println("Check i == len source j == len word")
		if i == len(source) && j == len(word) {
			fmt.Println("Third True")
			return true
		}

	}
	return false
}
