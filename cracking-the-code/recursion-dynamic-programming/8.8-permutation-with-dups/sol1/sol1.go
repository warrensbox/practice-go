package main

import "fmt"

func main() {
	arr := getPerms("abcd")
	for _, a := range arr {
		fmt.Println(a)
	}
}

func getPerms(str string) []string {
	fmt.Println("here")
	if str == "" {
		fmt.Println("null")
		return nil
	}

	permutations := []string{}
	if len(str) == 0 { //base case
		permutations = append(permutations, "")
		return permutations
	}

	first := str[0] //get first char
	fmt.Println("first", string(first))
	remainder := str[1:]
	fmt.Println("remainder", remainder)

	words := getPerms(remainder)
	fmt.Println("words", words)
	for _, word := range words {
		for j := 0; j < len(word); j++ {
			s := insertCharAt(word, first, j)
			permutations = append(permutations, s)
		}
	}
	return permutations
}

func insertCharAt(word string, c byte, i int) string {
	start := word[0:i]
	end := word[i:]
	return start + string(c) + end
}
