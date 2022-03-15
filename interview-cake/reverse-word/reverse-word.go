package main

import (
	"fmt"
	"strings"
)

/*
goal: reverse word
*/

func main() {

	//message := []rune{'c', 'a', 'k', 'e', ' ','p', 'o', 'u', 'n', 'd', ' ','s', 't', 'e', 'a', 'l'}

	message2 := []rune{'a', 'b', 'c', ' ', 'd', 'e', 'f'}

	//fmt.Println(reverseWords(message))
	reverseWords2(message2)
}

func reverseWords(message []rune) string {
	fmt.Println(message)
	var result strings.Builder
	//var stack []string
	var temp string
	for i := 0; i < len(message); i++ {

		if message[i] == rune(' ') {
			//stack = append(stack, result.String())
			temp = " " + result.String() + temp
			result.Reset()
			//stack = append(stack, "#")
		} else if i == len(message)-1 {
			result.WriteRune(message[i])
			temp = result.String() + temp
			//stack = append(stack, result.String())
		} else {
			result.WriteRune(message[i])
		}
	}

	for i := 0; i < len(message); i++ {
		message[i] = rune(temp[i])
	}
	fmt.Println(message)

	return temp
}

func reverseWords2(message []rune) {
	fmt.Println(message)
	start := 0
	end := len(message) - 1
	arr := reverser(message, start, end)
	fmt.Println("len(arr)-1", len(arr)-1)
	currIndex := 0
	for i := 0; i <= len(arr)-1; i++ {
		fmt.Println(i)
		if arr[i] == 32 {
			reverser(arr, currIndex, i-1)
			currIndex = i + 1
		} else if i == len(arr)-1 {
			reverser(arr, currIndex, i)
		}
	}
	fmt.Println(string(message))
}

func reverser(arr []rune, start, end int) []rune {

	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}

	return arr
}
