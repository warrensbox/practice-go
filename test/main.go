package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	fmt.Println(generate("cofe applo "))
}

func generate(str string) int {

	str = strings.Replace(str, " ", "", -1)
	pattern := "facebook"
	pattenChar := make([]int, 26)
	for _, val := range pattern {
		pattenChar[val-'a']++
	}

	textChar := make([]int, 26)
	for _, val := range str {
		textChar[val-'a']++
	}
	max := math.MinInt32
	for _, val := range str {
		if pattenChar[val-'a'] > 0 {
			num := textChar[val-'a'] / pattenChar[val-'a']
			max = Max(max, num)
		}
	}

	return max
}

func Max(x, y int) int {

	if x > y {
		return x
	}

	return y
}
