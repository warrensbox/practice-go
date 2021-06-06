package main

import (
	"fmt"
	"strings"
)

func main() {
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	str := "catsanddog"
	fmt.Println(word(str, wordDict))

}

func word(str string, wordDict []string) []string {
	return wordBreaker(str, wordDict)
}

func wordBreaker(str string, wordDict []string) []string {
	results := []string{}

	if len(str) == 0 {
		results = append(results, ", ")
		return results
	}

	for _, word := range wordDict {
		if strings.HasPrefix(str, word) {
			fmt.Println("WORD", word)

			sub := str[len(word):]
			fmt.Println("sub", sub)
			subStrings := wordBreaker(sub, wordDict)
			fmt.Println("subString", subStrings)
			for _, subString := range subStrings {
				fmt.Println("**word", word)
				fmt.Println("**subString", subString)
				results = append(results, word+" "+subString)
			}
		}
		fmt.Println("----")
	}

	return results
}
