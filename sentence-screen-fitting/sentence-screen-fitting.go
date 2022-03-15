package main

import (
	"fmt"
	"strings"
)

func wordsTyping(sentences []string, rows int, cols int) int {
	var res strings.Builder
	for _, s := range sentences {
		res.WriteString(s + " ")
	}
	sentence := res.String()
	sentenceLength := len(res.String())

	fmt.Println(sentenceLength)

	cursor := 0

	for row := 0; row < rows; row++ {
		cursor += cols
		fmt.Println(cursor)
		//fmt.Println(cursor % sentenceLength)
		fmt.Println("test", cursor%sentenceLength)
		if string(sentence[cursor%sentenceLength]) == " " {
			cursor++
			fmt.Println("cursor1", cursor)
		} else {
			for cursor >= 0 && string(sentence[cursor%sentenceLength]) != " " {
				cursor--
				fmt.Println("cursor2", cursor)
			}
			cursor++
		}
	}
	return cursor / sentenceLength
}
