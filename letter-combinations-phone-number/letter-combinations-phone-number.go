package main

import (
	"fmt"
	"strconv"
)

func main() {

	input := "23"
	output := t9(input)
	fmt.Println(output)
}

func t9(digits string) []string {

	var output []string
	var arr []string

	dict := make(map[int]string)
	dict[2] = "abc"
	dict[3] = "def"
	dict[4] = "ghi"
	dict[5] = "jkl"
	dict[6] = "mno"
	dict[7] = "pqrs"
	dict[8] = "tuv"
	dict[9] = "wxyz"
	if len(digits) != 0 {
		backtrack("", 0, dict, digits, &arr)
	}

	return output
}

func backtrack(prefix string, offset int, dict map[int]string, digit string, arr *[]string) {
	fmt.Println(offset)
	if offset == len(digit) {
		*arr = append(*arr, prefix)
		return
	}
	val, _ := strconv.Atoi(string(digit[offset]))
	letters := dict[val]
	fmt.Println(val)
	for i := 0; i < len(letters); i++ {
		letter := string(letters[i])
		backtrack(prefix+letter, offset+1, dict, digit, arr)
	}

}
