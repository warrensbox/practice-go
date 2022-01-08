package main

import "fmt"

/*


goal :find if word is a permutation of a palindrome

sol :
- each time we see new letter, we add that to hash
- if we see it again, we remove it

*/

func main() {
	fmt.Println(isPalindrome("civic"))
	fmt.Println(isPalindrome("ivicc"))
	fmt.Println(isPalindrome("civil"))
	fmt.Println(isPalindrome("livci"))
	fmt.Println(isPalindrome("haannnaah"))
	fmt.Println(isPalindrome("ahahanan"))

}

func isPalindrome(word string) bool {

	hash := make(map[rune]int)

	for _, letter := range word {
		if _, ok := hash[letter]; ok {
			delete(hash, letter)
		} else {
			hash[letter]++
		}
	}

	return len(hash) == 0 || len(hash) == 1
}
