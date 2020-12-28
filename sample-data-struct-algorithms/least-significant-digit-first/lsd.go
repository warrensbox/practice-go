package main

import (
	"fmt"
	"os"
)

func main() {

	words := []string{
		"dab",
		"add",
		"cab",
		"fad",
		"fee",
		"bad",
		"dad",
		"bee",
		"fed",
		"bed",
		"ebb",
		"ace",
	}

	W := 3 //length of word
	aux := make([]string, len(words))

	for d := W; d > 0; d-- {
		count := make([]int, 7)
		for i := 0; i < len(words); i++ {
			char := []rune(words[i][d-1 : d])
			count[(char[0]-'a')+1]++
		}
		fmt.Println(count)
		os.Exit(0)
		for i := 0; i < 6; i++ {
			count[i+1] += count[i] //compute cummulative
		}

		for i := 0; i < len(words); i++ {
			char := []rune(words[i][d-1 : d])
			index := count[char[0]-'a']
			aux[index] = words[i]
			count[char[0]-'a']++
		}

		for i := 0; i < len(words); i++ {
			words[i] = aux[i]
		}
	}

	fmt.Println(words)

}
