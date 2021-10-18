package main

import "fmt"

func main() {
	arr := getPerms("abc")
	// for _, a := range arr {
	// 	fmt.Println(a)
	// }
	fmt.Println(arr)
}

func getPerms(str string) []string {

	if len(str) <= 1 { //base case
		fmt.Println("str", str)
		permutations := []string{}
		permutations = append(permutations, str)
		return permutations
	}

	lastChar := str[len(str)-1] //get first char
	allCharsExceptLast := str[:len(str)-1]
	fmt.Println("lastChar", string(lastChar))
	fmt.Println("allCharsExceptLast", allCharsExceptLast)

	permutationsOfAllCharsExceptLast := getPerms(allCharsExceptLast)
	fmt.Println("LAST CHAR", string(lastChar))
	fmt.Println("permutationsOfAllCharsExceptLast : ", permutationsOfAllCharsExceptLast)
	var permutations2 []string
	for _, word := range permutationsOfAllCharsExceptLast {
		fmt.Println("word", word)
		for j := 0; j <= len(word); j++ {
			s := insertCharAt(word, lastChar, j)
			fmt.Println("s", s)
			permutations2 = append(permutations2, s)
		}
	}
	fmt.Println("FINAL:", permutations2)
	return permutations2
}

func insertCharAt(word string, char byte, i int) string {
	start := word[0:i]
	end := word[i:]
	fmt.Println("start + string(char) + end", start+string(char)+end)
	return start + string(char) + end
}
