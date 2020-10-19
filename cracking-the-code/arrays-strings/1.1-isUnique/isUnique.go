package main

import "fmt"

func main() {

	strA := "ABDEa"

	isUnique(strA)
}

func isUnique(str string) {

	hash := make(map[byte]bool)

	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
		if _, ok := hash[str[i]]; ok {
			fmt.Println("false")
			return
		}
		hash[str[i]] = true
	}

	fmt.Println("true")
}
