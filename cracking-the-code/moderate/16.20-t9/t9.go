package main

import "fmt"

func main() {
	fmt.Println(letterCombination([]int{8, 7, 3, 3}))
}

func letterCombination(digits []int) []string {

	result := []string{}
	if len(digits) == 0 {
		return []string{}
	}

	mapping := []string{
		"0",
		"1",
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"qprs",
		"tuv",
		"wxyz",
	}

	letterCombinationRecursion(&result, digits, "", 0, mapping)

	return result
}

func letterCombinationRecursion(result *[]string, digits []int, current string, index int, mapping []string) {

	if index == len(digits) { //purposely go over
		*result = append(*result, current)
		return
	}

	letters := mapping[digits[index]]
	for i := 0; i < len(letters); i++ {
		letterCombinationRecursion(result, digits, current+string(letters[i]), index+1, mapping)
	}

}
