package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		y = "wa"
		x = "terbottle"
		s1 = (y+x) = "waterbottle"
		s2 = = (x+y) ="terbottlewa"
	*/

	s1 := "waterbottle"
	s2 := "terbottlewa"

	isRotation(s1, s2)

}

func isRotation(str1, str2 string) {

	if len(str1) == len(str2) && len(str1) > 0 {
		concatStr := str2 + str2
		fmt.Println(concatStr)

		fmt.Println(strings.Contains(concatStr, str1))
	}
}
