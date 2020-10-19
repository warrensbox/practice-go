package main

import "fmt"

func main() {

	str1 := "god"
	str2 := "dog"

	isPalin(str1, str2)

}

func isPalin(str1, str2 string) {

	catalog := make([]int, 128)
	for i := 0; i < len(str1); i++ {
		catalog[str1[i]-'a']++
	}

	for i := 0; i < len(str2); i++ {
		if catalog[str2[i]-'a'] > 0 {
			catalog[str2[i]-'a']--
		} else {
			fmt.Println("False")
			return
		}
	}

	fmt.Println("True")
}
