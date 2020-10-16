package main

import (
	"fmt"
)

func main() {

	str1 := "tact cao"

	isPermutationOfPalindrome(str1)
}

func isPermutationOfPalindrome(str string) {

	catalog := make([]int, 256)

	for i := 0; i < len(str); i++ {
		catalog[str[i]-'a']++
	}

	foundOdd := false

	for i := 0; i < len(catalog); i++ {
		fmt.Println(catalog[i])
		if catalog[i]%2 == 1 {
			if foundOdd == true {
				fmt.Println("False")
				return
			}
			foundOdd = true
		}
	}

	fmt.Println("True")
	return
}
