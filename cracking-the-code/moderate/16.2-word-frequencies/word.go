package main

import "fmt"

func main() {
	book := []string{"jungle", "book", "jungle", "book"}

	table := freqCounter(book)

	fmt.Println(table["book"])
}

func freqCounter(book []string) map[string]int {

	hash := make(map[string]int)

	for _, key := range book {
		//	if _, ok := hash[key]; !ok {
		hash[key]++
		//	}
	}

	return hash
}
