package main

import "fmt"

func ladderLength(beginWord string, endWord string, wordList []string) int {

	//memorize the word list in a hash
	hashExist := make(map[string]bool)
	for _, val := range wordList {
		hashExist[val] = true
	}

	queue := []string{beginWord} //add first word to the queue
	delete(hashExist, beginWord) //remove the first word from hashExist
	level := 0                   //init the level
	for len(queue) > 0 {
		size := len(queue)
		level++
		for i := 0; i < size; i++ {

			front := queue[0]
			queue = queue[1:]
			if front == endWord {
				return level
			}

			//get the neighboring possible words
			possibleWords := neighbors(front)
			for _, neigh := range possibleWords {
				if hashExist[neigh] {
					delete(hashExist, neigh)
					queue = append(queue, neigh)
				}
			}

		}

	}

	return 0
}

func neighbors(front string) []string {

	words := []string{}
	for i := 0; i < len(front); i++ {
		for letter := 0; letter < 26; letter++ {
			char := string(letter + 'a')
			newword := front[:i] + char + front[i+1:]
			words = append(words, newword)
		}
	}
	fmt.Println(words)
	return words
}
