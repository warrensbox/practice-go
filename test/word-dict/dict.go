package main

import (
	"fmt"
	"strings"
)

/*
goal: find the number of times a word appears
sol:
- split word into array
- make all word to lowercase
- iterate and add to hash map
*/

func main() {

	sentence := "After beating the eggs, Dana read the next step:"
	counter(sentence)
}

func counter(sentence string) {

	hash := make(map[string]int)

	arr := strings.Split(sentence, " ")
	for _, val := range arr {
		val = strings.ToLower(val)
		hash[val]++
	}
	fmt.Println(hash)
}
