package main

import "fmt"

func main() {

	words := []string{"hello", "leetcode"} //true
	//words := []string{"word", "world", "row"} //false
	//words := []string{"apple", "app"} //false
	//words := []string{"kuvp", "q"} //true
	//words := []string{"l", "h"}          //false
	//words := []string{"iekm", "tpnhnbe"} //false

	order := "hlabcdefgijkmnopqrstuvwxyz" //true
	//order := "worldabcefghijkmnpqstuvxyz" //false
	//order := "abcdefghijklmnopqrstuvwxyz" //false
	//order := "ngxlkthsjuoqcpavbfdermiywz" //true
	//order := "xkbwnqozvterhpjifgualycmds" //true
	//order := "loxbzapnmstkhijfcuqdewyvrg"
	output := isAlienSorted(words, order)

	fmt.Println(output)

}

func isAlienSorted(words []string, order string) bool {

	charMap := [26]int{}

	for i := 0; i < len(order); i++ {
		charMap[order[i]-'a'] = i
	}

	for i := 1; i < len(words); i++ {
		if !compare(words[i-1], words[i], charMap) {
			return false
		}
	}

	return true
}

func compare(word1, word2 string, charMap [26]int) bool {

	if len(word1) == 0 {
		return true
	}

	if len(word2) == 0 {
		return false
	}

	for i := 0; i < len(word1) && i < len(word2); i++ {

		charAtWord1 := word1[i]
		charAtWord2 := word2[i]
		if charAtWord1 == charAtWord2 {
			continue
		}

		if charAtWord1 != charAtWord2 {
			if charMap[charAtWord1-'a'] > charMap[charAtWord2-'a'] {
				return false
			} else {
				return true
			}
		}

	}

	if len(word1) > len(word2) {

		return false
	}

	return true

}
