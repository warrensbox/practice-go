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
	if len(digits) != 0 {
		backtrack("", digits, output)
	}

	return output
}

func backtrack(combination string, nextDigit string, arr []string) {

	if len(nextDigit) == 0 {
		arr = append(arr, combination)
	} else {
		dict := make(map[string]string)
		dict["2"] = "abc"
		dict["3"] = "def"
		dict["4"] = "ghi"
		dict["5"] = "jkl"
		dict["6"] = "mno"
		dict["7"] = "pqrs"
		dict["8"] = "tuv"
		dict["9"] = "wxyz"

		digit := nextDigit[0:1]
		letters := dict[string(digit)]

		fmt.Println(letters)
		fmt.Println(digit)
		for i := 0; i < len(letters); i++ {
			letter := string(dict[digit][i : i+1])
			fmt.Println(letter)
			fmt.Println(string(nextDigit[1]))
			//backtrack(combination+letter, string(nextDigit[1]), arr)
		}
	}
}
