package main

import "fmt"

func main() {

	input := "aba"
	output := validPalindrome(input)

	fmt.Println(output)

}

func validPalindrome(s string) bool {

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	processedString := reg.ReplaceAllString(s, "")

	processedString = strings.ToLower(processedString)

	j := len(processedString) - 1

	for i := 0; i < len(processedString)/2; i++ {
		if processedString[i] != processedString[j] {
			fmt.Println("Test")
			break
		}
		j--
	}

	return false
}

func isPanlindromeRange(s string, i, j) bool{

	for k := i; k <= i+(j-1)/2 ; k++{
		if s[i] != s[j-k+i] {
			return false
		}
		return true
	}
}
