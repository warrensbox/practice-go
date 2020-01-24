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

	var arrStr []string

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

		arrStr = append(arrStr, dict[num])

	}

	return arrStr
}
