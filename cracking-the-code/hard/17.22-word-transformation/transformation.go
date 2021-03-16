package main

import "fmt"

func main() {

	Input := []string{"DAMP", "LIKE"}
	fmt.Println(Transform("DAMP", "LIKE", Input))

}

//still need to implement one word away

func Transform(start, stop string, words []string) []string {
	wildcardToWordList := createWildCardToMap(words)
	visited := make(map[string]bool)
	return transform(visited, start, stop, wildcardToWordList)
}

func transform(visited map[string]bool, start, stop string, wildcardToWordList map[string][]string) []string {

	if start == stop {
		path := make([]string, 0)
		path = append(path, start)
		return path
	} else if visited[start] {
		return nil
	}

	visited[start] = true
	words := getValidLinkedWords(start, wildcardToWordList)

	fmt.Println(words)
	for _, word := range words {
		path := transform(visited, word, stop, wildcardToWordList)
		if path != nil {
			path = append(path, start)
			return path
		}
	}
	return nil
}

func createWildCardToMap(words []string) map[string][]string {
	worldcardToWords := make(map[string][]string)
	for _, word := range words {
		fmt.Println("ttt", word)
		linked := getWilscardRoots(word)
		fmt.Println("linked", linked)
		for _, linkedWord := range linked {
			worldcardToWords[linkedWord] = append(worldcardToWords[linkedWord], word)
		}
	}
	fmt.Println("test", worldcardToWords)
	return worldcardToWords
}

func getWilscardRoots(w string) []string {
	words := make([]string, 0)
	for i := 1; i < len(w); i++ {
		word := string(w[0:i]) + "_" + string(w[i:])
		words = append(words, word)
	}
	return words
}

func getValidLinkedWords(word string, wildcardToWords map[string][]string) []string {

	wildcards := getWilscardRoots(word)
	fmt.Println(wildcards)
	linkedWords := make([]string, 0)
	for _, wildcard := range wildcards {
		words := wildcardToWords[wildcard]
		fmt.Println(words)
		for _, linkedWord := range words {
			if linkedWord != word {
				linkedWords = append(linkedWords, linkedWord)
			}
		}
	}
	fmt.Println(linkedWords)
	return linkedWords
}
