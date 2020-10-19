package main

import (
	"fmt"
)

func main() {

	str1 := "pppale"
	str2 := "pale"

	oneWay(str1, str2)
}

func oneWay(str1, str2 string) {

	catalog := make([]int, 128)

	for i := 0; i < len(str1); i++ {
		catalog[str1[i]-'a']++
	}

	totalStr1 := len(str1)
	totalStr2 := len(str2)
	for i := 0; i < len((str2)); i++ {
		if catalog[str2[i]-'a'] > 0 {
			totalStr1 = totalStr1 - 1
			totalStr2 = totalStr2 - 1
			catalog[str2[i]-'a']--
		}
	}

	if totalStr1 == 1 {
		fmt.Println("true")
		return
	} else if totalStr1 == 0 {
		if totalStr2 <= 1 {
			fmt.Println("true")
			return
		}
	}

	fmt.Println("false")
	return
}
