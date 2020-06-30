package main

import (
	"fmt"
	"strings"
)

func main() {
	output := toGoatLatin("I speak Goat Latin")
	//toGoatLatin("The quick brown fox jumped over the lazy dog")

	fmt.Println(output)
}

func toGoatLatin(S string) string {
	var result strings.Builder
	words := strings.Fields(S)
	add := "a"
	for _, word := range words {

		//fmt.Println(word[:1])
		switch word[:1] {
		case "a", "e", "i", "o", "u", "A", "E", "I", "O", "U":
			word = word + "ma"
		default:
			word = word[1:] + word[:1] + "ma"
		}

		result.WriteString(word + add + " ")
		add += "a"

	}
	temp := result.String()
	return temp[:len(temp)-1]
}
