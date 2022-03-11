package main

import (
	"fmt"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT

	arr := []string{"food", "fund", "fast", "dork", "ages", "book", "cook"}

	var sb strings.Builder
	for position := 0; position < 4; position++ {
		arrByte := []byte{}
		for i := 0; i < len(arr); i++ {
			word := arr[i]
			arrByte = append(arrByte, word[position])
		}
		letter := getHighestFref(arrByte)
		sb.WriteString(string(letter))
	}

	fmt.Println(sb.String())
}

func getHighestFref(arrByte []byte) byte {

	var output byte
	memoAlpha := make([]int, 26)
	for _, val := range arrByte {
		memoAlpha[val-'a']++
	}

	biggest := 0
	char := 0
	for i := 0; i < 26; i++ {
		if memoAlpha[i] > biggest {
			biggest = memoAlpha[i]
			char = i
		}
	}
	output = byte('a' + char)
	return output
}
