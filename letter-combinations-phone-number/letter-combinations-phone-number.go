package main

import (
	"fmt"
)

func main() {

	input := "23"
	output := threeSum(input)
	fmt.Println(output)
}

func threeSum(digits string) []string {

	var output []string
	var arr []string
	if len(digits) != 0 {
		backtrack("", digits, 0, arr)
	}

	return output
}

func backtrack(prefix string, digit string, offset int, arr []string) {

	if offset >= len(digit) {
		arr = append(arr, prefix)
	}
	dict := make(map[string]string)
	dict["2"] = "abc"
	dict["3"] = "def"
	dict["4"] = "ghi"
	dict["5"] = "jkl"
	dict["6"] = "mno"
	dict["7"] = "pqrs"
	dict["8"] = "tuv"
	dict["9"] = "wxyz"

	letters := dict[string(digit[offset])]

	for i := 0; i < len(letters); i++ {
		letter := string(letters[i])
		backtrack(prefix+letter, digit, offset+1, arr)
	}

}
