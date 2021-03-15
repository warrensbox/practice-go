package main

import (
	"fmt"
	"sort"
)

func main() {

	words := []string{"charley", "love", "god", "coconut", "dog", "bentley", "bentleylove", "charleylovedog", "coconutlovepupies"}

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	printLongestWord(words)
}

func printLongestWord(arr []string) {
	fmt.Println(arr)
	hash := make(map[string]bool)

	for _, val := range arr {
		hash[val] = true
	}

	for _, val := range arr {
		if canBuildWord(val, true, hash) {
			fmt.Println(val)
		}
	}
}

func canBuildWord(word string, isOriginalWord bool, hash map[string]bool) bool {

	if hash[word] && !isOriginalWord {
		val, _ := hash[word]
		return val
	}

	for i := 0; i < len(word); i++ {
		left := word[:i]
		right := word[i:]
		wordLeft, _ := hash[left]
		//wordRight, _ := hash[right]
		if hash[left] && wordLeft && canBuildWord(right, false, hash) {
			return true
		}
	}
	//hash[word] = false
	return false
}
