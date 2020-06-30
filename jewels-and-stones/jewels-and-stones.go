package main

import (
	"fmt"
)

func main() {
	J := "aA"
	S := "aAAbbbb"

	output := numJewelsInStones(J, S)
	fmt.Println(output)
}

func numJewelsInStones(J string, S string) int {

	count := 0
	dict := make(map[rune]bool)

	for _, val := range J {

		dict[val] = true
	}

	for _, val := range S {

		if ok, _ := dict[val]; ok {
			count++
		}
	}

	return count

}
