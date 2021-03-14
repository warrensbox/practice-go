package main

import (
	"fmt"
	"sort"
)

func main() {

	words := []string{"charley", "love", "god", "coconut", "dog", "bentley", "bentleylove", "charleylovedog", "coconutlovepupies"}
	printLongestWord(words)
}

func printLongestWord(arr []string) {

	hash := make(map[string]bool)

	for _, val := range arr {
		hash[val] = true
	}

	sort.Strings(arr)
	fmt.Println(arr)
	for _, val := range arr {
		if canBuildWord(val, true, hash) {

		}
	}
}

func canBuildWord(word string, isOriginalWord bool, hash map[string]bool) bool {

	if hash[word] && !isOriginalWord {
		val, _ := hash[word]
		return val
	}

	for i := 0; i < len(word); i++ {
		// left := word[:i]
		// right := word[i:]
		// fmt.Println("left", left)
		// fmt.Println("right", right)
	}
	return false
}
