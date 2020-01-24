package main

import (
	"fmt"
	"strings"
)

func main() {

	input := "23"
	output := threeSum(input)
	fmt.Println(output)
}

func threeSum(digits string) []string {

	var arrStr []string
	var combineAll strings.Builder
	//var combinePair strings.Builder

	dict := make(map[string]string)

	dict["2"] = "abc"
	dict["3"] = "def"
	dict["4"] = "ghi"
	dict["5"] = "jkl"
	dict["6"] = "mno"
	dict["7"] = "pqrs"
	dict["8"] = "tuv"
	dict["9"] = "wxyz"

	for _, val := range digits {

		num := string(val)
		combineAll.WriteString(dict[num])

	}

	allLetters := combineAll.String()
	fmt.Println(string(allLetters))
	track := make(map[string]string)

	for i := 0; i < len(allLetters); i++ {
		//fmt.Println(string(allLetters[i]))
		for j := 0; j < len(allLetters)-1; j++ {
			fmt.Println(string(allLetters[i]) + string(allLetters[j]))

			_, ok := track[string(allLetters[j])+string(allLetters[i])]
			if !ok {

			}
		}
	}

	return arrStr
}
